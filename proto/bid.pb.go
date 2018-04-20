// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bid.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BidNetwork struct {
	Overlay  bool `protobuf:"varint,1,opt,name=overlay" json:"overlay,omitempty"`
	Outbound bool `protobuf:"varint,2,opt,name=outbound" json:"outbound,omitempty"`
	Incoming bool `protobuf:"varint,3,opt,name=incoming" json:"incoming,omitempty"`
}

func (m *BidNetwork) Reset()                    { *m = BidNetwork{} }
func (m *BidNetwork) String() string            { return proto.CompactTextString(m) }
func (*BidNetwork) ProtoMessage()               {}
func (*BidNetwork) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BidNetwork) GetOverlay() bool {
	if m != nil {
		return m.Overlay
	}
	return false
}

func (m *BidNetwork) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func (m *BidNetwork) GetIncoming() bool {
	if m != nil {
		return m.Incoming
	}
	return false
}

type BidResources struct {
	Network    *BidNetwork       `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Benchmarks map[string]uint64 `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *BidResources) Reset()                    { *m = BidResources{} }
func (m *BidResources) String() string            { return proto.CompactTextString(m) }
func (*BidResources) ProtoMessage()               {}
func (*BidResources) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *BidResources) GetNetwork() *BidNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *BidResources) GetBenchmarks() map[string]uint64 {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

type BidOrder struct {
	ID        string        `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Duration  *Duration     `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
	Price     *Price        `protobuf:"bytes,3,opt,name=price" json:"price,omitempty"`
	Blacklist *EthAddress   `protobuf:"bytes,4,opt,name=blacklist" json:"blacklist,omitempty"`
	Identity  IdentityLevel `protobuf:"varint,5,opt,name=identity,enum=sonm.IdentityLevel" json:"identity,omitempty"`
	Tag       string        `protobuf:"bytes,6,opt,name=tag" json:"tag,omitempty"`
	Resources *BidResources `protobuf:"bytes,7,opt,name=resources" json:"resources,omitempty"`
}

func (m *BidOrder) Reset()                    { *m = BidOrder{} }
func (m *BidOrder) String() string            { return proto.CompactTextString(m) }
func (*BidOrder) ProtoMessage()               {}
func (*BidOrder) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *BidOrder) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *BidOrder) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *BidOrder) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *BidOrder) GetBlacklist() *EthAddress {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *BidOrder) GetIdentity() IdentityLevel {
	if m != nil {
		return m.Identity
	}
	return IdentityLevel_ANONIMOUS
}

func (m *BidOrder) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *BidOrder) GetResources() *BidResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*BidNetwork)(nil), "sonm.BidNetwork")
	proto.RegisterType((*BidResources)(nil), "sonm.BidResources")
	proto.RegisterType((*BidOrder)(nil), "sonm.BidOrder")
}

