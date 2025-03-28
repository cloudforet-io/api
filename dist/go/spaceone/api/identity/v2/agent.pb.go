// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/agent.proto

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

type AgentInfo_State int32

const (
	AgentInfo_STATE_NONE AgentInfo_State = 0
	AgentInfo_ENABLED    AgentInfo_State = 1
	AgentInfo_DISABLED   AgentInfo_State = 2
	AgentInfo_EXPIRED    AgentInfo_State = 3
)

// Enum value maps for AgentInfo_State.
var (
	AgentInfo_State_name = map[int32]string{
		0: "STATE_NONE",
		1: "ENABLED",
		2: "DISABLED",
		3: "EXPIRED",
	}
	AgentInfo_State_value = map[string]int32{
		"STATE_NONE": 0,
		"ENABLED":    1,
		"DISABLED":   2,
		"EXPIRED":    3,
	}
)

func (x AgentInfo_State) Enum() *AgentInfo_State {
	p := new(AgentInfo_State)
	*p = x
	return p
}

func (x AgentInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[0].Descriptor()
}

func (AgentInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[0]
}

func (x AgentInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo_State.Descriptor instead.
func (AgentInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2, 0}
}

type AgentInfo_RoleType int32

const (
	AgentInfo_ROLE_TYPE_NONE  AgentInfo_RoleType = 0
	AgentInfo_DOMAIN_ADMIN    AgentInfo_RoleType = 1
	AgentInfo_WORKSPACE_OWNER AgentInfo_RoleType = 2
)

// Enum value maps for AgentInfo_RoleType.
var (
	AgentInfo_RoleType_name = map[int32]string{
		0: "ROLE_TYPE_NONE",
		1: "DOMAIN_ADMIN",
		2: "WORKSPACE_OWNER",
	}
	AgentInfo_RoleType_value = map[string]int32{
		"ROLE_TYPE_NONE":  0,
		"DOMAIN_ADMIN":    1,
		"WORKSPACE_OWNER": 2,
	}
)

func (x AgentInfo_RoleType) Enum() *AgentInfo_RoleType {
	p := new(AgentInfo_RoleType)
	*p = x
	return p
}

func (x AgentInfo_RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo_RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[1].Descriptor()
}

func (AgentInfo_RoleType) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[1]
}

func (x AgentInfo_RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo_RoleType.Descriptor instead.
func (AgentInfo_RoleType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2, 1}
}

type AgentInfo_ResourceGroup int32

const (
	AgentInfo_GROUP_NONE AgentInfo_ResourceGroup = 0
	AgentInfo_DOMAIN     AgentInfo_ResourceGroup = 1
	AgentInfo_WORKSPACE  AgentInfo_ResourceGroup = 2
)

// Enum value maps for AgentInfo_ResourceGroup.
var (
	AgentInfo_ResourceGroup_name = map[int32]string{
		0: "GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	AgentInfo_ResourceGroup_value = map[string]int32{
		"GROUP_NONE": 0,
		"DOMAIN":     1,
		"WORKSPACE":  2,
	}
)

func (x AgentInfo_ResourceGroup) Enum() *AgentInfo_ResourceGroup {
	p := new(AgentInfo_ResourceGroup)
	*p = x
	return p
}

func (x AgentInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[2].Descriptor()
}

func (AgentInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[2]
}

func (x AgentInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentInfo_ResourceGroup.Descriptor instead.
func (AgentInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2, 2}
}

type AgentSearchQuery_State int32

const (
	AgentSearchQuery_STATE_NONE AgentSearchQuery_State = 0
	AgentSearchQuery_ENABLED    AgentSearchQuery_State = 1
	AgentSearchQuery_DISABLED   AgentSearchQuery_State = 2
	AgentSearchQuery_EXPIRED    AgentSearchQuery_State = 3
)

// Enum value maps for AgentSearchQuery_State.
var (
	AgentSearchQuery_State_name = map[int32]string{
		0: "STATE_NONE",
		1: "ENABLED",
		2: "DISABLED",
		3: "EXPIRED",
	}
	AgentSearchQuery_State_value = map[string]int32{
		"STATE_NONE": 0,
		"ENABLED":    1,
		"DISABLED":   2,
		"EXPIRED":    3,
	}
)

