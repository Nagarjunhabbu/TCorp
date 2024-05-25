// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: tcorp.proto

package protocorp

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

type SearchUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserModel `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchUserResponse) Reset() {
	*x = SearchUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResponse) ProtoMessage() {}

func (x *SearchUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResponse.ProtoReflect.Descriptor instead.
func (*SearchUserResponse) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{0}
}

func (x *SearchUserResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchUserResponse) GetData() []*UserModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type SearchUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter []*SearchFilter `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
}

func (x *SearchUserRequest) Reset() {
	*x = SearchUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserRequest) ProtoMessage() {}

func (x *SearchUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserRequest.ProtoReflect.Descriptor instead.
func (*SearchUserRequest) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{1}
}

func (x *SearchUserRequest) GetFilter() []*SearchFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIdRequest) Reset() {
	*x = GetUserByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRequest) ProtoMessage() {}

func (x *GetUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserByIdRespnose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserModel `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByIdRespnose) Reset() {
	*x = GetUserByIdRespnose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdRespnose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRespnose) ProtoMessage() {}

func (x *GetUserByIdRespnose) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRespnose.ProtoReflect.Descriptor instead.
func (*GetUserByIdRespnose) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByIdRespnose) GetUser() *UserModel {
	if x != nil {
		return x.User
	}
	return nil
}

type UserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City    string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Phone   string  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Height  float32 `protobuf:"fixed32,5,opt,name=height,proto3" json:"height,omitempty"`
	Married bool    `protobuf:"varint,6,opt,name=married,proto3" json:"married,omitempty"`
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{4}
}

func (x *UserModel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserModel) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserModel) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserModel) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UserModel) GetMarried() bool {
	if x != nil {
		return x.Married
	}
	return false
}

type GetUserByIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int32 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIdsRequest) Reset() {
	*x = GetUserByIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdsRequest) ProtoMessage() {}

func (x *GetUserByIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdsRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdsRequest) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserByIdsRequest) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetUserByIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*UserModel `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByIdsResponse) Reset() {
	*x = GetUserByIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdsResponse) ProtoMessage() {}

func (x *GetUserByIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdsResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIdsResponse) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserByIdsResponse) GetUser() []*UserModel {
	if x != nil {
		return x.User
	}
	return nil
}

type SearchFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SearchFilter) Reset() {
	*x = SearchFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tcorp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFilter) ProtoMessage() {}

func (x *SearchFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tcorp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFilter.ProtoReflect.Descriptor instead.
func (*SearchFilter) Descriptor() ([]byte, []int) {
	return file_tcorp_proto_rawDescGZIP(), []int{7}
}

func (x *SearchFilter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SearchFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_tcorp_proto protoreflect.FileDescriptor

var file_tcorp_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a,
	0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x11, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6e, 0x6f,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x65, 0x64,
	0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x36, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb8, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6e, 0x6f, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x14,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x17, 0x5a, 0x15, 0x54, 0x43, 0x6f, 0x72, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x72, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tcorp_proto_rawDescOnce sync.Once
	file_tcorp_proto_rawDescData = file_tcorp_proto_rawDesc
)

func file_tcorp_proto_rawDescGZIP() []byte {
	file_tcorp_proto_rawDescOnce.Do(func() {
		file_tcorp_proto_rawDescData = protoimpl.X.CompressGZIP(file_tcorp_proto_rawDescData)
	})
	return file_tcorp_proto_rawDescData
}

var file_tcorp_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tcorp_proto_goTypes = []interface{}{
	(*SearchUserResponse)(nil),   // 0: SearchUserResponse
	(*SearchUserRequest)(nil),    // 1: SearchUserRequest
	(*GetUserByIdRequest)(nil),   // 2: GetUserByIdRequest
	(*GetUserByIdRespnose)(nil),  // 3: GetUserByIdRespnose
	(*UserModel)(nil),            // 4: UserModel
	(*GetUserByIdsRequest)(nil),  // 5: GetUserByIdsRequest
	(*GetUserByIdsResponse)(nil), // 6: GetUserByIdsResponse
	(*SearchFilter)(nil),         // 7: SearchFilter
}
var file_tcorp_proto_depIdxs = []int32{
	4, // 0: SearchUserResponse.data:type_name -> UserModel
	7, // 1: SearchUserRequest.filter:type_name -> SearchFilter
	4, // 2: GetUserByIdRespnose.user:type_name -> UserModel
	4, // 3: GetUserByIdsResponse.user:type_name -> UserModel
	2, // 4: UserInfo.GetUserById:input_type -> GetUserByIdRequest
	5, // 5: UserInfo.GetUserByIds:input_type -> GetUserByIdsRequest
	1, // 6: UserInfo.SearchUser:input_type -> SearchUserRequest
	3, // 7: UserInfo.GetUserById:output_type -> GetUserByIdRespnose
	6, // 8: UserInfo.GetUserByIds:output_type -> GetUserByIdsResponse
	0, // 9: UserInfo.SearchUser:output_type -> SearchUserResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tcorp_proto_init() }
func file_tcorp_proto_init() {
	if File_tcorp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tcorp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserResponse); i {
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
		file_tcorp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserRequest); i {
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
		file_tcorp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdRequest); i {
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
		file_tcorp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdRespnose); i {
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
		file_tcorp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModel); i {
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
		file_tcorp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdsRequest); i {
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
		file_tcorp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIdsResponse); i {
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
		file_tcorp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFilter); i {
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
			RawDescriptor: file_tcorp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tcorp_proto_goTypes,
		DependencyIndexes: file_tcorp_proto_depIdxs,
		MessageInfos:      file_tcorp_proto_msgTypes,
	}.Build()
	File_tcorp_proto = out.File
	file_tcorp_proto_rawDesc = nil
	file_tcorp_proto_goTypes = nil
	file_tcorp_proto_depIdxs = nil
}
