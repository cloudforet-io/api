// A Region is a resource storing regional information from each cloud service provider. Regional data stored by the resource includes the latitude and longitude of the region.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/inventory_v2/v1/region.proto

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

type CreateRegionRequest_ResourceGroup int32

const (
	CreateRegionRequest_RESOURCE_GROUP_NONE CreateRegionRequest_ResourceGroup = 0
	CreateRegionRequest_DOMAIN              CreateRegionRequest_ResourceGroup = 1
	CreateRegionRequest_WORKSPACE           CreateRegionRequest_ResourceGroup = 2
)

// Enum value maps for CreateRegionRequest_ResourceGroup.
var (
	CreateRegionRequest_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	CreateRegionRequest_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"DOMAIN":              1,
		"WORKSPACE":           2,
	}
)

func (x CreateRegionRequest_ResourceGroup) Enum() *CreateRegionRequest_ResourceGroup {
	p := new(CreateRegionRequest_ResourceGroup)
	*p = x
	return p
}

func (x CreateRegionRequest_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateRegionRequest_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_inventory_v2_v1_region_proto_enumTypes[0].Descriptor()
}

func (CreateRegionRequest_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_inventory_v2_v1_region_proto_enumTypes[0]
}

func (x CreateRegionRequest_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateRegionRequest_ResourceGroup.Descriptor instead.
func (CreateRegionRequest_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{0, 0}
}

type RegionInfo_ResourceGroup int32

const (
	RegionInfo_RESOURCE_GROUP_NONE RegionInfo_ResourceGroup = 0
	RegionInfo_DOMAIN              RegionInfo_ResourceGroup = 1
	RegionInfo_WORKSPACE           RegionInfo_ResourceGroup = 2
)

// Enum value maps for RegionInfo_ResourceGroup.
var (
	RegionInfo_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	RegionInfo_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"DOMAIN":              1,
		"WORKSPACE":           2,
	}
)

func (x RegionInfo_ResourceGroup) Enum() *RegionInfo_ResourceGroup {
	p := new(RegionInfo_ResourceGroup)
	*p = x
	return p
}

func (x RegionInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegionInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_inventory_v2_v1_region_proto_enumTypes[1].Descriptor()
}

func (RegionInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_inventory_v2_v1_region_proto_enumTypes[1]
}

func (x RegionInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegionInfo_ResourceGroup.Descriptor instead.
func (RegionInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{4, 0}
}

type CreateRegionRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RegionCode string                 `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Provider   string                 `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	Tags          *_struct.Struct                   `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	ResourceGroup CreateRegionRequest_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.inventory_v2.v1.CreateRegionRequest_ResourceGroup" json:"resource_group,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRegionRequest) Reset() {
	*x = CreateRegionRequest{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegionRequest) ProtoMessage() {}

func (x *CreateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegionRequest.ProtoReflect.Descriptor instead.
func (*CreateRegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRegionRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CreateRegionRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateRegionRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateRegionRequest) GetResourceGroup() CreateRegionRequest_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return CreateRegionRequest_RESOURCE_GROUP_NONE
}

func (x *CreateRegionRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type UpdateRegionRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	RegionId string                 `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRegionRequest) Reset() {
	*x = UpdateRegionRequest{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegionRequest) ProtoMessage() {}

func (x *UpdateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegionRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRegionRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *UpdateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRegionRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type RegionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RegionId      string                 `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegionRequest) Reset() {
	*x = RegionRequest{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionRequest) ProtoMessage() {}

func (x *RegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionRequest.ProtoReflect.Descriptor instead.
func (*RegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{2}
}

func (x *RegionRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type RegionQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	RegionId string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	RegionCode string `protobuf:"bytes,4,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	ExistsOnly bool `protobuf:"varint,6,opt,name=exists_only,json=existsOnly,proto3" json:"exists_only,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegionQuery) Reset() {
	*x = RegionQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegionQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionQuery) ProtoMessage() {}

func (x *RegionQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionQuery.ProtoReflect.Descriptor instead.
func (*RegionQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{3}
}

func (x *RegionQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RegionQuery) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *RegionQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionQuery) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RegionQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *RegionQuery) GetExistsOnly() bool {
	if x != nil {
		return x.ExistsOnly
	}
	return false
}

