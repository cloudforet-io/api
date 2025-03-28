// description of dashboard

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/dashboard/v1/private_dashboard.proto

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

type CreatePrivateDashboardRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Layouts []*Layout `protobuf:"bytes,3,rep,name=layouts,proto3" json:"layouts,omitempty"`
	// +optional
	Vars *_struct.Struct `protobuf:"bytes,4,opt,name=vars,proto3" json:"vars,omitempty"`
	// +optional
	VarsSchema *_struct.Struct `protobuf:"bytes,5,opt,name=vars_schema,json=varsSchema,proto3" json:"vars_schema,omitempty"`
	// +optional
	Options *_struct.Struct `protobuf:"bytes,6,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	Variables *_struct.Struct `protobuf:"bytes,7,opt,name=variables,proto3" json:"variables,omitempty"`
	// +optional
	VariablesSchema *_struct.Struct `protobuf:"bytes,8,opt,name=variables_schema,json=variablesSchema,proto3" json:"variables_schema,omitempty"`
	// +optional
	Labels *_struct.ListValue `protobuf:"bytes,9,opt,name=labels,proto3" json:"labels,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	FolderId      string `protobuf:"bytes,22,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePrivateDashboardRequest) Reset() {
	*x = CreatePrivateDashboardRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePrivateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivateDashboardRequest) ProtoMessage() {}

