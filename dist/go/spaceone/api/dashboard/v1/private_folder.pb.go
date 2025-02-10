// description of dashboard

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/dashboard/v1/private_folder.proto

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

type CreatePrivateFolderRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	Dashboards []*_struct.Struct `protobuf:"bytes,3,rep,name=dashboards,proto3" json:"dashboards,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,4,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePrivateFolderRequest) Reset() {
	*x = CreatePrivateFolderRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePrivateFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivateFolderRequest) ProtoMessage() {}

func (x *CreatePrivateFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivateFolderRequest.ProtoReflect.Descriptor instead.
func (*CreatePrivateFolderRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePrivateFolderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePrivateFolderRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreatePrivateFolderRequest) GetDashboards() []*_struct.Struct {
	if x != nil {
		return x.Dashboards
	}
	return nil
}

func (x *CreatePrivateFolderRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type UpdatePrivateFolderRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	FolderId string                 `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePrivateFolderRequest) Reset() {
	*x = UpdatePrivateFolderRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePrivateFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrivateFolderRequest) ProtoMessage() {}

func (x *UpdatePrivateFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrivateFolderRequest.ProtoReflect.Descriptor instead.
func (*UpdatePrivateFolderRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePrivateFolderRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *UpdatePrivateFolderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePrivateFolderRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PrivateFolderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FolderId      string                 `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateFolderRequest) Reset() {
	*x = PrivateFolderRequest{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateFolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateFolderRequest) ProtoMessage() {}

func (x *PrivateFolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateFolderRequest.ProtoReflect.Descriptor instead.
func (*PrivateFolderRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateFolderRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

type PrivateFolderQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,4,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateFolderQuery) Reset() {
	*x = PrivateFolderQuery{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateFolderQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateFolderQuery) ProtoMessage() {}

func (x *PrivateFolderQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateFolderQuery.ProtoReflect.Descriptor instead.
func (*PrivateFolderQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{3}
}

func (x *PrivateFolderQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PrivateFolderQuery) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *PrivateFolderQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateFolderQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type PrivateFolderInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FolderId      string                 `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tags          *_struct.Struct        `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	UserId        string                 `protobuf:"bytes,23,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateFolderInfo) Reset() {
	*x = PrivateFolderInfo{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateFolderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateFolderInfo) ProtoMessage() {}

func (x *PrivateFolderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateFolderInfo.ProtoReflect.Descriptor instead.
func (*PrivateFolderInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{4}
}

func (x *PrivateFolderInfo) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *PrivateFolderInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateFolderInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PrivateFolderInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *PrivateFolderInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *PrivateFolderInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PrivateFolderInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PrivateFolderInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PrivateFoldersInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*PrivateFolderInfo   `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateFoldersInfo) Reset() {
	*x = PrivateFoldersInfo{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateFoldersInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateFoldersInfo) ProtoMessage() {}

func (x *PrivateFoldersInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateFoldersInfo.ProtoReflect.Descriptor instead.
func (*PrivateFoldersInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{5}
}

func (x *PrivateFoldersInfo) GetResults() []*PrivateFolderInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PrivateFoldersInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type PrivateFolderStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateFolderStatQuery) Reset() {
	*x = PrivateFolderStatQuery{}
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateFolderStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateFolderStatQuery) ProtoMessage() {}

func (x *PrivateFolderStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateFolderStatQuery.ProtoReflect.Descriptor instead.
func (*PrivateFolderStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP(), []int{6}
}

func (x *PrivateFolderStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_dashboard_v1_private_folder_proto protoreflect.FileDescriptor

var file_spaceone_api_dashboard_v1_private_folder_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x7a, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x33,
	0x0a, 0x14, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x88, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a, 0x12,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x16, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x32, 0xff, 0x06, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28,
	0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x12, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x92, 0x01, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x80, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01,
	0x2a, 0x22, 0x21, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_spaceone_api_dashboard_v1_private_folder_proto_rawDescOnce sync.Once
	file_spaceone_api_dashboard_v1_private_folder_proto_rawDescData []byte
)

func file_spaceone_api_dashboard_v1_private_folder_proto_rawDescGZIP() []byte {
	file_spaceone_api_dashboard_v1_private_folder_proto_rawDescOnce.Do(func() {
		file_spaceone_api_dashboard_v1_private_folder_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_dashboard_v1_private_folder_proto_rawDesc), len(file_spaceone_api_dashboard_v1_private_folder_proto_rawDesc)))
	})
	return file_spaceone_api_dashboard_v1_private_folder_proto_rawDescData
}

var file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_dashboard_v1_private_folder_proto_goTypes = []any{
	(*CreatePrivateFolderRequest)(nil), // 0: spaceone.api.dashboard.v1.CreatePrivateFolderRequest
	(*UpdatePrivateFolderRequest)(nil), // 1: spaceone.api.dashboard.v1.UpdatePrivateFolderRequest
	(*PrivateFolderRequest)(nil),       // 2: spaceone.api.dashboard.v1.PrivateFolderRequest
	(*PrivateFolderQuery)(nil),         // 3: spaceone.api.dashboard.v1.PrivateFolderQuery
	(*PrivateFolderInfo)(nil),          // 4: spaceone.api.dashboard.v1.PrivateFolderInfo
	(*PrivateFoldersInfo)(nil),         // 5: spaceone.api.dashboard.v1.PrivateFoldersInfo
	(*PrivateFolderStatQuery)(nil),     // 6: spaceone.api.dashboard.v1.PrivateFolderStatQuery
	(*_struct.Struct)(nil),             // 7: google.protobuf.Struct
	(*v2.Query)(nil),                   // 8: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),         // 9: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_spaceone_api_dashboard_v1_private_folder_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.dashboard.v1.CreatePrivateFolderRequest.tags:type_name -> google.protobuf.Struct
	7,  // 1: spaceone.api.dashboard.v1.CreatePrivateFolderRequest.dashboards:type_name -> google.protobuf.Struct
	7,  // 2: spaceone.api.dashboard.v1.UpdatePrivateFolderRequest.tags:type_name -> google.protobuf.Struct
	8,  // 3: spaceone.api.dashboard.v1.PrivateFolderQuery.query:type_name -> spaceone.api.core.v2.Query
	7,  // 4: spaceone.api.dashboard.v1.PrivateFolderInfo.tags:type_name -> google.protobuf.Struct
	4,  // 5: spaceone.api.dashboard.v1.PrivateFoldersInfo.results:type_name -> spaceone.api.dashboard.v1.PrivateFolderInfo
	9,  // 6: spaceone.api.dashboard.v1.PrivateFolderStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 7: spaceone.api.dashboard.v1.PrivateFolder.create:input_type -> spaceone.api.dashboard.v1.CreatePrivateFolderRequest
	1,  // 8: spaceone.api.dashboard.v1.PrivateFolder.update:input_type -> spaceone.api.dashboard.v1.UpdatePrivateFolderRequest
	2,  // 9: spaceone.api.dashboard.v1.PrivateFolder.delete:input_type -> spaceone.api.dashboard.v1.PrivateFolderRequest
	2,  // 10: spaceone.api.dashboard.v1.PrivateFolder.get:input_type -> spaceone.api.dashboard.v1.PrivateFolderRequest
	3,  // 11: spaceone.api.dashboard.v1.PrivateFolder.list:input_type -> spaceone.api.dashboard.v1.PrivateFolderQuery
	6,  // 12: spaceone.api.dashboard.v1.PrivateFolder.stat:input_type -> spaceone.api.dashboard.v1.PrivateFolderStatQuery
	4,  // 13: spaceone.api.dashboard.v1.PrivateFolder.create:output_type -> spaceone.api.dashboard.v1.PrivateFolderInfo
	4,  // 14: spaceone.api.dashboard.v1.PrivateFolder.update:output_type -> spaceone.api.dashboard.v1.PrivateFolderInfo
	10, // 15: spaceone.api.dashboard.v1.PrivateFolder.delete:output_type -> google.protobuf.Empty
	4,  // 16: spaceone.api.dashboard.v1.PrivateFolder.get:output_type -> spaceone.api.dashboard.v1.PrivateFolderInfo
	5,  // 17: spaceone.api.dashboard.v1.PrivateFolder.list:output_type -> spaceone.api.dashboard.v1.PrivateFoldersInfo
	7,  // 18: spaceone.api.dashboard.v1.PrivateFolder.stat:output_type -> google.protobuf.Struct
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_dashboard_v1_private_folder_proto_init() }
func file_spaceone_api_dashboard_v1_private_folder_proto_init() {
	if File_spaceone_api_dashboard_v1_private_folder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_dashboard_v1_private_folder_proto_rawDesc), len(file_spaceone_api_dashboard_v1_private_folder_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_dashboard_v1_private_folder_proto_goTypes,
		DependencyIndexes: file_spaceone_api_dashboard_v1_private_folder_proto_depIdxs,
		MessageInfos:      file_spaceone_api_dashboard_v1_private_folder_proto_msgTypes,
	}.Build()
	File_spaceone_api_dashboard_v1_private_folder_proto = out.File
	file_spaceone_api_dashboard_v1_private_folder_proto_goTypes = nil
	file_spaceone_api_dashboard_v1_private_folder_proto_depIdxs = nil
}