func (x *RegionQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type RegionInfo struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	RegionId      string                   `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Name          string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RegionCode    string                   `protobuf:"bytes,3,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Provider      string                   `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	Tags          *_struct.Struct          `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	ResourceGroup RegionInfo_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.inventory_v2.v1.RegionInfo_ResourceGroup" json:"resource_group,omitempty"`
	DomainId      string                   `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                   `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	CreatedAt     string                   `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                   `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegionInfo) Reset() {
	*x = RegionInfo{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionInfo) ProtoMessage() {}

func (x *RegionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionInfo.ProtoReflect.Descriptor instead.
func (*RegionInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{4}
}

func (x *RegionInfo) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *RegionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionInfo) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RegionInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *RegionInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RegionInfo) GetResourceGroup() RegionInfo_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return RegionInfo_RESOURCE_GROUP_NONE
}

func (x *RegionInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *RegionInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *RegionInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RegionInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type RegionsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*RegionInfo          `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegionsInfo) Reset() {
	*x = RegionsInfo{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegionsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionsInfo) ProtoMessage() {}

func (x *RegionsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionsInfo.ProtoReflect.Descriptor instead.
func (*RegionsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{5}
}

func (x *RegionsInfo) GetResults() []*RegionInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *RegionsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type RegionStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegionStatQuery) Reset() {
	*x = RegionStatQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegionStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionStatQuery) ProtoMessage() {}

func (x *RegionStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_region_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionStatQuery.ProtoReflect.Descriptor instead.
func (*RegionStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP(), []int{6}
}

func (x *RegionStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_inventory_v2_v1_region_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_v2_v1_region_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x66, 0x0a, 0x0e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x57,
	0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x22, 0x73, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22,
	0x2c, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xf2, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x22, 0xc9, 0x03, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x5d, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x22, 0x72,
	0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x32, 0xb0, 0x06, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x90, 0x01,
	0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a,
	0x22, 0x1e, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x90, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d,
	0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x84, 0x01,
	0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76,
	0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x2f, 0x67, 0x65, 0x74, 0x12, 0x85, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x76, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x77, 0x0a, 0x04,
	0x73, 0x74, 0x61, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x32,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2d, 0x76, 0x32, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_spaceone_api_inventory_v2_v1_region_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v2_v1_region_proto_rawDescData []byte
)

func file_spaceone_api_inventory_v2_v1_region_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v2_v1_region_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v2_v1_region_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v2_v1_region_proto_rawDesc), len(file_spaceone_api_inventory_v2_v1_region_proto_rawDesc)))
	})
	return file_spaceone_api_inventory_v2_v1_region_proto_rawDescData
}