func (x AgentSearchQuery_State) Enum() *AgentSearchQuery_State {
	p := new(AgentSearchQuery_State)
	*p = x
	return p
}

func (x AgentSearchQuery_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentSearchQuery_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_agent_proto_enumTypes[3].Descriptor()
}

func (AgentSearchQuery_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_agent_proto_enumTypes[3]
}

func (x AgentSearchQuery_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentSearchQuery_State.Descriptor instead.
func (AgentSearchQuery_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{3, 0}
}

//	{
//	 "service_account_id": "sa-a120f6d21c4e",
//	 "options": {
//	   "cluster_name": "k8s-prd-cluster",
//	   "kube_state_metrics": "false",
//	   "prometheus_node_exporter": "false"
//	 }
//	}
type CreateAgentRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ServiceAccountId string                 `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	Options          *_struct.Struct        `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateAgentRequest) Reset() {
	*x = CreateAgentRequest{}
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentRequest) ProtoMessage() {}

func (x *CreateAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentRequest.ProtoReflect.Descriptor instead.
func (*CreateAgentRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAgentRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CreateAgentRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

//	{
//	 "service_account_id": "sa-a120f6d21c4e"
//	}
type AgentRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	ServiceAccountId string                 `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AgentRequest) Reset() {
	*x = AgentRequest{}
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRequest) ProtoMessage() {}

func (x *AgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRequest.ProtoReflect.Descriptor instead.
func (*AgentRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

//	{
//	   "agent_id": "agent-5cb52dc61c70",
//	   "options": {
//	       "kube_state_metrics": false,
//	       "cluster_name": "k8s-prd-cluster",
//	       "prometheus_node_exporter": false
//	   },
//	   "client_secret": "client_secret_from_app",
//	   "state": "ENABLED",
//	   "is_managed": true,
//	   "role_type": "WORKSPACE_OWNER",
//	   "domain_id": "domain-116226a1516a",
//	   "workspace_id": "workspace-7a0aebcf4eb2",
//	   "project_id": "project-441975c8dfd8",
//	   "service_account_id": "sa-a120f6d21c4e",
//	   "app_id": "app-aa7bf47c98ea",
//	   "role_id": "managed-workspace-owner",
//	   "client_id": "client-36e1034b3512",
//	   "created_at": "2024-11-13T00:34:09.125Z",
//	   "expired_at": "2025-11-13T00:34:09.000Z"
//	}
type AgentInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AgentId          string                 `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Options          *_struct.Struct        `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	ClientSecret     string                 `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	State            AgentInfo_State        `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.identity.v2.AgentInfo_State" json:"state,omitempty"`
	IsManaged        bool                   `protobuf:"varint,5,opt,name=is_managed,json=isManaged,proto3" json:"is_managed,omitempty"`
	RoleType         AgentInfo_RoleType     `protobuf:"varint,6,opt,name=role_type,json=roleType,proto3,enum=spaceone.api.identity.v2.AgentInfo_RoleType" json:"role_type,omitempty"`
	DomainId         string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId        string                 `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceAccountId string                 `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	AppId            string                 `protobuf:"bytes,25,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	RoleId           string                 `protobuf:"bytes,26,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	ClientId         string                 `protobuf:"bytes,27,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CreatedAt        string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiredAt        string                 `protobuf:"bytes,32,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	LastAccessedAt   string                 `protobuf:"bytes,33,opt,name=last_accessed_at,json=lastAccessedAt,proto3" json:"last_accessed_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AgentInfo) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *AgentInfo) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *AgentInfo) GetState() AgentInfo_State {
	if x != nil {
		return x.State
	}
	return AgentInfo_STATE_NONE
}

func (x *AgentInfo) GetIsManaged() bool {
	if x != nil {
		return x.IsManaged
	}
	return false
}

func (x *AgentInfo) GetRoleType() AgentInfo_RoleType {
	if x != nil {
		return x.RoleType
	}
	return AgentInfo_ROLE_TYPE_NONE
}

func (x *AgentInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AgentInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *AgentInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *AgentInfo) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *AgentInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AgentInfo) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AgentInfo) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AgentInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AgentInfo) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *AgentInfo) GetLastAccessedAt() string {
	if x != nil {
		return x.LastAccessedAt
	}
	return ""
}

