// A Cost is a resource of raw cost data collected by the cost_analysis.DataSource.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/cost.proto

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
//	       "cost": 142,
//	       "usage_quantity": 84532,
//	       "usage_unit": "GB",
//	       "provider": "aws",
//	       "region_code": "ap-northeast-1",
//	       "product": "AWSDataTransfer",
//	       "account": "722069360300",
//	       "usage_type": "data-transfer.out",
//	       "additional_info": {
//	           "raw_usage_type": "APN1-DataTransfer-Out-Bytes"
//	       },
//	       "tags": {
//	           "Environment": "Dev"
//	       },
//	       "data_source_id": "ds-fcba92ca73b1"
//	}
type CreateCostRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Cost  float64                `protobuf:"fixed64,1,opt,name=cost,proto3" json:"cost,omitempty"`
	// +optional
	UsageQuantity float64 `protobuf:"fixed64,2,opt,name=usage_quantity,json=usageQuantity,proto3" json:"usage_quantity,omitempty"`
	// +optional
	UsageUnit float32 `protobuf:"fixed32,3,opt,name=usage_unit,json=usageUnit,proto3" json:"usage_unit,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	RegionCode string `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	Product string `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
	// +optional
	UsageType string `protobuf:"bytes,7,opt,name=usage_type,json=usageType,proto3" json:"usage_type,omitempty"`
	// +optional
	Resource string `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	AdditionalInfo *_struct.Struct `protobuf:"bytes,10,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	ProjectId      string          `protobuf:"bytes,21,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	ServiceAccountId string `protobuf:"bytes,22,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	DataSourceId     string `protobuf:"bytes,23,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	BilledDate       string `protobuf:"bytes,31,opt,name=billed_date,json=billedDate,proto3" json:"billed_date,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateCostRequest) Reset() {
	*x = CreateCostRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCostRequest) ProtoMessage() {}

func (x *CreateCostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCostRequest.ProtoReflect.Descriptor instead.
func (*CreateCostRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCostRequest) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *CreateCostRequest) GetUsageQuantity() float64 {
	if x != nil {
		return x.UsageQuantity
	}
	return 0
}

func (x *CreateCostRequest) GetUsageUnit() float32 {
	if x != nil {
		return x.UsageUnit
	}
	return 0
}

func (x *CreateCostRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateCostRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CreateCostRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CreateCostRequest) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *CreateCostRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *CreateCostRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateCostRequest) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *CreateCostRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateCostRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CreateCostRequest) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CreateCostRequest) GetBilledDate() string {
	if x != nil {
		return x.BilledDate
	}
	return ""
}