func init() { proto.RegisterFile("bid.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6e, 0xa3, 0x30,
	0x10, 0xc6, 0x05, 0xf9, 0x07, 0xc3, 0x2a, 0xc9, 0x7a, 0xf7, 0x60, 0x71, 0xca, 0x72, 0x8a, 0x72,
	0x60, 0x57, 0xec, 0x65, 0xb5, 0x52, 0x0f, 0x45, 0xc9, 0x21, 0x52, 0xd5, 0x56, 0x7e, 0x80, 0x4a,
	0x06, 0x5b, 0x89, 0x05, 0xb1, 0x23, 0x63, 0x52, 0xe5, 0x25, 0xfa, 0x46, 0x7d, 0xb7, 0x0a, 0x1c,
	0x48, 0xd4, 0x9e, 0x60, 0xbe, 0xef, 0x37, 0xcc, 0xcc, 0x27, 0xc0, 0xcf, 0x04, 0x8b, 0x8f, 0x5a,
	0x19, 0x85, 0x86, 0x95, 0x92, 0x87, 0x70, 0x26, 0x64, 0xf3, 0x94, 0x82, 0x5a, 0x39, 0xfc, 0x7e,
	0xa0, 0xba, 0xe0, 0xe6, 0x58, 0xd2, 0x9c, 0x5b, 0x29, 0x7a, 0x01, 0x48, 0x05, 0x7b, 0xe4, 0xe6,
	0x55, 0xe9, 0x02, 0x61, 0x98, 0xa8, 0x13, 0xd7, 0x25, 0x3d, 0x63, 0x67, 0xe1, 0x2c, 0x3d, 0xd2,
	0x95, 0x28, 0x04, 0x4f, 0xd5, 0x26, 0x53, 0xb5, 0x64, 0xd8, 0x6d, 0xad, 0xbe, 0x6e, 0x3c, 0x21,
	0x73, 0x75, 0x10, 0x72, 0x87, 0x07, 0xd6, 0xeb, 0xea, 0xe8, 0xdd, 0x81, 0x6f, 0xa9, 0x60, 0x84,
	0x57, 0xaa, 0xd6, 0x39, 0xaf, 0xd0, 0x0a, 0x26, 0xd2, 0x4e, 0x6b, 0x47, 0x04, 0xc9, 0x3c, 0x6e,
	0x96, 0x8c, 0xaf, 0x5b, 0x90, 0x0e, 0x40, 0x29, 0x40, 0xc6, 0x65, 0xbe, 0x6f, 0xd6, 0xae, 0xb0,
	0xbb, 0x18, 0x2c, 0x83, 0x24, 0xea, 0xf1, 0xfe, 0x9b, 0x71, 0xda, 0x43, 0x1b, 0x69, 0xf4, 0x99,
	0xdc, 0x74, 0x85, 0x77, 0x30, 0xfb, 0x64, 0xa3, 0x39, 0x0c, 0x0a, 0x6e, 0x2f, 0xf4, 0x49, 0xf3,
	0x8a, 0x7e, 0xc2, 0xe8, 0x44, 0xcb, 0x9a, 0xb7, 0xa7, 0x0d, 0x89, 0x2d, 0xfe, 0xbb, 0xff, 0x9c,
	0xe8, 0xcd, 0x05, 0x2f, 0x15, 0xec, 0x49, 0x33, 0xae, 0xd1, 0x14, 0xdc, 0xed, 0xfa, 0xd2, 0xe7,
	0x6e, 0xd7, 0x68, 0x05, 0x1e, 0xab, 0x35, 0x35, 0x42, 0xc9, 0xb6, 0x33, 0x48, 0xa6, 0x76, 0xbb,
	0xf5, 0x45, 0x25, 0xbd, 0x8f, 0x7e, 0xc1, 0xe8, 0xa8, 0x45, 0xce, 0xdb, 0x84, 0x82, 0x24, 0xb0,
	0xe0, 0x73, 0x23, 0x11, 0xeb, 0xa0, 0x18, 0xfc, 0xac, 0xa4, 0x79, 0x51, 0x8a, 0xca, 0xe0, 0xe1,
	0x6d, 0x38, 0x1b, 0xb3, 0xbf, 0x67, 0x4c, 0xf3, 0xaa, 0x22, 0x57, 0x04, 0xfd, 0x06, 0x4f, 0x30,
	0x2e, 0x8d, 0x30, 0x67, 0x3c, 0x5a, 0x38, 0xcb, 0x69, 0xf2, 0xc3, 0xe2, 0xdb, 0x8b, 0xfa, 0xc0,
	0x4f, 0xbc, 0x24, 0x3d, 0xd4, 0x1c, 0x6e, 0xe8, 0x0e, 0x8f, 0xed, 0xe1, 0x86, 0xee, 0xd0, 0x1f,
	0xf0, 0x75, 0x17, 0x23, 0x9e, 0xb4, 0x23, 0xd1, 0xd7, 0x80, 0xc9, 0x15, 0xca, 0xc6, 0xed, 0x7f,
	0xf3, 0xf7, 0x23, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xbd, 0xad, 0x63, 0x6e, 0x02, 0x00, 0x00,
}