//	{
//	 "query": {
//	   "page": {
//	     "start":1,
//	     "limit": 10
//	   }
//	 }
//	}
type AgentSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	AgentId string `protobuf:"bytes,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	// +optional
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// +optional
	State         AgentSearchQuery_State `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.identity.v2.AgentSearchQuery_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentSearchQuery) Reset() {
	*x = AgentSearchQuery{}
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentSearchQuery) ProtoMessage() {}

func (x *AgentSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentSearchQuery.ProtoReflect.Descriptor instead.
func (*AgentSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{3}
}

func (x *AgentSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *AgentSearchQuery) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentSearchQuery) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *AgentSearchQuery) GetState() AgentSearchQuery_State {
	if x != nil {
		return x.State
	}
	return AgentSearchQuery_STATE_NONE
}

//	{
//	 "results": [
//	   {
//	     "agent_id": "agent-5cb52dc61c70",
//	     "options": {
//	       "kube_state_metrics": false,
//	       "cluster_name": "k8s-prd-cluster",
//	       "prometheus_node_exporter": false
//	     },
//	     "client_secret": "client_secret_from_app",
//	     "state": "ENABLED",
//	     "is_managed": true,
//	     "role_type": "WORKSPACE_OWNER",
//	     "domain_id": "domain-116226a1516a",
//	     "workspace_id": "workspace-7a0aebcf4eb2",
//	     "project_id": "project-441975c8dfd8",
//	     "service_account_id": "sa-a120f6d21c4e",
//	     "app_id": "app-aa7bf47c98ea",
//	     "role_id": "managed-workspace-owner",
//	     "client_id": "client-36e1034b3512",
//	     "created_at": "2024-11-13T00:34:09.125Z",
//	     "expired_at": "2025-11-13T00:34:09.000Z"
//	   }
//	 ],
//	 "total_count": 1
//	}
type AgentsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*AgentInfo           `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentsInfo) Reset() {
	*x = AgentsInfo{}
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentsInfo) ProtoMessage() {}

func (x *AgentsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentsInfo.ProtoReflect.Descriptor instead.
func (*AgentsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_agent_proto_rawDescGZIP(), []int{4}
}

func (x *AgentsInfo) GetResults() []*AgentInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *AgentsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_identity_v2_agent_proto protoreflect.FileDescriptor

const file_spaceone_api_identity_v2_agent_proto_rawDesc = "" +
	"\n" +
	"$spaceone/api/identity/v2/agent.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a\"spaceone/api/identity/v2/app.proto\"u\n" +
	"\x12CreateAgentRequest\x12,\n" +
	"\x12service_account_id\x18\x01 \x01(\tR\x10serviceAccountId\x121\n" +
	"\aoptions\x18\x02 \x01(\v2\x17.google.protobuf.StructR\aoptions\"<\n" +
	"\fAgentRequest\x12,\n" +
	"\x12service_account_id\x18\x01 \x01(\tR\x10serviceAccountId\"\xaf\x06\n" +
	"\tAgentInfo\x12\x19\n" +
	"\bagent_id\x18\x01 \x01(\tR\aagentId\x121\n" +
	"\aoptions\x18\x02 \x01(\v2\x17.google.protobuf.StructR\aoptions\x12#\n" +
	"\rclient_secret\x18\x03 \x01(\tR\fclientSecret\x12?\n" +
	"\x05state\x18\x04 \x01(\x0e2).spaceone.api.identity.v2.AgentInfo.StateR\x05state\x12\x1d\n" +
	"\n" +
	"is_managed\x18\x05 \x01(\bR\tisManaged\x12I\n" +
	"\trole_type\x18\x06 \x01(\x0e2,.spaceone.api.identity.v2.AgentInfo.RoleTypeR\broleType\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\x12,\n" +
	"\x12service_account_id\x18\x18 \x01(\tR\x10serviceAccountId\x12\x15\n" +
	"\x06app_id\x18\x19 \x01(\tR\x05appId\x12\x17\n" +
	"\arole_id\x18\x1a \x01(\tR\x06roleId\x12\x1b\n" +
	"\tclient_id\x18\x1b \x01(\tR\bclientId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"expired_at\x18  \x01(\tR\texpiredAt\x12(\n" +
	"\x10last_accessed_at\x18! \x01(\tR\x0elastAccessedAt\"?\n" +
	"\x05State\x12\x0e\n" +
	"\n" +
	"STATE_NONE\x10\x00\x12\v\n" +
	"\aENABLED\x10\x01\x12\f\n" +
	"\bDISABLED\x10\x02\x12\v\n" +
	"\aEXPIRED\x10\x03\"E\n" +
	"\bRoleType\x12\x12\n" +
	"\x0eROLE_TYPE_NONE\x10\x00\x12\x10\n" +
	"\fDOMAIN_ADMIN\x10\x01\x12\x13\n" +
	"\x0fWORKSPACE_OWNER\x10\x02\":\n" +
	"\rResourceGroup\x12\x0e\n" +
	"\n" +
	"GROUP_NONE\x10\x00\x12\n" +
	"\n" +
	"\x06DOMAIN\x10\x01\x12\r\n" +
	"\tWORKSPACE\x10\x02\"\x97\x02\n" +
	"\x10AgentSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x19\n" +
	"\bagent_id\x18\x02 \x01(\tR\aagentId\x12,\n" +
	"\x12service_account_id\x18\x03 \x01(\tR\x10serviceAccountId\x12F\n" +
	"\x05state\x18\x04 \x01(\x0e20.spaceone.api.identity.v2.AgentSearchQuery.StateR\x05state\"?\n" +
	"\x05State\x12\x0e\n" +
	"\n" +
	"STATE_NONE\x10\x00\x12\v\n" +
	"\aENABLED\x10\x01\x12\f\n" +
	"\bDISABLED\x10\x02\x12\v\n" +
	"\aEXPIRED\x10\x03\"l\n" +
	"\n" +
	"AgentsInfo\x12=\n" +
	"\aresults\x18\x01 \x03(\v2#.spaceone.api.identity.v2.AgentInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount2\xf2\x06\n" +
	"\x05Agent\x12\x81\x01\n" +
	"\x06create\x12,.spaceone.api.identity.v2.CreateAgentRequest\x1a#.spaceone.api.identity.v2.AgentInfo\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/identity/v2/agent/create\x12{\n" +
	"\x06enable\x12&.spaceone.api.identity.v2.AgentRequest\x1a#.spaceone.api.identity.v2.AgentInfo\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/identity/v2/agent/enable\x12}\n" +
	"\adisable\x12&.spaceone.api.identity.v2.AgentRequest\x1a#.spaceone.api.identity.v2.AgentInfo\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/identity/v2/agent/disable\x12\x83\x01\n" +
	"\n" +
	"regenerate\x12&.spaceone.api.identity.v2.AgentRequest\x1a#.spaceone.api.identity.v2.AgentInfo\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/identity/v2/agent/regenerate\x12n\n" +
	"\x06delete\x12&.spaceone.api.identity.v2.AgentRequest\x1a\x16.google.protobuf.Empty\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/identity/v2/agent/delete\x12u\n" +
	"\x03get\x12&.spaceone.api.identity.v2.AgentRequest\x1a#.spaceone.api.identity.v2.AgentInfo\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/identity/v2/agent/get\x12|\n" +
	"\x04list\x12*.spaceone.api.identity.v2.AgentSearchQuery\x1a$.spaceone.api.identity.v2.AgentsInfo\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\"\x17/identity/v2/agent/listB?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3"

var (
	file_spaceone_api_identity_v2_agent_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_agent_proto_rawDescData []byte
)

func file_spaceone_api_identity_v2_agent_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_agent_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_agent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v2_agent_proto_rawDesc), len(file_spaceone_api_identity_v2_agent_proto_rawDesc)))
	})
	return file_spaceone_api_identity_v2_agent_proto_rawDescData
}

var file_spaceone_api_identity_v2_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_spaceone_api_identity_v2_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_identity_v2_agent_proto_goTypes = []any{
	(AgentInfo_State)(0),         // 0: spaceone.api.identity.v2.AgentInfo.State
	(AgentInfo_RoleType)(0),      // 1: spaceone.api.identity.v2.AgentInfo.RoleType
	(AgentInfo_ResourceGroup)(0), // 2: spaceone.api.identity.v2.AgentInfo.ResourceGroup
	(AgentSearchQuery_State)(0),  // 3: spaceone.api.identity.v2.AgentSearchQuery.State
	(*CreateAgentRequest)(nil),   // 4: spaceone.api.identity.v2.CreateAgentRequest
	(*AgentRequest)(nil),         // 5: spaceone.api.identity.v2.AgentRequest
	(*AgentInfo)(nil),            // 6: spaceone.api.identity.v2.AgentInfo
	(*AgentSearchQuery)(nil),     // 7: spaceone.api.identity.v2.AgentSearchQuery
	(*AgentsInfo)(nil),           // 8: spaceone.api.identity.v2.AgentsInfo
	(*_struct.Struct)(nil),       // 9: google.protobuf.Struct
	(*v2.Query)(nil),             // 10: spaceone.api.core.v2.Query
	(*empty.Empty)(nil),          // 11: google.protobuf.Empty
}
var file_spaceone_api_identity_v2_agent_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.identity.v2.CreateAgentRequest.options:type_name -> google.protobuf.Struct
	9,  // 1: spaceone.api.identity.v2.AgentInfo.options:type_name -> google.protobuf.Struct
	0,  // 2: spaceone.api.identity.v2.AgentInfo.state:type_name -> spaceone.api.identity.v2.AgentInfo.State
	1,  // 3: spaceone.api.identity.v2.AgentInfo.role_type:type_name -> spaceone.api.identity.v2.AgentInfo.RoleType
	10, // 4: spaceone.api.identity.v2.AgentSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	3,  // 5: spaceone.api.identity.v2.AgentSearchQuery.state:type_name -> spaceone.api.identity.v2.AgentSearchQuery.State
	6,  // 6: spaceone.api.identity.v2.AgentsInfo.results:type_name -> spaceone.api.identity.v2.AgentInfo
	4,  // 7: spaceone.api.identity.v2.Agent.create:input_type -> spaceone.api.identity.v2.CreateAgentRequest
	5,  // 8: spaceone.api.identity.v2.Agent.enable:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 9: spaceone.api.identity.v2.Agent.disable:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 10: spaceone.api.identity.v2.Agent.regenerate:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 11: spaceone.api.identity.v2.Agent.delete:input_type -> spaceone.api.identity.v2.AgentRequest
	5,  // 12: spaceone.api.identity.v2.Agent.get:input_type -> spaceone.api.identity.v2.AgentRequest
	7,  // 13: spaceone.api.identity.v2.Agent.list:input_type -> spaceone.api.identity.v2.AgentSearchQuery
	6,  // 14: spaceone.api.identity.v2.Agent.create:output_type -> spaceone.api.identity.v2.AgentInfo
	6,  // 15: spaceone.api.identity.v2.Agent.enable:output_type -> spaceone.api.identity.v2.AgentInfo
	6,  // 16: spaceone.api.identity.v2.Agent.disable:output_type -> spaceone.api.identity.v2.AgentInfo
	6,  // 17: spaceone.api.identity.v2.Agent.regenerate:output_type -> spaceone.api.identity.v2.AgentInfo
	11, // 18: spaceone.api.identity.v2.Agent.delete:output_type -> google.protobuf.Empty
	6,  // 19: spaceone.api.identity.v2.Agent.get:output_type -> spaceone.api.identity.v2.AgentInfo
	8,  // 20: spaceone.api.identity.v2.Agent.list:output_type -> spaceone.api.identity.v2.AgentsInfo
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_agent_proto_init() }
func file_spaceone_api_identity_v2_agent_proto_init() {
	if File_spaceone_api_identity_v2_agent_proto != nil {
		return
	}
	file_spaceone_api_identity_v2_app_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v2_agent_proto_rawDesc), len(file_spaceone_api_identity_v2_agent_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_agent_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_agent_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_agent_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_agent_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_agent_proto = out.File
	file_spaceone_api_identity_v2_agent_proto_goTypes = nil
	file_spaceone_api_identity_v2_agent_proto_depIdxs = nil
}
