// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: mq.proto

package mq

import (
	reflect "reflect"
	sync "sync"

	common "github.com/Andrew-M-C/trpc-go-demo/proto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestMQAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Count    int32            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Type     string           `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"` // json 或 string
}

func (x *TestMQAddRequest) Reset() {
	*x = TestMQAddRequest{}
	mi := &file_mq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestMQAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMQAddRequest) ProtoMessage() {}

func (x *TestMQAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMQAddRequest.ProtoReflect.Descriptor instead.
func (*TestMQAddRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{0}
}

func (x *TestMQAddRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TestMQAddRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *TestMQAddRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type TestMQAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int32  `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *TestMQAddResponse) Reset() {
	*x = TestMQAddResponse{}
	mi := &file_mq_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestMQAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMQAddResponse) ProtoMessage() {}

func (x *TestMQAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMQAddResponse.ProtoReflect.Descriptor instead.
func (*TestMQAddResponse) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{1}
}

func (x *TestMQAddResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *TestMQAddResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

var File_mq_proto protoreflect.FileDescriptor

var file_mq_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x6d, 0x71, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x74, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x51, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x51, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x32, 0x52,
	0x0a, 0x02, 0x4d, 0x51, 0x12, 0x4c, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x51, 0x41, 0x64,
	0x64, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x6d, 0x71,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x51, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x6d, 0x71,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x51, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x77, 0x2d, 0x4d, 0x2d, 0x43, 0x2f, 0x74, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mq_proto_rawDescOnce sync.Once
	file_mq_proto_rawDescData = file_mq_proto_rawDesc
)

func file_mq_proto_rawDescGZIP() []byte {
	file_mq_proto_rawDescOnce.Do(func() {
		file_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_proto_rawDescData)
	})
	return file_mq_proto_rawDescData
}

var file_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mq_proto_goTypes = []any{
	(*TestMQAddRequest)(nil),  // 0: trpc.demo.mq.TestMQAddRequest
	(*TestMQAddResponse)(nil), // 1: trpc.demo.mq.TestMQAddResponse
	(*common.Metadata)(nil),   // 2: trpc.demo.common.Metadata
}
var file_mq_proto_depIdxs = []int32{
	2, // 0: trpc.demo.mq.TestMQAddRequest.metadata:type_name -> trpc.demo.common.Metadata
	0, // 1: trpc.demo.mq.MQ.TestMQAdd:input_type -> trpc.demo.mq.TestMQAddRequest
	1, // 2: trpc.demo.mq.MQ.TestMQAdd:output_type -> trpc.demo.mq.TestMQAddResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mq_proto_init() }
func file_mq_proto_init() {
	if File_mq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mq_proto_goTypes,
		DependencyIndexes: file_mq_proto_depIdxs,
		MessageInfos:      file_mq_proto_msgTypes,
	}.Build()
	File_mq_proto = out.File
	file_mq_proto_rawDesc = nil
	file_mq_proto_goTypes = nil
	file_mq_proto_depIdxs = nil
}
