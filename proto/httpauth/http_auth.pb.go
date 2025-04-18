// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.27.3
// source: http_auth.proto

package httpauth

import (
	reflect "reflect"
	sync "sync"

	achieve "github.com/Andrew-M-C/trpc-go-demo/proto/achieve"
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

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *common.Metadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PasswordHash  string                 `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_http_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ErrCode       int32                  `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code"`
	ErrMsg        string                 `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Data          *LoginResponse_Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_http_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *LoginResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *LoginResponse) GetData() *LoginResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type SynchronizeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *common.Metadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SynchronizeRequest) Reset() {
	*x = SynchronizeRequest{}
	mi := &file_http_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SynchronizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeRequest) ProtoMessage() {}

func (x *SynchronizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeRequest.ProtoReflect.Descriptor instead.
func (*SynchronizeRequest) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SynchronizeRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type SynchronizeResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	ErrCode       int32                     `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code"`
	ErrMsg        string                    `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Data          *SynchronizeResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SynchronizeResponse) Reset() {
	*x = SynchronizeResponse{}
	mi := &file_http_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SynchronizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeResponse) ProtoMessage() {}

func (x *SynchronizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeResponse.ProtoReflect.Descriptor instead.
func (*SynchronizeResponse) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SynchronizeResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *SynchronizeResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SynchronizeResponse) GetData() *SynchronizeResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserSpaceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *common.Metadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSpaceRequest) Reset() {
	*x = UserSpaceRequest{}
	mi := &file_http_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpaceRequest) ProtoMessage() {}

func (x *UserSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpaceRequest.ProtoReflect.Descriptor instead.
func (*UserSpaceRequest) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserSpaceRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UserSpaceRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserSpaceResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	ErrCode       int32                   `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code"`
	ErrMsg        string                  `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Data          *UserSpaceResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSpaceResponse) Reset() {
	*x = UserSpaceResponse{}
	mi := &file_http_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpaceResponse) ProtoMessage() {}

func (x *UserSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpaceResponse.ProtoReflect.Descriptor instead.
func (*UserSpaceResponse) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UserSpaceResponse) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *UserSpaceResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *UserSpaceResponse) GetData() *UserSpaceResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type LoginResponse_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdTicket      string                 `protobuf:"bytes,1,opt,name=id_ticket,json=idTicket,proto3" json:"id_ticket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse_Data) Reset() {
	*x = LoginResponse_Data{}
	mi := &file_http_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse_Data) ProtoMessage() {}

func (x *LoginResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse_Data.ProtoReflect.Descriptor instead.
func (*LoginResponse_Data) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LoginResponse_Data) GetIdTicket() string {
	if x != nil {
		return x.IdTicket
	}
	return ""
}

type SynchronizeResponse_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TsMsec        int64                  `protobuf:"varint,1,opt,name=ts_msec,json=tsMsec,proto3" json:"ts_msec,omitempty"` // 毫秒级时间戳
	Desc          string                 `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`                    // 时间可视化描述
	Timezone      string                 `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone,omitempty"`            // 时区
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SynchronizeResponse_Data) Reset() {
	*x = SynchronizeResponse_Data{}
	mi := &file_http_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SynchronizeResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeResponse_Data) ProtoMessage() {}

func (x *SynchronizeResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeResponse_Data.ProtoReflect.Descriptor instead.
func (*SynchronizeResponse_Data) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SynchronizeResponse_Data) GetTsMsec() int64 {
	if x != nil {
		return x.TsMsec
	}
	return 0
}