func (x *CreatePrivateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivateDashboardRequest.ProtoReflect.Descriptor instead.
func (*CreatePrivateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePrivateDashboardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePrivateDashboardRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePrivateDashboardRequest) GetLayouts() []*Layout {
	if x != nil {
		return x.Layouts
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetVars() *_struct.Struct {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetVarsSchema() *_struct.Struct {
	if x != nil {
		return x.VarsSchema
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetVariables() *_struct.Struct {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetVariablesSchema() *_struct.Struct {
	if x != nil {
		return x.VariablesSchema
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetLabels() *_struct.ListValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreatePrivateDashboardRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreatePrivateDashboardRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

type UpdatePrivateDashboardRequest struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	DashboardId string                 `protobuf:"bytes,1,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Layouts []*Layout `protobuf:"bytes,4,rep,name=layouts,proto3" json:"layouts,omitempty"`
	// +optional
	Vars *_struct.Struct `protobuf:"bytes,5,opt,name=vars,proto3" json:"vars,omitempty"`
	// +optional
	VarsSchema *_struct.Struct `protobuf:"bytes,6,opt,name=vars_schema,json=varsSchema,proto3" json:"vars_schema,omitempty"`
	// +optional
	Options *_struct.Struct `protobuf:"bytes,7,opt,name=options,proto3" json:"options,omitempty"`
	// +optional
	Variables *_struct.Struct `protobuf:"bytes,8,opt,name=variables,proto3" json:"variables,omitempty"`
	// +optional
	VariablesSchema *_struct.Struct `protobuf:"bytes,9,opt,name=variables_schema,json=variablesSchema,proto3" json:"variables_schema,omitempty"`
	// +optional
	Labels *_struct.ListValue `protobuf:"bytes,10,opt,name=labels,proto3" json:"labels,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	FolderId      string `protobuf:"bytes,21,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePrivateDashboardRequest) Reset() {
	*x = UpdatePrivateDashboardRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePrivateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrivateDashboardRequest) ProtoMessage() {}

func (x *UpdatePrivateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrivateDashboardRequest.ProtoReflect.Descriptor instead.
func (*UpdatePrivateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePrivateDashboardRequest) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *UpdatePrivateDashboardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePrivateDashboardRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdatePrivateDashboardRequest) GetLayouts() []*Layout {
	if x != nil {
		return x.Layouts
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetVars() *_struct.Struct {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetVarsSchema() *_struct.Struct {
	if x != nil {
		return x.VarsSchema
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetVariables() *_struct.Struct {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetVariablesSchema() *_struct.Struct {
	if x != nil {
		return x.VariablesSchema
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetLabels() *_struct.ListValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdatePrivateDashboardRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

type ChangeFolderPrivateDashboardRequest struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	DashboardId string                 `protobuf:"bytes,1,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	// +optional
	FolderId      string `protobuf:"bytes,21,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeFolderPrivateDashboardRequest) Reset() {
	*x = ChangeFolderPrivateDashboardRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeFolderPrivateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeFolderPrivateDashboardRequest) ProtoMessage() {}

func (x *ChangeFolderPrivateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeFolderPrivateDashboardRequest.ProtoReflect.Descriptor instead.
func (*ChangeFolderPrivateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeFolderPrivateDashboardRequest) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *ChangeFolderPrivateDashboardRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

type PrivateDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DashboardId   string                 `protobuf:"bytes,1,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateDashboardRequest) Reset() {
	*x = PrivateDashboardRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateDashboardRequest) ProtoMessage() {}

func (x *PrivateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateDashboardRequest.ProtoReflect.Descriptor instead.
func (*PrivateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *PrivateDashboardRequest) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

type PrivateDashboardQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	DashboardId string `protobuf:"bytes,2,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	FolderId      string `protobuf:"bytes,22,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateDashboardQuery) Reset() {
	*x = PrivateDashboardQuery{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateDashboardQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateDashboardQuery) ProtoMessage() {}

func (x *PrivateDashboardQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateDashboardQuery.ProtoReflect.Descriptor instead.
func (*PrivateDashboardQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *PrivateDashboardQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PrivateDashboardQuery) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *PrivateDashboardQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateDashboardQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *PrivateDashboardQuery) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

type PrivateDashboardInfo struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	DashboardId     string                 `protobuf:"bytes,1,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Version         string                 `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Layouts         []*Layout              `protobuf:"bytes,5,rep,name=layouts,proto3" json:"layouts,omitempty"`
	Vars            *_struct.Struct        `protobuf:"bytes,6,opt,name=vars,proto3" json:"vars,omitempty"`
	VarsSchema      *_struct.Struct        `protobuf:"bytes,7,opt,name=vars_schema,json=varsSchema,proto3" json:"vars_schema,omitempty"`
	Options         *_struct.Struct        `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	Variables       *_struct.Struct        `protobuf:"bytes,9,opt,name=variables,proto3" json:"variables,omitempty"`
	VariablesSchema *_struct.Struct        `protobuf:"bytes,10,opt,name=variables_schema,json=variablesSchema,proto3" json:"variables_schema,omitempty"`
	Labels          *_struct.ListValue     `protobuf:"bytes,11,opt,name=labels,proto3" json:"labels,omitempty"`
	Tags            *_struct.Struct        `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId        string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId     string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	UserId          string                 `protobuf:"bytes,23,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FolderId        string                 `protobuf:"bytes,24,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt       string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PrivateDashboardInfo) Reset() {
	*x = PrivateDashboardInfo{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateDashboardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateDashboardInfo) ProtoMessage() {}

func (x *PrivateDashboardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateDashboardInfo.ProtoReflect.Descriptor instead.
func (*PrivateDashboardInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{5}
}

func (x *PrivateDashboardInfo) GetDashboardId() string {
	if x != nil {
		return x.DashboardId
	}
	return ""
}

func (x *PrivateDashboardInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateDashboardInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PrivateDashboardInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PrivateDashboardInfo) GetLayouts() []*Layout {
	if x != nil {
		return x.Layouts
	}
	return nil
}

func (x *PrivateDashboardInfo) GetVars() *_struct.Struct {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *PrivateDashboardInfo) GetVarsSchema() *_struct.Struct {
	if x != nil {
		return x.VarsSchema
	}
	return nil
}

func (x *PrivateDashboardInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PrivateDashboardInfo) GetVariables() *_struct.Struct {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *PrivateDashboardInfo) GetVariablesSchema() *_struct.Struct {
	if x != nil {
		return x.VariablesSchema
	}
	return nil
}

func (x *PrivateDashboardInfo) GetLabels() *_struct.ListValue {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *PrivateDashboardInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PrivateDashboardInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *PrivateDashboardInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *PrivateDashboardInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PrivateDashboardInfo) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *PrivateDashboardInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PrivateDashboardInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PrivateDashboardsInfo struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Results       []*PrivateDashboardInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                   `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateDashboardsInfo) Reset() {
	*x = PrivateDashboardsInfo{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateDashboardsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateDashboardsInfo) ProtoMessage() {}

func (x *PrivateDashboardsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateDashboardsInfo.ProtoReflect.Descriptor instead.
func (*PrivateDashboardsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{6}
}

func (x *PrivateDashboardsInfo) GetResults() []*PrivateDashboardInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PrivateDashboardsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type PrivateDashboardStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateDashboardStatQuery) Reset() {
	*x = PrivateDashboardStatQuery{}
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateDashboardStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateDashboardStatQuery) ProtoMessage() {}

func (x *PrivateDashboardStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateDashboardStatQuery.ProtoReflect.Descriptor instead.
func (*PrivateDashboardStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP(), []int{7}
}

func (x *PrivateDashboardStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_dashboard_v1_private_dashboard_proto protoreflect.FileDescriptor

const file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDesc = "" +
	"\n" +
	"1spaceone/api/dashboard/v1/private_dashboard.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a0spaceone/api/dashboard/v1/public_dashboard.proto\"\xc8\x04\n" +
	"\x1dCreatePrivateDashboardRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12;\n" +
	"\alayouts\x18\x03 \x03(\v2!.spaceone.api.dashboard.v1.LayoutR\alayouts\x12+\n" +
	"\x04vars\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x04vars\x128\n" +
	"\vvars_schema\x18\x05 \x01(\v2\x17.google.protobuf.StructR\n" +
	"varsSchema\x121\n" +
	"\aoptions\x18\x06 \x01(\v2\x17.google.protobuf.StructR\aoptions\x125\n" +
	"\tvariables\x18\a \x01(\v2\x17.google.protobuf.StructR\tvariables\x12B\n" +
	"\x10variables_schema\x18\b \x01(\v2\x17.google.protobuf.StructR\x0fvariablesSchema\x122\n" +
	"\x06labels\x18\t \x01(\v2\x1a.google.protobuf.ListValueR\x06labels\x12+\n" +
	"\x04tags\x18\n" +
	" \x01(\v2\x17.google.protobuf.StructR\x04tags\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12\x1b\n" +
	"\tfolder_id\x18\x16 \x01(\tR\bfolderId\"\xc8\x04\n" +
	"\x1dUpdatePrivateDashboardRequest\x12!\n" +
	"\fdashboard_id\x18\x01 \x01(\tR\vdashboardId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12;\n" +
	"\alayouts\x18\x04 \x03(\v2!.spaceone.api.dashboard.v1.LayoutR\alayouts\x12+\n" +
	"\x04vars\x18\x05 \x01(\v2\x17.google.protobuf.StructR\x04vars\x128\n" +
	"\vvars_schema\x18\x06 \x01(\v2\x17.google.protobuf.StructR\n" +
	"varsSchema\x121\n" +
	"\aoptions\x18\a \x01(\v2\x17.google.protobuf.StructR\aoptions\x125\n" +
	"\tvariables\x18\b \x01(\v2\x17.google.protobuf.StructR\tvariables\x12B\n" +
	"\x10variables_schema\x18\t \x01(\v2\x17.google.protobuf.StructR\x0fvariablesSchema\x122\n" +
	"\x06labels\x18\n" +
	" \x01(\v2\x1a.google.protobuf.ListValueR\x06labels\x12+\n" +
	"\x04tags\x18\v \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tfolder_id\x18\x15 \x01(\tR\bfolderId\"e\n" +
	"#ChangeFolderPrivateDashboardRequest\x12!\n" +
	"\fdashboard_id\x18\x01 \x01(\tR\vdashboardId\x12\x1b\n" +
	"\tfolder_id\x18\x15 \x01(\tR\bfolderId\"<\n" +
	"\x17PrivateDashboardRequest\x12!\n" +
	"\fdashboard_id\x18\x01 \x01(\tR\vdashboardId\"\xc1\x01\n" +
	"\x15PrivateDashboardQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12!\n" +
	"\fdashboard_id\x18\x02 \x01(\tR\vdashboardId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12\x1b\n" +
	"\tfolder_id\x18\x16 \x01(\tR\bfolderId\"\xf0\x05\n" +
	"\x14PrivateDashboardInfo\x12!\n" +
	"\fdashboard_id\x18\x01 \x01(\tR\vdashboardId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x18\n" +
	"\aversion\x18\x04 \x01(\tR\aversion\x12;\n" +
	"\alayouts\x18\x05 \x03(\v2!.spaceone.api.dashboard.v1.LayoutR\alayouts\x12+\n" +
	"\x04vars\x18\x06 \x01(\v2\x17.google.protobuf.StructR\x04vars\x128\n" +
	"\vvars_schema\x18\a \x01(\v2\x17.google.protobuf.StructR\n" +
	"varsSchema\x121\n" +
	"\aoptions\x18\b \x01(\v2\x17.google.protobuf.StructR\aoptions\x125\n" +
	"\tvariables\x18\t \x01(\v2\x17.google.protobuf.StructR\tvariables\x12B\n" +
	"\x10variables_schema\x18\n" +
	" \x01(\v2\x17.google.protobuf.StructR\x0fvariablesSchema\x122\n" +
	"\x06labels\x18\v \x01(\v2\x1a.google.protobuf.ListValueR\x06labels\x12+\n" +
	"\x04tags\x18\f \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x17\n" +
	"\auser_id\x18\x17 \x01(\tR\x06userId\x12\x1b\n" +
	"\tfolder_id\x18\x18 \x01(\tR\bfolderId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18  \x01(\tR\tupdatedAt\"\x83\x01\n" +
	"\x15PrivateDashboardsInfo\x12I\n" +
	"\aresults\x18\x01 \x03(\v2/.spaceone.api.dashboard.v1.PrivateDashboardInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"X\n" +
	"\x19PrivateDashboardStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\xef\b\n" +
	"\x10PrivateDashboard\x12\xa6\x01\n" +
	"\x06create\x128.spaceone.api.dashboard.v1.CreatePrivateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\"1\x82\xd3\xe4\x93\x02+:\x01*\"&/dashboard/v1/private-dashboard/create\x12\xa6\x01\n" +
	"\x06update\x128.spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\"1\x82\xd3\xe4\x93\x02+:\x01*\"&/dashboard/v1/private-dashboard/update\x12\xba\x01\n" +
	"\rchange_folder\x12>.spaceone.api.dashboard.v1.ChangeFolderPrivateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\"8\x82\xd3\xe4\x93\x022:\x01*\"-/dashboard/v1/private-dashboard/change-folder\x12\x87\x01\n" +
	"\x06delete\x122.spaceone.api.dashboard.v1.PrivateDashboardRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+:\x01*\"&/dashboard/v1/private-dashboard/delete\x12\x9a\x01\n" +
	"\x03get\x122.spaceone.api.dashboard.v1.PrivateDashboardRequest\x1a/.spaceone.api.dashboard.v1.PrivateDashboardInfo\".\x82\xd3\xe4\x93\x02(:\x01*\"#/dashboard/v1/private-dashboard/get\x12\x9b\x01\n" +
	"\x04list\x120.spaceone.api.dashboard.v1.PrivateDashboardQuery\x1a0.spaceone.api.dashboard.v1.PrivateDashboardsInfo\"/\x82\xd3\xe4\x93\x02):\x01*\"$/dashboard/v1/private-dashboard/list\x12\x86\x01\n" +
	"\x04stat\x124.spaceone.api.dashboard.v1.PrivateDashboardStatQuery\x1a\x17.google.protobuf.Struct\"/\x82\xd3\xe4\x93\x02):\x01*\"$/dashboard/v1/private-dashboard/statB@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3"

var (
	file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescOnce sync.Once
	file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescData []byte
)

func file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescGZIP() []byte {
	file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescOnce.Do(func() {
		file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDesc), len(file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDesc)))
	})
	return file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDescData
}

var file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_dashboard_v1_private_dashboard_proto_goTypes = []any{
	(*CreatePrivateDashboardRequest)(nil),       // 0: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest
	(*UpdatePrivateDashboardRequest)(nil),       // 1: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest
	(*ChangeFolderPrivateDashboardRequest)(nil), // 2: spaceone.api.dashboard.v1.ChangeFolderPrivateDashboardRequest
	(*PrivateDashboardRequest)(nil),             // 3: spaceone.api.dashboard.v1.PrivateDashboardRequest
	(*PrivateDashboardQuery)(nil),               // 4: spaceone.api.dashboard.v1.PrivateDashboardQuery
	(*PrivateDashboardInfo)(nil),                // 5: spaceone.api.dashboard.v1.PrivateDashboardInfo
	(*PrivateDashboardsInfo)(nil),               // 6: spaceone.api.dashboard.v1.PrivateDashboardsInfo
	(*PrivateDashboardStatQuery)(nil),           // 7: spaceone.api.dashboard.v1.PrivateDashboardStatQuery
	(*Layout)(nil),                              // 8: spaceone.api.dashboard.v1.Layout
	(*_struct.Struct)(nil),                      // 9: google.protobuf.Struct
	(*_struct.ListValue)(nil),                   // 10: google.protobuf.ListValue
	(*v2.Query)(nil),                            // 11: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),                  // 12: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                         // 13: google.protobuf.Empty
}
var file_spaceone_api_dashboard_v1_private_dashboard_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.layouts:type_name -> spaceone.api.dashboard.v1.Layout
	9,  // 1: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.vars:type_name -> google.protobuf.Struct
	9,  // 2: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.vars_schema:type_name -> google.protobuf.Struct
	9,  // 3: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.options:type_name -> google.protobuf.Struct
	9,  // 4: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.variables:type_name -> google.protobuf.Struct
	9,  // 5: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.variables_schema:type_name -> google.protobuf.Struct
	10, // 6: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.labels:type_name -> google.protobuf.ListValue
	9,  // 7: spaceone.api.dashboard.v1.CreatePrivateDashboardRequest.tags:type_name -> google.protobuf.Struct
	8,  // 8: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.layouts:type_name -> spaceone.api.dashboard.v1.Layout
	9,  // 9: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.vars:type_name -> google.protobuf.Struct
	9,  // 10: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.vars_schema:type_name -> google.protobuf.Struct
	9,  // 11: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.options:type_name -> google.protobuf.Struct
	9,  // 12: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.variables:type_name -> google.protobuf.Struct
	9,  // 13: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.variables_schema:type_name -> google.protobuf.Struct
	10, // 14: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.labels:type_name -> google.protobuf.ListValue
	9,  // 15: spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest.tags:type_name -> google.protobuf.Struct
	11, // 16: spaceone.api.dashboard.v1.PrivateDashboardQuery.query:type_name -> spaceone.api.core.v2.Query
	8,  // 17: spaceone.api.dashboard.v1.PrivateDashboardInfo.layouts:type_name -> spaceone.api.dashboard.v1.Layout
	9,  // 18: spaceone.api.dashboard.v1.PrivateDashboardInfo.vars:type_name -> google.protobuf.Struct
	9,  // 19: spaceone.api.dashboard.v1.PrivateDashboardInfo.vars_schema:type_name -> google.protobuf.Struct
	9,  // 20: spaceone.api.dashboard.v1.PrivateDashboardInfo.options:type_name -> google.protobuf.Struct
	9,  // 21: spaceone.api.dashboard.v1.PrivateDashboardInfo.variables:type_name -> google.protobuf.Struct
	9,  // 22: spaceone.api.dashboard.v1.PrivateDashboardInfo.variables_schema:type_name -> google.protobuf.Struct
	10, // 23: spaceone.api.dashboard.v1.PrivateDashboardInfo.labels:type_name -> google.protobuf.ListValue
	9,  // 24: spaceone.api.dashboard.v1.PrivateDashboardInfo.tags:type_name -> google.protobuf.Struct
	5,  // 25: spaceone.api.dashboard.v1.PrivateDashboardsInfo.results:type_name -> spaceone.api.dashboard.v1.PrivateDashboardInfo
	12, // 26: spaceone.api.dashboard.v1.PrivateDashboardStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 27: spaceone.api.dashboard.v1.PrivateDashboard.create:input_type -> spaceone.api.dashboard.v1.CreatePrivateDashboardRequest
	1,  // 28: spaceone.api.dashboard.v1.PrivateDashboard.update:input_type -> spaceone.api.dashboard.v1.UpdatePrivateDashboardRequest
	2,  // 29: spaceone.api.dashboard.v1.PrivateDashboard.change_folder:input_type -> spaceone.api.dashboard.v1.ChangeFolderPrivateDashboardRequest
	3,  // 30: spaceone.api.dashboard.v1.PrivateDashboard.delete:input_type -> spaceone.api.dashboard.v1.PrivateDashboardRequest
	3,  // 31: spaceone.api.dashboard.v1.PrivateDashboard.get:input_type -> spaceone.api.dashboard.v1.PrivateDashboardRequest
	4,  // 32: spaceone.api.dashboard.v1.PrivateDashboard.list:input_type -> spaceone.api.dashboard.v1.PrivateDashboardQuery
	7,  // 33: spaceone.api.dashboard.v1.PrivateDashboard.stat:input_type -> spaceone.api.dashboard.v1.PrivateDashboardStatQuery
	5,  // 34: spaceone.api.dashboard.v1.PrivateDashboard.create:output_type -> spaceone.api.dashboard.v1.PrivateDashboardInfo
	5,  // 35: spaceone.api.dashboard.v1.PrivateDashboard.update:output_type -> spaceone.api.dashboard.v1.PrivateDashboardInfo
	5,  // 36: spaceone.api.dashboard.v1.PrivateDashboard.change_folder:output_type -> spaceone.api.dashboard.v1.PrivateDashboardInfo
	13, // 37: spaceone.api.dashboard.v1.PrivateDashboard.delete:output_type -> google.protobuf.Empty
	5,  // 38: spaceone.api.dashboard.v1.PrivateDashboard.get:output_type -> spaceone.api.dashboard.v1.PrivateDashboardInfo
	6,  // 39: spaceone.api.dashboard.v1.PrivateDashboard.list:output_type -> spaceone.api.dashboard.v1.PrivateDashboardsInfo
	9,  // 40: spaceone.api.dashboard.v1.PrivateDashboard.stat:output_type -> google.protobuf.Struct
	34, // [34:41] is the sub-list for method output_type
	27, // [27:34] is the sub-list for method input_type
	27, // [27:27] is the sub-list for extension type_name
	27, // [27:27] is the sub-list for extension extendee
	0,  // [0:27] is the sub-list for field type_name
}

func init() { file_spaceone_api_dashboard_v1_private_dashboard_proto_init() }
func file_spaceone_api_dashboard_v1_private_dashboard_proto_init() {
	if File_spaceone_api_dashboard_v1_private_dashboard_proto != nil {
		return
	}
	file_spaceone_api_dashboard_v1_public_dashboard_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDesc), len(file_spaceone_api_dashboard_v1_private_dashboard_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_dashboard_v1_private_dashboard_proto_goTypes,
		DependencyIndexes: file_spaceone_api_dashboard_v1_private_dashboard_proto_depIdxs,
		MessageInfos:      file_spaceone_api_dashboard_v1_private_dashboard_proto_msgTypes,
	}.Build()
	File_spaceone_api_dashboard_v1_private_dashboard_proto = out.File
	file_spaceone_api_dashboard_v1_private_dashboard_proto_goTypes = nil
	file_spaceone_api_dashboard_v1_private_dashboard_proto_depIdxs = nil
}
