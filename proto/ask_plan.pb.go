// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ask_plan.proto

/*
Package sonm is a generated protocol buffer package.

It is generated from these files:
	ask_plan.proto
	benchmarks.proto
	bid.proto
	bigint.proto
	capabilities.proto
	container.proto
	deal.proto
	dwh.proto
	hub.proto
	insonmnia.proto
	marketplace.proto
	miner.proto
	net.proto
	node.proto
	relay.proto
	rendezvous.proto
	timestamp.proto
	volume.proto

It has these top-level messages:
	AskPlanCPU
	AskPlanGPU
	AskPlanRAM
	AskPlanStorage
	AskPlanNetwork
	AskPlanResources
	AskPlan
	Benchmark
	BidNetwork
	BidResources
	BidOrder
	Resources
	BigInt
	Capabilities
	CPUDevice
	CPU
	RAMDevice
	RAM
	GPUDevice
	GPU
	NetworkDevice
	Network
	StorageDevice
	Storage
	NetworkSpec
	Container
	Deal
	DealsRequest
	DealsReply
	DWHDeal
	DealConditionsRequest
	DealConditionsReply
	OrdersRequest
	MatchingOrdersRequest
	OrdersReply
	DWHOrder
	DealCondition
	WorkerAnnouncement
	ProfilesRequest
	ProfilesReply
	Profile
	BlacklistReply
	ValidatorsRequest
	ValidatorsReply
	Validator
	DealChangeRequestsReply
	WorkersRequest
	WorkersReply
	Certificate
	DWHBenchmarks
	DWHBenchmarkConditions
	MaxMinUint64
	MaxMinBig
	StartTaskRequest
	HubJoinNetworkRequest
	StartTaskReply
	HubStatusReply
	DealRequest
	ApproveDealRequest
	AskPlansReply
	TaskListReply
	DevicesReply
	PullTaskRequest
	DealInfoReply
	Empty
	ID
	TaskID
	PingReply
	CPUUsage
	MemoryUsage
	NetworkUsage
	ResourceUsage
	InfoReply
	TaskStatusReply
	AvailableResources
	StatusMapReply
	ContainerRestartPolicy
	TaskLogsRequest
	TaskLogsChunk
	DiscoverHubRequest
	TaskResourceRequirements
	Chunk
	Progress
	Duration
	EthAddress
	DataSize
	DataSizeRate
	Price
	GetOrdersRequest
	GetOrdersReply
	GetProcessingReply
	TouchOrdersRequest
	MarketOrder
	MarketDeal
	MinerStartRequest
	MinerStartReply
	TaskInfo
	Endpoints
	SaveRequest
	Addr
	SocketAddr
	JoinNetworkRequest
	TaskListRequest
	DealListRequest
	DealListReply
	DealStatusReply
	Worker
	WorkerListReply
	HandshakeRequest
	DiscoverResponse
	HandshakeResponse
	RelayClusterReply
	RelayMetrics
	NetMetrics
	ConnectRequest
	PublishRequest
	RendezvousReply
	RendezvousState
	RendezvousMeeting
	ResolveMetaReply
	Timestamp
	Volume
*/
package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AskPlanCPU struct {
	CorePercents uint64 `protobuf:"varint,1,opt,name=core_percents,json=corePercents" json:"core_percents,omitempty"`
}

func (m *AskPlanCPU) Reset()                    { *m = AskPlanCPU{} }
func (m *AskPlanCPU) String() string            { return proto.CompactTextString(m) }
func (*AskPlanCPU) ProtoMessage()               {}
func (*AskPlanCPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AskPlanCPU) GetCorePercents() uint64 {
	if m != nil {
		return m.CorePercents
	}
	return 0
}