//	{
//	   "cost_id": "cost-2ad052ed03d7"
//	}
type CostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CostId        string                 `protobuf:"bytes,1,opt,name=cost_id,json=costId,proto3" json:"cost_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CostRequest) Reset() {
	*x = CostRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostRequest) ProtoMessage() {}

func (x *CostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostRequest.ProtoReflect.Descriptor instead.
func (*CostRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{1}
}

func (x *CostRequest) GetCostId() string {
	if x != nil {
		return x.CostId
	}
	return ""
}

//	{
//	   "query": {}
//	}
type CostQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query        *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DataSourceId string    `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	// +optional
	CostId string `protobuf:"bytes,3,opt,name=cost_id,json=costId,proto3" json:"cost_id,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	RegionCode string `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	RegionKey string `protobuf:"bytes,6,opt,name=region_key,json=regionKey,proto3" json:"region_key,omitempty"`
	// +optional
	Product string `protobuf:"bytes,7,opt,name=product,proto3" json:"product,omitempty"`
	// +optional
	UsageType string `protobuf:"bytes,8,opt,name=usage_type,json=usageType,proto3" json:"usage_type,omitempty"`
	// +optional
	Resource string `protobuf:"bytes,9,opt,name=resource,proto3" json:"resource,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,22,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	ProjectGroupId string `protobuf:"bytes,23,opt,name=project_group_id,json=projectGroupId,proto3" json:"project_group_id,omitempty"`
	// +optional
	ServiceAccountId string `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// +optional
	BilledYear string `protobuf:"bytes,31,opt,name=billed_year,json=billedYear,proto3" json:"billed_year,omitempty"`
	// +optional
	BilledMonth string `protobuf:"bytes,32,opt,name=billed_month,json=billedMonth,proto3" json:"billed_month,omitempty"`
	// +optional
	BilledDate    string `protobuf:"bytes,33,opt,name=billed_date,json=billedDate,proto3" json:"billed_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CostQuery) Reset() {
	*x = CostQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostQuery) ProtoMessage() {}

func (x *CostQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostQuery.ProtoReflect.Descriptor instead.
func (*CostQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{2}
}

func (x *CostQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CostQuery) GetCostId() string {
	if x != nil {
		return x.CostId
	}
	return ""
}

func (x *CostQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CostQuery) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CostQuery) GetRegionKey() string {
	if x != nil {
		return x.RegionKey
	}
	return ""
}

func (x *CostQuery) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CostQuery) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *CostQuery) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *CostQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CostQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CostQuery) GetProjectGroupId() string {
	if x != nil {
		return x.ProjectGroupId
	}
	return ""
}

func (x *CostQuery) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CostQuery) GetBilledYear() string {
	if x != nil {
		return x.BilledYear
	}
	return ""
}

func (x *CostQuery) GetBilledMonth() string {
	if x != nil {
		return x.BilledMonth
	}
	return ""
}

func (x *CostQuery) GetBilledDate() string {
	if x != nil {
		return x.BilledDate
	}
	return ""
}

//	{
//	       "cost_id": "cost-c5aae7712ec9",
//	       "cost": 142,
//	       "usage_quantity": 84532,
//	       "usage_unit": "GB",
//	       "provider": "aws",
//	       "region_code": "ap-northeast-1",
//	       "product": "AWSDataTransfer",
//	       "usage_type": "data-transfer.out",
//	       "additional_info": {
//	           "raw_usage_type": "APN1-DataTransfer-Out-Bytes"
//	       },
//	       "tags": {
//	           "Environment": "Dev"
//	       },
//	       "data_source_id": "ds-fcba92ca73b1"
//	       "domain_id": "domain-58010aa2e451",
//	       "billed_year": "2022",
//	       "billed_month": "2022-07",
//	       "billed_date": "2022-07-19"
//	}
type CostInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	CostId           string                 `protobuf:"bytes,1,opt,name=cost_id,json=costId,proto3" json:"cost_id,omitempty"`
	Cost             float64                `protobuf:"fixed64,2,opt,name=cost,proto3" json:"cost,omitempty"`
	UsageQuantity    float64                `protobuf:"fixed64,3,opt,name=usage_quantity,json=usageQuantity,proto3" json:"usage_quantity,omitempty"`
	UsageUnit        string                 `protobuf:"bytes,4,opt,name=usage_unit,json=usageUnit,proto3" json:"usage_unit,omitempty"`
	Provider         string                 `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	RegionCode       string                 `protobuf:"bytes,6,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	RegionKey        string                 `protobuf:"bytes,7,opt,name=region_key,json=regionKey,proto3" json:"region_key,omitempty"`
	Product          string                 `protobuf:"bytes,8,opt,name=product,proto3" json:"product,omitempty"`
	UsageType        string                 `protobuf:"bytes,9,opt,name=usage_type,json=usageType,proto3" json:"usage_type,omitempty"`
	Resource         string                 `protobuf:"bytes,10,opt,name=resource,proto3" json:"resource,omitempty"`
	Tags             *_struct.Struct        `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	AdditionalInfo   *_struct.Struct        `protobuf:"bytes,12,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	Data             *_struct.Struct        `protobuf:"bytes,13,opt,name=data,proto3" json:"data,omitempty"`
	DomainId         string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId        string                 `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceAccountId string                 `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	DataSourceId     string                 `protobuf:"bytes,25,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	AccountId        string                 `protobuf:"bytes,26,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	BilledYear       string                 `protobuf:"bytes,31,opt,name=billed_year,json=billedYear,proto3" json:"billed_year,omitempty"`
	BilledMonth      string                 `protobuf:"bytes,32,opt,name=billed_month,json=billedMonth,proto3" json:"billed_month,omitempty"`
	BilledDate       string                 `protobuf:"bytes,33,opt,name=billed_date,json=billedDate,proto3" json:"billed_date,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CostInfo) Reset() {
	*x = CostInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostInfo) ProtoMessage() {}

func (x *CostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostInfo.ProtoReflect.Descriptor instead.
func (*CostInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{3}
}

func (x *CostInfo) GetCostId() string {
	if x != nil {
		return x.CostId
	}
	return ""
}

func (x *CostInfo) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *CostInfo) GetUsageQuantity() float64 {
	if x != nil {
		return x.UsageQuantity
	}
	return 0
}

func (x *CostInfo) GetUsageUnit() string {
	if x != nil {
		return x.UsageUnit
	}
	return ""
}

func (x *CostInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CostInfo) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CostInfo) GetRegionKey() string {
	if x != nil {
		return x.RegionKey
	}
	return ""
}

func (x *CostInfo) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CostInfo) GetUsageType() string {
	if x != nil {
		return x.UsageType
	}
	return ""
}

func (x *CostInfo) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *CostInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CostInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *CostInfo) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CostInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CostInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CostInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CostInfo) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CostInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CostInfo) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CostInfo) GetBilledYear() string {
	if x != nil {
		return x.BilledYear
	}
	return ""
}

func (x *CostInfo) GetBilledMonth() string {
	if x != nil {
		return x.BilledMonth
	}
	return ""
}

func (x *CostInfo) GetBilledDate() string {
	if x != nil {
		return x.BilledDate
	}
	return ""
}

//	{
//	       "results": [
//	           {
//	               "cost_id": "cost-c5aae7712ec9",
//	               "cost": 142,
//	               "usage_quantity": 84532,
//	               "usage_unit": "GB",
//	               "provider": "aws",
//	               "region_code": "ap-northeast-1",
//	               "product": "AWSDataTransfer",
//	               "usage_type": "data-transfer.out",
//	               "additional_info": {
//	               "raw_usage_type": "APN1-DataTransfer-Out-Bytes"
//	               },
//	               "tags": {
//	               "Environment": "Dev"
//	               },
//	               "data_source_id": "ds-fcba92ca73b1"
//	               "domain_id": "domain-58010aa2e451",
//	               "billed_year": "2022",
//	               "billed_month": "2022-07",
//	               "billed_date": "2022-07-19"
//	           },
//	           {
//	               "cost_id": "cost-1d5e1b0dbf82",
//	               "cost": 78,
//	               "usage_quantity": 34523,
//	               "usage_unit": "Count",
//	               "provider": "aws",
//	               "region_code": "ap-northeast-1",
//	               "product": "AWSQueueService"
//	               "additional_info": {
//	                   "raw_usage_type": "APN1-Requests-Tier1"
//	               },
//	               "tags": {},
//	               "data_source_id": "ds-fcba92ca73b1",
//	               "domain_id": "domain-58010aa2e451",
//	               "billed_year": "2022",
//	               "billed_month": "2022-07",
//	               "billed_date": "2022-07-20"
//	           }
//	       ],
//	       "total_count": 307066
//	}
type CostsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*CostInfo            `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CostsInfo) Reset() {
	*x = CostsInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostsInfo) ProtoMessage() {}

