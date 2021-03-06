// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/hcs.proto

package orderer

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HcsMessageRegular_Class int32

const (
	HcsMessageRegular_NORMAL HcsMessageRegular_Class = 0
	HcsMessageRegular_CONFIG HcsMessageRegular_Class = 1
)

var HcsMessageRegular_Class_name = map[int32]string{
	0: "NORMAL",
	1: "CONFIG",
}

var HcsMessageRegular_Class_value = map[string]int32{
	"NORMAL": 0,
	"CONFIG": 1,
}

func (x HcsMessageRegular_Class) String() string {
	return proto.EnumName(HcsMessageRegular_Class_name, int32(x))
}

func (HcsMessageRegular_Class) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{1, 0}
}

// HcsMessage is a wrapper type for the messages that the HCS-based
// orderer deals with.
type HcsMessage struct {
	// Types that are valid to be assigned to Type:
	//	*HcsMessage_Regular
	//	*HcsMessage_TimeToCut
	//	*HcsMessage_OrdererStarted
	Type                 isHcsMessage_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HcsMessage) Reset()         { *m = HcsMessage{} }
func (m *HcsMessage) String() string { return proto.CompactTextString(m) }
func (*HcsMessage) ProtoMessage()    {}
func (*HcsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{0}
}

func (m *HcsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessage.Unmarshal(m, b)
}
func (m *HcsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessage.Marshal(b, m, deterministic)
}
func (m *HcsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessage.Merge(m, src)
}
func (m *HcsMessage) XXX_Size() int {
	return xxx_messageInfo_HcsMessage.Size(m)
}
func (m *HcsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessage proto.InternalMessageInfo

type isHcsMessage_Type interface {
	isHcsMessage_Type()
}

type HcsMessage_Regular struct {
	Regular *HcsMessageRegular `protobuf:"bytes,1,opt,name=regular,proto3,oneof"`
}

type HcsMessage_TimeToCut struct {
	TimeToCut *HcsMessageTimeToCut `protobuf:"bytes,2,opt,name=time_to_cut,json=timeToCut,proto3,oneof"`
}

type HcsMessage_OrdererStarted struct {
	OrdererStarted *HcsMessageOrdererStarted `protobuf:"bytes,3,opt,name=orderer_started,json=ordererStarted,proto3,oneof"`
}

func (*HcsMessage_Regular) isHcsMessage_Type() {}

func (*HcsMessage_TimeToCut) isHcsMessage_Type() {}

func (*HcsMessage_OrdererStarted) isHcsMessage_Type() {}

func (m *HcsMessage) GetType() isHcsMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *HcsMessage) GetRegular() *HcsMessageRegular {
	if x, ok := m.GetType().(*HcsMessage_Regular); ok {
		return x.Regular
	}
	return nil
}

func (m *HcsMessage) GetTimeToCut() *HcsMessageTimeToCut {
	if x, ok := m.GetType().(*HcsMessage_TimeToCut); ok {
		return x.TimeToCut
	}
	return nil
}

func (m *HcsMessage) GetOrdererStarted() *HcsMessageOrdererStarted {
	if x, ok := m.GetType().(*HcsMessage_OrdererStarted); ok {
		return x.OrdererStarted
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HcsMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HcsMessage_Regular)(nil),
		(*HcsMessage_TimeToCut)(nil),
		(*HcsMessage_OrdererStarted)(nil),
	}
}

// HcsMessageRegular wraps a marshalled envelope.
type HcsMessageRegular struct {
	Payload              []byte                  `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ConfigSeq            uint64                  `protobuf:"varint,2,opt,name=config_seq,json=configSeq,proto3" json:"config_seq,omitempty"`
	Class                HcsMessageRegular_Class `protobuf:"varint,3,opt,name=class,proto3,enum=orderer.HcsMessageRegular_Class" json:"class,omitempty"`
	OriginalSeq          uint64                  `protobuf:"varint,4,opt,name=original_seq,json=originalSeq,proto3" json:"original_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HcsMessageRegular) Reset()         { *m = HcsMessageRegular{} }
func (m *HcsMessageRegular) String() string { return proto.CompactTextString(m) }
func (*HcsMessageRegular) ProtoMessage()    {}
func (*HcsMessageRegular) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{1}
}

