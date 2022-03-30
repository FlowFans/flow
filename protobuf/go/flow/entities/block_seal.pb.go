// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: flow/entities/block_seal.proto

package entities

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockSeal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId                    []byte                 `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	ExecutionReceiptId         []byte                 `protobuf:"bytes,2,opt,name=execution_receipt_id,json=executionReceiptId,proto3" json:"execution_receipt_id,omitempty"`
	ExecutionReceiptSignatures [][]byte               `protobuf:"bytes,3,rep,name=execution_receipt_signatures,json=executionReceiptSignatures,proto3" json:"execution_receipt_signatures,omitempty"`
	ResultApprovalSignatures   [][]byte               `protobuf:"bytes,4,rep,name=result_approval_signatures,json=resultApprovalSignatures,proto3" json:"result_approval_signatures,omitempty"`
	FinalState                 []byte                 `protobuf:"bytes,5,opt,name=final_state,json=finalState,proto3" json:"final_state,omitempty"`
	ResultId                   []byte                 `protobuf:"bytes,6,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	AggregatedApprovalSigs     []*AggregatedSignature `protobuf:"bytes,7,rep,name=aggregated_approval_sigs,json=aggregatedApprovalSigs,proto3" json:"aggregated_approval_sigs,omitempty"`
}

func (x *BlockSeal) Reset() {
	*x = BlockSeal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_entities_block_seal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockSeal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockSeal) ProtoMessage() {}

func (x *BlockSeal) ProtoReflect() protoreflect.Message {
	mi := &file_flow_entities_block_seal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockSeal.ProtoReflect.Descriptor instead.
func (*BlockSeal) Descriptor() ([]byte, []int) {
	return file_flow_entities_block_seal_proto_rawDescGZIP(), []int{0}
}

func (x *BlockSeal) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *BlockSeal) GetExecutionReceiptId() []byte {
	if x != nil {
		return x.ExecutionReceiptId
	}
	return nil
}

func (x *BlockSeal) GetExecutionReceiptSignatures() [][]byte {
	if x != nil {
		return x.ExecutionReceiptSignatures
	}
	return nil
}

func (x *BlockSeal) GetResultApprovalSignatures() [][]byte {
	if x != nil {
		return x.ResultApprovalSignatures
	}
	return nil
}

func (x *BlockSeal) GetFinalState() []byte {
	if x != nil {
		return x.FinalState
	}
	return nil
}

func (x *BlockSeal) GetResultId() []byte {
	if x != nil {
		return x.ResultId
	}
	return nil
}

func (x *BlockSeal) GetAggregatedApprovalSigs() []*AggregatedSignature {
	if x != nil {
		return x.AggregatedApprovalSigs
	}
	return nil
}

type AggregatedSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifierSignatures [][]byte `protobuf:"bytes,1,rep,name=verifier_signatures,json=verifierSignatures,proto3" json:"verifier_signatures,omitempty"`
	SignerIds          [][]byte `protobuf:"bytes,2,rep,name=signer_ids,json=signerIds,proto3" json:"signer_ids,omitempty"`
}

func (x *AggregatedSignature) Reset() {
	*x = AggregatedSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_entities_block_seal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedSignature) ProtoMessage() {}

func (x *AggregatedSignature) ProtoReflect() protoreflect.Message {
	mi := &file_flow_entities_block_seal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedSignature.ProtoReflect.Descriptor instead.
func (*AggregatedSignature) Descriptor() ([]byte, []int) {
	return file_flow_entities_block_seal_proto_rawDescGZIP(), []int{1}
}

func (x *AggregatedSignature) GetVerifierSignatures() [][]byte {
	if x != nil {
		return x.VerifierSignatures
	}
	return nil
}

func (x *AggregatedSignature) GetSignerIds() [][]byte {
	if x != nil {
		return x.SignerIds
	}
	return nil
}

var File_flow_entities_block_seal_proto protoreflect.FileDescriptor

var file_flow_entities_block_seal_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0xf4, 0x02, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x61, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x1c, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x1a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x1a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x18, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x18, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x16,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x53, 0x69, 0x67, 0x73, 0x22, 0x65, 0x0a, 0x13, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73, 0x42, 0x50, 0x0a,
	0x1c, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_entities_block_seal_proto_rawDescOnce sync.Once
	file_flow_entities_block_seal_proto_rawDescData = file_flow_entities_block_seal_proto_rawDesc
)

func file_flow_entities_block_seal_proto_rawDescGZIP() []byte {
	file_flow_entities_block_seal_proto_rawDescOnce.Do(func() {
		file_flow_entities_block_seal_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_entities_block_seal_proto_rawDescData)
	})
	return file_flow_entities_block_seal_proto_rawDescData
}

var file_flow_entities_block_seal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flow_entities_block_seal_proto_goTypes = []interface{}{
	(*BlockSeal)(nil),           // 0: flow.entities.BlockSeal
	(*AggregatedSignature)(nil), // 1: flow.entities.AggregatedSignature
}
var file_flow_entities_block_seal_proto_depIdxs = []int32{
	1, // 0: flow.entities.BlockSeal.aggregated_approval_sigs:type_name -> flow.entities.AggregatedSignature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flow_entities_block_seal_proto_init() }
func file_flow_entities_block_seal_proto_init() {
	if File_flow_entities_block_seal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_entities_block_seal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockSeal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flow_entities_block_seal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregatedSignature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flow_entities_block_seal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_entities_block_seal_proto_goTypes,
		DependencyIndexes: file_flow_entities_block_seal_proto_depIdxs,
		MessageInfos:      file_flow_entities_block_seal_proto_msgTypes,
	}.Build()
	File_flow_entities_block_seal_proto = out.File
	file_flow_entities_block_seal_proto_rawDesc = nil
	file_flow_entities_block_seal_proto_goTypes = nil
	file_flow_entities_block_seal_proto_depIdxs = nil
}
