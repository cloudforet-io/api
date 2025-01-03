// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: spaceone/api/identity/plugin/external_auth.proto

package plugin

import (
	_ "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type UserInfo_State int32

const (
	UserInfo_NONE         UserInfo_State = 0
	UserInfo_ENABLED      UserInfo_State = 1
	UserInfo_DISABLED     UserInfo_State = 2
	UserInfo_UNIDENTIFIED UserInfo_State = 3
)

// Enum value maps for UserInfo_State.
var (
	UserInfo_State_name = map[int32]string{
		0: "NONE",
		1: "ENABLED",
		2: "DISABLED",
		3: "UNIDENTIFIED",
	}
	UserInfo_State_value = map[string]int32{
		"NONE":         0,
		"ENABLED":      1,
		"DISABLED":     2,
		"UNIDENTIFIED": 3,
	}
)

func (x UserInfo_State) Enum() *UserInfo_State {
	p := new(UserInfo_State)
	*p = x
	return p
}

func (x UserInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_plugin_external_auth_proto_enumTypes[0].Descriptor()
}

func (UserInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_plugin_external_auth_proto_enumTypes[0]
}

func (x UserInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInfo_State.Descriptor instead.
func (UserInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_external_auth_proto_rawDescGZIP(), []int{2, 0}
}

type InitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Options       *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_external_auth_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *InitRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AuthorizeRequest struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Options *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	SchemaId string `protobuf:"bytes,2,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// +optional
	SecretData  *_struct.Struct `protobuf:"bytes,3,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	Credentials *_struct.Struct `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// +optional
	Metadata      *_struct.Struct `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	DomainId      string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthorizeRequest) Reset() {
	*x = AuthorizeRequest{}
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequest) ProtoMessage() {}

func (x *AuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_external_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *AuthorizeRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *AuthorizeRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *AuthorizeRequest) GetCredentials() *_struct.Struct {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *AuthorizeRequest) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AuthorizeRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Mobile        string                 `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Group         string                 `protobuf:"bytes,5,opt,name=group,proto3" json:"group,omitempty"`
	State         UserInfo_State         `protobuf:"varint,6,opt,name=state,proto3,enum=spaceone.api.identity.plugin.UserInfo_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_external_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserInfo) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *UserInfo) GetState() UserInfo_State {
	if x != nil {
		return x.State
	}
	return UserInfo_NONE
}

type PluginInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *_struct.Struct        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_plugin_external_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_plugin_external_auth_proto_rawDescGZIP(), []int{3}
}

func (x *PluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spaceone_api_identity_plugin_external_auth_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_plugin_external_auth_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0b, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xa9, 0x02, 0x0a, 0x10, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xff, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x03, 0x22, 0x41, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd4, 0x01, 0x0a, 0x0c,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x12, 0x5d, 0x0a, 0x04,
	0x69, 0x6e, 0x69, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_plugin_external_auth_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_plugin_external_auth_proto_rawDescData = file_spaceone_api_identity_plugin_external_auth_proto_rawDesc
)

func file_spaceone_api_identity_plugin_external_auth_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_plugin_external_auth_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_plugin_external_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_plugin_external_auth_proto_rawDescData)
	})
	return file_spaceone_api_identity_plugin_external_auth_proto_rawDescData
}

var file_spaceone_api_identity_plugin_external_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_identity_plugin_external_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spaceone_api_identity_plugin_external_auth_proto_goTypes = []any{
	(UserInfo_State)(0),      // 0: spaceone.api.identity.plugin.UserInfo.State
	(*InitRequest)(nil),      // 1: spaceone.api.identity.plugin.InitRequest
	(*AuthorizeRequest)(nil), // 2: spaceone.api.identity.plugin.AuthorizeRequest
	(*UserInfo)(nil),         // 3: spaceone.api.identity.plugin.UserInfo
	(*PluginInfo)(nil),       // 4: spaceone.api.identity.plugin.PluginInfo
	(*_struct.Struct)(nil),   // 5: google.protobuf.Struct
}
var file_spaceone_api_identity_plugin_external_auth_proto_depIdxs = []int32{
	5, // 0: spaceone.api.identity.plugin.InitRequest.options:type_name -> google.protobuf.Struct
	5, // 1: spaceone.api.identity.plugin.AuthorizeRequest.options:type_name -> google.protobuf.Struct
	5, // 2: spaceone.api.identity.plugin.AuthorizeRequest.secret_data:type_name -> google.protobuf.Struct
	5, // 3: spaceone.api.identity.plugin.AuthorizeRequest.credentials:type_name -> google.protobuf.Struct
	5, // 4: spaceone.api.identity.plugin.AuthorizeRequest.metadata:type_name -> google.protobuf.Struct
	0, // 5: spaceone.api.identity.plugin.UserInfo.state:type_name -> spaceone.api.identity.plugin.UserInfo.State
	5, // 6: spaceone.api.identity.plugin.PluginInfo.metadata:type_name -> google.protobuf.Struct
	1, // 7: spaceone.api.identity.plugin.ExternalAuth.init:input_type -> spaceone.api.identity.plugin.InitRequest
	2, // 8: spaceone.api.identity.plugin.ExternalAuth.authorize:input_type -> spaceone.api.identity.plugin.AuthorizeRequest
	4, // 9: spaceone.api.identity.plugin.ExternalAuth.init:output_type -> spaceone.api.identity.plugin.PluginInfo
	3, // 10: spaceone.api.identity.plugin.ExternalAuth.authorize:output_type -> spaceone.api.identity.plugin.UserInfo
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_plugin_external_auth_proto_init() }
func file_spaceone_api_identity_plugin_external_auth_proto_init() {
	if File_spaceone_api_identity_plugin_external_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_identity_plugin_external_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_plugin_external_auth_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_plugin_external_auth_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_plugin_external_auth_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_plugin_external_auth_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_plugin_external_auth_proto = out.File
	file_spaceone_api_identity_plugin_external_auth_proto_rawDesc = nil
	file_spaceone_api_identity_plugin_external_auth_proto_goTypes = nil
	file_spaceone_api_identity_plugin_external_auth_proto_depIdxs = nil
}