func (m *HcsMessageRegular) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageRegular.Unmarshal(m, b)
}
func (m *HcsMessageRegular) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageRegular.Marshal(b, m, deterministic)
}
func (m *HcsMessageRegular) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageRegular.Merge(m, src)
}
func (m *HcsMessageRegular) XXX_Size() int {
	return xxx_messageInfo_HcsMessageRegular.Size(m)
}
func (m *HcsMessageRegular) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageRegular.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageRegular proto.InternalMessageInfo

func (m *HcsMessageRegular) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HcsMessageRegular) GetConfigSeq() uint64 {
	if m != nil {
		return m.ConfigSeq
	}
	return 0
}

func (m *HcsMessageRegular) GetClass() HcsMessageRegular_Class {
	if m != nil {
		return m.Class
	}
	return HcsMessageRegular_NORMAL
}

func (m *HcsMessageRegular) GetOriginalSeq() uint64 {
	if m != nil {
		return m.OriginalSeq
	}
	return 0
}

// TODO: a fragment key is needed to identify the subject so as to handle
// the case that different messages are fragmented by different subjects with
// the same fragment_id
type HcsMessageFragment struct {
	Fragment             []byte   `protobuf:"bytes,1,opt,name=fragment,proto3" json:"fragment,omitempty"`
	FragmentKey          []byte   `protobuf:"bytes,2,opt,name=fragment_key,json=fragmentKey,proto3" json:"fragment_key,omitempty"`
	FragmentId           uint64   `protobuf:"fixed64,3,opt,name=fragment_id,json=fragmentId,proto3" json:"fragment_id,omitempty"`
	Sequence             uint32   `protobuf:"fixed32,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	TotalFragments       uint32   `protobuf:"fixed32,5,opt,name=total_fragments,json=totalFragments,proto3" json:"total_fragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsMessageFragment) Reset()         { *m = HcsMessageFragment{} }
func (m *HcsMessageFragment) String() string { return proto.CompactTextString(m) }
func (*HcsMessageFragment) ProtoMessage()    {}
func (*HcsMessageFragment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{2}
}

func (m *HcsMessageFragment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageFragment.Unmarshal(m, b)
}
func (m *HcsMessageFragment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageFragment.Marshal(b, m, deterministic)
}
func (m *HcsMessageFragment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageFragment.Merge(m, src)
}
func (m *HcsMessageFragment) XXX_Size() int {
	return xxx_messageInfo_HcsMessageFragment.Size(m)
}
func (m *HcsMessageFragment) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageFragment.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageFragment proto.InternalMessageInfo

func (m *HcsMessageFragment) GetFragment() []byte {
	if m != nil {
		return m.Fragment
	}
	return nil
}

func (m *HcsMessageFragment) GetFragmentKey() []byte {
	if m != nil {
		return m.FragmentKey
	}
	return nil
}

func (m *HcsMessageFragment) GetFragmentId() uint64 {
	if m != nil {
		return m.FragmentId
	}
	return 0
}

func (m *HcsMessageFragment) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *HcsMessageFragment) GetTotalFragments() uint32 {
	if m != nil {
		return m.TotalFragments
	}
	return 0
}

