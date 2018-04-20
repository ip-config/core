// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deal.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Deprecated: TODO: please migrate to the new API.
type DeprecatedDealStatus int32

const (
	DeprecatedDealStatus_ANY_STATUS DeprecatedDealStatus = 0
	DeprecatedDealStatus_PENDING    DeprecatedDealStatus = 1
	DeprecatedDealStatus_ACCEPTED   DeprecatedDealStatus = 2
	DeprecatedDealStatus_CLOSED     DeprecatedDealStatus = 3
)

var DeprecatedDealStatus_name = map[int32]string{
	0: "ANY_STATUS",
	1: "PENDING",
	2: "ACCEPTED",
	3: "CLOSED",
}
var DeprecatedDealStatus_value = map[string]int32{
	"ANY_STATUS": 0,
	"PENDING":    1,
	"ACCEPTED":   2,
	"CLOSED":     3,
}

func (x DeprecatedDealStatus) String() string {
	return proto.EnumName(DeprecatedDealStatus_name, int32(x))
}
func (DeprecatedDealStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// Deprecated: TODO: please migrate to the new API.
type DeprecatedDeal struct {
	BuyerID           string               `protobuf:"bytes,1,opt,name=BuyerID" json:"BuyerID,omitempty"`
	SupplierID        string               `protobuf:"bytes,2,opt,name=SupplierID" json:"SupplierID,omitempty"`
	Status            DeprecatedDealStatus `protobuf:"varint,3,opt,name=status,enum=sonm.DeprecatedDealStatus" json:"status,omitempty"`
	Price             *BigInt              `protobuf:"bytes,4,opt,name=price" json:"price,omitempty"`
	StartTime         *Timestamp           `protobuf:"bytes,5,opt,name=startTime" json:"startTime,omitempty"`
	EndTime           *Timestamp           `protobuf:"bytes,6,opt,name=endTime" json:"endTime,omitempty"`
	SpecificationHash string               `protobuf:"bytes,7,opt,name=SpecificationHash" json:"SpecificationHash,omitempty"`
	WorkTime          uint64               `protobuf:"varint,8,opt,name=workTime" json:"workTime,omitempty"`
	Id                string               `protobuf:"bytes,9,opt,name=id" json:"id,omitempty"`
	Benchmarks        []uint64             `protobuf:"varint,10,rep,packed,name=benchmarks" json:"benchmarks,omitempty"`
	AskID             string               `protobuf:"bytes,11,opt,name=AskID" json:"AskID,omitempty"`
	BidID             string               `protobuf:"bytes,12,opt,name=BidID" json:"BidID,omitempty"`
}

func (m *DeprecatedDeal) Reset()                    { *m = DeprecatedDeal{} }
func (m *DeprecatedDeal) String() string            { return proto.CompactTextString(m) }
func (*DeprecatedDeal) ProtoMessage()               {}
func (*DeprecatedDeal) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *DeprecatedDeal) GetBuyerID() string {
	if m != nil {
		return m.BuyerID
	}
	return ""
}

func (m *DeprecatedDeal) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *DeprecatedDeal) GetStatus() DeprecatedDealStatus {
	if m != nil {
		return m.Status
	}
	return DeprecatedDealStatus_ANY_STATUS
}

func (m *DeprecatedDeal) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *DeprecatedDeal) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *DeprecatedDeal) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *DeprecatedDeal) GetSpecificationHash() string {
	if m != nil {
		return m.SpecificationHash
	}
	return ""
}

func (m *DeprecatedDeal) GetWorkTime() uint64 {
	if m != nil {
		return m.WorkTime
	}
	return 0
}

func (m *DeprecatedDeal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeprecatedDeal) GetBenchmarks() []uint64 {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *DeprecatedDeal) GetAskID() string {
	if m != nil {
		return m.AskID
	}
	return ""
}

func (m *DeprecatedDeal) GetBidID() string {
	if m != nil {
		return m.BidID
	}
	return ""
}

func init() {
	proto.RegisterType((*DeprecatedDeal)(nil), "sonm.DeprecatedDeal")
	proto.RegisterEnum("sonm.DeprecatedDealStatus", DeprecatedDealStatus_name, DeprecatedDealStatus_value)
}

