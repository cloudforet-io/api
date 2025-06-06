// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/file_manager/v1/file.proto

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

type CreateFileRequest_ResourceGroup int32

const (
	CreateFileRequest_RESOURCE_GROUP_NONE CreateFileRequest_ResourceGroup = 0
	CreateFileRequest_SYSTEM              CreateFileRequest_ResourceGroup = 1
	CreateFileRequest_DOMAIN              CreateFileRequest_ResourceGroup = 2
	CreateFileRequest_WORKSPACE           CreateFileRequest_ResourceGroup = 3
	CreateFileRequest_PROJECT             CreateFileRequest_ResourceGroup = 4
)

// Enum value maps for CreateFileRequest_ResourceGroup.
var (
	CreateFileRequest_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "SYSTEM",
		2: "DOMAIN",
		3: "WORKSPACE",
		4: "PROJECT",
	}
	CreateFileRequest_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"SYSTEM":              1,
		"DOMAIN":              2,
		"WORKSPACE":           3,
		"PROJECT":             4,
	}
)

func (x CreateFileRequest_ResourceGroup) Enum() *CreateFileRequest_ResourceGroup {
	p := new(CreateFileRequest_ResourceGroup)
	*p = x
	return p
}

func (x CreateFileRequest_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateFileRequest_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_file_manager_v1_file_proto_enumTypes[0].Descriptor()
}

func (CreateFileRequest_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_file_manager_v1_file_proto_enumTypes[0]
}

func (x CreateFileRequest_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateFileRequest_ResourceGroup.Descriptor instead.
func (CreateFileRequest_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{1, 0}
}

type FileInfo_ResourceGroup int32

const (
	FileInfo_RESOURCE_GROUP_NONE FileInfo_ResourceGroup = 0
	FileInfo_SYSTEM              FileInfo_ResourceGroup = 1
	FileInfo_DOMAIN              FileInfo_ResourceGroup = 2
	FileInfo_WORKSPACE           FileInfo_ResourceGroup = 3
	FileInfo_PROJECT             FileInfo_ResourceGroup = 4
)

// Enum value maps for FileInfo_ResourceGroup.
var (
	FileInfo_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "SYSTEM",
		2: "DOMAIN",
		3: "WORKSPACE",
		4: "PROJECT",
	}
	FileInfo_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"SYSTEM":              1,
		"DOMAIN":              2,
		"WORKSPACE":           3,
		"PROJECT":             4,
	}
)

func (x FileInfo_ResourceGroup) Enum() *FileInfo_ResourceGroup {
	p := new(FileInfo_ResourceGroup)
	*p = x
	return p
}

func (x FileInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_file_manager_v1_file_proto_enumTypes[1].Descriptor()
}

func (FileInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_file_manager_v1_file_proto_enumTypes[1]
}

func (x FileInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileInfo_ResourceGroup.Descriptor instead.
func (FileInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{5, 0}
}

type FileReference struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceType  string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId    string                 `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileReference) Reset() {
	*x = FileReference{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileReference) ProtoMessage() {}

func (x *FileReference) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileReference.ProtoReflect.Descriptor instead.
func (*FileReference) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileReference) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *FileReference) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type CreateFileRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	Reference     *FileReference                  `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	ResourceGroup CreateFileRequest_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.file_manager.v1.CreateFileRequest_ResourceGroup" json:"resource_group,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId     string `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFileRequest) Reset() {
	*x = CreateFileRequest{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequest) ProtoMessage() {}

func (x *CreateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequest.ProtoReflect.Descriptor instead.
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFileRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateFileRequest) GetReference() *FileReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *CreateFileRequest) GetResourceGroup() CreateFileRequest_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return CreateFileRequest_RESOURCE_GROUP_NONE
}

func (x *CreateFileRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CreateFileRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreateFileRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type UpdateFileRequest struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	FileId string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	Reference *FileReference `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	// +optional
	ProjectId     string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UpdateFileRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateFileRequest) GetReference() *FileReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *UpdateFileRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type FileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{3}
}

