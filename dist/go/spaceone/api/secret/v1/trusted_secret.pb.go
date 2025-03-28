// Trusted secret is a resource that stores and manages credentials.
// Trusted secret is merged with linked secret and used to access data in other microservices.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/secret/v1/trusted_secret.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTrustedSecretRequest_ResourceGroup int32

const (
	CreateTrustedSecretRequest_RESOURCE_GROUP_NONE CreateTrustedSecretRequest_ResourceGroup = 0
	CreateTrustedSecretRequest_DOMAIN              CreateTrustedSecretRequest_ResourceGroup = 1
	CreateTrustedSecretRequest_WORKSPACE           CreateTrustedSecretRequest_ResourceGroup = 2
)

// Enum value maps for CreateTrustedSecretRequest_ResourceGroup.
var (
	CreateTrustedSecretRequest_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	CreateTrustedSecretRequest_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"DOMAIN":              1,
		"WORKSPACE":           2,
	}
)

func (x CreateTrustedSecretRequest_ResourceGroup) Enum() *CreateTrustedSecretRequest_ResourceGroup {
	p := new(CreateTrustedSecretRequest_ResourceGroup)
	*p = x
	return p
}

func (x CreateTrustedSecretRequest_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateTrustedSecretRequest_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_secret_v1_trusted_secret_proto_enumTypes[0].Descriptor()
}

func (CreateTrustedSecretRequest_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_secret_v1_trusted_secret_proto_enumTypes[0]
}

func (x CreateTrustedSecretRequest_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateTrustedSecretRequest_ResourceGroup.Descriptor instead.
func (CreateTrustedSecretRequest_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{0, 0}
}

type TrustedSecretInfo_ResourceGroup int32

const (
	TrustedSecretInfo_RESOURCE_GROUP_NONE TrustedSecretInfo_ResourceGroup = 0
	TrustedSecretInfo_DOMAIN              TrustedSecretInfo_ResourceGroup = 1
	TrustedSecretInfo_WORKSPACE           TrustedSecretInfo_ResourceGroup = 2
)

// Enum value maps for TrustedSecretInfo_ResourceGroup.
var (
	TrustedSecretInfo_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	TrustedSecretInfo_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"DOMAIN":              1,
		"WORKSPACE":           2,
	}
)

func (x TrustedSecretInfo_ResourceGroup) Enum() *TrustedSecretInfo_ResourceGroup {
	p := new(TrustedSecretInfo_ResourceGroup)
	*p = x
	return p
}

func (x TrustedSecretInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrustedSecretInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_secret_v1_trusted_secret_proto_enumTypes[1].Descriptor()
}

func (TrustedSecretInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_secret_v1_trusted_secret_proto_enumTypes[1]
}

