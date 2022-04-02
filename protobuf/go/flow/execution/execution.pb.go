// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/execution/execution.proto

package execution

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	entities "github.com/onflow/flow/protobuf/go/flow/entities"
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

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

type GetAccountAtBlockIDRequest struct {
	BlockId              []byte   `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountAtBlockIDRequest) Reset()         { *m = GetAccountAtBlockIDRequest{} }
func (m *GetAccountAtBlockIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountAtBlockIDRequest) ProtoMessage()    {}
func (*GetAccountAtBlockIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{2}
}

func (m *GetAccountAtBlockIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountAtBlockIDRequest.Unmarshal(m, b)
}
func (m *GetAccountAtBlockIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountAtBlockIDRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountAtBlockIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountAtBlockIDRequest.Merge(m, src)
}
func (m *GetAccountAtBlockIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountAtBlockIDRequest.Size(m)
}
func (m *GetAccountAtBlockIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountAtBlockIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountAtBlockIDRequest proto.InternalMessageInfo

func (m *GetAccountAtBlockIDRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetAccountAtBlockIDRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type GetAccountAtBlockIDResponse struct {
	Account              *entities.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetAccountAtBlockIDResponse) Reset()         { *m = GetAccountAtBlockIDResponse{} }
func (m *GetAccountAtBlockIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountAtBlockIDResponse) ProtoMessage()    {}
func (*GetAccountAtBlockIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{3}
}

func (m *GetAccountAtBlockIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountAtBlockIDResponse.Unmarshal(m, b)
}
func (m *GetAccountAtBlockIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountAtBlockIDResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountAtBlockIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountAtBlockIDResponse.Merge(m, src)
}
func (m *GetAccountAtBlockIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountAtBlockIDResponse.Size(m)
}
func (m *GetAccountAtBlockIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountAtBlockIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountAtBlockIDResponse proto.InternalMessageInfo

func (m *GetAccountAtBlockIDResponse) GetAccount() *entities.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ExecuteScriptAtBlockIDRequest struct {
	BlockId              []byte   `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Script               []byte   `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Arguments            [][]byte `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteScriptAtBlockIDRequest) Reset()         { *m = ExecuteScriptAtBlockIDRequest{} }
func (m *ExecuteScriptAtBlockIDRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteScriptAtBlockIDRequest) ProtoMessage()    {}
func (*ExecuteScriptAtBlockIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{4}
}

func (m *ExecuteScriptAtBlockIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteScriptAtBlockIDRequest.Unmarshal(m, b)
}
func (m *ExecuteScriptAtBlockIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteScriptAtBlockIDRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteScriptAtBlockIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteScriptAtBlockIDRequest.Merge(m, src)
}
func (m *ExecuteScriptAtBlockIDRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteScriptAtBlockIDRequest.Size(m)
}
func (m *ExecuteScriptAtBlockIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteScriptAtBlockIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteScriptAtBlockIDRequest proto.InternalMessageInfo

func (m *ExecuteScriptAtBlockIDRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *ExecuteScriptAtBlockIDRequest) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *ExecuteScriptAtBlockIDRequest) GetArguments() [][]byte {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type ExecuteScriptAtBlockIDResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteScriptAtBlockIDResponse) Reset()         { *m = ExecuteScriptAtBlockIDResponse{} }
func (m *ExecuteScriptAtBlockIDResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteScriptAtBlockIDResponse) ProtoMessage()    {}
func (*ExecuteScriptAtBlockIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{5}
}

func (m *ExecuteScriptAtBlockIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteScriptAtBlockIDResponse.Unmarshal(m, b)
}
func (m *ExecuteScriptAtBlockIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteScriptAtBlockIDResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteScriptAtBlockIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteScriptAtBlockIDResponse.Merge(m, src)
}
func (m *ExecuteScriptAtBlockIDResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteScriptAtBlockIDResponse.Size(m)
}
func (m *ExecuteScriptAtBlockIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteScriptAtBlockIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteScriptAtBlockIDResponse proto.InternalMessageInfo

func (m *ExecuteScriptAtBlockIDResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetEventsForBlockIDsResponse struct {
	Results              []*GetEventsForBlockIDsResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *GetEventsForBlockIDsResponse) Reset()         { *m = GetEventsForBlockIDsResponse{} }
func (m *GetEventsForBlockIDsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventsForBlockIDsResponse) ProtoMessage()    {}
func (*GetEventsForBlockIDsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{6}
}

func (m *GetEventsForBlockIDsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsForBlockIDsResponse.Unmarshal(m, b)
}
func (m *GetEventsForBlockIDsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsForBlockIDsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventsForBlockIDsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsForBlockIDsResponse.Merge(m, src)
}
func (m *GetEventsForBlockIDsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventsForBlockIDsResponse.Size(m)
}
func (m *GetEventsForBlockIDsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsForBlockIDsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsForBlockIDsResponse proto.InternalMessageInfo

func (m *GetEventsForBlockIDsResponse) GetResults() []*GetEventsForBlockIDsResponse_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type GetEventsForBlockIDsResponse_Result struct {
	BlockId              []byte            `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	BlockHeight          uint64            `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Events               []*entities.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetEventsForBlockIDsResponse_Result) Reset()         { *m = GetEventsForBlockIDsResponse_Result{} }
func (m *GetEventsForBlockIDsResponse_Result) String() string { return proto.CompactTextString(m) }
func (*GetEventsForBlockIDsResponse_Result) ProtoMessage()    {}
func (*GetEventsForBlockIDsResponse_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{6, 0}
}

func (m *GetEventsForBlockIDsResponse_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsForBlockIDsResponse_Result.Unmarshal(m, b)
}
func (m *GetEventsForBlockIDsResponse_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsForBlockIDsResponse_Result.Marshal(b, m, deterministic)
}
func (m *GetEventsForBlockIDsResponse_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsForBlockIDsResponse_Result.Merge(m, src)
}
func (m *GetEventsForBlockIDsResponse_Result) XXX_Size() int {
	return xxx_messageInfo_GetEventsForBlockIDsResponse_Result.Size(m)
}
func (m *GetEventsForBlockIDsResponse_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsForBlockIDsResponse_Result.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsForBlockIDsResponse_Result proto.InternalMessageInfo

func (m *GetEventsForBlockIDsResponse_Result) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetEventsForBlockIDsResponse_Result) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *GetEventsForBlockIDsResponse_Result) GetEvents() []*entities.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetEventsForBlockIDsRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	BlockIds             [][]byte `protobuf:"bytes,2,rep,name=block_ids,json=blockIds,proto3" json:"block_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsForBlockIDsRequest) Reset()         { *m = GetEventsForBlockIDsRequest{} }
func (m *GetEventsForBlockIDsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsForBlockIDsRequest) ProtoMessage()    {}
func (*GetEventsForBlockIDsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{7}
}

func (m *GetEventsForBlockIDsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsForBlockIDsRequest.Unmarshal(m, b)
}
func (m *GetEventsForBlockIDsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsForBlockIDsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventsForBlockIDsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsForBlockIDsRequest.Merge(m, src)
}
func (m *GetEventsForBlockIDsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsForBlockIDsRequest.Size(m)
}
func (m *GetEventsForBlockIDsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsForBlockIDsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsForBlockIDsRequest proto.InternalMessageInfo

func (m *GetEventsForBlockIDsRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetEventsForBlockIDsRequest) GetBlockIds() [][]byte {
	if m != nil {
		return m.BlockIds
	}
	return nil
}

type GetTransactionResultRequest struct {
	BlockId              []byte   `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	TransactionId        []byte   `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionResultRequest) Reset()         { *m = GetTransactionResultRequest{} }
func (m *GetTransactionResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionResultRequest) ProtoMessage()    {}
func (*GetTransactionResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{8}
}

func (m *GetTransactionResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionResultRequest.Unmarshal(m, b)
}
func (m *GetTransactionResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionResultRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionResultRequest.Merge(m, src)
}
func (m *GetTransactionResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionResultRequest.Size(m)
}
func (m *GetTransactionResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionResultRequest proto.InternalMessageInfo

func (m *GetTransactionResultRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetTransactionResultRequest) GetTransactionId() []byte {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

type GetTransactionByIndexRequest struct {
	BlockId              []byte   `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionByIndexRequest) Reset()         { *m = GetTransactionByIndexRequest{} }
func (m *GetTransactionByIndexRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionByIndexRequest) ProtoMessage()    {}
func (*GetTransactionByIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{9}
}

func (m *GetTransactionByIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionByIndexRequest.Unmarshal(m, b)
}
func (m *GetTransactionByIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionByIndexRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionByIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionByIndexRequest.Merge(m, src)
}
func (m *GetTransactionByIndexRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionByIndexRequest.Size(m)
}
func (m *GetTransactionByIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionByIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionByIndexRequest proto.InternalMessageInfo

func (m *GetTransactionByIndexRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetTransactionByIndexRequest) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type GetTransactionResultResponse struct {
	StatusCode           uint32            `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	ErrorMessage         string            `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Events               []*entities.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetTransactionResultResponse) Reset()         { *m = GetTransactionResultResponse{} }
func (m *GetTransactionResultResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionResultResponse) ProtoMessage()    {}
func (*GetTransactionResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{10}
}

func (m *GetTransactionResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionResultResponse.Unmarshal(m, b)
}
func (m *GetTransactionResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionResultResponse.Marshal(b, m, deterministic)
}
func (m *GetTransactionResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionResultResponse.Merge(m, src)
}
func (m *GetTransactionResultResponse) XXX_Size() int {
	return xxx_messageInfo_GetTransactionResultResponse.Size(m)
}
func (m *GetTransactionResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionResultResponse proto.InternalMessageInfo

func (m *GetTransactionResultResponse) GetStatusCode() uint32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *GetTransactionResultResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *GetTransactionResultResponse) GetEvents() []*entities.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetRegisterAtBlockIDRequest struct {
	BlockId              []byte   `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	RegisterOwner        []byte   `protobuf:"bytes,2,opt,name=register_owner,json=registerOwner,proto3" json:"register_owner,omitempty"`
	RegisterController   []byte   `protobuf:"bytes,3,opt,name=register_controller,json=registerController,proto3" json:"register_controller,omitempty"`
	RegisterKey          []byte   `protobuf:"bytes,4,opt,name=register_key,json=registerKey,proto3" json:"register_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegisterAtBlockIDRequest) Reset()         { *m = GetRegisterAtBlockIDRequest{} }
func (m *GetRegisterAtBlockIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegisterAtBlockIDRequest) ProtoMessage()    {}
func (*GetRegisterAtBlockIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{11}
}

func (m *GetRegisterAtBlockIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegisterAtBlockIDRequest.Unmarshal(m, b)
}
func (m *GetRegisterAtBlockIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegisterAtBlockIDRequest.Marshal(b, m, deterministic)
}
func (m *GetRegisterAtBlockIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegisterAtBlockIDRequest.Merge(m, src)
}
func (m *GetRegisterAtBlockIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegisterAtBlockIDRequest.Size(m)
}
func (m *GetRegisterAtBlockIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegisterAtBlockIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegisterAtBlockIDRequest proto.InternalMessageInfo

func (m *GetRegisterAtBlockIDRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetRegisterAtBlockIDRequest) GetRegisterOwner() []byte {
	if m != nil {
		return m.RegisterOwner
	}
	return nil
}

func (m *GetRegisterAtBlockIDRequest) GetRegisterController() []byte {
	if m != nil {
		return m.RegisterController
	}
	return nil
}

func (m *GetRegisterAtBlockIDRequest) GetRegisterKey() []byte {
	if m != nil {
		return m.RegisterKey
	}
	return nil
}

type GetRegisterAtBlockIDResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegisterAtBlockIDResponse) Reset()         { *m = GetRegisterAtBlockIDResponse{} }
func (m *GetRegisterAtBlockIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetRegisterAtBlockIDResponse) ProtoMessage()    {}
func (*GetRegisterAtBlockIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{12}
}

func (m *GetRegisterAtBlockIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegisterAtBlockIDResponse.Unmarshal(m, b)
}
func (m *GetRegisterAtBlockIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegisterAtBlockIDResponse.Marshal(b, m, deterministic)
}
func (m *GetRegisterAtBlockIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegisterAtBlockIDResponse.Merge(m, src)
}
func (m *GetRegisterAtBlockIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetRegisterAtBlockIDResponse.Size(m)
}
func (m *GetRegisterAtBlockIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegisterAtBlockIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegisterAtBlockIDResponse proto.InternalMessageInfo

func (m *GetRegisterAtBlockIDResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetLatestBlockHeaderRequest struct {
	IsSealed             bool     `protobuf:"varint,1,opt,name=is_sealed,json=isSealed,proto3" json:"is_sealed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestBlockHeaderRequest) Reset()         { *m = GetLatestBlockHeaderRequest{} }
func (m *GetLatestBlockHeaderRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestBlockHeaderRequest) ProtoMessage()    {}
func (*GetLatestBlockHeaderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{13}
}

func (m *GetLatestBlockHeaderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLatestBlockHeaderRequest.Unmarshal(m, b)
}
func (m *GetLatestBlockHeaderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLatestBlockHeaderRequest.Marshal(b, m, deterministic)
}
func (m *GetLatestBlockHeaderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestBlockHeaderRequest.Merge(m, src)
}
func (m *GetLatestBlockHeaderRequest) XXX_Size() int {
	return xxx_messageInfo_GetLatestBlockHeaderRequest.Size(m)
}
func (m *GetLatestBlockHeaderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestBlockHeaderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestBlockHeaderRequest proto.InternalMessageInfo

func (m *GetLatestBlockHeaderRequest) GetIsSealed() bool {
	if m != nil {
		return m.IsSealed
	}
	return false
}

type GetBlockHeaderByIDRequest struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockHeaderByIDRequest) Reset()         { *m = GetBlockHeaderByIDRequest{} }
func (m *GetBlockHeaderByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockHeaderByIDRequest) ProtoMessage()    {}
func (*GetBlockHeaderByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{14}
}

func (m *GetBlockHeaderByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockHeaderByIDRequest.Unmarshal(m, b)
}
func (m *GetBlockHeaderByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockHeaderByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockHeaderByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockHeaderByIDRequest.Merge(m, src)
}
func (m *GetBlockHeaderByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockHeaderByIDRequest.Size(m)
}
func (m *GetBlockHeaderByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockHeaderByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockHeaderByIDRequest proto.InternalMessageInfo

func (m *GetBlockHeaderByIDRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type BlockHeaderResponse struct {
	Block                *entities.BlockHeader `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BlockHeaderResponse) Reset()         { *m = BlockHeaderResponse{} }
func (m *BlockHeaderResponse) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderResponse) ProtoMessage()    {}
func (*BlockHeaderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_699624211ed46c53, []int{15}
}

func (m *BlockHeaderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeaderResponse.Unmarshal(m, b)
}
func (m *BlockHeaderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeaderResponse.Marshal(b, m, deterministic)
}
func (m *BlockHeaderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderResponse.Merge(m, src)
}
func (m *BlockHeaderResponse) XXX_Size() int {
	return xxx_messageInfo_BlockHeaderResponse.Size(m)
}
func (m *BlockHeaderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderResponse proto.InternalMessageInfo

func (m *BlockHeaderResponse) GetBlock() *entities.BlockHeader {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "flow.execution.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "flow.execution.PingResponse")
	proto.RegisterType((*GetAccountAtBlockIDRequest)(nil), "flow.execution.GetAccountAtBlockIDRequest")
	proto.RegisterType((*GetAccountAtBlockIDResponse)(nil), "flow.execution.GetAccountAtBlockIDResponse")
	proto.RegisterType((*ExecuteScriptAtBlockIDRequest)(nil), "flow.execution.ExecuteScriptAtBlockIDRequest")
	proto.RegisterType((*ExecuteScriptAtBlockIDResponse)(nil), "flow.execution.ExecuteScriptAtBlockIDResponse")
	proto.RegisterType((*GetEventsForBlockIDsResponse)(nil), "flow.execution.GetEventsForBlockIDsResponse")
	proto.RegisterType((*GetEventsForBlockIDsResponse_Result)(nil), "flow.execution.GetEventsForBlockIDsResponse.Result")
	proto.RegisterType((*GetEventsForBlockIDsRequest)(nil), "flow.execution.GetEventsForBlockIDsRequest")
	proto.RegisterType((*GetTransactionResultRequest)(nil), "flow.execution.GetTransactionResultRequest")
	proto.RegisterType((*GetTransactionByIndexRequest)(nil), "flow.execution.GetTransactionByIndexRequest")
	proto.RegisterType((*GetTransactionResultResponse)(nil), "flow.execution.GetTransactionResultResponse")
	proto.RegisterType((*GetRegisterAtBlockIDRequest)(nil), "flow.execution.GetRegisterAtBlockIDRequest")
	proto.RegisterType((*GetRegisterAtBlockIDResponse)(nil), "flow.execution.GetRegisterAtBlockIDResponse")
	proto.RegisterType((*GetLatestBlockHeaderRequest)(nil), "flow.execution.GetLatestBlockHeaderRequest")
	proto.RegisterType((*GetBlockHeaderByIDRequest)(nil), "flow.execution.GetBlockHeaderByIDRequest")
	proto.RegisterType((*BlockHeaderResponse)(nil), "flow.execution.BlockHeaderResponse")
}

func init() { proto.RegisterFile("flow/execution/execution.proto", fileDescriptor_699624211ed46c53) }

var fileDescriptor_699624211ed46c53 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x72, 0xda, 0x46,
	0x14, 0x1e, 0x6c, 0xcc, 0xcf, 0xe1, 0xe7, 0x62, 0xcd, 0x78, 0xb0, 0x70, 0x52, 0xaa, 0x4c, 0x66,
	0xdc, 0x3a, 0x15, 0x2d, 0xe9, 0xf4, 0xa2, 0x77, 0x90, 0xba, 0x84, 0x69, 0x53, 0xa7, 0x4a, 0xaf,
	0x7a, 0xc3, 0x08, 0xe9, 0x44, 0x68, 0x82, 0xb5, 0x64, 0x77, 0x15, 0x87, 0xd7, 0xe8, 0xab, 0xf4,
	0x7d, 0x7a, 0xd3, 0x17, 0xc9, 0x68, 0x77, 0xc5, 0xef, 0x82, 0xe1, 0x86, 0x61, 0xcf, 0x7e, 0xe7,
	0x7c, 0xe7, 0x67, 0xcf, 0x37, 0x82, 0xa7, 0xef, 0xa7, 0xf4, 0xa1, 0x83, 0x9f, 0xd1, 0x4f, 0x44,
	0x44, 0xe3, 0xe5, 0x3f, 0x67, 0xc6, 0xa8, 0xa0, 0xa4, 0x9e, 0xde, 0x3b, 0x0b, 0xab, 0xd5, 0x52,
	0xf8, 0x58, 0x44, 0x22, 0x42, 0xde, 0xf1, 0x7c, 0x9f, 0x26, 0xb1, 0x50, 0x60, 0xab, 0xbd, 0x7e,
	0x39, 0x9e, 0x52, 0xff, 0xc3, 0x68, 0x82, 0x5e, 0x80, 0x4c, 0x23, 0x2e, 0xd7, 0x11, 0xf8, 0x09,
	0x33, 0x67, 0xbb, 0x06, 0x95, 0xb7, 0x51, 0x1c, 0xba, 0xf8, 0x31, 0x41, 0x2e, 0xec, 0x3a, 0x54,
	0xd5, 0x91, 0xcf, 0x68, 0xcc, 0xd1, 0xfe, 0x13, 0xac, 0x01, 0x8a, 0x9e, 0xe2, 0xeb, 0x89, 0x7e,
	0x1a, 0x7b, 0xf8, 0x8b, 0x46, 0x93, 0x4b, 0x28, 0x29, 0xb6, 0x28, 0x68, 0xe6, 0xda, 0xb9, 0xeb,
	0xaa, 0x5b, 0x94, 0xe7, 0x61, 0x40, 0x9a, 0x50, 0xf4, 0x82, 0x80, 0x21, 0xe7, 0xcd, 0x13, 0x75,
	0xa3, 0x8f, 0xf6, 0x1d, 0xb4, 0x8c, 0x21, 0x15, 0x23, 0xf9, 0x1e, 0x8a, 0xba, 0x3c, 0x19, 0xb2,
	0xd2, 0xbd, 0x70, 0x54, 0x33, 0x74, 0xf6, 0x8e, 0xf6, 0x74, 0x33, 0x98, 0x3d, 0x83, 0x27, 0xb7,
	0xb2, 0x53, 0xf8, 0xce, 0x67, 0xd1, 0xec, 0xa8, 0x34, 0x2f, 0xa0, 0xc0, 0xa5, 0x93, 0xce, 0x52,
	0x9f, 0xc8, 0x15, 0x94, 0x3d, 0x16, 0x26, 0xf7, 0x18, 0x0b, 0xde, 0x3c, 0x6d, 0x9f, 0x5e, 0x57,
	0xdd, 0xa5, 0xc1, 0xfe, 0x09, 0x9e, 0xee, 0x62, 0xd4, 0x55, 0x34, 0xe0, 0xec, 0x93, 0x37, 0x4d,
	0x50, 0xf3, 0xa9, 0x83, 0xfd, 0x7f, 0x0e, 0xae, 0x06, 0x28, 0x6e, 0xd3, 0xfe, 0xf3, 0x5f, 0x29,
	0xd3, 0x5e, 0x7c, 0xe1, 0xf6, 0x06, 0x8a, 0x0c, 0x79, 0x32, 0x15, 0xbc, 0x99, 0x6b, 0x9f, 0x5e,
	0x57, 0xba, 0x2f, 0x9d, 0xf5, 0x97, 0xe0, 0xec, 0x73, 0x77, 0x5c, 0xe9, 0xeb, 0x66, 0x31, 0x2c,
	0x01, 0x05, 0x65, 0xda, 0xd7, 0x82, 0xaf, 0xa1, 0x9a, 0x3d, 0x99, 0x28, 0x9c, 0xa8, 0x46, 0xe4,
	0xdd, 0x8a, 0xb4, 0xbd, 0x96, 0x26, 0xf2, 0x02, 0x0a, 0xf2, 0xcd, 0xa8, 0x56, 0x54, 0xba, 0x8d,
	0x8d, 0x91, 0xc8, 0x8c, 0x5c, 0x8d, 0xb1, 0xff, 0x90, 0x03, 0x36, 0x64, 0xa9, 0xa6, 0x41, 0x20,
	0x2f, 0xe6, 0x33, 0xd5, 0x99, 0xb2, 0x2b, 0xff, 0x93, 0x16, 0x94, 0xb3, 0xf4, 0xd2, 0xf7, 0x92,
	0xb6, 0xbb, 0xa4, 0xf3, 0xe3, 0xf6, 0x48, 0xc6, 0xfb, 0x8b, 0x79, 0x31, 0xf7, 0xfc, 0xb4, 0x09,
	0xba, 0xcc, 0xc7, 0xa7, 0xfb, 0x1c, 0xea, 0x62, 0xe9, 0x96, 0x02, 0xd4, 0x94, 0x6b, 0x2b, 0xd6,
	0x61, 0x60, 0xdf, 0xc9, 0xa9, 0xac, 0x10, 0xf4, 0xe7, 0xc3, 0x38, 0xc0, 0xcf, 0x07, 0x30, 0x34,
	0xe0, 0x2c, 0x4a, 0xa1, 0x32, 0x70, 0xcd, 0x55, 0x07, 0xfb, 0x9f, 0xdc, 0x66, 0xc4, 0x2c, 0x65,
	0x3d, 0xe7, 0xaf, 0xa0, 0xc2, 0x85, 0x27, 0x12, 0x3e, 0xf2, 0x69, 0xa0, 0x5a, 0x51, 0x73, 0x41,
	0x99, 0x5e, 0xd1, 0x00, 0xc9, 0x33, 0xa8, 0x21, 0x63, 0x94, 0x8d, 0xee, 0x91, 0x73, 0x2f, 0x44,
	0x19, 0xbf, 0xec, 0x56, 0xa5, 0xf1, 0x8d, 0xb2, 0x1d, 0x39, 0x96, 0x7f, 0x73, 0xb2, 0x8f, 0x2e,
	0x86, 0x11, 0x17, 0xc8, 0x8e, 0xd9, 0x92, 0xe7, 0x50, 0x67, 0xda, 0x6d, 0x44, 0x1f, 0x62, 0x64,
	0x59, 0x1f, 0x33, 0xeb, 0x5d, 0x6a, 0x24, 0x1d, 0x38, 0x5f, 0xc0, 0x7c, 0x1a, 0x0b, 0x46, 0xa7,
	0x53, 0x64, 0xcd, 0x53, 0x89, 0x25, 0xd9, 0xd5, 0xab, 0xc5, 0x4d, 0xfa, 0xf4, 0x16, 0x0e, 0x1f,
	0x70, 0xde, 0xcc, 0x4b, 0x64, 0x25, 0xb3, 0xfd, 0x86, 0x73, 0xfb, 0x47, 0xd9, 0x49, 0x43, 0xd2,
	0x7b, 0x17, 0xed, 0x67, 0x59, 0xea, 0xef, 0x9e, 0x40, 0xae, 0x3c, 0x5e, 0x4b, 0x39, 0xcc, 0x4a,
	0x6d, 0x41, 0x39, 0xe2, 0x23, 0x8e, 0xde, 0x14, 0x55, 0xad, 0x25, 0xb7, 0x14, 0xf1, 0x77, 0xf2,
	0x6c, 0xdf, 0xc0, 0xe5, 0x00, 0x57, 0xbd, 0xfa, 0xf3, 0x65, 0x93, 0xea, 0x70, 0xb2, 0x68, 0xcf,
	0x49, 0x14, 0xd8, 0x03, 0x38, 0x5f, 0x8b, 0xbf, 0x10, 0xb1, 0x33, 0xd9, 0x3b, 0x2d, 0x61, 0xd6,
	0xc6, 0x60, 0x56, 0x5d, 0x14, 0xb0, 0xfb, 0x5f, 0x11, 0xaa, 0xb7, 0xd9, 0x96, 0xf7, 0xde, 0x0e,
	0x49, 0x0f, 0xf2, 0xa9, 0x12, 0x93, 0xd6, 0xa6, 0x02, 0xac, 0xc8, 0xb5, 0x75, 0x65, 0xbe, 0xd4,
	0x59, 0xc4, 0x70, 0x6e, 0x50, 0x5a, 0xf2, 0xad, 0x41, 0x53, 0x76, 0x28, 0xbc, 0x75, 0x73, 0x10,
	0x56, 0xf3, 0x3d, 0xc0, 0x85, 0x59, 0x16, 0xc9, 0x77, 0x9b, 0x61, 0xf6, 0x0a, 0xb6, 0xe5, 0x1c,
	0x0a, 0xd7, 0xc4, 0x1f, 0xa1, 0x61, 0x52, 0x1c, 0x72, 0x73, 0x98, 0x7a, 0x2a, 0xd2, 0x17, 0xc7,
	0x48, 0xad, 0xa6, 0xdc, 0xda, 0x70, 0x23, 0xe5, 0x2e, 0xe9, 0x32, 0x52, 0xee, 0x16, 0x8d, 0xb9,
	0x59, 0x07, 0xb5, 0x58, 0x91, 0x47, 0x82, 0xad, 0x6b, 0xda, 0x91, 0xd4, 0xaa, 0xda, 0xad, 0x2d,
	0x34, 0x56, 0xbb, 0x4b, 0x60, 0x8c, 0x94, 0xbb, 0x17, 0x7b, 0x22, 0x29, 0xb7, 0x56, 0xd8, 0x48,
	0xb9, 0x6b, 0xd1, 0xad, 0x67, 0x9b, 0x60, 0xd3, 0xb2, 0x06, 0x40, 0xb6, 0x17, 0x9e, 0x7c, 0x63,
	0xe0, 0x31, 0x8b, 0xc2, 0x41, 0x2c, 0x7d, 0x17, 0x9e, 0x50, 0x16, 0x3a, 0x34, 0x96, 0x58, 0xf9,
	0xf1, 0x35, 0x4e, 0xde, 0x2f, 0x9d, 0xfe, 0xfe, 0x21, 0x8c, 0xc4, 0x24, 0x19, 0x3b, 0x3e, 0xbd,
	0xef, 0x28, 0x54, 0x47, 0xfe, 0x64, 0xd0, 0x4e, 0x48, 0x3b, 0xeb, 0x9f, 0x8d, 0xe3, 0x82, 0xbc,
	0x7b, 0xf9, 0x25, 0x00, 0x00, 0xff, 0xff, 0x76, 0xa3, 0x86, 0x7b, 0x4f, 0x0a, 0x00, 0x00,
}