func (x *FileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type FileSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// / +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	ResourceType string `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// +optional
	ResourceId string `protobuf:"bytes,5,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId     string `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileSearchQuery) Reset() {
	*x = FileSearchQuery{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSearchQuery) ProtoMessage() {}

func (x *FileSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSearchQuery.ProtoReflect.Descriptor instead.
func (*FileSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{4}
}

func (x *FileSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *FileSearchQuery) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileSearchQuery) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *FileSearchQuery) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *FileSearchQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *FileSearchQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *FileSearchQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DownloadUrl   string                 `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	Reference     *FileReference         `protobuf:"bytes,4,opt,name=reference,proto3" json:"reference,omitempty"`
	Tags          *_struct.Struct        `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	ResourceGroup FileInfo_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.file_manager.v1.FileInfo_ResourceGroup" json:"resource_group,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId     string                 `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{5}
}

func (x *FileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *FileInfo) GetReference() *FileReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *FileInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *FileInfo) GetResourceGroup() FileInfo_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return FileInfo_RESOURCE_GROUP_NONE
}

func (x *FileInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *FileInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *FileInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *FileInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type FilesInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*FileInfo            `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilesInfo) Reset() {
	*x = FilesInfo{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesInfo) ProtoMessage() {}

func (x *FilesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesInfo.ProtoReflect.Descriptor instead.
func (*FilesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{6}
}

func (x *FilesInfo) GetResults() []*FileInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *FilesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type FileStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileStatQuery) Reset() {
	*x = FileStatQuery{}
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStatQuery) ProtoMessage() {}

func (x *FileStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_file_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStatQuery.ProtoReflect.Descriptor instead.
func (*FileStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP(), []int{7}
}

func (x *FileStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_file_manager_v1_file_proto protoreflect.FileDescriptor

const file_spaceone_api_file_manager_v1_file_proto_rawDesc = "" +
	"\n" +
	"'spaceone/api/file_manager/v1/file.proto\x12\x1cspaceone.api.file_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"U\n" +
	"\rFileReference\x12#\n" +
	"\rresource_type\x18\x01 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x02 \x01(\tR\n" +
	"resourceId\"\xc2\x03\n" +
	"\x11CreateFileRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12+\n" +
	"\x04tags\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12I\n" +
	"\treference\x18\x03 \x01(\v2+.spaceone.api.file_manager.v1.FileReferenceR\treference\x12d\n" +
	"\x0eresource_group\x18\x14 \x01(\x0e2=.spaceone.api.file_manager.v1.CreateFileRequest.ResourceGroupR\rresourceGroup\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\"\\\n" +
	"\rResourceGroup\x12\x17\n" +
	"\x13RESOURCE_GROUP_NONE\x10\x00\x12\n" +
	"\n" +
	"\x06SYSTEM\x10\x01\x12\n" +
	"\n" +
	"\x06DOMAIN\x10\x02\x12\r\n" +
	"\tWORKSPACE\x10\x03\x12\v\n" +
	"\aPROJECT\x10\x04\"\xc3\x01\n" +
	"\x11UpdateFileRequest\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12+\n" +
	"\x04tags\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12I\n" +
	"\treference\x18\x03 \x01(\v2+.spaceone.api.file_manager.v1.FileReferenceR\treference\x12\x1d\n" +
	"\n" +
	"project_id\x18\x04 \x01(\tR\tprojectId\"&\n" +
	"\vFileRequest\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\"\x96\x02\n" +
	"\x0fFileSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x17\n" +
	"\afile_id\x18\x02 \x01(\tR\x06fileId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12#\n" +
	"\rresource_type\x18\x04 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x05 \x01(\tR\n" +
	"resourceId\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\"\x8b\x04\n" +
	"\bFileInfo\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12!\n" +
	"\fdownload_url\x18\x03 \x01(\tR\vdownloadUrl\x12I\n" +
	"\treference\x18\x04 \x01(\v2+.spaceone.api.file_manager.v1.FileReferenceR\treference\x12+\n" +
	"\x04tags\x18\x05 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12[\n" +
	"\x0eresource_group\x18\x14 \x01(\x0e24.spaceone.api.file_manager.v1.FileInfo.ResourceGroupR\rresourceGroup\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\"\\\n" +
	"\rResourceGroup\x12\x17\n" +
	"\x13RESOURCE_GROUP_NONE\x10\x00\x12\n" +
	"\n" +
	"\x06SYSTEM\x10\x01\x12\n" +
	"\n" +
	"\x06DOMAIN\x10\x02\x12\r\n" +
	"\tWORKSPACE\x10\x03\x12\v\n" +
	"\aPROJECT\x10\x04\"n\n" +
	"\tFilesInfo\x12@\n" +
	"\aresults\x18\x01 \x03(\v2&.spaceone.api.file_manager.v1.FileInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"L\n" +
	"\rFileStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\x86\x05\n" +
	"\x04File\x12\x8a\x01\n" +
	"\x06update\x12/.spaceone.api.file_manager.v1.UpdateFileRequest\x1a&.spaceone.api.file_manager.v1.FileInfo\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/file-manager/v1/file/update\x12t\n" +
	"\x06delete\x12).spaceone.api.file_manager.v1.FileRequest\x1a\x16.google.protobuf.Empty\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/file-manager/v1/file/delete\x12~\n" +
	"\x03get\x12).spaceone.api.file_manager.v1.FileRequest\x1a&.spaceone.api.file_manager.v1.FileInfo\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/file-manager/v1/file/get\x12\x85\x01\n" +
	"\x04list\x12-.spaceone.api.file_manager.v1.FileSearchQuery\x1a'.spaceone.api.file_manager.v1.FilesInfo\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/file-manager/v1/file/list\x12s\n" +
	"\x04stat\x12+.spaceone.api.file_manager.v1.FileStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/file-manager/v1/file/statBCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/file_manager/v1b\x06proto3"

var (
	file_spaceone_api_file_manager_v1_file_proto_rawDescOnce sync.Once
	file_spaceone_api_file_manager_v1_file_proto_rawDescData []byte
)

func file_spaceone_api_file_manager_v1_file_proto_rawDescGZIP() []byte {
	file_spaceone_api_file_manager_v1_file_proto_rawDescOnce.Do(func() {
		file_spaceone_api_file_manager_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_file_manager_v1_file_proto_rawDesc), len(file_spaceone_api_file_manager_v1_file_proto_rawDesc)))
	})
	return file_spaceone_api_file_manager_v1_file_proto_rawDescData
}