func (x *CostsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostsInfo.ProtoReflect.Descriptor instead.
func (*CostsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{4}
}

func (x *CostsInfo) GetResults() []*CostInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CostsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CostAnalyzeQuery struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Query         *v2.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DataSourceId  string                     `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CostAnalyzeQuery) Reset() {
	*x = CostAnalyzeQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostAnalyzeQuery) ProtoMessage() {}

func (x *CostAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*CostAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{5}
}

func (x *CostAnalyzeQuery) GetQuery() *v2.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostAnalyzeQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

type CostStatQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Query *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	DataSourceId  string `protobuf:"bytes,2,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CostStatQuery) Reset() {
	*x = CostStatQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CostStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostStatQuery) ProtoMessage() {}

func (x *CostStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostStatQuery.ProtoReflect.Descriptor instead.
func (*CostStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP(), []int{6}
}

func (x *CostStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CostStatQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

var File_spaceone_api_cost_analysis_v1_cost_proto protoreflect.FileDescriptor

const file_spaceone_api_cost_analysis_v1_cost_proto_rawDesc = "" +
	"\n" +
	"(spaceone/api/cost_analysis/v1/cost.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x82\x04\n" +
	"\x11CreateCostRequest\x12\x12\n" +
	"\x04cost\x18\x01 \x01(\x01R\x04cost\x12%\n" +
	"\x0eusage_quantity\x18\x02 \x01(\x01R\rusageQuantity\x12\x1d\n" +
	"\n" +
	"usage_unit\x18\x03 \x01(\x02R\tusageUnit\x12\x1a\n" +
	"\bprovider\x18\x04 \x01(\tR\bprovider\x12\x1f\n" +
	"\vregion_code\x18\x05 \x01(\tR\n" +
	"regionCode\x12\x18\n" +
	"\aproduct\x18\x06 \x01(\tR\aproduct\x12\x1d\n" +
	"\n" +
	"usage_type\x18\a \x01(\tR\tusageType\x12\x1a\n" +
	"\bresource\x18\b \x01(\tR\bresource\x12+\n" +
	"\x04tags\x18\t \x01(\v2\x17.google.protobuf.StructR\x04tags\x12@\n" +
	"\x0fadditional_info\x18\n" +
	" \x01(\v2\x17.google.protobuf.StructR\x0eadditionalInfo\x12\x1d\n" +
	"\n" +
	"project_id\x18\x15 \x01(\tR\tprojectId\x12,\n" +
	"\x12service_account_id\x18\x16 \x01(\tR\x10serviceAccountId\x12$\n" +
	"\x0edata_source_id\x18\x17 \x01(\tR\fdataSourceId\x12\x1f\n" +
	"\vbilled_date\x18\x1f \x01(\tR\n" +
	"billedDate\"&\n" +
	"\vCostRequest\x12\x17\n" +
	"\acost_id\x18\x01 \x01(\tR\x06costId\"\xad\x04\n" +
	"\tCostQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12$\n" +
	"\x0edata_source_id\x18\x02 \x01(\tR\fdataSourceId\x12\x17\n" +
	"\acost_id\x18\x03 \x01(\tR\x06costId\x12\x1a\n" +
	"\bprovider\x18\x04 \x01(\tR\bprovider\x12\x1f\n" +
	"\vregion_code\x18\x05 \x01(\tR\n" +
	"regionCode\x12\x1d\n" +
	"\n" +
	"region_key\x18\x06 \x01(\tR\tregionKey\x12\x18\n" +
	"\aproduct\x18\a \x01(\tR\aproduct\x12\x1d\n" +
	"\n" +
	"usage_type\x18\b \x01(\tR\tusageType\x12\x1a\n" +
	"\bresource\x18\t \x01(\tR\bresource\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x16 \x01(\tR\tprojectId\x12(\n" +
	"\x10project_group_id\x18\x17 \x01(\tR\x0eprojectGroupId\x12,\n" +
	"\x12service_account_id\x18\x18 \x01(\tR\x10serviceAccountId\x12\x1f\n" +
	"\vbilled_year\x18\x1f \x01(\tR\n" +
	"billedYear\x12!\n" +
	"\fbilled_month\x18  \x01(\tR\vbilledMonth\x12\x1f\n" +
	"\vbilled_date\x18! \x01(\tR\n" +
	"billedDate\"\x81\x06\n" +
	"\bCostInfo\x12\x17\n" +
	"\acost_id\x18\x01 \x01(\tR\x06costId\x12\x12\n" +
	"\x04cost\x18\x02 \x01(\x01R\x04cost\x12%\n" +
	"\x0eusage_quantity\x18\x03 \x01(\x01R\rusageQuantity\x12\x1d\n" +
	"\n" +
	"usage_unit\x18\x04 \x01(\tR\tusageUnit\x12\x1a\n" +
	"\bprovider\x18\x05 \x01(\tR\bprovider\x12\x1f\n" +
	"\vregion_code\x18\x06 \x01(\tR\n" +
	"regionCode\x12\x1d\n" +
	"\n" +
	"region_key\x18\a \x01(\tR\tregionKey\x12\x18\n" +
	"\aproduct\x18\b \x01(\tR\aproduct\x12\x1d\n" +
	"\n" +
	"usage_type\x18\t \x01(\tR\tusageType\x12\x1a\n" +
	"\bresource\x18\n" +
	" \x01(\tR\bresource\x12+\n" +
	"\x04tags\x18\v \x01(\v2\x17.google.protobuf.StructR\x04tags\x12@\n" +
	"\x0fadditional_info\x18\f \x01(\v2\x17.google.protobuf.StructR\x0eadditionalInfo\x12+\n" +
	"\x04data\x18\r \x01(\v2\x17.google.protobuf.StructR\x04data\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\x12,\n" +
	"\x12service_account_id\x18\x18 \x01(\tR\x10serviceAccountId\x12$\n" +
	"\x0edata_source_id\x18\x19 \x01(\tR\fdataSourceId\x12\x1d\n" +
	"\n" +
	"account_id\x18\x1a \x01(\tR\taccountId\x12\x1f\n" +
	"\vbilled_year\x18\x1f \x01(\tR\n" +
	"billedYear\x12!\n" +
	"\fbilled_month\x18  \x01(\tR\vbilledMonth\x12\x1f\n" +
	"\vbilled_date\x18! \x01(\tR\n" +
	"billedDate\"o\n" +
	"\tCostsInfo\x12A\n" +
	"\aresults\x18\x01 \x03(\v2'.spaceone.api.cost_analysis.v1.CostInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"|\n" +
	"\x10CostAnalyzeQuery\x12B\n" +
	"\x05query\x18\x01 \x01(\v2,.spaceone.api.core.v2.TimeSeriesAnalyzeQueryR\x05query\x12$\n" +
	"\x0edata_source_id\x18\x02 \x01(\tR\fdataSourceId\"r\n" +
	"\rCostStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query\x12$\n" +
	"\x0edata_source_id\x18\x02 \x01(\tR\fdataSourceId2\x8e\x06\n" +
	"\x04Cost\x12\x8d\x01\n" +
	"\x06create\x120.spaceone.api.cost_analysis.v1.CreateCostRequest\x1a'.spaceone.api.cost_analysis.v1.CostInfo\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/cost-analysis/v1/cost/create\x12v\n" +
	"\x06delete\x12*.spaceone.api.cost_analysis.v1.CostRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\":\x01*\"\x1d/cost-analysis/v1/cost/delete\x12\x81\x01\n" +
	"\x03get\x12*.spaceone.api.cost_analysis.v1.CostRequest\x1a'.spaceone.api.cost_analysis.v1.CostInfo\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/cost-analysis/v1/cost/get\x12\x82\x01\n" +
	"\x04list\x12(.spaceone.api.cost_analysis.v1.CostQuery\x1a(.spaceone.api.cost_analysis.v1.CostsInfo\"&\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/cost-analysis/v1/cost/list\x12~\n" +
	"\aanalyze\x12/.spaceone.api.cost_analysis.v1.CostAnalyzeQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/cost-analysis/v1/cost/analyze\x12u\n" +
	"\x04stat\x12,.spaceone.api.cost_analysis.v1.CostStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/cost-analysis/v1/cost/statBDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3"

var (
	file_spaceone_api_cost_analysis_v1_cost_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_cost_proto_rawDescData []byte
)

func file_spaceone_api_cost_analysis_v1_cost_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_cost_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_cost_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_cost_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_cost_proto_rawDesc)))
	})
	return file_spaceone_api_cost_analysis_v1_cost_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_cost_analysis_v1_cost_proto_goTypes = []any{
	(*CreateCostRequest)(nil),         // 0: spaceone.api.cost_analysis.v1.CreateCostRequest
	(*CostRequest)(nil),               // 1: spaceone.api.cost_analysis.v1.CostRequest
	(*CostQuery)(nil),                 // 2: spaceone.api.cost_analysis.v1.CostQuery
	(*CostInfo)(nil),                  // 3: spaceone.api.cost_analysis.v1.CostInfo
	(*CostsInfo)(nil),                 // 4: spaceone.api.cost_analysis.v1.CostsInfo
	(*CostAnalyzeQuery)(nil),          // 5: spaceone.api.cost_analysis.v1.CostAnalyzeQuery
	(*CostStatQuery)(nil),             // 6: spaceone.api.cost_analysis.v1.CostStatQuery
	(*_struct.Struct)(nil),            // 7: google.protobuf.Struct
	(*v2.Query)(nil),                  // 8: spaceone.api.core.v2.Query
	(*v2.TimeSeriesAnalyzeQuery)(nil), // 9: spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	(*v2.StatisticsQuery)(nil),        // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_spaceone_api_cost_analysis_v1_cost_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.cost_analysis.v1.CreateCostRequest.tags:type_name -> google.protobuf.Struct
	7,  // 1: spaceone.api.cost_analysis.v1.CreateCostRequest.additional_info:type_name -> google.protobuf.Struct
	8,  // 2: spaceone.api.cost_analysis.v1.CostQuery.query:type_name -> spaceone.api.core.v2.Query
	7,  // 3: spaceone.api.cost_analysis.v1.CostInfo.tags:type_name -> google.protobuf.Struct
	7,  // 4: spaceone.api.cost_analysis.v1.CostInfo.additional_info:type_name -> google.protobuf.Struct
	7,  // 5: spaceone.api.cost_analysis.v1.CostInfo.data:type_name -> google.protobuf.Struct
	3,  // 6: spaceone.api.cost_analysis.v1.CostsInfo.results:type_name -> spaceone.api.cost_analysis.v1.CostInfo
	9,  // 7: spaceone.api.cost_analysis.v1.CostAnalyzeQuery.query:type_name -> spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	10, // 8: spaceone.api.cost_analysis.v1.CostStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 9: spaceone.api.cost_analysis.v1.Cost.create:input_type -> spaceone.api.cost_analysis.v1.CreateCostRequest
	1,  // 10: spaceone.api.cost_analysis.v1.Cost.delete:input_type -> spaceone.api.cost_analysis.v1.CostRequest
	1,  // 11: spaceone.api.cost_analysis.v1.Cost.get:input_type -> spaceone.api.cost_analysis.v1.CostRequest
	2,  // 12: spaceone.api.cost_analysis.v1.Cost.list:input_type -> spaceone.api.cost_analysis.v1.CostQuery
	5,  // 13: spaceone.api.cost_analysis.v1.Cost.analyze:input_type -> spaceone.api.cost_analysis.v1.CostAnalyzeQuery
	6,  // 14: spaceone.api.cost_analysis.v1.Cost.stat:input_type -> spaceone.api.cost_analysis.v1.CostStatQuery
	3,  // 15: spaceone.api.cost_analysis.v1.Cost.create:output_type -> spaceone.api.cost_analysis.v1.CostInfo
	11, // 16: spaceone.api.cost_analysis.v1.Cost.delete:output_type -> google.protobuf.Empty
	3,  // 17: spaceone.api.cost_analysis.v1.Cost.get:output_type -> spaceone.api.cost_analysis.v1.CostInfo
	4,  // 18: spaceone.api.cost_analysis.v1.Cost.list:output_type -> spaceone.api.cost_analysis.v1.CostsInfo
	7,  // 19: spaceone.api.cost_analysis.v1.Cost.analyze:output_type -> google.protobuf.Struct
	7,  // 20: spaceone.api.cost_analysis.v1.Cost.stat:output_type -> google.protobuf.Struct
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_cost_proto_init() }
func file_spaceone_api_cost_analysis_v1_cost_proto_init() {
	if File_spaceone_api_cost_analysis_v1_cost_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_cost_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_cost_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_cost_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_cost_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_cost_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_cost_proto = out.File
	file_spaceone_api_cost_analysis_v1_cost_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_cost_proto_depIdxs = nil
}
