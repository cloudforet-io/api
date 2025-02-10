// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/inventory_v2/v1/namespace_group.proto

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

// {
//
// }
type CreateNamespaceGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	NamespaceGroupId string `protobuf:"bytes,1,opt,name=namespace_group_id,json=namespaceGroupId,proto3" json:"namespace_group_id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon             string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	// +optional
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	ResourceGroup string          `protobuf:"bytes,20,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNamespaceGroupRequest) Reset() {
	*x = CreateNamespaceGroupRequest{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNamespaceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceGroupRequest) ProtoMessage() {}

func (x *CreateNamespaceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNamespaceGroupRequest) GetNamespaceGroupId() string {
	if x != nil {
		return x.NamespaceGroupId
	}
	return ""
}

func (x *CreateNamespaceGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceGroupRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CreateNamespaceGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateNamespaceGroupRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateNamespaceGroupRequest) GetResourceGroup() string {
	if x != nil {
		return x.ResourceGroup
	}
	return ""
}

func (x *CreateNamespaceGroupRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

// {
//
// }
type UpdateNamespaceGroupRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	NamespaceGroupId string                 `protobuf:"bytes,1,opt,name=namespace_group_id,json=namespaceGroupId,proto3" json:"namespace_group_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	// +optional
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNamespaceGroupRequest) Reset() {
	*x = UpdateNamespaceGroupRequest{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNamespaceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceGroupRequest) ProtoMessage() {}

func (x *UpdateNamespaceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNamespaceGroupRequest) GetNamespaceGroupId() string {
	if x != nil {
		return x.NamespaceGroupId
	}
	return ""
}

func (x *UpdateNamespaceGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNamespaceGroupRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UpdateNamespaceGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateNamespaceGroupRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

// {
//
// }
type NamespaceGroupRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	NamespaceGroupId string                 `protobuf:"bytes,1,opt,name=namespace_group_id,json=namespaceGroupId,proto3" json:"namespace_group_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *NamespaceGroupRequest) Reset() {
	*x = NamespaceGroupRequest{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceGroupRequest) ProtoMessage() {}

func (x *NamespaceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceGroupRequest.ProtoReflect.Descriptor instead.
func (*NamespaceGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceGroupRequest) GetNamespaceGroupId() string {
	if x != nil {
		return x.NamespaceGroupId
	}
	return ""
}

// {
//
// }
type NamespaceGroupQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	NamespaceGroupId string `protobuf:"bytes,2,opt,name=namespace_group_id,json=namespaceGroupId,proto3" json:"namespace_group_id,omitempty"`
	// +optional
	ExistOnly bool `protobuf:"varint,3,opt,name=exist_only,json=existOnly,proto3" json:"exist_only,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamespaceGroupQuery) Reset() {
	*x = NamespaceGroupQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceGroupQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceGroupQuery) ProtoMessage() {}

func (x *NamespaceGroupQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceGroupQuery.ProtoReflect.Descriptor instead.
func (*NamespaceGroupQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{3}
}

func (x *NamespaceGroupQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *NamespaceGroupQuery) GetNamespaceGroupId() string {
	if x != nil {
		return x.NamespaceGroupId
	}
	return ""
}

func (x *NamespaceGroupQuery) GetExistOnly() bool {
	if x != nil {
		return x.ExistOnly
	}
	return false
}

func (x *NamespaceGroupQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

// {
//
// }
type NamespaceGroupInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	NamespaceGroupId string                 `protobuf:"bytes,1,opt,name=namespace_group_id,json=namespaceGroupId,proto3" json:"namespace_group_id,omitempty"`
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon             string                 `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	// +optional
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	IsManaged     bool            `protobuf:"varint,6,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	ResourceGroup string          `protobuf:"bytes,20,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	DomainId      string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string          `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	CreatedAt     string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string          `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamespaceGroupInfo) Reset() {
	*x = NamespaceGroupInfo{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceGroupInfo) ProtoMessage() {}

func (x *NamespaceGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceGroupInfo.ProtoReflect.Descriptor instead.
func (*NamespaceGroupInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{4}
}

func (x *NamespaceGroupInfo) GetNamespaceGroupId() string {
	if x != nil {
		return x.NamespaceGroupId
	}
	return ""
}

func (x *NamespaceGroupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamespaceGroupInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *NamespaceGroupInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NamespaceGroupInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *NamespaceGroupInfo) GetIsManaged() bool {
	if x != nil {
		return x.IsManaged
	}
	return false
}

func (x *NamespaceGroupInfo) GetResourceGroup() string {
	if x != nil {
		return x.ResourceGroup
	}
	return ""
}

func (x *NamespaceGroupInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *NamespaceGroupInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *NamespaceGroupInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NamespaceGroupInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type NamespaceGroupsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*NamespaceGroupInfo  `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamespaceGroupsInfo) Reset() {
	*x = NamespaceGroupsInfo{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceGroupsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceGroupsInfo) ProtoMessage() {}

func (x *NamespaceGroupsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceGroupsInfo.ProtoReflect.Descriptor instead.
func (*NamespaceGroupsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{5}
}

func (x *NamespaceGroupsInfo) GetResults() []*NamespaceGroupInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *NamespaceGroupsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type NamespaceGroupStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamespaceGroupStatQuery) Reset() {
	*x = NamespaceGroupStatQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceGroupStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceGroupStatQuery) ProtoMessage() {}

func (x *NamespaceGroupStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceGroupStatQuery.ProtoReflect.Descriptor instead.
func (*NamespaceGroupStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP(), []int{6}
}

func (x *NamespaceGroupStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_inventory_v2_v1_namespace_group_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDesc = string([]byte{
	0x0a, 0x32, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a,
	0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x45, 0x0a, 0x15, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x78, 0x69, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x22, 0xfd, 0x02, 0x0a, 0x12, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x17, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x32, 0xc0, 0x07, 0x0a, 0x0e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0xa9, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x39,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0xa9, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x39, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76,
	0x32, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a,
	0x01, 0x2a, 0x22, 0x27, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76,
	0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x06,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12,
	0x33, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01,
	0x2a, 0x22, 0x24, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01,
	0x2a, 0x22, 0x25, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x88, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61,
	0x74, 0x12, 0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescData []byte
)

func file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDesc), len(file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDesc)))
	})
	return file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDescData
}

var file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_inventory_v2_v1_namespace_group_proto_goTypes = []any{
	(*CreateNamespaceGroupRequest)(nil), // 0: spaceone.api.inventory_v2.v1.CreateNamespaceGroupRequest
	(*UpdateNamespaceGroupRequest)(nil), // 1: spaceone.api.inventory_v2.v1.UpdateNamespaceGroupRequest
	(*NamespaceGroupRequest)(nil),       // 2: spaceone.api.inventory_v2.v1.NamespaceGroupRequest
	(*NamespaceGroupQuery)(nil),         // 3: spaceone.api.inventory_v2.v1.NamespaceGroupQuery
	(*NamespaceGroupInfo)(nil),          // 4: spaceone.api.inventory_v2.v1.NamespaceGroupInfo
	(*NamespaceGroupsInfo)(nil),         // 5: spaceone.api.inventory_v2.v1.NamespaceGroupsInfo
	(*NamespaceGroupStatQuery)(nil),     // 6: spaceone.api.inventory_v2.v1.NamespaceGroupStatQuery
	(*_struct.Struct)(nil),              // 7: google.protobuf.Struct
	(*v2.Query)(nil),                    // 8: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),          // 9: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_spaceone_api_inventory_v2_v1_namespace_group_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.inventory_v2.v1.CreateNamespaceGroupRequest.tags:type_name -> google.protobuf.Struct
	7,  // 1: spaceone.api.inventory_v2.v1.UpdateNamespaceGroupRequest.tags:type_name -> google.protobuf.Struct
	8,  // 2: spaceone.api.inventory_v2.v1.NamespaceGroupQuery.query:type_name -> spaceone.api.core.v2.Query
	7,  // 3: spaceone.api.inventory_v2.v1.NamespaceGroupInfo.tags:type_name -> google.protobuf.Struct
	4,  // 4: spaceone.api.inventory_v2.v1.NamespaceGroupsInfo.results:type_name -> spaceone.api.inventory_v2.v1.NamespaceGroupInfo
	9,  // 5: spaceone.api.inventory_v2.v1.NamespaceGroupStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 6: spaceone.api.inventory_v2.v1.NamespaceGroup.create:input_type -> spaceone.api.inventory_v2.v1.CreateNamespaceGroupRequest
	1,  // 7: spaceone.api.inventory_v2.v1.NamespaceGroup.update:input_type -> spaceone.api.inventory_v2.v1.UpdateNamespaceGroupRequest
	2,  // 8: spaceone.api.inventory_v2.v1.NamespaceGroup.delete:input_type -> spaceone.api.inventory_v2.v1.NamespaceGroupRequest
	2,  // 9: spaceone.api.inventory_v2.v1.NamespaceGroup.get:input_type -> spaceone.api.inventory_v2.v1.NamespaceGroupRequest
	3,  // 10: spaceone.api.inventory_v2.v1.NamespaceGroup.list:input_type -> spaceone.api.inventory_v2.v1.NamespaceGroupQuery
	6,  // 11: spaceone.api.inventory_v2.v1.NamespaceGroup.stat:input_type -> spaceone.api.inventory_v2.v1.NamespaceGroupStatQuery
	4,  // 12: spaceone.api.inventory_v2.v1.NamespaceGroup.create:output_type -> spaceone.api.inventory_v2.v1.NamespaceGroupInfo
	4,  // 13: spaceone.api.inventory_v2.v1.NamespaceGroup.update:output_type -> spaceone.api.inventory_v2.v1.NamespaceGroupInfo
	10, // 14: spaceone.api.inventory_v2.v1.NamespaceGroup.delete:output_type -> google.protobuf.Empty
	4,  // 15: spaceone.api.inventory_v2.v1.NamespaceGroup.get:output_type -> spaceone.api.inventory_v2.v1.NamespaceGroupInfo
	5,  // 16: spaceone.api.inventory_v2.v1.NamespaceGroup.list:output_type -> spaceone.api.inventory_v2.v1.NamespaceGroupsInfo
	7,  // 17: spaceone.api.inventory_v2.v1.NamespaceGroup.stat:output_type -> google.protobuf.Struct
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v2_v1_namespace_group_proto_init() }
func file_spaceone_api_inventory_v2_v1_namespace_group_proto_init() {
	if File_spaceone_api_inventory_v2_v1_namespace_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDesc), len(file_spaceone_api_inventory_v2_v1_namespace_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v2_v1_namespace_group_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v2_v1_namespace_group_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_v2_v1_namespace_group_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v2_v1_namespace_group_proto = out.File
	file_spaceone_api_inventory_v2_v1_namespace_group_proto_goTypes = nil
	file_spaceone_api_inventory_v2_v1_namespace_group_proto_depIdxs = nil
}
