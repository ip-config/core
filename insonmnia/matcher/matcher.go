package matcher

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/noxiouz/zapctx/ctxlog"
	errors "github.com/pkg/errors"
	"github.com/sonm-io/core/blockchain"
	"github.com/sonm-io/core/insonmnia/dwh"
	"github.com/sonm-io/core/proto"
	pb "github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
	"go.uber.org/zap"
)

type Matcher interface {
	CreateDealByOrder(ctx context.Context, order *sonm.Order) (*sonm.Deal, error)
}

// YAMLConfig is embeddable config that can be integrated with
// another component's config.
type YAMLConfig struct {
	PollDelay  time.Duration `yaml:"poll_delay"`
	QueryLimit uint64        `yaml:"query_limit"`
}

type Config struct {
	Key        *ecdsa.PrivateKey
	PollDelay  time.Duration
	DWH        sonm.DWHClient
	Eth        blockchain.API
	QueryLimit uint64
}

func (c *Config) validate() error {
	err := &multierror.Error{ErrorFormat: util.MultierrFormat()}

	if c.QueryLimit == 0 {
		c.QueryLimit = dwh.MaxLimit
	}

	if c.Key == nil {
		err = multierror.Append(err, errors.New("private key is required"))
	}
	if c.PollDelay < time.Second {
		err = multierror.Append(err, errors.New("poll delay is too small"))
	}
	if c.DWH == nil {
		err = multierror.Append(err, errors.New("DWH client is required"))
	}
	if c.Eth == nil {
		err = multierror.Append(err, errors.New("blockchain market client is required"))
	}

	return err.ErrorOrNil()
}

type matcher struct {
	cfg *Config
}

func NewMatcher(cfg *Config) (Matcher, error) {
	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "invalid matcher config")
	}

	return &matcher{cfg: cfg}, nil
}

func (m *matcher) CreateDealByOrder(ctx context.Context, order *sonm.Order) (*sonm.Deal, error) {
	id := order.GetId().Unwrap()
	ctxlog.G(ctx).Debug("starting matcher", zap.String("orderID", id.String()))

	tk := util.NewImmediateTicker(m.cfg.PollDelay)
	defer tk.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-tk.C:
			if err := m.checkIfOrderExists(ctx, id); err != nil {
				return nil, err
			}

			matchingOrders, err := m.getMatchingOrders(ctx, id)
			if err != nil {
				// dwh failure is not critical, we must survive it
				ctxlog.S(ctx).Debugf("failed to get matching orders from DWH: %s", err)
				continue

			}

			if len(matchingOrders) == 0 {
				continue
			}

			// 3. iterate over sorted orders
			for _, dealWithMe := range matchingOrders {
				bid, ask, err := m.reorderOrders(order, dealWithMe)
				if err != nil {
					return nil, err
				}

				// 4. try to open deal
				deal, err := m.openDeal(ctx, bid, ask)
				if err == nil {
					ctxlog.G(ctx).Debug("deal is opened",
						zap.String("bid", bid.GetId().Unwrap().String()),
						zap.String("ask", ask.GetId().Unwrap().String()),
						zap.String("deal", deal.GetId().Unwrap().String()))
					return deal, nil
				}

				// 5. if deal is not created - wait for timeout and goto 1
				ctxlog.G(ctx).Warn("cannot open deal",
					zap.Error(err),
					zap.String("bid", bid.GetId().Unwrap().String()),
					zap.String("ask", ask.GetId().Unwrap().String()))
			}
		}
	}
}

func (m *matcher) checkIfOrderExists(ctx context.Context, id *big.Int) error {
	order, err := m.cfg.Eth.Market().GetOrderInfo(ctx, id)
	if err != nil {
		return err
	}

	if order.GetOrderStatus() != sonm.OrderStatus_ORDER_ACTIVE {
		return errors.New("unable to match inactive order")
	}

	return nil
}

func (m *matcher) getMatchingOrders(ctx context.Context, id *big.Int) ([]*sonm.Order, error) {
	dwhReply, err := m.cfg.DWH.GetMatchingOrders(ctx, &sonm.MatchingOrdersRequest{
		Id:    pb.NewBigInt(id),
		Limit: m.cfg.QueryLimit,
	})

	if err != nil {
		return nil, err
	}

	orders := make([]*sonm.Order, len(dwhReply.GetOrders()))
	for idx, ord := range dwhReply.GetOrders() {
		orders[idx] = ord.GetOrder()
	}

	return orders, nil
}

func (m *matcher) openDeal(ctx context.Context, bid, ask *sonm.Order) (*sonm.Deal, error) {
	askID := ask.GetId().Unwrap()
	bidID := bid.GetId().Unwrap()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-m.cfg.Eth.Market().OpenDeal(ctx, m.cfg.Key, askID, bidID):
		return res.Deal, res.Err
	}
}

func (m *matcher) reorderOrders(one, two *sonm.Order) (bid, ask *sonm.Order, err error) {
	// just a sanity check, orders must have different types
	if one.GetOrderType() == two.GetOrderType() {
		return nil, nil, errors.New("orders must have different types")
	}

	if one.GetOrderType() == sonm.OrderType_BID {
		bid, ask = one, two
	} else {
		ask, bid = one, two
	}

	return
}

type disabledMatcher struct{}

// NewDisabledMatcher return `Matcher` interface implementation that does nothing.
func NewDisabledMatcher() Matcher {
	return &disabledMatcher{}
}

func (disabledMatcher) CreateDealByOrder(context.Context, *sonm.Order) (*sonm.Deal, error) {
	return nil, errors.New("matcher disabled")
}
