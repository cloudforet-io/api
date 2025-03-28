// A DataSourceAccount is a resource that for routing cost data from a specific account to a workspace, project, service account.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/data_source_account.proto

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

//	{
//	   "data_source_id": "ds-faaa11aa1111",
//	   "account_id": "111069360300",
//	   "workspace_id": "ws-aaaa11aa1111"
//	}
type UpdateDataSourceAccountRequest struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	DataSourceId string                 `protobuf:"bytes,1,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// account_id is the unique identifier of each CSP account.(e.g. Azure Tenant ID)
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// +optional
	WorkspaceId   string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDataSourceAccountRequest) Reset() {
	*x = UpdateDataSourceAccountRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDataSourceAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDataSourceAccountRequest) ProtoMessage() {}

func (x *UpdateDataSourceAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDataSourceAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateDataSourceAccountRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateDataSourceAccountRequest) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *UpdateDataSourceAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *UpdateDataSourceAccountRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

//	{
//	   "data_source_id": "ds-faaa11aa1111",
//	   "account_id": "111069360300",
//	   "sync_state": true
//	}
type ResetDataSourceAccountRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// data_source_id is the unique identifier of each data source.
	DataSourceId string `protobuf:"bytes,1,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// +optional
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// if sync_state is true, it will reset the sync state of the data source account.
	// +optional
	ResetSync     bool `protobuf:"varint,3,opt,name=reset_sync,json=resetSync,proto3" json:"reset_sync,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResetDataSourceAccountRequest) Reset() {
	*x = ResetDataSourceAccountRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetDataSourceAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetDataSourceAccountRequest) ProtoMessage() {}

func (x *ResetDataSourceAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetDataSourceAccountRequest.ProtoReflect.Descriptor instead.
func (*ResetDataSourceAccountRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{1}
}

func (x *ResetDataSourceAccountRequest) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *ResetDataSourceAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *ResetDataSourceAccountRequest) GetResetSync() bool {
	if x != nil {
		return x.ResetSync
	}
	return false
}

type DataSourceAccountRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	DataSourceAccountId string                 `protobuf:"bytes,1,opt,name=data_source_account_id,json=dataSourceAccountId,proto3" json:"data_source_account_id,omitempty"`
	AccountId           string                 `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *DataSourceAccountRequest) Reset() {
	*x = DataSourceAccountRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceAccountRequest) ProtoMessage() {}

func (x *DataSourceAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceAccountRequest.ProtoReflect.Descriptor instead.
func (*DataSourceAccountRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{2}
}

func (x *DataSourceAccountRequest) GetDataSourceAccountId() string {
	if x != nil {
		return x.DataSourceAccountId
	}
	return ""
}

func (x *DataSourceAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

//	{
//	   "query": {}
//	}
type DataSourceAccountQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	DataSourceId string `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// +optional
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,22,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	ServiceAccountId string `protobuf:"bytes,23,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DataSourceAccountQuery) Reset() {
	*x = DataSourceAccountQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceAccountQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceAccountQuery) ProtoMessage() {}

func (x *DataSourceAccountQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceAccountQuery.ProtoReflect.Descriptor instead.
func (*DataSourceAccountQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{3}
}

func (x *DataSourceAccountQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *DataSourceAccountQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *DataSourceAccountQuery) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *DataSourceAccountQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *DataSourceAccountQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *DataSourceAccountQuery) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

type DataSourceAccountInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	DataSourceId  string                 `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	IsSync        bool                   `protobuf:"varint,4,opt,name=is_sync,json=isSync,proto3" json:"is_sync,omitempty"`
	IsLinked      bool                   `protobuf:"varint,5,opt,name=is_linked,json=isLinked,proto3" json:"is_linked,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	VWorkspaceId  string                 `protobuf:"bytes,23,opt,name=v_workspace_id,json=vWorkspaceId,proto3" json:"v_workspace_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	LastSyncedAt  string                 `protobuf:"bytes,33,opt,name=last_synced_at,json=lastSyncedAt,proto3" json:"last_synced_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataSourceAccountInfo) Reset() {
	*x = DataSourceAccountInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceAccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceAccountInfo) ProtoMessage() {}

func (x *DataSourceAccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceAccountInfo.ProtoReflect.Descriptor instead.
func (*DataSourceAccountInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{4}
}

func (x *DataSourceAccountInfo) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *DataSourceAccountInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *DataSourceAccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSourceAccountInfo) GetIsSync() bool {
	if x != nil {
		return x.IsSync
	}
	return false
}

func (x *DataSourceAccountInfo) GetIsLinked() bool {
	if x != nil {
		return x.IsLinked
	}
	return false
}

func (x *DataSourceAccountInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *DataSourceAccountInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *DataSourceAccountInfo) GetVWorkspaceId() string {
	if x != nil {
		return x.VWorkspaceId
	}
	return ""
}

func (x *DataSourceAccountInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DataSourceAccountInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DataSourceAccountInfo) GetLastSyncedAt() string {
	if x != nil {
		return x.LastSyncedAt
	}
	return ""
}

type DataSourceAccountsInfo struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Results       []*DataSourceAccountInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                    `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataSourceAccountsInfo) Reset() {
	*x = DataSourceAccountsInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceAccountsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceAccountsInfo) ProtoMessage() {}

