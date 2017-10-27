// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetOrdersRequest struct {
	// Slot describe resuorces to search for
	Slot *Slot `protobuf:"bytes,1,opt,name=slot" json:"slot,omitempty"`
	// OrderType describe a type of orders to search
	OrderType OrderType `protobuf:"varint,2,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	// Count describe how namy results must be returned (order by price)
	Count uint64 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *GetOrdersRequest) Reset()                    { *m = GetOrdersRequest{} }
func (m *GetOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersRequest) ProtoMessage()               {}
func (*GetOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *GetOrdersRequest) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func (m *GetOrdersRequest) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *GetOrdersRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetOrdersReply struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersReply) Reset()                    { *m = GetOrdersReply{} }
func (m *GetOrdersReply) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersReply) ProtoMessage()               {}
func (*GetOrdersReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *GetOrdersReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOrdersRequest)(nil), "sonm.GetOrdersRequest")
	proto.RegisterType((*GetOrdersReply)(nil), "sonm.GetOrdersReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Market service

type MarketClient interface {
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error)
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error)
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	CancelOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Empty, error)
}

type marketClient struct {
	cc *grpc.ClientConn
}

func NewMarketClient(cc *grpc.ClientConn) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrderByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Market service

type MarketServer interface {
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersReply, error)
	GetOrderByID(context.Context, *ID) (*Order, error)
	CreateOrder(context.Context, *Order) (*Order, error)
	CancelOrder(context.Context, *Order) (*Empty, error)
}

func RegisterMarketServer(s *grpc.Server, srv MarketServer) {
	s.RegisterService(&_Market_serviceDesc, srv)
}

func _Market_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _Market_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Market_GetOrders_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Market_GetOrderByID_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Market_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Market_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x59, 0x8b, 0x44, 0x06, 0xd3, 0xea, 0xa4, 0x31, 0x84, 0x83, 0x21, 0x78, 0x10, 0x0f,
	0x72, 0xc0, 0x78, 0xf2, 0xa6, 0x35, 0xa6, 0x07, 0xd3, 0x04, 0x7d, 0x01, 0x4a, 0xe7, 0x40, 0x5c,
	0xd8, 0x15, 0xb6, 0x31, 0xfb, 0x7e, 0x3e, 0x98, 0x61, 0xd7, 0x2a, 0xf1, 0xcf, 0x71, 0x7f, 0xf3,
	0xfb, 0x66, 0xf2, 0x2d, 0x1c, 0x37, 0x65, 0xf7, 0x42, 0x4a, 0xf2, 0xb2, 0xa2, 0x4c, 0x76, 0x42,
	0x09, 0x74, 0x7b, 0xd1, 0x36, 0x91, 0xbf, 0xae, 0x37, 0x16, 0x44, 0xb3, 0xba, 0x1d, 0x50, 0x5b,
	0x97, 0x16, 0x24, 0x6f, 0x70, 0xf4, 0x40, 0x6a, 0xd5, 0x6d, 0xa8, 0xeb, 0x0b, 0x7a, 0xdd, 0x52,
	0xaf, 0xf0, 0x14, 0xdc, 0x9e, 0x0b, 0x15, 0xb2, 0x98, 0xa5, 0x41, 0x0e, 0xd9, 0x90, 0xc8, 0x9e,
	0xb8, 0x50, 0x85, 0xe1, 0x78, 0x09, 0xbe, 0x18, 0x02, 0xcf, 0x5a, 0x52, 0xb8, 0x17, 0xb3, 0x74,
	0x9a, 0xcf, 0xac, 0xb4, 0xda, 0xe1, 0xe2, 0xdb, 0xc0, 0x39, 0xec, 0x57, 0x62, 0xdb, 0xaa, 0x70,
	0x12, 0xb3, 0xd4, 0x2d, 0xec, 0x23, 0xb9, 0x86, 0xe9, 0xe8, 0xb0, 0xe4, 0x1a, 0xcf, 0xc0, 0x33,
	0xa1, 0x3e, 0x64, 0xf1, 0x24, 0x0d, 0xf2, 0x60, 0xb4, 0xb3, 0xf8, 0x1c, 0xe5, 0xef, 0x0c, 0xbc,
	0x47, 0xd3, 0x13, 0x6f, 0xc0, 0xff, 0xda, 0x80, 0x27, 0x56, 0xfe, 0xd9, 0x25, 0x9a, 0xff, 0xe2,
	0x92, 0xeb, 0xc4, 0xc1, 0x73, 0x38, 0xdc, 0xb1, 0x5b, 0xbd, 0x5c, 0xe0, 0x81, 0xf5, 0x96, 0x8b,
	0x68, 0x7c, 0x36, 0x71, 0xf0, 0x02, 0x82, 0xbb, 0x8e, 0x4a, 0x45, 0x06, 0xe0, 0x78, 0xfa, 0x97,
	0x5a, 0xb6, 0x15, 0xf1, 0xff, 0xd5, 0xfb, 0x46, 0x2a, 0x9d, 0x38, 0x6b, 0xcf, 0xfc, 0xfe, 0xd5,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0xe1, 0x61, 0xcc, 0xb4, 0x01, 0x00, 0x00,
}