func (x *SynchronizeResponse_Data) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *SynchronizeResponse_Data) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type UserSpaceResponse_Data struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	UserInfo      *UserSpaceResponse_UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Reputation    *achieve.Reputation         `protobuf:"bytes,2,opt,name=reputation,proto3" json:"reputation,omitempty"`
	Badges        []*achieve.Badge            `protobuf:"bytes,3,rep,name=badges,proto3" json:"badges,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSpaceResponse_Data) Reset() {
	*x = UserSpaceResponse_Data{}
	mi := &file_http_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSpaceResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpaceResponse_Data) ProtoMessage() {}

func (x *UserSpaceResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpaceResponse_Data.ProtoReflect.Descriptor instead.
func (*UserSpaceResponse_Data) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{5, 0}
}

func (x *UserSpaceResponse_Data) GetUserInfo() *UserSpaceResponse_UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *UserSpaceResponse_Data) GetReputation() *achieve.Reputation {
	if x != nil {
		return x.Reputation
	}
	return nil
}

func (x *UserSpaceResponse_Data) GetBadges() []*achieve.Badge {
	if x != nil {
		return x.Badges
	}
	return nil
}

type UserSpaceResponse_UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Nickname      string                 `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	CreateTsSec   int64                  `protobuf:"varint,4,opt,name=create_ts_sec,json=createTsSec,proto3" json:"create_ts_sec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSpaceResponse_UserInfo) Reset() {
	*x = UserSpaceResponse_UserInfo{}
	mi := &file_http_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSpaceResponse_UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpaceResponse_UserInfo) ProtoMessage() {}

func (x *UserSpaceResponse_UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_http_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpaceResponse_UserInfo.ProtoReflect.Descriptor instead.
func (*UserSpaceResponse_UserInfo) Descriptor() ([]byte, []int) {
	return file_http_auth_proto_rawDescGZIP(), []int{5, 1}
}

func (x *UserSpaceResponse_UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserSpaceResponse_UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSpaceResponse_UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserSpaceResponse_UserInfo) GetCreateTsSec() int64 {
	if x != nil {
		return x.CreateTsSec
	}
	return 0
}

var File_http_auth_proto protoreflect.FileDescriptor

var file_http_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x22, 0xa4, 0x01,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x23, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x5f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x4c, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x40,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x4f, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x73, 0x5f, 0x6d,
	0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x73, 0x4d, 0x73, 0x65,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x22, 0x66, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcf, 0x03, 0x0a, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x3e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0xc4, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x61, 0x64, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x2e, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x52, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x1a, 0x7f, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x53, 0x65, 0x63, 0x32, 0x8e, 0x02, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x4c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x2e,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x65, 0x12, 0x26, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x24, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65,
	0x77, 0x2d, 0x4d, 0x2d, 0x43, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x61, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_auth_proto_rawDescOnce sync.Once
	file_http_auth_proto_rawDescData = file_http_auth_proto_rawDesc
)

func file_http_auth_proto_rawDescGZIP() []byte {
	file_http_auth_proto_rawDescOnce.Do(func() {
		file_http_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_auth_proto_rawDescData)
	})
	return file_http_auth_proto_rawDescData
}

var file_http_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_http_auth_proto_goTypes = []any{
	(*LoginRequest)(nil),               // 0: trpc.demo.httpauth.LoginRequest
	(*LoginResponse)(nil),              // 1: trpc.demo.httpauth.LoginResponse
	(*SynchronizeRequest)(nil),         // 2: trpc.demo.httpauth.SynchronizeRequest
	(*SynchronizeResponse)(nil),        // 3: trpc.demo.httpauth.SynchronizeResponse
	(*UserSpaceRequest)(nil),           // 4: trpc.demo.httpauth.UserSpaceRequest
	(*UserSpaceResponse)(nil),          // 5: trpc.demo.httpauth.UserSpaceResponse
	(*LoginResponse_Data)(nil),         // 6: trpc.demo.httpauth.LoginResponse.Data
	(*SynchronizeResponse_Data)(nil),   // 7: trpc.demo.httpauth.SynchronizeResponse.Data
	(*UserSpaceResponse_Data)(nil),     // 8: trpc.demo.httpauth.UserSpaceResponse.Data
	(*UserSpaceResponse_UserInfo)(nil), // 9: trpc.demo.httpauth.UserSpaceResponse.UserInfo
	(*common.Metadata)(nil),            // 10: trpc.demo.common.Metadata
	(*achieve.Reputation)(nil),         // 11: trpc.demo.achieve.Reputation
	(*achieve.Badge)(nil),              // 12: trpc.demo.achieve.Badge
}
var file_http_auth_proto_depIdxs = []int32{
	10, // 0: trpc.demo.httpauth.LoginRequest.metadata:type_name -> trpc.demo.common.Metadata
	6,  // 1: trpc.demo.httpauth.LoginResponse.data:type_name -> trpc.demo.httpauth.LoginResponse.Data
	10, // 2: trpc.demo.httpauth.SynchronizeRequest.metadata:type_name -> trpc.demo.common.Metadata
	7,  // 3: trpc.demo.httpauth.SynchronizeResponse.data:type_name -> trpc.demo.httpauth.SynchronizeResponse.Data
	10, // 4: trpc.demo.httpauth.UserSpaceRequest.metadata:type_name -> trpc.demo.common.Metadata
	8,  // 5: trpc.demo.httpauth.UserSpaceResponse.data:type_name -> trpc.demo.httpauth.UserSpaceResponse.Data
	9,  // 6: trpc.demo.httpauth.UserSpaceResponse.Data.user_info:type_name -> trpc.demo.httpauth.UserSpaceResponse.UserInfo
	11, // 7: trpc.demo.httpauth.UserSpaceResponse.Data.reputation:type_name -> trpc.demo.achieve.Reputation
	12, // 8: trpc.demo.httpauth.UserSpaceResponse.Data.badges:type_name -> trpc.demo.achieve.Badge
	0,  // 9: trpc.demo.httpauth.Auth.Login:input_type -> trpc.demo.httpauth.LoginRequest
	2,  // 10: trpc.demo.httpauth.Auth.Synchronize:input_type -> trpc.demo.httpauth.SynchronizeRequest
	4,  // 11: trpc.demo.httpauth.Auth.UserSpace:input_type -> trpc.demo.httpauth.UserSpaceRequest
	1,  // 12: trpc.demo.httpauth.Auth.Login:output_type -> trpc.demo.httpauth.LoginResponse
	3,  // 13: trpc.demo.httpauth.Auth.Synchronize:output_type -> trpc.demo.httpauth.SynchronizeResponse
	5,  // 14: trpc.demo.httpauth.Auth.UserSpace:output_type -> trpc.demo.httpauth.UserSpaceResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_http_auth_proto_init() }
func file_http_auth_proto_init() {
	if File_http_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_http_auth_proto_goTypes,
		DependencyIndexes: file_http_auth_proto_depIdxs,
		MessageInfos:      file_http_auth_proto_msgTypes,
	}.Build()
	File_http_auth_proto = out.File
	file_http_auth_proto_rawDesc = nil
	file_http_auth_proto_goTypes = nil
	file_http_auth_proto_depIdxs = nil
}