func (x *DataSourceAccountsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceAccountsInfo.ProtoReflect.Descriptor instead.
func (*DataSourceAccountsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{5}
}

func (x *DataSourceAccountsInfo) GetResults() []*DataSourceAccountInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *DataSourceAccountsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type DataSourceAccountAnalyzeQuery struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Query         *v2.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DataSourceId  string                     `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	AccountId     string                     `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	WorkspaceId   string                     `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataSourceAccountAnalyzeQuery) Reset() {
	*x = DataSourceAccountAnalyzeQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceAccountAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceAccountAnalyzeQuery) ProtoMessage() {}

func (x *DataSourceAccountAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceAccountAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*DataSourceAccountAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{6}
}

func (x *DataSourceAccountAnalyzeQuery) GetQuery() *v2.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *DataSourceAccountAnalyzeQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *DataSourceAccountAnalyzeQuery) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *DataSourceAccountAnalyzeQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type DataSourceAccountStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataSourceAccountStatQuery) Reset() {
	*x = DataSourceAccountStatQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSourceAccountStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceAccountStatQuery) ProtoMessage() {}

func (x *DataSourceAccountStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceAccountStatQuery.ProtoReflect.Descriptor instead.
func (*DataSourceAccountStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP(), []int{7}
}

func (x *DataSourceAccountStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_cost_analysis_v1_data_source_account_proto protoreflect.FileDescriptor

const file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDesc = "" +
	"\n" +
	"7spaceone/api/cost_analysis/v1/data_source_account.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x88\x01\n" +
	"\x1eUpdateDataSourceAccountRequest\x12$\n" +
	"\x0edata_source_id\x18\x01 \x01(\tR\fdataSourceId\x12\x1d\n" +
	"\n" +
	"account_id\x18\x02 \x01(\tR\taccountId\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\"\x83\x01\n" +
	"\x1dResetDataSourceAccountRequest\x12$\n" +
	"\x0edata_source_id\x18\x01 \x01(\tR\fdataSourceId\x12\x1d\n" +
	"\n" +
	"account_id\x18\x02 \x01(\tR\taccountId\x12\x1d\n" +
	"\n" +
	"reset_sync\x18\x03 \x01(\bR\tresetSync\"n\n" +
	"\x18DataSourceAccountRequest\x123\n" +
	"\x16data_source_account_id\x18\x01 \x01(\tR\x13dataSourceAccountId\x12\x1d\n" +
	"\n" +
	"account_id\x18\x02 \x01(\tR\taccountId\"\x80\x02\n" +
	"\x16DataSourceAccountQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12$\n" +
	"\x0edata_source_id\x18\x02 \x01(\tR\fdataSourceId\x12\x1d\n" +
	"\n" +
	"account_id\x18\x03 \x01(\tR\taccountId\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x16 \x01(\tR\tprojectId\x12,\n" +
	"\x12service_account_id\x18\x17 \x01(\tR\x10serviceAccountId\"\xf0\x02\n" +
	"\x15DataSourceAccountInfo\x12\x1d\n" +
	"\n" +
	"account_id\x18\x01 \x01(\tR\taccountId\x12$\n" +
	"\x0edata_source_id\x18\x02 \x01(\tR\fdataSourceId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x17\n" +
	"\ais_sync\x18\x04 \x01(\bR\x06isSync\x12\x1b\n" +
	"\tis_linked\x18\x05 \x01(\bR\bisLinked\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12$\n" +
	"\x0ev_workspace_id\x18\x17 \x01(\tR\fvWorkspaceId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18  \x01(\tR\tupdatedAt\x12$\n" +
	"\x0elast_synced_at\x18! \x01(\tR\flastSyncedAt\"\x89\x01\n" +
	"\x16DataSourceAccountsInfo\x12N\n" +
	"\aresults\x18\x01 \x03(\v24.spaceone.api.cost_analysis.v1.DataSourceAccountInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"\xcb\x01\n" +
	"\x1dDataSourceAccountAnalyzeQuery\x12B\n" +
	"\x05query\x18\x01 \x01(\v2,.spaceone.api.core.v2.TimeSeriesAnalyzeQueryR\x05query\x12$\n" +
	"\x0edata_source_id\x18\x02 \x01(\tR\fdataSourceId\x12\x1d\n" +
	"\n" +
	"account_id\x18\x03 \x01(\tR\taccountId\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\"Y\n" +
	"\x1aDataSourceAccountStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\xf0\a\n" +
	"\x11DataSourceAccount\x12\xb6\x01\n" +
	"\x06update\x12=.spaceone.api.cost_analysis.v1.UpdateDataSourceAccountRequest\x1a4.spaceone.api.cost_analysis.v1.DataSourceAccountInfo\"7\x82\xd3\xe4\x93\x021:\x01*\",/cost-analysis/v1/data-source-account/update\x12\x95\x01\n" +
	"\x05reset\x12<.spaceone.api.cost_analysis.v1.ResetDataSourceAccountRequest\x1a\x16.google.protobuf.Empty\"6\x82\xd3\xe4\x93\x020:\x01*\"+/cost-analysis/v1/data-source-account/reset\x12\xaa\x01\n" +
	"\x03get\x127.spaceone.api.cost_analysis.v1.DataSourceAccountRequest\x1a4.spaceone.api.cost_analysis.v1.DataSourceAccountInfo\"4\x82\xd3\xe4\x93\x02.:\x01*\")/cost-analysis/v1/data-source-account/get\x12\xab\x01\n" +
	"\x04list\x125.spaceone.api.cost_analysis.v1.DataSourceAccountQuery\x1a5.spaceone.api.cost_analysis.v1.DataSourceAccountsInfo\"5\x82\xd3\xe4\x93\x02/:\x01*\"*/cost-analysis/v1/data-source-account/list\x12\x9a\x01\n" +
	"\aanalyze\x12<.spaceone.api.cost_analysis.v1.DataSourceAccountAnalyzeQuery\x1a\x17.google.protobuf.Struct\"8\x82\xd3\xe4\x93\x022:\x01*\"-/cost-analysis/v1/data-source-account/analyze\x12\x91\x01\n" +
	"\x04stat\x129.spaceone.api.cost_analysis.v1.DataSourceAccountStatQuery\x1a\x17.google.protobuf.Struct\"5\x82\xd3\xe4\x93\x02/:\x01*\"*/cost-analysis/v1/data-source-account/statBDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3"

var (
	file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescData []byte
)

func file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDesc)))
	})
	return file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_cost_analysis_v1_data_source_account_proto_goTypes = []any{
	(*UpdateDataSourceAccountRequest)(nil), // 0: spaceone.api.cost_analysis.v1.UpdateDataSourceAccountRequest
	(*ResetDataSourceAccountRequest)(nil),  // 1: spaceone.api.cost_analysis.v1.ResetDataSourceAccountRequest
	(*DataSourceAccountRequest)(nil),       // 2: spaceone.api.cost_analysis.v1.DataSourceAccountRequest
	(*DataSourceAccountQuery)(nil),         // 3: spaceone.api.cost_analysis.v1.DataSourceAccountQuery
	(*DataSourceAccountInfo)(nil),          // 4: spaceone.api.cost_analysis.v1.DataSourceAccountInfo
	(*DataSourceAccountsInfo)(nil),         // 5: spaceone.api.cost_analysis.v1.DataSourceAccountsInfo
	(*DataSourceAccountAnalyzeQuery)(nil),  // 6: spaceone.api.cost_analysis.v1.DataSourceAccountAnalyzeQuery
	(*DataSourceAccountStatQuery)(nil),     // 7: spaceone.api.cost_analysis.v1.DataSourceAccountStatQuery
	(*v2.Query)(nil),                       // 8: spaceone.api.core.v2.Query
	(*v2.TimeSeriesAnalyzeQuery)(nil),      // 9: spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	(*v2.StatisticsQuery)(nil),             // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                    // 11: google.protobuf.Empty
	(*_struct.Struct)(nil),                 // 12: google.protobuf.Struct
}
var file_spaceone_api_cost_analysis_v1_data_source_account_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.cost_analysis.v1.DataSourceAccountQuery.query:type_name -> spaceone.api.core.v2.Query
	4,  // 1: spaceone.api.cost_analysis.v1.DataSourceAccountsInfo.results:type_name -> spaceone.api.cost_analysis.v1.DataSourceAccountInfo
	9,  // 2: spaceone.api.cost_analysis.v1.DataSourceAccountAnalyzeQuery.query:type_name -> spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	10, // 3: spaceone.api.cost_analysis.v1.DataSourceAccountStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 4: spaceone.api.cost_analysis.v1.DataSourceAccount.update:input_type -> spaceone.api.cost_analysis.v1.UpdateDataSourceAccountRequest
	1,  // 5: spaceone.api.cost_analysis.v1.DataSourceAccount.reset:input_type -> spaceone.api.cost_analysis.v1.ResetDataSourceAccountRequest
	2,  // 6: spaceone.api.cost_analysis.v1.DataSourceAccount.get:input_type -> spaceone.api.cost_analysis.v1.DataSourceAccountRequest
	3,  // 7: spaceone.api.cost_analysis.v1.DataSourceAccount.list:input_type -> spaceone.api.cost_analysis.v1.DataSourceAccountQuery
	6,  // 8: spaceone.api.cost_analysis.v1.DataSourceAccount.analyze:input_type -> spaceone.api.cost_analysis.v1.DataSourceAccountAnalyzeQuery
	7,  // 9: spaceone.api.cost_analysis.v1.DataSourceAccount.stat:input_type -> spaceone.api.cost_analysis.v1.DataSourceAccountStatQuery
	4,  // 10: spaceone.api.cost_analysis.v1.DataSourceAccount.update:output_type -> spaceone.api.cost_analysis.v1.DataSourceAccountInfo
	11, // 11: spaceone.api.cost_analysis.v1.DataSourceAccount.reset:output_type -> google.protobuf.Empty
	4,  // 12: spaceone.api.cost_analysis.v1.DataSourceAccount.get:output_type -> spaceone.api.cost_analysis.v1.DataSourceAccountInfo
	5,  // 13: spaceone.api.cost_analysis.v1.DataSourceAccount.list:output_type -> spaceone.api.cost_analysis.v1.DataSourceAccountsInfo
	12, // 14: spaceone.api.cost_analysis.v1.DataSourceAccount.analyze:output_type -> google.protobuf.Struct
	12, // 15: spaceone.api.cost_analysis.v1.DataSourceAccount.stat:output_type -> google.protobuf.Struct
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_data_source_account_proto_init() }
func file_spaceone_api_cost_analysis_v1_data_source_account_proto_init() {
	if File_spaceone_api_cost_analysis_v1_data_source_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_data_source_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_data_source_account_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_data_source_account_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_data_source_account_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_data_source_account_proto = out.File
	file_spaceone_api_cost_analysis_v1_data_source_account_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_data_source_account_proto_depIdxs = nil
}