var file_spaceone_api_file_manager_v1_file_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_file_manager_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_file_manager_v1_file_proto_goTypes = []any{
	(CreateFileRequest_ResourceGroup)(0), // 0: spaceone.api.file_manager.v1.CreateFileRequest.ResourceGroup
	(FileInfo_ResourceGroup)(0),          // 1: spaceone.api.file_manager.v1.FileInfo.ResourceGroup
	(*FileReference)(nil),                // 2: spaceone.api.file_manager.v1.FileReference
	(*CreateFileRequest)(nil),            // 3: spaceone.api.file_manager.v1.CreateFileRequest
	(*UpdateFileRequest)(nil),            // 4: spaceone.api.file_manager.v1.UpdateFileRequest
	(*FileRequest)(nil),                  // 5: spaceone.api.file_manager.v1.FileRequest
	(*FileSearchQuery)(nil),              // 6: spaceone.api.file_manager.v1.FileSearchQuery
	(*FileInfo)(nil),                     // 7: spaceone.api.file_manager.v1.FileInfo
	(*FilesInfo)(nil),                    // 8: spaceone.api.file_manager.v1.FilesInfo
	(*FileStatQuery)(nil),                // 9: spaceone.api.file_manager.v1.FileStatQuery
	(*_struct.Struct)(nil),               // 10: google.protobuf.Struct
	(*v2.Query)(nil),                     // 11: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),           // 12: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                  // 13: google.protobuf.Empty
}
var file_spaceone_api_file_manager_v1_file_proto_depIdxs = []int32{
	10, // 0: spaceone.api.file_manager.v1.CreateFileRequest.tags:type_name -> google.protobuf.Struct
	2,  // 1: spaceone.api.file_manager.v1.CreateFileRequest.reference:type_name -> spaceone.api.file_manager.v1.FileReference
	0,  // 2: spaceone.api.file_manager.v1.CreateFileRequest.resource_group:type_name -> spaceone.api.file_manager.v1.CreateFileRequest.ResourceGroup
	10, // 3: spaceone.api.file_manager.v1.UpdateFileRequest.tags:type_name -> google.protobuf.Struct
	2,  // 4: spaceone.api.file_manager.v1.UpdateFileRequest.reference:type_name -> spaceone.api.file_manager.v1.FileReference
	11, // 5: spaceone.api.file_manager.v1.FileSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	2,  // 6: spaceone.api.file_manager.v1.FileInfo.reference:type_name -> spaceone.api.file_manager.v1.FileReference
	10, // 7: spaceone.api.file_manager.v1.FileInfo.tags:type_name -> google.protobuf.Struct
	1,  // 8: spaceone.api.file_manager.v1.FileInfo.resource_group:type_name -> spaceone.api.file_manager.v1.FileInfo.ResourceGroup
	7,  // 9: spaceone.api.file_manager.v1.FilesInfo.results:type_name -> spaceone.api.file_manager.v1.FileInfo
	12, // 10: spaceone.api.file_manager.v1.FileStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	4,  // 11: spaceone.api.file_manager.v1.File.update:input_type -> spaceone.api.file_manager.v1.UpdateFileRequest
	5,  // 12: spaceone.api.file_manager.v1.File.delete:input_type -> spaceone.api.file_manager.v1.FileRequest
	5,  // 13: spaceone.api.file_manager.v1.File.get:input_type -> spaceone.api.file_manager.v1.FileRequest
	6,  // 14: spaceone.api.file_manager.v1.File.list:input_type -> spaceone.api.file_manager.v1.FileSearchQuery
	9,  // 15: spaceone.api.file_manager.v1.File.stat:input_type -> spaceone.api.file_manager.v1.FileStatQuery
	7,  // 16: spaceone.api.file_manager.v1.File.update:output_type -> spaceone.api.file_manager.v1.FileInfo
	13, // 17: spaceone.api.file_manager.v1.File.delete:output_type -> google.protobuf.Empty
	7,  // 18: spaceone.api.file_manager.v1.File.get:output_type -> spaceone.api.file_manager.v1.FileInfo
	8,  // 19: spaceone.api.file_manager.v1.File.list:output_type -> spaceone.api.file_manager.v1.FilesInfo
	10, // 20: spaceone.api.file_manager.v1.File.stat:output_type -> google.protobuf.Struct
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_spaceone_api_file_manager_v1_file_proto_init() }
func file_spaceone_api_file_manager_v1_file_proto_init() {
	if File_spaceone_api_file_manager_v1_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_file_manager_v1_file_proto_rawDesc), len(file_spaceone_api_file_manager_v1_file_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_file_manager_v1_file_proto_goTypes,
		DependencyIndexes: file_spaceone_api_file_manager_v1_file_proto_depIdxs,
		EnumInfos:         file_spaceone_api_file_manager_v1_file_proto_enumTypes,
		MessageInfos:      file_spaceone_api_file_manager_v1_file_proto_msgTypes,
	}.Build()
	File_spaceone_api_file_manager_v1_file_proto = out.File
	file_spaceone_api_file_manager_v1_file_proto_goTypes = nil
	file_spaceone_api_file_manager_v1_file_proto_depIdxs = nil
}
