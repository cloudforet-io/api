// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/user_group.proto

package v2

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

type CreateUserGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserGroupRequest) Reset() {
	*x = CreateUserGroupRequest{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserGroupRequest) ProtoMessage() {}

func (x *CreateUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateUserGroupRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateUserGroupRequest struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	UserGroupId string                 `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserGroupRequest) Reset() {
	*x = UpdateUserGroupRequest{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserGroupRequest) ProtoMessage() {}

func (x *UpdateUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserGroupRequest) GetUserGroupId() string {
	if x != nil {
		return x.UserGroupId
	}
	return ""
}

func (x *UpdateUserGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateUserGroupRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UserGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserGroupId   string                 `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGroupRequest) Reset() {
	*x = UserGroupRequest{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupRequest) ProtoMessage() {}

func (x *UserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupRequest.ProtoReflect.Descriptor instead.
func (*UserGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{2}
}

func (x *UserGroupRequest) GetUserGroupId() string {
	if x != nil {
		return x.UserGroupId
	}
	return ""
}

type UsersUserGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserGroupId   string                 `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
	Users         []string               `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersUserGroupRequest) Reset() {
	*x = UsersUserGroupRequest{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersUserGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersUserGroupRequest) ProtoMessage() {}

func (x *UsersUserGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersUserGroupRequest.ProtoReflect.Descriptor instead.
func (*UsersUserGroupRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{3}
}

func (x *UsersUserGroupRequest) GetUserGroupId() string {
	if x != nil {
		return x.UserGroupId
	}
	return ""
}

func (x *UsersUserGroupRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserGroupInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserGroupId   string                 `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Users         []string               `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	Tags          *_struct.Struct        `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGroupInfo) Reset() {
	*x = UserGroupInfo{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupInfo) ProtoMessage() {}

func (x *UserGroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupInfo.ProtoReflect.Descriptor instead.
func (*UserGroupInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{4}
}

func (x *UserGroupInfo) GetUserGroupId() string {
	if x != nil {
		return x.UserGroupId
	}
	return ""
}

func (x *UserGroupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserGroupInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserGroupInfo) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UserGroupInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UserGroupInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *UserGroupInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *UserGroupInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UserGroupSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	UserGroupId string `protobuf:"bytes,2,opt,name=user_group_id,json=userGroupId,proto3" json:"user_group_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	UserId        string `protobuf:"bytes,21,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGroupSearchQuery) Reset() {
	*x = UserGroupSearchQuery{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupSearchQuery) ProtoMessage() {}

func (x *UserGroupSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupSearchQuery.ProtoReflect.Descriptor instead.
func (*UserGroupSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{5}
}

func (x *UserGroupSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *UserGroupSearchQuery) GetUserGroupId() string {
	if x != nil {
		return x.UserGroupId
	}
	return ""
}

func (x *UserGroupSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserGroupSearchQuery) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserGroupsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*UserGroupInfo       `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGroupsInfo) Reset() {
	*x = UserGroupsInfo{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupsInfo) ProtoMessage() {}

func (x *UserGroupsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupsInfo.ProtoReflect.Descriptor instead.
func (*UserGroupsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{6}
}

func (x *UserGroupsInfo) GetResults() []*UserGroupInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *UserGroupsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type UserGroupStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserGroupStatQuery) Reset() {
	*x = UserGroupStatQuery{}
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserGroupStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupStatQuery) ProtoMessage() {}

func (x *UserGroupStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_group_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupStatQuery.ProtoReflect.Descriptor instead.
func (*UserGroupStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP(), []int{7}
}

func (x *UserGroupStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_identity_v2_user_group_proto protoreflect.FileDescriptor

const file_spaceone_api_identity_v2_user_group_proto_rawDesc = "" +
	"\n" +
	")spaceone/api/identity/v2/user_group.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"{\n" +
	"\x16CreateUserGroupRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12+\n" +
	"\x04tags\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04tags\"\x9f\x01\n" +
	"\x16UpdateUserGroupRequest\x12\"\n" +
	"\ruser_group_id\x18\x01 \x01(\tR\vuserGroupId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12+\n" +
	"\x04tags\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x04tags\"6\n" +
	"\x10UserGroupRequest\x12\"\n" +
	"\ruser_group_id\x18\x01 \x01(\tR\vuserGroupId\"Q\n" +
	"\x15UsersUserGroupRequest\x12\"\n" +
	"\ruser_group_id\x18\x01 \x01(\tR\vuserGroupId\x12\x14\n" +
	"\x05users\x18\x02 \x03(\tR\x05users\"\x8b\x02\n" +
	"\rUserGroupInfo\x12\"\n" +
	"\ruser_group_id\x18\x01 \x01(\tR\vuserGroupId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05users\x18\x04 \x03(\tR\x05users\x12+\n" +
	"\x04tags\x18\x05 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\"\x9a\x01\n" +
	"\x14UserGroupSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\"\n" +
	"\ruser_group_id\x18\x02 \x01(\tR\vuserGroupId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x17\n" +
	"\auser_id\x18\x15 \x01(\tR\x06userId\"t\n" +
	"\x0eUserGroupsInfo\x12A\n" +
	"\aresults\x18\x01 \x03(\v2'.spaceone.api.identity.v2.UserGroupInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"Q\n" +
	"\x12UserGroupStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\xe1\b\n" +
	"\tUserGroup\x12\x8e\x01\n" +
	"\x06create\x120.spaceone.api.identity.v2.CreateUserGroupRequest\x1a'.spaceone.api.identity.v2.UserGroupInfo\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/identity/v2/user-group/create\x12\x8e\x01\n" +
	"\x06update\x120.spaceone.api.identity.v2.UpdateUserGroupRequest\x1a'.spaceone.api.identity.v2.UserGroupInfo\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/identity/v2/user-group/update\x12w\n" +
	"\x06delete\x12*.spaceone.api.identity.v2.UserGroupRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/identity/v2/user-group/delete\x12\x93\x01\n" +
	"\tadd_users\x12/.spaceone.api.identity.v2.UsersUserGroupRequest\x1a'.spaceone.api.identity.v2.UserGroupInfo\",\x82\xd3\xe4\x93\x02&:\x01*\"!/identity/v2/user-group/add-users\x12\x99\x01\n" +
	"\fremove_users\x12/.spaceone.api.identity.v2.UsersUserGroupRequest\x1a'.spaceone.api.identity.v2.UserGroupInfo\"/\x82\xd3\xe4\x93\x02):\x01*\"$/identity/v2/user-group/remove-users\x12\x82\x01\n" +
	"\x03get\x12*.spaceone.api.identity.v2.UserGroupRequest\x1a'.spaceone.api.identity.v2.UserGroupInfo\"&\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/identity/v2/user-group/get\x12\x89\x01\n" +
	"\x04list\x12..spaceone.api.identity.v2.UserGroupSearchQuery\x1a(.spaceone.api.identity.v2.UserGroupsInfo\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/identity/v2/user-group/list\x12v\n" +
	"\x04stat\x12,.spaceone.api.identity.v2.UserGroupStatQuery\x1a\x17.google.protobuf.Struct\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/identity/v1/user-group/statB?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3"

var (
	file_spaceone_api_identity_v2_user_group_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_user_group_proto_rawDescData []byte
)

func file_spaceone_api_identity_v2_user_group_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_user_group_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_user_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v2_user_group_proto_rawDesc), len(file_spaceone_api_identity_v2_user_group_proto_rawDesc)))
	})
	return file_spaceone_api_identity_v2_user_group_proto_rawDescData
}

var file_spaceone_api_identity_v2_user_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_identity_v2_user_group_proto_goTypes = []any{
	(*CreateUserGroupRequest)(nil), // 0: spaceone.api.identity.v2.CreateUserGroupRequest
	(*UpdateUserGroupRequest)(nil), // 1: spaceone.api.identity.v2.UpdateUserGroupRequest
	(*UserGroupRequest)(nil),       // 2: spaceone.api.identity.v2.UserGroupRequest
	(*UsersUserGroupRequest)(nil),  // 3: spaceone.api.identity.v2.UsersUserGroupRequest
	(*UserGroupInfo)(nil),          // 4: spaceone.api.identity.v2.UserGroupInfo
	(*UserGroupSearchQuery)(nil),   // 5: spaceone.api.identity.v2.UserGroupSearchQuery
	(*UserGroupsInfo)(nil),         // 6: spaceone.api.identity.v2.UserGroupsInfo
	(*UserGroupStatQuery)(nil),     // 7: spaceone.api.identity.v2.UserGroupStatQuery
	(*_struct.Struct)(nil),         // 8: google.protobuf.Struct
	(*v2.Query)(nil),               // 9: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),     // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_spaceone_api_identity_v2_user_group_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.identity.v2.CreateUserGroupRequest.tags:type_name -> google.protobuf.Struct
	8,  // 1: spaceone.api.identity.v2.UpdateUserGroupRequest.tags:type_name -> google.protobuf.Struct
	8,  // 2: spaceone.api.identity.v2.UserGroupInfo.tags:type_name -> google.protobuf.Struct
	9,  // 3: spaceone.api.identity.v2.UserGroupSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	4,  // 4: spaceone.api.identity.v2.UserGroupsInfo.results:type_name -> spaceone.api.identity.v2.UserGroupInfo
	10, // 5: spaceone.api.identity.v2.UserGroupStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 6: spaceone.api.identity.v2.UserGroup.create:input_type -> spaceone.api.identity.v2.CreateUserGroupRequest
	1,  // 7: spaceone.api.identity.v2.UserGroup.update:input_type -> spaceone.api.identity.v2.UpdateUserGroupRequest
	2,  // 8: spaceone.api.identity.v2.UserGroup.delete:input_type -> spaceone.api.identity.v2.UserGroupRequest
	3,  // 9: spaceone.api.identity.v2.UserGroup.add_users:input_type -> spaceone.api.identity.v2.UsersUserGroupRequest
	3,  // 10: spaceone.api.identity.v2.UserGroup.remove_users:input_type -> spaceone.api.identity.v2.UsersUserGroupRequest
	2,  // 11: spaceone.api.identity.v2.UserGroup.get:input_type -> spaceone.api.identity.v2.UserGroupRequest
	5,  // 12: spaceone.api.identity.v2.UserGroup.list:input_type -> spaceone.api.identity.v2.UserGroupSearchQuery
	7,  // 13: spaceone.api.identity.v2.UserGroup.stat:input_type -> spaceone.api.identity.v2.UserGroupStatQuery
	4,  // 14: spaceone.api.identity.v2.UserGroup.create:output_type -> spaceone.api.identity.v2.UserGroupInfo
	4,  // 15: spaceone.api.identity.v2.UserGroup.update:output_type -> spaceone.api.identity.v2.UserGroupInfo
	11, // 16: spaceone.api.identity.v2.UserGroup.delete:output_type -> google.protobuf.Empty
	4,  // 17: spaceone.api.identity.v2.UserGroup.add_users:output_type -> spaceone.api.identity.v2.UserGroupInfo
	4,  // 18: spaceone.api.identity.v2.UserGroup.remove_users:output_type -> spaceone.api.identity.v2.UserGroupInfo
	4,  // 19: spaceone.api.identity.v2.UserGroup.get:output_type -> spaceone.api.identity.v2.UserGroupInfo
	6,  // 20: spaceone.api.identity.v2.UserGroup.list:output_type -> spaceone.api.identity.v2.UserGroupsInfo
	8,  // 21: spaceone.api.identity.v2.UserGroup.stat:output_type -> google.protobuf.Struct
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_user_group_proto_init() }
func file_spaceone_api_identity_v2_user_group_proto_init() {
	if File_spaceone_api_identity_v2_user_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v2_user_group_proto_rawDesc), len(file_spaceone_api_identity_v2_user_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_user_group_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_user_group_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v2_user_group_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_user_group_proto = out.File
	file_spaceone_api_identity_v2_user_group_proto_goTypes = nil
	file_spaceone_api_identity_v2_user_group_proto_depIdxs = nil
}