func init() { proto.RegisterFile("deal.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x8e, 0xda, 0x30,
	0x14, 0x85, 0x9b, 0x1f, 0x12, 0xb8, 0x44, 0x21, 0xb5, 0x58, 0x58, 0x59, 0x54, 0x11, 0xab, 0xb4,
	0x6a, 0x59, 0xd0, 0x27, 0x08, 0x38, 0x6a, 0x23, 0xb5, 0x14, 0x25, 0xe9, 0xa2, 0xab, 0xca, 0x24,
	0x1e, 0xb0, 0x20, 0x3f, 0x8a, 0x8d, 0x46, 0xf3, 0x3a, 0xf3, 0xa4, 0x23, 0x9c, 0x61, 0x86, 0xd1,
	0xb0, 0x3c, 0xe7, 0x7c, 0xe7, 0xca, 0xf7, 0x1a, 0xa0, 0x64, 0xf4, 0x38, 0x6f, 0xbb, 0x46, 0x36,
	0xc8, 0x14, 0x4d, 0x5d, 0xf9, 0xce, 0x96, 0xef, 0x78, 0x2d, 0x7b, 0xcf, 0x9f, 0x48, 0x5e, 0x31,
	0x21, 0x69, 0xd5, 0xf6, 0xc6, 0xec, 0xd1, 0x00, 0x97, 0xb0, 0xb6, 0x63, 0x05, 0x95, 0xac, 0x24,
	0x8c, 0x1e, 0x11, 0x06, 0x7b, 0x79, 0x7a, 0x60, 0x5d, 0x42, 0xb0, 0x16, 0x68, 0xe1, 0x28, 0xbd,
	0x48, 0xf4, 0x09, 0x20, 0x3b, 0xb5, 0xed, 0x91, 0xab, 0x50, 0x57, 0xe1, 0x95, 0x83, 0x16, 0x60,
	0x09, 0x49, 0xe5, 0x49, 0x60, 0x23, 0xd0, 0x42, 0x77, 0xe1, 0xcf, 0xcf, 0x4f, 0x98, 0xbf, 0x9d,
	0x9f, 0x29, 0x22, 0x7d, 0x26, 0xd1, 0x0c, 0x06, 0x6d, 0xc7, 0x0b, 0x86, 0xcd, 0x40, 0x0b, 0xc7,
	0x0b, 0xa7, 0xaf, 0x2c, 0xf9, 0x2e, 0xa9, 0x65, 0xda, 0x47, 0xe8, 0x1b, 0x8c, 0x84, 0xa4, 0x9d,
	0xcc, 0x79, 0xc5, 0xf0, 0x40, 0x71, 0x93, 0x9e, 0xcb, 0x2f, 0xeb, 0xa4, 0xaf, 0x04, 0xfa, 0x0c,
	0x36, 0xab, 0x4b, 0x05, 0x5b, 0xb7, 0xe1, 0x4b, 0x8e, 0xbe, 0xc2, 0xc7, 0xac, 0x65, 0x05, 0xbf,
	0xe3, 0x05, 0x95, 0xbc, 0xa9, 0x7f, 0x52, 0xb1, 0xc7, 0xb6, 0x5a, 0xec, 0x7d, 0x80, 0x7c, 0x18,
	0xde, 0x37, 0xdd, 0x41, 0x4d, 0x1e, 0x06, 0x5a, 0x68, 0xa6, 0x2f, 0x1a, 0xb9, 0xa0, 0xf3, 0x12,
	0x8f, 0x54, 0x55, 0xe7, 0xe5, 0xf9, 0x56, 0x5b, 0x56, 0x17, 0xfb, 0x8a, 0x76, 0x07, 0x81, 0x21,
	0x30, 0x42, 0x33, 0xbd, 0x72, 0xd0, 0x14, 0x06, 0x91, 0x38, 0x24, 0x04, 0x8f, 0x55, 0xa5, 0x17,
	0x67, 0x77, 0xc9, 0xcb, 0x84, 0x60, 0xa7, 0x77, 0x95, 0xf8, 0xf2, 0x1b, 0xa6, 0xb7, 0x6e, 0x88,
	0x5c, 0x80, 0x68, 0xfd, 0xef, 0x7f, 0x96, 0x47, 0xf9, 0xdf, 0xcc, 0xfb, 0x80, 0xc6, 0x60, 0x6f,
	0xe2, 0x35, 0x49, 0xd6, 0x3f, 0x3c, 0x0d, 0x39, 0x30, 0x8c, 0x56, 0xab, 0x78, 0x93, 0xc7, 0xc4,
	0xd3, 0x11, 0x80, 0xb5, 0xfa, 0xf5, 0x27, 0x8b, 0x89, 0x67, 0x6c, 0x2d, 0xf5, 0xf5, 0xdf, 0x9f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x33, 0xd8, 0x5d, 0xde, 0x2d, 0x02, 0x00, 0x00,
}