type AskPlanGPU struct {
	Indexes []string `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	Hashes  []string `protobuf:"bytes,2,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *AskPlanGPU) Reset()                    { *m = AskPlanGPU{} }
func (m *AskPlanGPU) String() string            { return proto.CompactTextString(m) }
func (*AskPlanGPU) ProtoMessage()               {}
func (*AskPlanGPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AskPlanGPU) GetIndexes() []string {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *AskPlanGPU) GetHashes() []string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type AskPlanRAM struct {
	Size *DataSize `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
}

func (m *AskPlanRAM) Reset()                    { *m = AskPlanRAM{} }
func (m *AskPlanRAM) String() string            { return proto.CompactTextString(m) }
func (*AskPlanRAM) ProtoMessage()               {}
func (*AskPlanRAM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AskPlanRAM) GetSize() *DataSize {
	if m != nil {
		return m.Size
	}
	return nil
}

type AskPlanStorage struct {
	Size *DataSize `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
}

func (m *AskPlanStorage) Reset()                    { *m = AskPlanStorage{} }
func (m *AskPlanStorage) String() string            { return proto.CompactTextString(m) }
func (*AskPlanStorage) ProtoMessage()               {}
func (*AskPlanStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AskPlanStorage) GetSize() *DataSize {
	if m != nil {
		return m.Size
	}
	return nil
}

type AskPlanNetwork struct {
	ThroughputIn  *DataSizeRate `protobuf:"bytes,1,opt,name=throughputIn" json:"throughputIn,omitempty"`
	ThroughputOut *DataSizeRate `protobuf:"bytes,2,opt,name=throughputOut" json:"throughputOut,omitempty"`
	Overlay       bool          `protobuf:"varint,3,opt,name=overlay" json:"overlay,omitempty"`
	Outbound      bool          `protobuf:"varint,4,opt,name=outbound" json:"outbound,omitempty"`
	Incoming      bool          `protobuf:"varint,5,opt,name=incoming" json:"incoming,omitempty"`
}

func (m *AskPlanNetwork) Reset()                    { *m = AskPlanNetwork{} }
func (m *AskPlanNetwork) String() string            { return proto.CompactTextString(m) }
func (*AskPlanNetwork) ProtoMessage()               {}
func (*AskPlanNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AskPlanNetwork) GetThroughputIn() *DataSizeRate {
	if m != nil {
		return m.ThroughputIn
	}
	return nil
}

func (m *AskPlanNetwork) GetThroughputOut() *DataSizeRate {
	if m != nil {
		return m.ThroughputOut
	}
	return nil
}

func (m *AskPlanNetwork) GetOverlay() bool {
	if m != nil {
		return m.Overlay
	}
	return false
}

func (m *AskPlanNetwork) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func (m *AskPlanNetwork) GetIncoming() bool {
	if m != nil {
		return m.Incoming
	}
	return false
}

type AskPlanResources struct {
	CPU     *AskPlanCPU     `protobuf:"bytes,1,opt,name=CPU" json:"CPU,omitempty"`
	RAM     *AskPlanRAM     `protobuf:"bytes,2,opt,name=RAM" json:"RAM,omitempty"`
	Storage *AskPlanStorage `protobuf:"bytes,3,opt,name=storage" json:"storage,omitempty"`
	GPU     *AskPlanGPU     `protobuf:"bytes,4,opt,name=GPU" json:"GPU,omitempty"`
	Network *AskPlanNetwork `protobuf:"bytes,5,opt,name=network" json:"network,omitempty"`
}

func (m *AskPlanResources) Reset()                    { *m = AskPlanResources{} }
func (m *AskPlanResources) String() string            { return proto.CompactTextString(m) }
func (*AskPlanResources) ProtoMessage()               {}
func (*AskPlanResources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AskPlanResources) GetCPU() *AskPlanCPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *AskPlanResources) GetRAM() *AskPlanRAM {
	if m != nil {
		return m.RAM
	}
	return nil
}

func (m *AskPlanResources) GetStorage() *AskPlanStorage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *AskPlanResources) GetGPU() *AskPlanGPU {
	if m != nil {
		return m.GPU
	}
	return nil
}

func (m *AskPlanResources) GetNetwork() *AskPlanNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