func (x TrustedSecretInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrustedSecretInfo_ResourceGroup.Descriptor instead.
func (TrustedSecretInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{7, 0}
}

//	{
//	   "name": "Cloudforet Broker Account - Managed",
//	   "data": "********",
//	   "schema_id": "aws_access_key",
//	   "trusted_account_id": "trusted-sa-123456789012",
//	   "tags": {}
//	   "resource_group": "DOMAIN"
//	}
type CreateTrustedSecretRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data  *_struct.Struct        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// +optional
	SchemaId string `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// +optional
	Tags          *_struct.Struct                          `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	ResourceGroup CreateTrustedSecretRequest_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.secret.v1.CreateTrustedSecretRequest_ResourceGroup" json:"resource_group,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	TrustedAccountId string `protobuf:"bytes,22,opt,name=trusted_account_id,json=trustedAccountId,proto3" json:"trusted_account_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateTrustedSecretRequest) Reset() {
	*x = CreateTrustedSecretRequest{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTrustedSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrustedSecretRequest) ProtoMessage() {}

func (x *CreateTrustedSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrustedSecretRequest.ProtoReflect.Descriptor instead.
func (*CreateTrustedSecretRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTrustedSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTrustedSecretRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateTrustedSecretRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *CreateTrustedSecretRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateTrustedSecretRequest) GetResourceGroup() CreateTrustedSecretRequest_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return CreateTrustedSecretRequest_RESOURCE_GROUP_NONE
}

func (x *CreateTrustedSecretRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreateTrustedSecretRequest) GetTrustedAccountId() string {
	if x != nil {
		return x.TrustedAccountId
	}
	return ""
}

//	{
//	   "trusted_secret_id": "trusted-secret-123456789012",
//	   "name": "aws-dev2",
//	   "tags": { "a": "b"}
//	}
type UpdateTrustedSecretRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TrustedSecretId string                 `protobuf:"bytes,1,opt,name=trusted_secret_id,json=trustedSecretId,proto3" json:"trusted_secret_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTrustedSecretRequest) Reset() {
	*x = UpdateTrustedSecretRequest{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTrustedSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrustedSecretRequest) ProtoMessage() {}

func (x *UpdateTrustedSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrustedSecretRequest.ProtoReflect.Descriptor instead.
func (*UpdateTrustedSecretRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTrustedSecretRequest) GetTrustedSecretId() string {
	if x != nil {
		return x.TrustedSecretId
	}
	return ""
}

func (x *UpdateTrustedSecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTrustedSecretRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

//	{
//	   "trusted_secret_id": "trusted-secret-123456789012"
//	}
type TrustedSecretRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TrustedSecretId string                 `protobuf:"bytes,1,opt,name=trusted_secret_id,json=trustedSecretId,proto3" json:"trusted_secret_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TrustedSecretRequest) Reset() {
	*x = TrustedSecretRequest{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustedSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedSecretRequest) ProtoMessage() {}

func (x *TrustedSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedSecretRequest.ProtoReflect.Descriptor instead.
func (*TrustedSecretRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{2}
}

func (x *TrustedSecretRequest) GetTrustedSecretId() string {
	if x != nil {
		return x.TrustedSecretId
	}
	return ""
}

//	{
//	   "trusted_secret_id": "ta-123456789012",
//	   "domain_id": "domain-12345abcde"
//	}
type GetTrustedSecretDataRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TrustedSecretId string                 `protobuf:"bytes,1,opt,name=trusted_secret_id,json=trustedSecretId,proto3" json:"trusted_secret_id,omitempty"`
	DomainId        string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetTrustedSecretDataRequest) Reset() {
	*x = GetTrustedSecretDataRequest{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTrustedSecretDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrustedSecretDataRequest) ProtoMessage() {}

func (x *GetTrustedSecretDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrustedSecretDataRequest.ProtoReflect.Descriptor instead.
func (*GetTrustedSecretDataRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{3}
}

func (x *GetTrustedSecretDataRequest) GetTrustedSecretId() string {
	if x != nil {
		return x.TrustedSecretId
	}
	return ""
}

func (x *GetTrustedSecretDataRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "trusted_secret_id": "trusted-secret-12345abcde",
//	   "data": "********",
//	}
type UpdateTrustedSecretDataRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TrustedSecretId string                 `protobuf:"bytes,1,opt,name=trusted_secret_id,json=trustedSecretId,proto3" json:"trusted_secret_id,omitempty"`
	SchemaId        string                 `protobuf:"bytes,2,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Data            *_struct.Struct        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateTrustedSecretDataRequest) Reset() {
	*x = UpdateTrustedSecretDataRequest{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTrustedSecretDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrustedSecretDataRequest) ProtoMessage() {}

func (x *UpdateTrustedSecretDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrustedSecretDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateTrustedSecretDataRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTrustedSecretDataRequest) GetTrustedSecretId() string {
	if x != nil {
		return x.TrustedSecretId
	}
	return ""
}

func (x *UpdateTrustedSecretDataRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *UpdateTrustedSecretDataRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

//	{
//	   "query": {}
//	}
type TrustedSecretQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	TrustedSecretId string `protobuf:"bytes,2,opt,name=trusted_secret_id,json=trustedSecretId,proto3" json:"trusted_secret_id,omitempty"`
	// / +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	SchemaId string `protobuf:"bytes,4,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	TrustedAccountId string `protobuf:"bytes,22,opt,name=trusted_account_id,json=trustedAccountId,proto3" json:"trusted_account_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TrustedSecretQuery) Reset() {
	*x = TrustedSecretQuery{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustedSecretQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedSecretQuery) ProtoMessage() {}

func (x *TrustedSecretQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedSecretQuery.ProtoReflect.Descriptor instead.
func (*TrustedSecretQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{5}
}

func (x *TrustedSecretQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *TrustedSecretQuery) GetTrustedSecretId() string {
	if x != nil {
		return x.TrustedSecretId
	}
	return ""
}

func (x *TrustedSecretQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrustedSecretQuery) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *TrustedSecretQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *TrustedSecretQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *TrustedSecretQuery) GetTrustedAccountId() string {
	if x != nil {
		return x.TrustedAccountId
	}
	return ""
}

type TrustedSecretDataInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Encrypted      bool                   `protobuf:"varint,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	EncryptOptions *_struct.Struct        `protobuf:"bytes,2,opt,name=encrypt_options,json=encryptOptions,proto3" json:"encrypt_options,omitempty"`
	Data           *_struct.Struct        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TrustedSecretDataInfo) Reset() {
	*x = TrustedSecretDataInfo{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustedSecretDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedSecretDataInfo) ProtoMessage() {}

func (x *TrustedSecretDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedSecretDataInfo.ProtoReflect.Descriptor instead.
func (*TrustedSecretDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{6}
}

func (x *TrustedSecretDataInfo) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

func (x *TrustedSecretDataInfo) GetEncryptOptions() *_struct.Struct {
	if x != nil {
		return x.EncryptOptions
	}
	return nil
}

func (x *TrustedSecretDataInfo) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

//	{
//	   "trusted_secret_id": "trusted-secret-123456789012",
//	   "name": "aws-dev",
//	   "tags": {},
//	   "schema": "aws_access_key",
//	   "provider": "aws",
//	   "resource_group": "DOMAIN",
//	   "trusted_account_id": "ta-123456789012",
//	   "domain_id": "domain-123456789012",
//	   "created_at": "2022-01-01T06:10:14.851Z"
//	}
type TrustedSecretInfo struct {
	state            protoimpl.MessageState          `protogen:"open.v1"`
	TrustedSecretId  string                          `protobuf:"bytes,1,opt,name=trusted_secret_id,json=trustedSecretId,proto3" json:"trusted_secret_id,omitempty"`
	Name             string                          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SchemaId         string                          `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Tags             *_struct.Struct                 `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Provider         string                          `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	ResourceGroup    TrustedSecretInfo_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.secret.v1.TrustedSecretInfo_ResourceGroup" json:"resource_group,omitempty"`
	DomainId         string                          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string                          `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	TrustedAccountId string                          `protobuf:"bytes,23,opt,name=trusted_account_id,json=trustedAccountId,proto3" json:"trusted_account_id,omitempty"`
	CreatedAt        string                          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TrustedSecretInfo) Reset() {
	*x = TrustedSecretInfo{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustedSecretInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedSecretInfo) ProtoMessage() {}

func (x *TrustedSecretInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedSecretInfo.ProtoReflect.Descriptor instead.
func (*TrustedSecretInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{7}
}

func (x *TrustedSecretInfo) GetTrustedSecretId() string {
	if x != nil {
		return x.TrustedSecretId
	}
	return ""
}

func (x *TrustedSecretInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrustedSecretInfo) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *TrustedSecretInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TrustedSecretInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *TrustedSecretInfo) GetResourceGroup() TrustedSecretInfo_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return TrustedSecretInfo_RESOURCE_GROUP_NONE
}

func (x *TrustedSecretInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *TrustedSecretInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *TrustedSecretInfo) GetTrustedAccountId() string {
	if x != nil {
		return x.TrustedAccountId
	}
	return ""
}

func (x *TrustedSecretInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "trusted_secret_id": "trusted-secret-12345abcde",
//	           "name": "Cloudforet Broker Account - Managed",
//	           "schema_id": "aws-secret-access-key",
//	           "tags": {"foo": "bar"},
//	           "provider": "aws",
//	           "resource_group": "DOMAIN",
//	           "trusted_account_id": "ta-12345abcde",
//	           "domain_id": "domain-12345abcde",
//	           "created_at": "2022-01-01T06:10:14Z"
//	       },
//	       {
//	           "trusted_secret_id": "trusted-secret-56789abcde",
//	           "name": "Customer Broker Account",
//	           "schema_id": "aws-secret-access-key",
//	           "provider": "aws",
//	           "resource_group": "WORKSPACE",
//	           "trusted_account_id": "ta-56789abcde",
//	           "domain_id": "domain-12345abcde",
//	           "workspace_id": "workspace-12345abcde",
//	           "created_at": "2023-11-04T00:00:00Z"
//	       }
//	   ],
//	   "total_count": 2
//	}
type TrustedSecretsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*TrustedSecretInfo   `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrustedSecretsInfo) Reset() {
	*x = TrustedSecretsInfo{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustedSecretsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedSecretsInfo) ProtoMessage() {}

func (x *TrustedSecretsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedSecretsInfo.ProtoReflect.Descriptor instead.
func (*TrustedSecretsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{8}
}

func (x *TrustedSecretsInfo) GetResults() []*TrustedSecretInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *TrustedSecretsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type TrustedSecretStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId      string                 `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrustedSecretStatQuery) Reset() {
	*x = TrustedSecretStatQuery{}
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustedSecretStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustedSecretStatQuery) ProtoMessage() {}

func (x *TrustedSecretStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustedSecretStatQuery.ProtoReflect.Descriptor instead.
func (*TrustedSecretStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP(), []int{9}
}

func (x *TrustedSecretStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *TrustedSecretStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_secret_v1_trusted_secret_proto protoreflect.FileDescriptor

const file_spaceone_api_secret_v1_trusted_secret_proto_rawDesc = "" +
	"\n" +
	"+spaceone/api/secret/v1/trusted_secret.proto\x12\x16spaceone.api.secret.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xa6\x03\n" +
	"\x1aCreateTrustedSecretRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12+\n" +
	"\x04data\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04data\x12\x1b\n" +
	"\tschema_id\x18\x03 \x01(\tR\bschemaId\x12+\n" +
	"\x04tags\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12g\n" +
	"\x0eresource_group\x18\x14 \x01(\x0e2@.spaceone.api.secret.v1.CreateTrustedSecretRequest.ResourceGroupR\rresourceGroup\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12,\n" +
	"\x12trusted_account_id\x18\x16 \x01(\tR\x10trustedAccountId\"C\n" +
	"\rResourceGroup\x12\x17\n" +
	"\x13RESOURCE_GROUP_NONE\x10\x00\x12\n" +
	"\n" +
	"\x06DOMAIN\x10\x01\x12\r\n" +
	"\tWORKSPACE\x10\x02\"\x89\x01\n" +
	"\x1aUpdateTrustedSecretRequest\x12*\n" +
	"\x11trusted_secret_id\x18\x01 \x01(\tR\x0ftrustedSecretId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12+\n" +
	"\x04tags\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04tags\"B\n" +
	"\x14TrustedSecretRequest\x12*\n" +
	"\x11trusted_secret_id\x18\x01 \x01(\tR\x0ftrustedSecretId\"f\n" +
	"\x1bGetTrustedSecretDataRequest\x12*\n" +
	"\x11trusted_secret_id\x18\x01 \x01(\tR\x0ftrustedSecretId\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\"\x96\x01\n" +
	"\x1eUpdateTrustedSecretDataRequest\x12*\n" +
	"\x11trusted_secret_id\x18\x01 \x01(\tR\x0ftrustedSecretId\x12\x1b\n" +
	"\tschema_id\x18\x02 \x01(\tR\bschemaId\x12+\n" +
	"\x04data\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04data\"\x91\x02\n" +
	"\x12TrustedSecretQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12*\n" +
	"\x11trusted_secret_id\x18\x02 \x01(\tR\x0ftrustedSecretId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1b\n" +
	"\tschema_id\x18\x04 \x01(\tR\bschemaId\x12\x1a\n" +
	"\bprovider\x18\x05 \x01(\tR\bprovider\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12,\n" +
	"\x12trusted_account_id\x18\x16 \x01(\tR\x10trustedAccountId\"\xa4\x01\n" +
	"\x15TrustedSecretDataInfo\x12\x1c\n" +
	"\tencrypted\x18\x01 \x01(\bR\tencrypted\x12@\n" +
	"\x0fencrypt_options\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x0eencryptOptions\x12+\n" +
	"\x04data\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04data\"\xeb\x03\n" +
	"\x11TrustedSecretInfo\x12*\n" +
	"\x11trusted_secret_id\x18\x01 \x01(\tR\x0ftrustedSecretId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1b\n" +
	"\tschema_id\x18\x03 \x01(\tR\bschemaId\x12+\n" +
	"\x04tags\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1a\n" +
	"\bprovider\x18\x05 \x01(\tR\bprovider\x12^\n" +
	"\x0eresource_group\x18\x14 \x01(\x0e27.spaceone.api.secret.v1.TrustedSecretInfo.ResourceGroupR\rresourceGroup\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12,\n" +
	"\x12trusted_account_id\x18\x17 \x01(\tR\x10trustedAccountId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\"C\n" +
	"\rResourceGroup\x12\x17\n" +
	"\x13RESOURCE_GROUP_NONE\x10\x00\x12\n" +
	"\n" +
	"\x06DOMAIN\x10\x01\x12\r\n" +
	"\tWORKSPACE\x10\x02\"z\n" +
	"\x12TrustedSecretsInfo\x12C\n" +
	"\aresults\x18\x01 \x03(\v2).spaceone.api.secret.v1.TrustedSecretInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"r\n" +
	"\x16TrustedSecretStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query\x12\x1b\n" +
	"\tdomain_id\x18\x02 \x01(\tR\bdomainId2\xd1\b\n" +
	"\rTrustedSecret\x12\x94\x01\n" +
	"\x06create\x122.spaceone.api.secret.v1.CreateTrustedSecretRequest\x1a).spaceone.api.secret.v1.TrustedSecretInfo\"+\x82\xd3\xe4\x93\x02%:\x01*\" /secret/v1/trusted-secret/create\x12\x94\x01\n" +
	"\x06update\x122.spaceone.api.secret.v1.UpdateTrustedSecretRequest\x1a).spaceone.api.secret.v1.TrustedSecretInfo\"+\x82\xd3\xe4\x93\x02%:\x01*\" /secret/v1/trusted-secret/update\x12{\n" +
	"\x06delete\x12,.spaceone.api.secret.v1.TrustedSecretRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%:\x01*\" /secret/v1/trusted-secret/delete\x12\x8f\x01\n" +
	"\vupdate_data\x126.spaceone.api.secret.v1.UpdateTrustedSecretDataRequest\x1a\x16.google.protobuf.Empty\"0\x82\xd3\xe4\x93\x02*:\x01*\"%/secret/v1/trusted-secret/update-data\x12p\n" +
	"\bget_data\x123.spaceone.api.secret.v1.GetTrustedSecretDataRequest\x1a-.spaceone.api.secret.v1.TrustedSecretDataInfo\"\x00\x12\x88\x01\n" +
	"\x03get\x12,.spaceone.api.secret.v1.TrustedSecretRequest\x1a).spaceone.api.secret.v1.TrustedSecretInfo\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/secret/v1/trusted-secret/get\x12\x89\x01\n" +
	"\x04list\x12*.spaceone.api.secret.v1.TrustedSecretQuery\x1a*.spaceone.api.secret.v1.TrustedSecretsInfo\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/secret/v1/trusted-secret/list\x12z\n" +
	"\x04stat\x12..spaceone.api.secret.v1.TrustedSecretStatQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/secret/v1/trusted-secret/statB=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/secret/v1b\x06proto3"

var (
	file_spaceone_api_secret_v1_trusted_secret_proto_rawDescOnce sync.Once
	file_spaceone_api_secret_v1_trusted_secret_proto_rawDescData []byte
)

func file_spaceone_api_secret_v1_trusted_secret_proto_rawDescGZIP() []byte {
	file_spaceone_api_secret_v1_trusted_secret_proto_rawDescOnce.Do(func() {
		file_spaceone_api_secret_v1_trusted_secret_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_secret_v1_trusted_secret_proto_rawDesc), len(file_spaceone_api_secret_v1_trusted_secret_proto_rawDesc)))
	})
	return file_spaceone_api_secret_v1_trusted_secret_proto_rawDescData
}

var file_spaceone_api_secret_v1_trusted_secret_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_spaceone_api_secret_v1_trusted_secret_proto_goTypes = []any{
	(CreateTrustedSecretRequest_ResourceGroup)(0), // 0: spaceone.api.secret.v1.CreateTrustedSecretRequest.ResourceGroup
	(TrustedSecretInfo_ResourceGroup)(0),          // 1: spaceone.api.secret.v1.TrustedSecretInfo.ResourceGroup
	(*CreateTrustedSecretRequest)(nil),            // 2: spaceone.api.secret.v1.CreateTrustedSecretRequest
	(*UpdateTrustedSecretRequest)(nil),            // 3: spaceone.api.secret.v1.UpdateTrustedSecretRequest
	(*TrustedSecretRequest)(nil),                  // 4: spaceone.api.secret.v1.TrustedSecretRequest
	(*GetTrustedSecretDataRequest)(nil),           // 5: spaceone.api.secret.v1.GetTrustedSecretDataRequest
	(*UpdateTrustedSecretDataRequest)(nil),        // 6: spaceone.api.secret.v1.UpdateTrustedSecretDataRequest
	(*TrustedSecretQuery)(nil),                    // 7: spaceone.api.secret.v1.TrustedSecretQuery
	(*TrustedSecretDataInfo)(nil),                 // 8: spaceone.api.secret.v1.TrustedSecretDataInfo
	(*TrustedSecretInfo)(nil),                     // 9: spaceone.api.secret.v1.TrustedSecretInfo
	(*TrustedSecretsInfo)(nil),                    // 10: spaceone.api.secret.v1.TrustedSecretsInfo
	(*TrustedSecretStatQuery)(nil),                // 11: spaceone.api.secret.v1.TrustedSecretStatQuery
	(*_struct.Struct)(nil),                        // 12: google.protobuf.Struct
	(*v2.Query)(nil),                              // 13: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),                    // 14: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                           // 15: google.protobuf.Empty
}
var file_spaceone_api_secret_v1_trusted_secret_proto_depIdxs = []int32{
	12, // 0: spaceone.api.secret.v1.CreateTrustedSecretRequest.data:type_name -> google.protobuf.Struct
	12, // 1: spaceone.api.secret.v1.CreateTrustedSecretRequest.tags:type_name -> google.protobuf.Struct
	0,  // 2: spaceone.api.secret.v1.CreateTrustedSecretRequest.resource_group:type_name -> spaceone.api.secret.v1.CreateTrustedSecretRequest.ResourceGroup
	12, // 3: spaceone.api.secret.v1.UpdateTrustedSecretRequest.tags:type_name -> google.protobuf.Struct
	12, // 4: spaceone.api.secret.v1.UpdateTrustedSecretDataRequest.data:type_name -> google.protobuf.Struct
	13, // 5: spaceone.api.secret.v1.TrustedSecretQuery.query:type_name -> spaceone.api.core.v2.Query
	12, // 6: spaceone.api.secret.v1.TrustedSecretDataInfo.encrypt_options:type_name -> google.protobuf.Struct
	12, // 7: spaceone.api.secret.v1.TrustedSecretDataInfo.data:type_name -> google.protobuf.Struct
	12, // 8: spaceone.api.secret.v1.TrustedSecretInfo.tags:type_name -> google.protobuf.Struct
	1,  // 9: spaceone.api.secret.v1.TrustedSecretInfo.resource_group:type_name -> spaceone.api.secret.v1.TrustedSecretInfo.ResourceGroup
	9,  // 10: spaceone.api.secret.v1.TrustedSecretsInfo.results:type_name -> spaceone.api.secret.v1.TrustedSecretInfo
	14, // 11: spaceone.api.secret.v1.TrustedSecretStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	2,  // 12: spaceone.api.secret.v1.TrustedSecret.create:input_type -> spaceone.api.secret.v1.CreateTrustedSecretRequest
	3,  // 13: spaceone.api.secret.v1.TrustedSecret.update:input_type -> spaceone.api.secret.v1.UpdateTrustedSecretRequest
	4,  // 14: spaceone.api.secret.v1.TrustedSecret.delete:input_type -> spaceone.api.secret.v1.TrustedSecretRequest
	6,  // 15: spaceone.api.secret.v1.TrustedSecret.update_data:input_type -> spaceone.api.secret.v1.UpdateTrustedSecretDataRequest
	5,  // 16: spaceone.api.secret.v1.TrustedSecret.get_data:input_type -> spaceone.api.secret.v1.GetTrustedSecretDataRequest
	4,  // 17: spaceone.api.secret.v1.TrustedSecret.get:input_type -> spaceone.api.secret.v1.TrustedSecretRequest
	7,  // 18: spaceone.api.secret.v1.TrustedSecret.list:input_type -> spaceone.api.secret.v1.TrustedSecretQuery
	11, // 19: spaceone.api.secret.v1.TrustedSecret.stat:input_type -> spaceone.api.secret.v1.TrustedSecretStatQuery
	9,  // 20: spaceone.api.secret.v1.TrustedSecret.create:output_type -> spaceone.api.secret.v1.TrustedSecretInfo
	9,  // 21: spaceone.api.secret.v1.TrustedSecret.update:output_type -> spaceone.api.secret.v1.TrustedSecretInfo
	15, // 22: spaceone.api.secret.v1.TrustedSecret.delete:output_type -> google.protobuf.Empty
	15, // 23: spaceone.api.secret.v1.TrustedSecret.update_data:output_type -> google.protobuf.Empty
	8,  // 24: spaceone.api.secret.v1.TrustedSecret.get_data:output_type -> spaceone.api.secret.v1.TrustedSecretDataInfo
	9,  // 25: spaceone.api.secret.v1.TrustedSecret.get:output_type -> spaceone.api.secret.v1.TrustedSecretInfo
	10, // 26: spaceone.api.secret.v1.TrustedSecret.list:output_type -> spaceone.api.secret.v1.TrustedSecretsInfo
	12, // 27: spaceone.api.secret.v1.TrustedSecret.stat:output_type -> google.protobuf.Struct
	20, // [20:28] is the sub-list for method output_type
	12, // [12:20] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_spaceone_api_secret_v1_trusted_secret_proto_init() }
func file_spaceone_api_secret_v1_trusted_secret_proto_init() {
	if File_spaceone_api_secret_v1_trusted_secret_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_secret_v1_trusted_secret_proto_rawDesc), len(file_spaceone_api_secret_v1_trusted_secret_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_secret_v1_trusted_secret_proto_goTypes,
		DependencyIndexes: file_spaceone_api_secret_v1_trusted_secret_proto_depIdxs,
		EnumInfos:         file_spaceone_api_secret_v1_trusted_secret_proto_enumTypes,
		MessageInfos:      file_spaceone_api_secret_v1_trusted_secret_proto_msgTypes,
	}.Build()
	File_spaceone_api_secret_v1_trusted_secret_proto = out.File
	file_spaceone_api_secret_v1_trusted_secret_proto_goTypes = nil
	file_spaceone_api_secret_v1_trusted_secret_proto_depIdxs = nil
}
