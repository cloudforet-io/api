// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/identity/v1/role_binding.proto

package v1

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
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

type CreateRoleBindingRequest struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	ResourceType string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId   string                 `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RoleId       string                 `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	ProjectGroupId string `protobuf:"bytes,5,opt,name=project_group_id,json=projectGroupId,proto3" json:"project_group_id,omitempty"`
	// +optional
	Labels *_struct.ListValue `protobuf:"bytes,6,opt,name=labels,proto3" json:"labels,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,7,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId      string          `protobuf:"bytes,8,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoleBindingRequest) Reset() {
	*x = CreateRoleBindingRequest{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleBindingRequest) ProtoMessage() {}

func (x *CreateRoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleBindingRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoleBindingRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *CreateRoleBindingRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CreateRoleBindingRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *CreateRoleBindingRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateRoleBindingRequest) GetProjectGroupId() string {
	if x != nil {
		return x.ProjectGroupId
	}
	return ""
}

func (x *CreateRoleBindingRequest) GetLabels() *_struct.ListValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateRoleBindingRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateRoleBindingRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type UpdateRoleBindingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleBindingId string                 `protobuf:"bytes,1,opt,name=role_binding_id,json=roleBindingId,proto3" json:"role_binding_id,omitempty"`
	// +optional
	Labels *_struct.ListValue `protobuf:"bytes,2,opt,name=labels,proto3" json:"labels,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId      string          `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRoleBindingRequest) Reset() {
	*x = UpdateRoleBindingRequest{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleBindingRequest) ProtoMessage() {}

func (x *UpdateRoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleBindingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRoleBindingRequest) GetRoleBindingId() string {
	if x != nil {
		return x.RoleBindingId
	}
	return ""
}

func (x *UpdateRoleBindingRequest) GetLabels() *_struct.ListValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *UpdateRoleBindingRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateRoleBindingRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type RoleBindingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleBindingId string                 `protobuf:"bytes,1,opt,name=role_binding_id,json=roleBindingId,proto3" json:"role_binding_id,omitempty"`
	DomainId      string                 `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleBindingRequest) Reset() {
	*x = RoleBindingRequest{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingRequest) ProtoMessage() {}

func (x *RoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingRequest.ProtoReflect.Descriptor instead.
func (*RoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{2}
}

func (x *RoleBindingRequest) GetRoleBindingId() string {
	if x != nil {
		return x.RoleBindingId
	}
	return ""
}

func (x *RoleBindingRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type GetRoleBindingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleBindingId string                 `protobuf:"bytes,1,opt,name=role_binding_id,json=roleBindingId,proto3" json:"role_binding_id,omitempty"`
	DomainId      string                 `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	Only          []string `protobuf:"bytes,3,rep,name=only,proto3" json:"only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoleBindingRequest) Reset() {
	*x = GetRoleBindingRequest{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleBindingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleBindingRequest) ProtoMessage() {}

func (x *GetRoleBindingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleBindingRequest.ProtoReflect.Descriptor instead.
func (*GetRoleBindingRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoleBindingRequest) GetRoleBindingId() string {
	if x != nil {
		return x.RoleBindingId
	}
	return ""
}

func (x *GetRoleBindingRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GetRoleBindingRequest) GetOnly() []string {
	if x != nil {
		return x.Only
	}
	return nil
}

type RoleBindingInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	RoleBindingId    string                 `protobuf:"bytes,1,opt,name=role_binding_id,json=roleBindingId,proto3" json:"role_binding_id,omitempty"`
	ResourceType     string                 `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId       string                 `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	RoleInfo         *RoleInfo              `protobuf:"bytes,4,opt,name=role_info,json=roleInfo,proto3" json:"role_info,omitempty"`
	ProjectInfo      *ProjectInfo           `protobuf:"bytes,5,opt,name=project_info,json=projectInfo,proto3" json:"project_info,omitempty"`
	ProjectGroupInfo *ProjectGroupInfo      `protobuf:"bytes,6,opt,name=project_group_info,json=projectGroupInfo,proto3" json:"project_group_info,omitempty"`
	Labels           *_struct.ListValue     `protobuf:"bytes,7,opt,name=labels,proto3" json:"labels,omitempty"`
	Tags             *_struct.Struct        `protobuf:"bytes,8,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId         string                 `protobuf:"bytes,11,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt        string                 `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *RoleBindingInfo) Reset() {
	*x = RoleBindingInfo{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleBindingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingInfo) ProtoMessage() {}

func (x *RoleBindingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingInfo.ProtoReflect.Descriptor instead.
func (*RoleBindingInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{4}
}

func (x *RoleBindingInfo) GetRoleBindingId() string {
	if x != nil {
		return x.RoleBindingId
	}
	return ""
}

func (x *RoleBindingInfo) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *RoleBindingInfo) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RoleBindingInfo) GetRoleInfo() *RoleInfo {
	if x != nil {
		return x.RoleInfo
	}
	return nil
}

func (x *RoleBindingInfo) GetProjectInfo() *ProjectInfo {
	if x != nil {
		return x.ProjectInfo
	}
	return nil
}

func (x *RoleBindingInfo) GetProjectGroupInfo() *ProjectGroupInfo {
	if x != nil {
		return x.ProjectGroupInfo
	}
	return nil
}

func (x *RoleBindingInfo) GetLabels() *_struct.ListValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RoleBindingInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RoleBindingInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *RoleBindingInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type RoleBindingQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	RoleBindingId string `protobuf:"bytes,2,opt,name=role_binding_id,json=roleBindingId,proto3" json:"role_binding_id,omitempty"`
	// +optional
	ResourceType string `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// +optional
	ResourceId string `protobuf:"bytes,4,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// +optional
	RoleId string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// +optional
	RoleType string `protobuf:"bytes,6,opt,name=role_type,json=roleType,proto3" json:"role_type,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,7,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	ProjectGroupId string `protobuf:"bytes,8,opt,name=project_group_id,json=projectGroupId,proto3" json:"project_group_id,omitempty"`
	DomainId       string `protobuf:"bytes,9,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RoleBindingQuery) Reset() {
	*x = RoleBindingQuery{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleBindingQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingQuery) ProtoMessage() {}

func (x *RoleBindingQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingQuery.ProtoReflect.Descriptor instead.
func (*RoleBindingQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{5}
}

func (x *RoleBindingQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RoleBindingQuery) GetRoleBindingId() string {
	if x != nil {
		return x.RoleBindingId
	}
	return ""
}

func (x *RoleBindingQuery) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *RoleBindingQuery) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RoleBindingQuery) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleBindingQuery) GetRoleType() string {
	if x != nil {
		return x.RoleType
	}
	return ""
}

func (x *RoleBindingQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *RoleBindingQuery) GetProjectGroupId() string {
	if x != nil {
		return x.ProjectGroupId
	}
	return ""
}

func (x *RoleBindingQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type RoleBindingsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*RoleBindingInfo     `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleBindingsInfo) Reset() {
	*x = RoleBindingsInfo{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleBindingsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingsInfo) ProtoMessage() {}

func (x *RoleBindingsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingsInfo.ProtoReflect.Descriptor instead.
func (*RoleBindingsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{6}
}

func (x *RoleBindingsInfo) GetResults() []*RoleBindingInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *RoleBindingsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type RoleBindingStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v1.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId      string                 `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleBindingStatQuery) Reset() {
	*x = RoleBindingStatQuery{}
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleBindingStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingStatQuery) ProtoMessage() {}

func (x *RoleBindingStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_role_binding_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingStatQuery.ProtoReflect.Descriptor instead.
func (*RoleBindingStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP(), []int{7}
}

func (x *RoleBindingStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RoleBindingStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_identity_v1_role_binding_proto protoreflect.FileDescriptor

const file_spaceone_api_identity_v1_role_binding_proto_rawDesc = "" +
	"\n" +
	"+spaceone/api/identity/v1/role_binding.proto\x12\x18spaceone.api.identity.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v1/query.proto\x1a&spaceone/api/identity/v1/project.proto\x1a,spaceone/api/identity/v1/project_group.proto\x1a#spaceone/api/identity/v1/role.proto\"\xc0\x02\n" +
	"\x18CreateRoleBindingRequest\x12#\n" +
	"\rresource_type\x18\x01 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x02 \x01(\tR\n" +
	"resourceId\x12\x17\n" +
	"\arole_id\x18\x03 \x01(\tR\x06roleId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x04 \x01(\tR\tprojectId\x12(\n" +
	"\x10project_group_id\x18\x05 \x01(\tR\x0eprojectGroupId\x122\n" +
	"\x06labels\x18\x06 \x01(\v2\x1a.google.protobuf.ListValueR\x06labels\x12+\n" +
	"\x04tags\x18\a \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\b \x01(\tR\bdomainId\"\xc0\x01\n" +
	"\x18UpdateRoleBindingRequest\x12&\n" +
	"\x0frole_binding_id\x18\x01 \x01(\tR\rroleBindingId\x122\n" +
	"\x06labels\x18\x02 \x01(\v2\x1a.google.protobuf.ListValueR\x06labels\x12+\n" +
	"\x04tags\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\x04 \x01(\tR\bdomainId\"Y\n" +
	"\x12RoleBindingRequest\x12&\n" +
	"\x0frole_binding_id\x18\x01 \x01(\tR\rroleBindingId\x12\x1b\n" +
	"\tdomain_id\x18\x02 \x01(\tR\bdomainId\"p\n" +
	"\x15GetRoleBindingRequest\x12&\n" +
	"\x0frole_binding_id\x18\x01 \x01(\tR\rroleBindingId\x12\x1b\n" +
	"\tdomain_id\x18\x02 \x01(\tR\bdomainId\x12\x12\n" +
	"\x04only\x18\x03 \x03(\tR\x04only\"\x81\x04\n" +
	"\x0fRoleBindingInfo\x12&\n" +
	"\x0frole_binding_id\x18\x01 \x01(\tR\rroleBindingId\x12#\n" +
	"\rresource_type\x18\x02 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x03 \x01(\tR\n" +
	"resourceId\x12?\n" +
	"\trole_info\x18\x04 \x01(\v2\".spaceone.api.identity.v1.RoleInfoR\broleInfo\x12H\n" +
	"\fproject_info\x18\x05 \x01(\v2%.spaceone.api.identity.v1.ProjectInfoR\vprojectInfo\x12X\n" +
	"\x12project_group_info\x18\x06 \x01(\v2*.spaceone.api.identity.v1.ProjectGroupInfoR\x10projectGroupInfo\x122\n" +
	"\x06labels\x18\a \x01(\v2\x1a.google.protobuf.ListValueR\x06labels\x12+\n" +
	"\x04tags\x18\b \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\v \x01(\tR\bdomainId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x15 \x01(\tR\tcreatedAt\"\xcf\x02\n" +
	"\x10RoleBindingQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v1.QueryR\x05query\x12&\n" +
	"\x0frole_binding_id\x18\x02 \x01(\tR\rroleBindingId\x12#\n" +
	"\rresource_type\x18\x03 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x04 \x01(\tR\n" +
	"resourceId\x12\x17\n" +
	"\arole_id\x18\x05 \x01(\tR\x06roleId\x12\x1b\n" +
	"\trole_type\x18\x06 \x01(\tR\broleType\x12\x1d\n" +
	"\n" +
	"project_id\x18\a \x01(\tR\tprojectId\x12(\n" +
	"\x10project_group_id\x18\b \x01(\tR\x0eprojectGroupId\x12\x1b\n" +
	"\tdomain_id\x18\t \x01(\tR\bdomainId\"x\n" +
	"\x10RoleBindingsInfo\x12C\n" +
	"\aresults\x18\x01 \x03(\v2).spaceone.api.identity.v1.RoleBindingInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"p\n" +
	"\x14RoleBindingStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v1.StatisticsQueryR\x05query\x12\x1b\n" +
	"\tdomain_id\x18\x02 \x01(\tR\bdomainId2\xce\x06\n" +
	"\vRoleBinding\x12\x94\x01\n" +
	"\x06create\x122.spaceone.api.identity.v1.CreateRoleBindingRequest\x1a).spaceone.api.identity.v1.RoleBindingInfo\"+\x82\xd3\xe4\x93\x02%:\x01*\" /identity/v1/role-binding/create\x12\x94\x01\n" +
	"\x06update\x122.spaceone.api.identity.v1.UpdateRoleBindingRequest\x1a).spaceone.api.identity.v1.RoleBindingInfo\"+\x82\xd3\xe4\x93\x02%:\x01*\" /identity/v1/role-binding/update\x12{\n" +
	"\x06delete\x12,.spaceone.api.identity.v1.RoleBindingRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%:\x01*\" /identity/v1/role-binding/delete\x12\x8b\x01\n" +
	"\x03get\x12/.spaceone.api.identity.v1.GetRoleBindingRequest\x1a).spaceone.api.identity.v1.RoleBindingInfo\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/identity/v1/role-binding/get\x12\x89\x01\n" +
	"\x04list\x12*.spaceone.api.identity.v1.RoleBindingQuery\x1a*.spaceone.api.identity.v1.RoleBindingsInfo\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/identity/v1/role-binding/list\x12z\n" +
	"\x04stat\x12..spaceone.api.identity.v1.RoleBindingStatQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/identity/v1/role-binding/statB?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1b\x06proto3"

var (
	file_spaceone_api_identity_v1_role_binding_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v1_role_binding_proto_rawDescData []byte
)

func file_spaceone_api_identity_v1_role_binding_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v1_role_binding_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v1_role_binding_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v1_role_binding_proto_rawDesc), len(file_spaceone_api_identity_v1_role_binding_proto_rawDesc)))
	})
	return file_spaceone_api_identity_v1_role_binding_proto_rawDescData
}

var file_spaceone_api_identity_v1_role_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_identity_v1_role_binding_proto_goTypes = []any{
	(*CreateRoleBindingRequest)(nil), // 0: spaceone.api.identity.v1.CreateRoleBindingRequest
	(*UpdateRoleBindingRequest)(nil), // 1: spaceone.api.identity.v1.UpdateRoleBindingRequest
	(*RoleBindingRequest)(nil),       // 2: spaceone.api.identity.v1.RoleBindingRequest
	(*GetRoleBindingRequest)(nil),    // 3: spaceone.api.identity.v1.GetRoleBindingRequest
	(*RoleBindingInfo)(nil),          // 4: spaceone.api.identity.v1.RoleBindingInfo
	(*RoleBindingQuery)(nil),         // 5: spaceone.api.identity.v1.RoleBindingQuery
	(*RoleBindingsInfo)(nil),         // 6: spaceone.api.identity.v1.RoleBindingsInfo
	(*RoleBindingStatQuery)(nil),     // 7: spaceone.api.identity.v1.RoleBindingStatQuery
	(*_struct.ListValue)(nil),        // 8: google.protobuf.ListValue
	(*_struct.Struct)(nil),           // 9: google.protobuf.Struct
	(*RoleInfo)(nil),                 // 10: spaceone.api.identity.v1.RoleInfo
	(*ProjectInfo)(nil),              // 11: spaceone.api.identity.v1.ProjectInfo
	(*ProjectGroupInfo)(nil),         // 12: spaceone.api.identity.v1.ProjectGroupInfo
	(*v1.Query)(nil),                 // 13: spaceone.api.core.v1.Query
	(*v1.StatisticsQuery)(nil),       // 14: spaceone.api.core.v1.StatisticsQuery
	(*empty.Empty)(nil),              // 15: google.protobuf.Empty
}
var file_spaceone_api_identity_v1_role_binding_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.identity.v1.CreateRoleBindingRequest.labels:type_name -> google.protobuf.ListValue
	9,  // 1: spaceone.api.identity.v1.CreateRoleBindingRequest.tags:type_name -> google.protobuf.Struct
	8,  // 2: spaceone.api.identity.v1.UpdateRoleBindingRequest.labels:type_name -> google.protobuf.ListValue
	9,  // 3: spaceone.api.identity.v1.UpdateRoleBindingRequest.tags:type_name -> google.protobuf.Struct
	10, // 4: spaceone.api.identity.v1.RoleBindingInfo.role_info:type_name -> spaceone.api.identity.v1.RoleInfo
	11, // 5: spaceone.api.identity.v1.RoleBindingInfo.project_info:type_name -> spaceone.api.identity.v1.ProjectInfo
	12, // 6: spaceone.api.identity.v1.RoleBindingInfo.project_group_info:type_name -> spaceone.api.identity.v1.ProjectGroupInfo
	8,  // 7: spaceone.api.identity.v1.RoleBindingInfo.labels:type_name -> google.protobuf.ListValue
	9,  // 8: spaceone.api.identity.v1.RoleBindingInfo.tags:type_name -> google.protobuf.Struct
	13, // 9: spaceone.api.identity.v1.RoleBindingQuery.query:type_name -> spaceone.api.core.v1.Query
	4,  // 10: spaceone.api.identity.v1.RoleBindingsInfo.results:type_name -> spaceone.api.identity.v1.RoleBindingInfo
	14, // 11: spaceone.api.identity.v1.RoleBindingStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	0,  // 12: spaceone.api.identity.v1.RoleBinding.create:input_type -> spaceone.api.identity.v1.CreateRoleBindingRequest
	1,  // 13: spaceone.api.identity.v1.RoleBinding.update:input_type -> spaceone.api.identity.v1.UpdateRoleBindingRequest
	2,  // 14: spaceone.api.identity.v1.RoleBinding.delete:input_type -> spaceone.api.identity.v1.RoleBindingRequest
	3,  // 15: spaceone.api.identity.v1.RoleBinding.get:input_type -> spaceone.api.identity.v1.GetRoleBindingRequest
	5,  // 16: spaceone.api.identity.v1.RoleBinding.list:input_type -> spaceone.api.identity.v1.RoleBindingQuery
	7,  // 17: spaceone.api.identity.v1.RoleBinding.stat:input_type -> spaceone.api.identity.v1.RoleBindingStatQuery
	4,  // 18: spaceone.api.identity.v1.RoleBinding.create:output_type -> spaceone.api.identity.v1.RoleBindingInfo
	4,  // 19: spaceone.api.identity.v1.RoleBinding.update:output_type -> spaceone.api.identity.v1.RoleBindingInfo
	15, // 20: spaceone.api.identity.v1.RoleBinding.delete:output_type -> google.protobuf.Empty
	4,  // 21: spaceone.api.identity.v1.RoleBinding.get:output_type -> spaceone.api.identity.v1.RoleBindingInfo
	6,  // 22: spaceone.api.identity.v1.RoleBinding.list:output_type -> spaceone.api.identity.v1.RoleBindingsInfo
	9,  // 23: spaceone.api.identity.v1.RoleBinding.stat:output_type -> google.protobuf.Struct
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v1_role_binding_proto_init() }
func file_spaceone_api_identity_v1_role_binding_proto_init() {
	if File_spaceone_api_identity_v1_role_binding_proto != nil {
		return
	}
	file_spaceone_api_identity_v1_project_proto_init()
	file_spaceone_api_identity_v1_project_group_proto_init()
	file_spaceone_api_identity_v1_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v1_role_binding_proto_rawDesc), len(file_spaceone_api_identity_v1_role_binding_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v1_role_binding_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v1_role_binding_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v1_role_binding_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v1_role_binding_proto = out.File
	file_spaceone_api_identity_v1_role_binding_proto_goTypes = nil
	file_spaceone_api_identity_v1_role_binding_proto_depIdxs = nil
}
