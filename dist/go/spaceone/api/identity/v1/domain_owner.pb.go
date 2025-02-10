// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/identity/v1/domain_owner.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type CreateDomainOwner struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	OwnerId  string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// +optional
	Language string `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// +optional
	Timezone      string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	DomainId      string `protobuf:"bytes,10,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDomainOwner) Reset() {
	*x = CreateDomainOwner{}
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDomainOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDomainOwner) ProtoMessage() {}

func (x *CreateDomainOwner) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDomainOwner.ProtoReflect.Descriptor instead.
func (*CreateDomainOwner) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_domain_owner_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDomainOwner) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateDomainOwner) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateDomainOwner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDomainOwner) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateDomainOwner) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CreateDomainOwner) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *CreateDomainOwner) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type UpdateDomainOwner struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	OwnerId string                 `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// +optional
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// +optional
	Language string `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// +optional
	Timezone      string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	DomainId      string `protobuf:"bytes,10,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDomainOwner) Reset() {
	*x = UpdateDomainOwner{}
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDomainOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainOwner) ProtoMessage() {}

func (x *UpdateDomainOwner) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainOwner.ProtoReflect.Descriptor instead.
func (*UpdateDomainOwner) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_domain_owner_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDomainOwner) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *UpdateDomainOwner) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateDomainOwner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDomainOwner) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateDomainOwner) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *UpdateDomainOwner) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *UpdateDomainOwner) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type DomainOwnerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DomainId      string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	OwnerId       string                 `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DomainOwnerRequest) Reset() {
	*x = DomainOwnerRequest{}
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DomainOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainOwnerRequest) ProtoMessage() {}

func (x *DomainOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainOwnerRequest.ProtoReflect.Descriptor instead.
func (*DomainOwnerRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_domain_owner_proto_rawDescGZIP(), []int{2}
}

func (x *DomainOwnerRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *DomainOwnerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetDomainOwnerRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	DomainId string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	OwnerId  string                 `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// +optional
	Only          []string `protobuf:"bytes,3,rep,name=only,proto3" json:"only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDomainOwnerRequest) Reset() {
	*x = GetDomainOwnerRequest{}
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDomainOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainOwnerRequest) ProtoMessage() {}

func (x *GetDomainOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainOwnerRequest.ProtoReflect.Descriptor instead.
func (*GetDomainOwnerRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_domain_owner_proto_rawDescGZIP(), []int{3}
}

func (x *GetDomainOwnerRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GetDomainOwnerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *GetDomainOwnerRequest) GetOnly() []string {
	if x != nil {
		return x.Only
	}
	return nil
}

type DomainOwnerInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OwnerId        string                 `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email          string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Language       string                 `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Timezone       string                 `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	LastAccessedAt string                 `protobuf:"bytes,11,opt,name=last_accessed_at,json=lastAccessedAt,proto3" json:"last_accessed_at,omitempty"`
	CreatedAt      string                 `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DomainId       string                 `protobuf:"bytes,13,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DomainOwnerInfo) Reset() {
	*x = DomainOwnerInfo{}
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DomainOwnerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainOwnerInfo) ProtoMessage() {}

func (x *DomainOwnerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_domain_owner_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainOwnerInfo.ProtoReflect.Descriptor instead.
func (*DomainOwnerInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_domain_owner_proto_rawDescGZIP(), []int{4}
}

func (x *DomainOwnerInfo) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *DomainOwnerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DomainOwnerInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DomainOwnerInfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *DomainOwnerInfo) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *DomainOwnerInfo) GetLastAccessedAt() string {
	if x != nil {
		return x.LastAccessedAt
	}
	return ""
}

func (x *DomainOwnerInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DomainOwnerInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_identity_v1_domain_owner_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v1_domain_owner_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xc9,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x6e, 0x6c, 0x79, 0x22, 0xf4, 0x01,
	0x0a, 0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x32, 0xe1, 0x03, 0x0a, 0x0b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x1a, 0x29, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x8d, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x1a, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2d, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x03, 0x67,
	0x65, 0x74, 0x12, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65,
	0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_spaceone_api_identity_v1_domain_owner_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v1_domain_owner_proto_rawDescData []byte
)

func file_spaceone_api_identity_v1_domain_owner_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v1_domain_owner_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v1_domain_owner_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v1_domain_owner_proto_rawDesc), len(file_spaceone_api_identity_v1_domain_owner_proto_rawDesc)))
	})
	return file_spaceone_api_identity_v1_domain_owner_proto_rawDescData
}

var file_spaceone_api_identity_v1_domain_owner_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_identity_v1_domain_owner_proto_goTypes = []any{
	(*CreateDomainOwner)(nil),     // 0: spaceone.api.identity.v1.CreateDomainOwner
	(*UpdateDomainOwner)(nil),     // 1: spaceone.api.identity.v1.UpdateDomainOwner
	(*DomainOwnerRequest)(nil),    // 2: spaceone.api.identity.v1.DomainOwnerRequest
	(*GetDomainOwnerRequest)(nil), // 3: spaceone.api.identity.v1.GetDomainOwnerRequest
	(*DomainOwnerInfo)(nil),       // 4: spaceone.api.identity.v1.DomainOwnerInfo
	(*empty.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_spaceone_api_identity_v1_domain_owner_proto_depIdxs = []int32{
	0, // 0: spaceone.api.identity.v1.DomainOwner.create:input_type -> spaceone.api.identity.v1.CreateDomainOwner
	1, // 1: spaceone.api.identity.v1.DomainOwner.update:input_type -> spaceone.api.identity.v1.UpdateDomainOwner
	2, // 2: spaceone.api.identity.v1.DomainOwner.delete:input_type -> spaceone.api.identity.v1.DomainOwnerRequest
	3, // 3: spaceone.api.identity.v1.DomainOwner.get:input_type -> spaceone.api.identity.v1.GetDomainOwnerRequest
	4, // 4: spaceone.api.identity.v1.DomainOwner.create:output_type -> spaceone.api.identity.v1.DomainOwnerInfo
	4, // 5: spaceone.api.identity.v1.DomainOwner.update:output_type -> spaceone.api.identity.v1.DomainOwnerInfo
	5, // 6: spaceone.api.identity.v1.DomainOwner.delete:output_type -> google.protobuf.Empty
	4, // 7: spaceone.api.identity.v1.DomainOwner.get:output_type -> spaceone.api.identity.v1.DomainOwnerInfo
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v1_domain_owner_proto_init() }
func file_spaceone_api_identity_v1_domain_owner_proto_init() {
	if File_spaceone_api_identity_v1_domain_owner_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v1_domain_owner_proto_rawDesc), len(file_spaceone_api_identity_v1_domain_owner_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v1_domain_owner_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v1_domain_owner_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v1_domain_owner_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v1_domain_owner_proto = out.File
	file_spaceone_api_identity_v1_domain_owner_proto_goTypes = nil
	file_spaceone_api_identity_v1_domain_owner_proto_depIdxs = nil
}