var file_spaceone_api_inventory_v2_v1_region_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_inventory_v2_v1_region_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_inventory_v2_v1_region_proto_goTypes = []any{
	(CreateRegionRequest_ResourceGroup)(0), // 0: spaceone.api.inventory_v2.v1.CreateRegionRequest.ResourceGroup
	(RegionInfo_ResourceGroup)(0),          // 1: spaceone.api.inventory_v2.v1.RegionInfo.ResourceGroup
	(*CreateRegionRequest)(nil),            // 2: spaceone.api.inventory_v2.v1.CreateRegionRequest
	(*UpdateRegionRequest)(nil),            // 3: spaceone.api.inventory_v2.v1.UpdateRegionRequest
	(*RegionRequest)(nil),                  // 4: spaceone.api.inventory_v2.v1.RegionRequest
	(*RegionQuery)(nil),                    // 5: spaceone.api.inventory_v2.v1.RegionQuery
	(*RegionInfo)(nil),                     // 6: spaceone.api.inventory_v2.v1.RegionInfo
	(*RegionsInfo)(nil),                    // 7: spaceone.api.inventory_v2.v1.RegionsInfo
	(*RegionStatQuery)(nil),                // 8: spaceone.api.inventory_v2.v1.RegionStatQuery
	(*_struct.Struct)(nil),                 // 9: google.protobuf.Struct
	(*v2.Query)(nil),                       // 10: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),             // 11: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                    // 12: google.protobuf.Empty
}
var file_spaceone_api_inventory_v2_v1_region_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.inventory_v2.v1.CreateRegionRequest.tags:type_name -> google.protobuf.Struct
	0,  // 1: spaceone.api.inventory_v2.v1.CreateRegionRequest.resource_group:type_name -> spaceone.api.inventory_v2.v1.CreateRegionRequest.ResourceGroup
	9,  // 2: spaceone.api.inventory_v2.v1.UpdateRegionRequest.tags:type_name -> google.protobuf.Struct
	10, // 3: spaceone.api.inventory_v2.v1.RegionQuery.query:type_name -> spaceone.api.core.v2.Query
	9,  // 4: spaceone.api.inventory_v2.v1.RegionInfo.tags:type_name -> google.protobuf.Struct
	1,  // 5: spaceone.api.inventory_v2.v1.RegionInfo.resource_group:type_name -> spaceone.api.inventory_v2.v1.RegionInfo.ResourceGroup
	6,  // 6: spaceone.api.inventory_v2.v1.RegionsInfo.results:type_name -> spaceone.api.inventory_v2.v1.RegionInfo
	11, // 7: spaceone.api.inventory_v2.v1.RegionStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	2,  // 8: spaceone.api.inventory_v2.v1.Region.create:input_type -> spaceone.api.inventory_v2.v1.CreateRegionRequest
	3,  // 9: spaceone.api.inventory_v2.v1.Region.update:input_type -> spaceone.api.inventory_v2.v1.UpdateRegionRequest
	4,  // 10: spaceone.api.inventory_v2.v1.Region.delete:input_type -> spaceone.api.inventory_v2.v1.RegionRequest
	4,  // 11: spaceone.api.inventory_v2.v1.Region.get:input_type -> spaceone.api.inventory_v2.v1.RegionRequest
	5,  // 12: spaceone.api.inventory_v2.v1.Region.list:input_type -> spaceone.api.inventory_v2.v1.RegionQuery
	8,  // 13: spaceone.api.inventory_v2.v1.Region.stat:input_type -> spaceone.api.inventory_v2.v1.RegionStatQuery
	6,  // 14: spaceone.api.inventory_v2.v1.Region.create:output_type -> spaceone.api.inventory_v2.v1.RegionInfo
	6,  // 15: spaceone.api.inventory_v2.v1.Region.update:output_type -> spaceone.api.inventory_v2.v1.RegionInfo
	12, // 16: spaceone.api.inventory_v2.v1.Region.delete:output_type -> google.protobuf.Empty
	6,  // 17: spaceone.api.inventory_v2.v1.Region.get:output_type -> spaceone.api.inventory_v2.v1.RegionInfo
	7,  // 18: spaceone.api.inventory_v2.v1.Region.list:output_type -> spaceone.api.inventory_v2.v1.RegionsInfo
	9,  // 19: spaceone.api.inventory_v2.v1.Region.stat:output_type -> google.protobuf.Struct
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v2_v1_region_proto_init() }
func file_spaceone_api_inventory_v2_v1_region_proto_init() {
	if File_spaceone_api_inventory_v2_v1_region_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v2_v1_region_proto_rawDesc), len(file_spaceone_api_inventory_v2_v1_region_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v2_v1_region_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v2_v1_region_proto_depIdxs,
		EnumInfos:         file_spaceone_api_inventory_v2_v1_region_proto_enumTypes,
		MessageInfos:      file_spaceone_api_inventory_v2_v1_region_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v2_v1_region_proto = out.File
	file_spaceone_api_inventory_v2_v1_region_proto_goTypes = nil
	file_spaceone_api_inventory_v2_v1_region_proto_depIdxs = nil
}
