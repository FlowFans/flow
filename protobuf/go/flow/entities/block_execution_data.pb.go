// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/block_execution_data.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// BlockExecutionData represents the collection of data produced while execiting the block.
type BlockExecutionData struct {
	// Block ID of the block that was executed.
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Ordered list of ChunkExecutionData produced while executing the block.
	//
	// Note: there will be one ChunkExecutionData per collection in the block, plus one for the
	// service chunk. The service chunk is executed last and is always the last chunk in the list.
	ChunkExecutionData   []*ChunkExecutionData `protobuf:"bytes,2,rep,name=chunk_execution_data,json=chunkExecutionData,proto3" json:"chunk_execution_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BlockExecutionData) Reset()         { *m = BlockExecutionData{} }
func (m *BlockExecutionData) String() string { return proto.CompactTextString(m) }
func (*BlockExecutionData) ProtoMessage()    {}
func (*BlockExecutionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{0}
}

func (m *BlockExecutionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockExecutionData.Unmarshal(m, b)
}
func (m *BlockExecutionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockExecutionData.Marshal(b, m, deterministic)
}
func (m *BlockExecutionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockExecutionData.Merge(m, src)
}
func (m *BlockExecutionData) XXX_Size() int {
	return xxx_messageInfo_BlockExecutionData.Size(m)
}
func (m *BlockExecutionData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockExecutionData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockExecutionData proto.InternalMessageInfo

func (m *BlockExecutionData) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *BlockExecutionData) GetChunkExecutionData() []*ChunkExecutionData {
	if m != nil {
		return m.ChunkExecutionData
	}
	return nil
}

// ChunkExecutionData represents the collection of data produced while executing a chunk.
type ChunkExecutionData struct {
	// Ordered list of transactions included in the collection that was executed in the chunk.
	Collection *ExecutionDataCollection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// Events emitted by transactions in the collection.
	//
	// Note: events listed in the last ChunkExecutionData in the BlockExecutionData were emitted by
	// the service transaction. Some, but not all, of these events are service events.
	Events []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	// TrieUpdate produced by executing the collection.
	//
	// TrieUpdates contain a list of registers that were modified during chunk execution. The value
	// included is the new value of the register.
	TrieUpdate *TrieUpdate `protobuf:"bytes,3,opt,name=trieUpdate,proto3" json:"trieUpdate,omitempty"`
	// Transaction results produced by executing the collection.
	//
	// Note: these are not the same type of results returned by other RPCs. These results are sepcific
	// to execution data. The most notable difference is they only include a boolean value to indicate
	// whether or not an error was encountered during execution, not the error itself.
	TransactionResults   []*ExecutionDataTransactionResult `protobuf:"bytes,4,rep,name=transaction_results,json=transactionResults,proto3" json:"transaction_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ChunkExecutionData) Reset()         { *m = ChunkExecutionData{} }
func (m *ChunkExecutionData) String() string { return proto.CompactTextString(m) }
func (*ChunkExecutionData) ProtoMessage()    {}
func (*ChunkExecutionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{1}
}

func (m *ChunkExecutionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkExecutionData.Unmarshal(m, b)
}
func (m *ChunkExecutionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkExecutionData.Marshal(b, m, deterministic)
}
func (m *ChunkExecutionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkExecutionData.Merge(m, src)
}
func (m *ChunkExecutionData) XXX_Size() int {
	return xxx_messageInfo_ChunkExecutionData.Size(m)
}
func (m *ChunkExecutionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkExecutionData.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkExecutionData proto.InternalMessageInfo

func (m *ChunkExecutionData) GetCollection() *ExecutionDataCollection {
	if m != nil {
		return m.Collection
	}
	return nil
}

func (m *ChunkExecutionData) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ChunkExecutionData) GetTrieUpdate() *TrieUpdate {
	if m != nil {
		return m.TrieUpdate
	}
	return nil
}

func (m *ChunkExecutionData) GetTransactionResults() []*ExecutionDataTransactionResult {
	if m != nil {
		return m.TransactionResults
	}
	return nil
}

// ExecutionDataCollection represents the collection of transactions that were executed within a chunk.
//
// Note: this is not the same type as the entities.Collection.
type ExecutionDataCollection struct {
	// List of transactions included in the collection.
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecutionDataCollection) Reset()         { *m = ExecutionDataCollection{} }
func (m *ExecutionDataCollection) String() string { return proto.CompactTextString(m) }
func (*ExecutionDataCollection) ProtoMessage()    {}
func (*ExecutionDataCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{2}
}

func (m *ExecutionDataCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionDataCollection.Unmarshal(m, b)
}
func (m *ExecutionDataCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionDataCollection.Marshal(b, m, deterministic)
}
func (m *ExecutionDataCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionDataCollection.Merge(m, src)
}
func (m *ExecutionDataCollection) XXX_Size() int {
	return xxx_messageInfo_ExecutionDataCollection.Size(m)
}
func (m *ExecutionDataCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionDataCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionDataCollection proto.InternalMessageInfo

func (m *ExecutionDataCollection) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// TrieUpdate produced by executing the collection.
//
// TrieUpdates contain a list of registers that were modified during chunk execution. The value
// included is the new value of the register.
type TrieUpdate struct {
	// RootHash is the root hash of the trie that included the contained updates.
	RootHash []byte `protobuf:"bytes,1,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// List of register paths updated.
	//
	// Note: paths and payloads map 1:1 with eachother. i.e. for each element in path, the value in
	// payloads at the same index is the value of the register at that path.
	Paths [][]byte `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	// List of register values updated.
	//
	// Note: paths and payloads map 1:1 with eachother. i.e. for each element in path, the value in
	// payloads at the same index is the value of the register at that path.
	Payloads             []*Payload `protobuf:"bytes,3,rep,name=payloads,proto3" json:"payloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TrieUpdate) Reset()         { *m = TrieUpdate{} }
func (m *TrieUpdate) String() string { return proto.CompactTextString(m) }
func (*TrieUpdate) ProtoMessage()    {}
func (*TrieUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{3}
}

func (m *TrieUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrieUpdate.Unmarshal(m, b)
}
func (m *TrieUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrieUpdate.Marshal(b, m, deterministic)
}
func (m *TrieUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrieUpdate.Merge(m, src)
}
func (m *TrieUpdate) XXX_Size() int {
	return xxx_messageInfo_TrieUpdate.Size(m)
}
func (m *TrieUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_TrieUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_TrieUpdate proto.InternalMessageInfo

func (m *TrieUpdate) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *TrieUpdate) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *TrieUpdate) GetPayloads() []*Payload {
	if m != nil {
		return m.Payloads
	}
	return nil
}

// Payload represents the value of a register.
type Payload struct {
	// List of key parts that make up the register key.
	KeyPart []*KeyPart `protobuf:"bytes,1,rep,name=keyPart,proto3" json:"keyPart,omitempty"`
	// Value of the register.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{4}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetKeyPart() []*KeyPart {
	if m != nil {
		return m.KeyPart
	}
	return nil
}

func (m *Payload) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// KeyPart represents a part of a register key.
type KeyPart struct {
	// Type of the key part.
	Type uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Value of the key part.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPart) Reset()         { *m = KeyPart{} }
func (m *KeyPart) String() string { return proto.CompactTextString(m) }
func (*KeyPart) ProtoMessage()    {}
func (*KeyPart) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{5}
}

func (m *KeyPart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPart.Unmarshal(m, b)
}
func (m *KeyPart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPart.Marshal(b, m, deterministic)
}
func (m *KeyPart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPart.Merge(m, src)
}
func (m *KeyPart) XXX_Size() int {
	return xxx_messageInfo_KeyPart.Size(m)
}
func (m *KeyPart) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPart.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPart proto.InternalMessageInfo

func (m *KeyPart) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *KeyPart) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// ExecutionDataTransactionResult represents the result of executing a transaction.
type ExecutionDataTransactionResult struct {
	// Transaction ID of the transaction that was executed.
	TransactionId []byte `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Boolean indicating whether or not the transaction's execution failed with an error.
	Failed bool `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	// Amount of computation used during execution.
	ComputationUsed      uint64   `protobuf:"varint,3,opt,name=computation_used,json=computationUsed,proto3" json:"computation_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionDataTransactionResult) Reset()         { *m = ExecutionDataTransactionResult{} }
func (m *ExecutionDataTransactionResult) String() string { return proto.CompactTextString(m) }
func (*ExecutionDataTransactionResult) ProtoMessage()    {}
func (*ExecutionDataTransactionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dbfd6581e1d1c72, []int{6}
}

func (m *ExecutionDataTransactionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionDataTransactionResult.Unmarshal(m, b)
}
func (m *ExecutionDataTransactionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionDataTransactionResult.Marshal(b, m, deterministic)
}
func (m *ExecutionDataTransactionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionDataTransactionResult.Merge(m, src)
}
func (m *ExecutionDataTransactionResult) XXX_Size() int {
	return xxx_messageInfo_ExecutionDataTransactionResult.Size(m)
}
func (m *ExecutionDataTransactionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionDataTransactionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionDataTransactionResult proto.InternalMessageInfo

func (m *ExecutionDataTransactionResult) GetTransactionId() []byte {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

func (m *ExecutionDataTransactionResult) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *ExecutionDataTransactionResult) GetComputationUsed() uint64 {
	if m != nil {
		return m.ComputationUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockExecutionData)(nil), "flow.entities.BlockExecutionData")
	proto.RegisterType((*ChunkExecutionData)(nil), "flow.entities.ChunkExecutionData")
	proto.RegisterType((*ExecutionDataCollection)(nil), "flow.entities.ExecutionDataCollection")
	proto.RegisterType((*TrieUpdate)(nil), "flow.entities.TrieUpdate")
	proto.RegisterType((*Payload)(nil), "flow.entities.Payload")
	proto.RegisterType((*KeyPart)(nil), "flow.entities.KeyPart")
	proto.RegisterType((*ExecutionDataTransactionResult)(nil), "flow.entities.ExecutionDataTransactionResult")
}

func init() {
	proto.RegisterFile("flow/entities/block_execution_data.proto", fileDescriptor_0dbfd6581e1d1c72)
}

var fileDescriptor_0dbfd6581e1d1c72 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x6f, 0xd3, 0x4e,
	0x10, 0xc5, 0xe5, 0x24, 0xff, 0x24, 0xff, 0x69, 0x02, 0x68, 0x88, 0x8a, 0x13, 0x10, 0x14, 0x4b,
	0xa0, 0x20, 0x81, 0x5d, 0xa5, 0x27, 0x2e, 0x1c, 0x5a, 0x40, 0x54, 0x5c, 0xc2, 0xd2, 0x1e, 0xe0,
	0x40, 0xb4, 0xb1, 0x37, 0xb1, 0x55, 0xd7, 0x6b, 0x79, 0xd7, 0x85, 0xdc, 0x39, 0xf1, 0x3d, 0xf8,
	0x9e, 0xc8, 0x63, 0xc7, 0xb1, 0x1d, 0xb5, 0x97, 0x28, 0x3b, 0xf3, 0x7b, 0xfb, 0xde, 0x8e, 0x77,
	0x61, 0xba, 0x0a, 0xe5, 0x4f, 0x47, 0x44, 0x3a, 0xd0, 0x81, 0x50, 0xce, 0x32, 0x94, 0xee, 0xd5,
	0x42, 0xfc, 0x12, 0x6e, 0xaa, 0x03, 0x19, 0x2d, 0x3c, 0xae, 0xb9, 0x1d, 0x27, 0x52, 0x4b, 0x1c,
	0x66, 0xa4, 0xbd, 0x25, 0x27, 0xe3, 0xba, 0x50, 0xdc, 0x88, 0x48, 0xe7, 0xe4, 0xe4, 0x59, 0xbd,
	0xa5, 0x13, 0x1e, 0x29, 0xee, 0x66, 0xfb, 0xe5, 0x80, 0xf5, 0xdb, 0x00, 0x3c, 0xcd, 0x9c, 0x3e,
	0x6c, 0x8d, 0xde, 0x73, 0xcd, 0x71, 0x0c, 0xfd, 0xdc, 0x3f, 0xf0, 0x4c, 0xe3, 0xc8, 0x98, 0x0e,
	0x58, 0x8f, 0xd6, 0xe7, 0x1e, 0x7e, 0x85, 0x91, 0xeb, 0xa7, 0x51, 0x33, 0x9a, 0xd9, 0x3a, 0x6a,
	0x4f, 0x0f, 0x66, 0xcf, 0xed, 0x5a, 0x36, 0xfb, 0x2c, 0x43, 0x6b, 0x7b, 0x33, 0x74, 0xf7, 0x6a,
	0xd6, 0xdf, 0x16, 0xe0, 0x3e, 0x8a, 0x1f, 0x01, 0x5c, 0x19, 0x86, 0x82, 0x12, 0x53, 0x90, 0x83,
	0xd9, 0xcb, 0x86, 0x43, 0x4d, 0x71, 0x56, 0xd2, 0xac, 0xa2, 0xc4, 0xd7, 0xd0, 0xa5, 0xa9, 0xa8,
	0x22, 0xe5, 0xa8, 0xb9, 0x47, 0xd6, 0x64, 0x05, 0x83, 0x6f, 0x01, 0x74, 0x12, 0x88, 0xcb, 0xd8,
	0xe3, 0x5a, 0x98, 0x6d, 0x72, 0x1d, 0x37, 0x14, 0x17, 0x25, 0xc0, 0x2a, 0x30, 0xfe, 0x80, 0x87,
	0x95, 0x19, 0x2f, 0x12, 0xa1, 0xd2, 0x50, 0x2b, 0xb3, 0x43, 0xae, 0x6f, 0xee, 0x4a, 0x7e, 0xb1,
	0x93, 0x31, 0x52, 0x31, 0xd4, 0xcd, 0x92, 0xb2, 0xbe, 0xc1, 0xa3, 0x5b, 0xce, 0x8b, 0xef, 0x60,
	0x50, 0x11, 0x28, 0xd3, 0x20, 0xcf, 0xc9, 0x5e, 0xee, 0xdd, 0x9e, 0x35, 0xde, 0x52, 0x00, 0xbb,
	0x43, 0xe1, 0x63, 0xf8, 0x3f, 0x91, 0x52, 0x2f, 0x7c, 0xae, 0xfc, 0xe2, 0x06, 0xf4, 0xb3, 0xc2,
	0x27, 0xae, 0x7c, 0x1c, 0xc1, 0x7f, 0x31, 0xd7, 0x7e, 0x3e, 0xcd, 0x01, 0xcb, 0x17, 0x38, 0x83,
	0x7e, 0xcc, 0x37, 0xa1, 0xe4, 0x9e, 0x32, 0xdb, 0x64, 0x7e, 0xd8, 0x30, 0x9f, 0xe7, 0x6d, 0x56,
	0x72, 0xd6, 0x17, 0xe8, 0x15, 0x45, 0x3c, 0x86, 0xde, 0x95, 0xd8, 0xcc, 0x79, 0xa2, 0x8b, 0xe8,
	0x4d, 0xf5, 0xe7, 0xbc, 0xcb, 0xb6, 0x58, 0x16, 0xe3, 0x86, 0x87, 0xa9, 0x30, 0x5b, 0x94, 0x2f,
	0x5f, 0x58, 0x27, 0xd0, 0x2b, 0x48, 0x44, 0xe8, 0xe8, 0x4d, 0x2c, 0x28, 0xff, 0x90, 0xd1, 0xff,
	0x5b, 0x44, 0x7f, 0x0c, 0x78, 0x7a, 0xf7, 0xe7, 0xc0, 0x17, 0x70, 0xaf, 0xfa, 0x69, 0xcb, 0x87,
	0x31, 0xac, 0x54, 0xcf, 0x3d, 0x3c, 0x84, 0xee, 0x8a, 0x07, 0xa1, 0xf0, 0xc8, 0xa0, 0xcf, 0x8a,
	0x15, 0xbe, 0x82, 0x07, 0xae, 0xbc, 0x8e, 0x53, 0xcd, 0x49, 0x9e, 0x2a, 0xe1, 0xd1, 0xd5, 0xea,
	0xb0, 0xfb, 0x95, 0xfa, 0xa5, 0x12, 0xde, 0xe9, 0x1c, 0x9e, 0xc8, 0x64, 0x6d, 0xcb, 0x88, 0xce,
	0x4f, 0xef, 0x74, 0x99, 0xae, 0xca, 0x41, 0x7c, 0x3f, 0x5e, 0x07, 0xda, 0x4f, 0x97, 0xb6, 0x2b,
	0xaf, 0x9d, 0x1c, 0x72, 0xe8, 0x67, 0x4b, 0x3a, 0x6b, 0xe9, 0xd4, 0xde, 0xfd, 0xb2, 0x4b, 0xad,
	0x93, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x8b, 0x44, 0xf1, 0x63, 0x04, 0x00, 0x00,
}