type AskPlan struct {
	ID        string            `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	MarketID  string            `protobuf:"bytes,2,opt,name=marketID" json:"marketID,omitempty"`
	Duration  *Duration         `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	Price     *Price            `protobuf:"bytes,4,opt,name=price" json:"price,omitempty"`
	Blacklist *EthAddress       `protobuf:"bytes,5,opt,name=blacklist" json:"blacklist,omitempty"`
	Resources *AskPlanResources `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
}

func (m *AskPlan) Reset()                    { *m = AskPlan{} }
func (m *AskPlan) String() string            { return proto.CompactTextString(m) }
func (*AskPlan) ProtoMessage()               {}
func (*AskPlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AskPlan) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AskPlan) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *AskPlan) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AskPlan) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *AskPlan) GetBlacklist() *EthAddress {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *AskPlan) GetResources() *AskPlanResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*AskPlanCPU)(nil), "sonm.AskPlanCPU")
	proto.RegisterType((*AskPlanGPU)(nil), "sonm.AskPlanGPU")
	proto.RegisterType((*AskPlanRAM)(nil), "sonm.AskPlanRAM")
	proto.RegisterType((*AskPlanStorage)(nil), "sonm.AskPlanStorage")
	proto.RegisterType((*AskPlanNetwork)(nil), "sonm.AskPlanNetwork")
	proto.RegisterType((*AskPlanResources)(nil), "sonm.AskPlanResources")
	proto.RegisterType((*AskPlan)(nil), "sonm.AskPlan")
}

func init() { proto.RegisterFile("ask_plan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8a, 0xdb, 0x30,
	0x18, 0xc4, 0x49, 0x36, 0x89, 0xbf, 0xec, 0xa6, 0x8b, 0x28, 0x8b, 0xd9, 0xd3, 0xd6, 0xbd, 0x84,
	0x1e, 0x4c, 0xbb, 0x5d, 0x4a, 0x4f, 0x05, 0xb3, 0x29, 0x21, 0x87, 0xb4, 0x46, 0x4b, 0xce, 0x8b,
	0x62, 0x8b, 0x58, 0xd8, 0x91, 0x8c, 0x24, 0xf7, 0x27, 0xcf, 0xd9, 0x73, 0x1f, 0xa0, 0x4f, 0x51,
	0x24, 0xcb, 0x0e, 0x0e, 0x14, 0x7a, 0x0a, 0xf3, 0xcd, 0x8c, 0x3c, 0xdf, 0x48, 0x81, 0x39, 0x51,
	0xc5, 0x73, 0x55, 0x12, 0x1e, 0x55, 0x52, 0x68, 0x81, 0x46, 0x4a, 0xf0, 0xc3, 0xed, 0x0b, 0xc6,
	0xcd, 0x2f, 0x67, 0xa4, 0x19, 0x87, 0xef, 0x00, 0x62, 0x55, 0x24, 0x25, 0xe1, 0x8f, 0xc9, 0x16,
	0xbd, 0x86, 0xab, 0x54, 0x48, 0xfa, 0x5c, 0x51, 0x99, 0x52, 0xae, 0x55, 0xe0, 0xdd, 0x79, 0x8b,
	0x11, 0xbe, 0x34, 0xc3, 0xc4, 0xcd, 0xc2, 0x4f, 0x9d, 0x65, 0x95, 0x6c, 0x51, 0x00, 0x13, 0xc6,
	0x33, 0xfa, 0x83, 0x1a, 0xf1, 0x70, 0xe1, 0xe3, 0x16, 0xa2, 0x1b, 0x18, 0xe7, 0x44, 0xe5, 0x54,
	0x05, 0x03, 0x4b, 0x38, 0x14, 0xbe, 0xed, 0xfc, 0x38, 0xde, 0xa0, 0x10, 0x46, 0x8a, 0x1d, 0xa9,
	0xfd, 0xd2, 0xec, 0x7e, 0x1e, 0x99, 0x78, 0xd1, 0x92, 0x68, 0xf2, 0xc4, 0x8e, 0x14, 0x5b, 0x2e,
	0x7c, 0x80, 0xb9, 0x73, 0x3c, 0x69, 0x21, 0xc9, 0x9e, 0xfe, 0x97, 0xeb, 0x97, 0xd7, 0xd9, 0xbe,
	0x50, 0xfd, 0x5d, 0xc8, 0x02, 0x7d, 0x80, 0x4b, 0x9d, 0x4b, 0x51, 0xef, 0xf3, 0xaa, 0xd6, 0x6b,
	0xee, 0xec, 0xe8, 0xcc, 0x4e, 0x34, 0xc5, 0x3d, 0x1d, 0xfa, 0x08, 0x57, 0x27, 0xfc, 0xb5, 0xd6,
	0xc1, 0xe0, 0x9f, 0xc6, 0xbe, 0xd0, 0xd4, 0x23, 0xbe, 0x51, 0x59, 0x92, 0x9f, 0xc1, 0xf0, 0xce,
	0x5b, 0x4c, 0x71, 0x0b, 0xd1, 0x2d, 0x4c, 0x45, 0xad, 0x77, 0xa2, 0xe6, 0x59, 0x30, 0xb2, 0x54,
	0x87, 0x0d, 0xc7, 0x78, 0x2a, 0x0e, 0x8c, 0xef, 0x83, 0x8b, 0x86, 0x6b, 0x71, 0xf8, 0xdb, 0x83,
	0xeb, 0xb6, 0x3f, 0xaa, 0x44, 0x2d, 0x53, 0xaa, 0x50, 0x08, 0xc3, 0xc7, 0x64, 0xeb, 0xf6, 0xb9,
	0x6e, 0x62, 0x9d, 0xee, 0x15, 0x1b, 0xd2, 0x68, 0x70, 0xbc, 0x71, 0xd1, 0xfb, 0x1a, 0x1c, 0x6f,
	0xb0, 0x21, 0x51, 0x04, 0x13, 0xd5, 0x54, 0x6c, 0xe3, 0xce, 0xee, 0x5f, 0xf6, 0x74, 0xae, 0x7e,
	0xdc, 0x8a, 0xcc, 0x99, 0xab, 0x64, 0x6b, 0xf3, 0x9f, 0x9f, 0xb9, 0x32, 0xdf, 0x35, 0x2f, 0x24,
	0x82, 0x09, 0x6f, 0xfa, 0xb7, 0xbb, 0x9c, 0x9f, 0xe9, 0xee, 0x06, 0xb7, 0xa2, 0xf0, 0x8f, 0x07,
	0x13, 0xc7, 0xa1, 0x39, 0x0c, 0xd6, 0x4b, 0xbb, 0x96, 0x8f, 0x07, 0xeb, 0xa5, 0x29, 0xe6, 0x40,
	0x64, 0x41, 0xf5, 0x7a, 0x69, 0x17, 0xf1, 0x71, 0x87, 0xd1, 0x1b, 0x98, 0x66, 0xb5, 0x24, 0x9a,
	0x09, 0xee, 0xc2, 0xb7, 0xef, 0xc2, 0x4d, 0x71, 0xc7, 0xa3, 0x57, 0x70, 0x51, 0x49, 0x96, 0x52,
	0x97, 0x7c, 0xd6, 0x08, 0x13, 0x33, 0xc2, 0x0d, 0x83, 0x22, 0xf0, 0x77, 0x25, 0x49, 0x8b, 0x92,
	0x29, 0xed, 0x82, 0xbb, 0x05, 0x3f, 0xeb, 0x3c, 0xce, 0x32, 0x49, 0x95, 0xc2, 0x27, 0x09, 0x7a,
	0x00, 0x5f, 0xb6, 0xf7, 0x11, 0x8c, 0xad, 0xfe, 0xa6, 0x5f, 0x72, 0xcb, 0xe2, 0x93, 0x70, 0x37,
	0xb6, 0x7f, 0xc3, 0xf7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x02, 0xb8, 0x75, 0x33, 0xaf, 0x03,
	0x00, 0x00,
}