// HcsMessageTimeToCut is used to signal to the orderers that it is time to
// cut block <block_number>.
type HcsMessageTimeToCut struct {
	BlockNumber          uint64   `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsMessageTimeToCut) Reset()         { *m = HcsMessageTimeToCut{} }
func (m *HcsMessageTimeToCut) String() string { return proto.CompactTextString(m) }
func (*HcsMessageTimeToCut) ProtoMessage()    {}
func (*HcsMessageTimeToCut) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{3}
}

func (m *HcsMessageTimeToCut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageTimeToCut.Unmarshal(m, b)
}
func (m *HcsMessageTimeToCut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageTimeToCut.Marshal(b, m, deterministic)
}
func (m *HcsMessageTimeToCut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageTimeToCut.Merge(m, src)
}
func (m *HcsMessageTimeToCut) XXX_Size() int {
	return xxx_messageInfo_HcsMessageTimeToCut.Size(m)
}
func (m *HcsMessageTimeToCut) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageTimeToCut.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageTimeToCut proto.InternalMessageInfo

func (m *HcsMessageTimeToCut) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

type HcsMessageOrdererStarted struct {
	OrdererIdentity      []byte   `protobuf:"bytes,1,opt,name=orderer_identity,json=ordererIdentity,proto3" json:"orderer_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsMessageOrdererStarted) Reset()         { *m = HcsMessageOrdererStarted{} }
func (m *HcsMessageOrdererStarted) String() string { return proto.CompactTextString(m) }
func (*HcsMessageOrdererStarted) ProtoMessage()    {}
func (*HcsMessageOrdererStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{4}
}

func (m *HcsMessageOrdererStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMessageOrdererStarted.Unmarshal(m, b)
}
func (m *HcsMessageOrdererStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMessageOrdererStarted.Marshal(b, m, deterministic)
}
func (m *HcsMessageOrdererStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMessageOrdererStarted.Merge(m, src)
}
func (m *HcsMessageOrdererStarted) XXX_Size() int {
	return xxx_messageInfo_HcsMessageOrdererStarted.Size(m)
}
func (m *HcsMessageOrdererStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMessageOrdererStarted.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMessageOrdererStarted proto.InternalMessageInfo

func (m *HcsMessageOrdererStarted) GetOrdererIdentity() []byte {
	if m != nil {
		return m.OrdererIdentity
	}
	return nil
}

// HcsMetadata is encoded into the ORDERER block to keep track of HCS-related
// metadata associated with this block.
type HcsMetadata struct {
	// lastTimestampPersisted is the used to keep track of the timestamp of
	// the last ordered message in the last block so when an HCS-based orderer
	// boots up, it knows from when to retrieve ordererd messages
	LastConsensusTimestampPersisted *timestamp.Timestamp `protobuf:"bytes,1,opt,name=last_consensus_timestamp_persisted,json=lastConsensusTimestampPersisted,proto3" json:"last_consensus_timestamp_persisted,omitempty"`
	LastOriginalSequenceProcessed   uint64               `protobuf:"varint,2,opt,name=last_original_sequence_processed,json=lastOriginalSequenceProcessed,proto3" json:"last_original_sequence_processed,omitempty"`
	LastResubmittedConfigSequence   uint64               `protobuf:"varint,3,opt,name=last_resubmitted_config_sequence,json=lastResubmittedConfigSequence,proto3" json:"last_resubmitted_config_sequence,omitempty"`
	// last fragment free consensus timestamp is the consensus timestamp of the last
	// fragment free block
	LastFragmentFreeConsensusTimestampPersisted *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_fragment_free_consensus_timestamp_persisted,json=lastFragmentFreeConsensusTimestampPersisted,proto3" json:"last_fragment_free_consensus_timestamp_persisted,omitempty"`
	XXX_NoUnkeyedLiteral                        struct{}             `json:"-"`
	XXX_unrecognized                            []byte               `json:"-"`
	XXX_sizecache                               int32                `json:"-"`
}

func (m *HcsMetadata) Reset()         { *m = HcsMetadata{} }
func (m *HcsMetadata) String() string { return proto.CompactTextString(m) }
func (*HcsMetadata) ProtoMessage()    {}
func (*HcsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{5}
}

func (m *HcsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsMetadata.Unmarshal(m, b)
}
func (m *HcsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsMetadata.Marshal(b, m, deterministic)
}
func (m *HcsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsMetadata.Merge(m, src)
}
func (m *HcsMetadata) XXX_Size() int {
	return xxx_messageInfo_HcsMetadata.Size(m)
}
func (m *HcsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HcsMetadata proto.InternalMessageInfo

func (m *HcsMetadata) GetLastConsensusTimestampPersisted() *timestamp.Timestamp {
	if m != nil {
		return m.LastConsensusTimestampPersisted
	}
	return nil
}

func (m *HcsMetadata) GetLastOriginalSequenceProcessed() uint64 {
	if m != nil {
		return m.LastOriginalSequenceProcessed
	}
	return 0
}

func (m *HcsMetadata) GetLastResubmittedConfigSequence() uint64 {
	if m != nil {
		return m.LastResubmittedConfigSequence
	}
	return 0
}

func (m *HcsMetadata) GetLastFragmentFreeConsensusTimestampPersisted() *timestamp.Timestamp {
	if m != nil {
		return m.LastFragmentFreeConsensusTimestampPersisted
	}
	return nil
}

type HcsConfigMetadata struct {
	// HCS topic ID in the format of shard.realm.num
	TopicID              string   `protobuf:"bytes,1,opt,name=TopicID,proto3" json:"TopicID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HcsConfigMetadata) Reset()         { *m = HcsConfigMetadata{} }
func (m *HcsConfigMetadata) String() string { return proto.CompactTextString(m) }
func (*HcsConfigMetadata) ProtoMessage()    {}
func (*HcsConfigMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d490350899af2cbc, []int{6}
}

func (m *HcsConfigMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HcsConfigMetadata.Unmarshal(m, b)
}
func (m *HcsConfigMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HcsConfigMetadata.Marshal(b, m, deterministic)
}
func (m *HcsConfigMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HcsConfigMetadata.Merge(m, src)
}
func (m *HcsConfigMetadata) XXX_Size() int {
	return xxx_messageInfo_HcsConfigMetadata.Size(m)
}
func (m *HcsConfigMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HcsConfigMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HcsConfigMetadata proto.InternalMessageInfo

func (m *HcsConfigMetadata) GetTopicID() string {
	if m != nil {
		return m.TopicID
	}
	return ""
}

func init() {
	proto.RegisterEnum("orderer.HcsMessageRegular_Class", HcsMessageRegular_Class_name, HcsMessageRegular_Class_value)
	proto.RegisterType((*HcsMessage)(nil), "orderer.HcsMessage")
	proto.RegisterType((*HcsMessageRegular)(nil), "orderer.HcsMessageRegular")
	proto.RegisterType((*HcsMessageFragment)(nil), "orderer.HcsMessageFragment")
	proto.RegisterType((*HcsMessageTimeToCut)(nil), "orderer.HcsMessageTimeToCut")
	proto.RegisterType((*HcsMessageOrdererStarted)(nil), "orderer.HcsMessageOrdererStarted")
	proto.RegisterType((*HcsMetadata)(nil), "orderer.HcsMetadata")
	proto.RegisterType((*HcsConfigMetadata)(nil), "orderer.HcsConfigMetadata")
}

func init() { proto.RegisterFile("orderer/hcs.proto", fileDescriptor_d490350899af2cbc) }

var fileDescriptor_d490350899af2cbc = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0x93, 0x4b, 0x48, 0x2e, 0x27, 0x88, 0x8f, 0xb9, 0x9b, 0x08, 0x5d, 0x14, 0xb0, 0x54,
	0xb5, 0x55, 0x85, 0x5d, 0x51, 0x09, 0x75, 0x55, 0xa9, 0xa4, 0x85, 0x44, 0x05, 0x82, 0x86, 0x2c,
	0xaa, 0x6e, 0xac, 0xb1, 0x7d, 0x62, 0x2c, 0x6c, 0x8f, 0x33, 0x33, 0x5e, 0x64, 0xd9, 0x7d, 0x5f,
	0xa8, 0x6f, 0xd0, 0x67, 0xe8, 0xd3, 0x54, 0x9e, 0xf1, 0x38, 0x54, 0x85, 0xb2, 0x9b, 0xf3, 0xd7,
	0x6f, 0xfe, 0x73, 0xe6, 0x7c, 0xc0, 0x2e, 0x17, 0x11, 0x0a, 0x14, 0xde, 0x6d, 0x28, 0xdd, 0x42,
	0x70, 0xc5, 0x49, 0xaf, 0x96, 0xf6, 0x86, 0x31, 0xe7, 0x71, 0x8a, 0x9e, 0x96, 0x83, 0x72, 0xee,
	0xa9, 0x24, 0x43, 0xa9, 0x58, 0x56, 0x18, 0xd2, 0xf9, 0xd9, 0x06, 0x18, 0x87, 0xf2, 0x12, 0xa5,
	0x64, 0x31, 0x92, 0x13, 0xe8, 0x09, 0x8c, 0xcb, 0x94, 0x89, 0x41, 0xfb, 0xa0, 0xfd, 0xa2, 0x7f,
	0xbc, 0xe7, 0xd6, 0x56, 0xee, 0x8a, 0xa2, 0x86, 0x18, 0xb7, 0xa8, 0x85, 0xc9, 0x3b, 0xe8, 0x57,
	0xce, 0xbe, 0xe2, 0x7e, 0x58, 0xaa, 0xc1, 0x3f, 0xfa, 0xee, 0xff, 0x0f, 0xdc, 0x9d, 0x25, 0x19,
	0xce, 0xf8, 0xa8, 0x54, 0xe3, 0x16, 0xdd, 0x50, 0x36, 0x20, 0x17, 0xb0, 0x5d, 0xb3, 0xbe, 0x54,
	0x4c, 0x28, 0x8c, 0x06, 0x6b, 0xda, 0xe3, 0xf0, 0x01, 0x8f, 0xa9, 0x51, 0x6e, 0x0c, 0x38, 0x6e,
	0xd1, 0x2d, 0xfe, 0x9b, 0x72, 0xda, 0x85, 0xce, 0x6c, 0x59, 0xa0, 0xf3, 0xa3, 0x0d, 0xbb, 0x7f,
	0xa4, 0x4d, 0x06, 0xd0, 0x2b, 0xd8, 0x32, 0xe5, 0x2c, 0xd2, 0x7f, 0xdc, 0xa4, 0x36, 0x24, 0xfb,
	0x00, 0x21, 0xcf, 0xe7, 0x49, 0xec, 0x4b, 0x5c, 0xe8, 0x4f, 0x74, 0xe8, 0x86, 0x51, 0x6e, 0x70,
	0x41, 0x4e, 0x60, 0x3d, 0x4c, 0x99, 0x94, 0x3a, 0xb5, 0xad, 0xe3, 0x83, 0xc7, 0x4b, 0xe3, 0x8e,
	0x2a, 0x8e, 0x1a, 0x9c, 0x1c, 0xc2, 0x26, 0x17, 0x49, 0x9c, 0xe4, 0x2c, 0xd5, 0xc6, 0x1d, 0x6d,
	0xdc, 0xb7, 0xda, 0x0d, 0x2e, 0x9c, 0x21, 0xac, 0xeb, 0x2b, 0x04, 0xa0, 0x7b, 0x35, 0xa5, 0x97,
	0xef, 0x2f, 0x76, 0x5a, 0xd5, 0x79, 0x34, 0xbd, 0x3a, 0x9b, 0x9c, 0xef, 0xb4, 0x9d, 0xef, 0x6d,
	0x20, 0xab, 0x67, 0xce, 0x04, 0x8b, 0x33, 0xcc, 0x15, 0xd9, 0x83, 0x7f, 0xe7, 0xf5, 0xb9, 0xfe,
	0x4c, 0x13, 0x57, 0xcf, 0xda, 0xb3, 0x7f, 0x87, 0x4b, 0xfd, 0x9f, 0x4d, 0xda, 0xb7, 0xda, 0x27,
	0x5c, 0x92, 0x21, 0x34, 0xa1, 0x9f, 0x98, 0x92, 0x77, 0x29, 0x58, 0x69, 0x12, 0x55, 0xfe, 0x12,
	0x17, 0x25, 0xe6, 0x21, 0xea, 0xb4, 0x7b, 0xb4, 0x89, 0xc9, 0x73, 0xd8, 0x56, 0x5c, 0xb1, 0xd4,
	0xb7, 0xbc, 0x1c, 0xac, 0x6b, 0x64, 0x4b, 0xcb, 0x36, 0x47, 0xe9, 0xbc, 0x85, 0xff, 0x1e, 0x18,
	0x80, 0x2a, 0xbf, 0x20, 0xe5, 0xe1, 0x9d, 0x9f, 0x97, 0x59, 0x80, 0x66, 0xe0, 0x3a, 0xb4, 0xaf,
	0xb5, 0x2b, 0x2d, 0x39, 0x1f, 0x61, 0xf0, 0x58, 0xdb, 0xc9, 0x4b, 0xd8, 0xb1, 0x23, 0x93, 0x44,
	0x98, 0xab, 0x44, 0x2d, 0xeb, 0x12, 0xd8, 0x51, 0x9a, 0xd4, 0xb2, 0xf3, 0x6d, 0x0d, 0xfa, 0xda,
	0x47, 0xb1, 0x88, 0x29, 0x46, 0x62, 0x70, 0x52, 0x26, 0x95, 0x1f, 0xf2, 0x5c, 0x62, 0x2e, 0x4b,
	0xe9, 0x37, 0x6b, 0xe1, 0x17, 0x28, 0x64, 0x22, 0xab, 0x01, 0xb4, 0x0b, 0x60, 0x56, 0xc8, 0xb5,
	0x2b, 0xe4, 0xce, 0x2c, 0x4b, 0x87, 0x95, 0xcb, 0xc8, 0x9a, 0x34, 0xfa, 0xb5, 0xb5, 0x20, 0xe7,
	0x70, 0xa0, 0x1f, 0xba, 0xdf, 0x7e, 0x5d, 0x3c, 0xbf, 0x10, 0x3c, 0x44, 0x29, 0x31, 0xaa, 0xc7,
	0x6c, 0xbf, 0xe2, 0xa6, 0xab, 0x89, 0xd0, 0xd4, 0xb5, 0x85, 0x1a, 0x23, 0x81, 0xb2, 0x0c, 0xb2,
	0x44, 0x29, 0x8c, 0xfc, 0xd5, 0xa8, 0x9a, 0xfe, 0xac, 0xad, 0x8c, 0xe8, 0x0a, 0x1b, 0xd9, 0xf1,
	0x35, 0x4d, 0xfb, 0xda, 0x86, 0xd7, 0xda, 0xa9, 0xe9, 0xfb, 0x5c, 0x20, 0x3e, 0x51, 0x89, 0xce,
	0x93, 0x95, 0x78, 0x55, 0x79, 0xda, 0x8e, 0x9f, 0x09, 0xc4, 0xbf, 0x54, 0xc5, 0x39, 0xd2, 0x5b,
	0x69, 0x12, 0x6b, 0x7a, 0x32, 0x80, 0xde, 0x8c, 0x17, 0x49, 0x38, 0xf9, 0xa0, 0x0b, 0xbf, 0x41,
	0x6d, 0x78, 0xfa, 0x19, 0x9e, 0x71, 0x11, 0xbb, 0xb7, 0xcb, 0x02, 0x45, 0x8a, 0x51, 0x8c, 0xc2,
	0x9d, 0xb3, 0x40, 0x24, 0xa1, 0x49, 0x44, 0xda, 0x35, 0xfc, 0xe2, 0xc5, 0x89, 0xba, 0x2d, 0x03,
	0x37, 0xe4, 0x99, 0x77, 0x8f, 0xf6, 0x0c, 0x7d, 0x64, 0xe8, 0xa3, 0x98, 0x7b, 0xf5, 0x85, 0xa0,
	0xab, 0xa5, 0x37, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x75, 0x50, 0x29, 0x42, 0x05, 0x00,
	0x00,
}
