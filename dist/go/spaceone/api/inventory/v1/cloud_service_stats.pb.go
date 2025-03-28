// A CloudServiceStats is statistics data created from from cloud service query sets.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/inventory/v1/cloud_service_stats.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
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
//	   "query": <SearchQuery>,
//	   "query_set_id": "query-set-abcd1234",
//	   "provider": "aws",
//	   "cloud_service_group": "EC2",
//	   "cloud_service_type": "Instance",
//	   "region_code": "us-east-1",
//	   "account": "aws-account-id",
//	   "project_id": "project-abcd1234"
//	}
type CloudServiceStatsQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query      *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	QuerySetId string    `protobuf:"bytes,2,opt,name=query_set_id,json=querySetId,proto3" json:"query_set_id,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	CloudServiceGroup string `protobuf:"bytes,4,opt,name=cloud_service_group,json=cloudServiceGroup,proto3" json:"cloud_service_group,omitempty"`
	// +optional
	CloudServiceType string `protobuf:"bytes,5,opt,name=cloud_service_type,json=cloudServiceType,proto3" json:"cloud_service_type,omitempty"`
	// +optional
	RegionCode string `protobuf:"bytes,6,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	Account string `protobuf:"bytes,7,opt,name=account,proto3" json:"account,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId     string `protobuf:"bytes,22,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloudServiceStatsQuery) Reset() {
	*x = CloudServiceStatsQuery{}
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudServiceStatsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsQuery) ProtoMessage() {}

func (x *CloudServiceStatsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsQuery.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{0}
}

func (x *CloudServiceStatsQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CloudServiceStatsQuery) GetQuerySetId() string {
	if x != nil {
		return x.QuerySetId
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetCloudServiceGroup() string {
	if x != nil {
		return x.CloudServiceGroup
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetCloudServiceType() string {
	if x != nil {
		return x.CloudServiceType
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

//	{
//	   "query_set_id": "query-set-abcd1234",
//	   "data": {
//	       "Disk Size": 1040,
//	       "Memory Size": 1024,
//	       "CPU": 2
//	   },
//	   "unit": {
//	       "Disk Size": "GB",
//	       "Memory": "GB",
//	       "CPU": "Core"
//	   },
//	   "provider": "aws",
//	   "cloud_service_group": "EC2",
//	   "cloud_service_type": "Instance",
//	   "region_code": "us-east-1",
//	   "account": "aws-account-id",
//	   "additional_info": {
//	       "instance_type": "t2.micro"
//	   },
//	   "project_id": "project-abcd1234",
//	   "domain_id": "domain-58010aa2e451",
//	   "created_date": "2022-06-22"
//	}
type CloudServiceStatInfo struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	QuerySetId        string                 `protobuf:"bytes,1,opt,name=query_set_id,json=querySetId,proto3" json:"query_set_id,omitempty"`
	Data              *_struct.Struct        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Unit              *_struct.Struct        `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Provider          string                 `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	CloudServiceGroup string                 `protobuf:"bytes,5,opt,name=cloud_service_group,json=cloudServiceGroup,proto3" json:"cloud_service_group,omitempty"`
	CloudServiceType  string                 `protobuf:"bytes,6,opt,name=cloud_service_type,json=cloudServiceType,proto3" json:"cloud_service_type,omitempty"`
	RegionCode        string                 `protobuf:"bytes,7,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Account           string                 `protobuf:"bytes,8,opt,name=account,proto3" json:"account,omitempty"`
	AdditionalInfo    *_struct.Struct        `protobuf:"bytes,9,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	DomainId          string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId       string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId         string                 `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CreatedDate       string                 `protobuf:"bytes,31,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CloudServiceStatInfo) Reset() {
	*x = CloudServiceStatInfo{}
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudServiceStatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatInfo) ProtoMessage() {}

func (x *CloudServiceStatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatInfo.ProtoReflect.Descriptor instead.
func (*CloudServiceStatInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{1}
}

func (x *CloudServiceStatInfo) GetQuerySetId() string {
	if x != nil {
		return x.QuerySetId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CloudServiceStatInfo) GetUnit() *_struct.Struct {
	if x != nil {
		return x.Unit
	}
	return nil
}

func (x *CloudServiceStatInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CloudServiceStatInfo) GetCloudServiceGroup() string {
	if x != nil {
		return x.CloudServiceGroup
	}
	return ""
}

func (x *CloudServiceStatInfo) GetCloudServiceType() string {
	if x != nil {
		return x.CloudServiceType
	}
	return ""
}

func (x *CloudServiceStatInfo) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CloudServiceStatInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CloudServiceStatInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *CloudServiceStatInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "query_set_id": "query-set-abcd1234",
//	           "key": "Disk Size",
//	           "data": {
//	               "Disk Size": 1040,
//	               "Memory Size": 1024,
//	               "CPU": 2
//	           },
//	           "unit": {
//	               "Disk Size": "GB",
//	               "Memory": "GB",
//	               "CPU": "Core"
//	           },
//	           "provider": "aws",
//	           "cloud_service_group": "EC2",
//	           "cloud_service_type": "Instance",
//	           "region_code": "us-east-1",
//	           "account": "aws-account-id",
//	           "additional_info": {
//	               "instance_type": "t2.micro"
//	           },
//	           "project_id": "project-abcd1234",
//	           "domain_id": "domain-58010aa2e451",
//	           "created_date": "2022-06-22"
//	       },
//	       {...}
//	   ],
//	   "total_count": 2
//	}
type CloudServiceStatsInfo struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Results       []*CloudServiceStatInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                   `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloudServiceStatsInfo) Reset() {
	*x = CloudServiceStatsInfo{}
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudServiceStatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsInfo) ProtoMessage() {}

func (x *CloudServiceStatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsInfo.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{2}
}

func (x *CloudServiceStatsInfo) GetResults() []*CloudServiceStatInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CloudServiceStatsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CloudServiceStatsAnalyzeQuery struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Query         *v2.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	QuerySetId    string                     `protobuf:"bytes,2,opt,name=query_set_id,json=querySetId,proto3" json:"query_set_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloudServiceStatsAnalyzeQuery) Reset() {
	*x = CloudServiceStatsAnalyzeQuery{}
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudServiceStatsAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsAnalyzeQuery) ProtoMessage() {}

func (x *CloudServiceStatsAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{3}
}

func (x *CloudServiceStatsAnalyzeQuery) GetQuery() *v2.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CloudServiceStatsAnalyzeQuery) GetQuerySetId() string {
	if x != nil {
		return x.QuerySetId
	}
	return ""
}

type CloudServiceStatsStatQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Query *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	QuerySetId    string `protobuf:"bytes,2,opt,name=query_set_id,json=querySetId,proto3" json:"query_set_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloudServiceStatsStatQuery) Reset() {
	*x = CloudServiceStatsStatQuery{}
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloudServiceStatsStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsStatQuery) ProtoMessage() {}

func (x *CloudServiceStatsStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsStatQuery.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{4}
}

func (x *CloudServiceStatsStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CloudServiceStatsStatQuery) GetQuerySetId() string {
	if x != nil {
		return x.QuerySetId
	}
	return ""
}

var File_spaceone_api_inventory_v1_cloud_service_stats_proto protoreflect.FileDescriptor

const file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc = "" +
	"\n" +
	"3spaceone/api/inventory/v1/cloud_service_stats.proto\x12\x19spaceone.api.inventory.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xe4\x02\n" +
	"\x16CloudServiceStatsQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12 \n" +
	"\fquery_set_id\x18\x02 \x01(\tR\n" +
	"querySetId\x12\x1a\n" +
	"\bprovider\x18\x03 \x01(\tR\bprovider\x12.\n" +
	"\x13cloud_service_group\x18\x04 \x01(\tR\x11cloudServiceGroup\x12,\n" +
	"\x12cloud_service_type\x18\x05 \x01(\tR\x10cloudServiceType\x12\x1f\n" +
	"\vregion_code\x18\x06 \x01(\tR\n" +
	"regionCode\x12\x18\n" +
	"\aaccount\x18\a \x01(\tR\aaccount\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x16 \x01(\tR\tprojectId\"\x8b\x04\n" +
	"\x14CloudServiceStatInfo\x12 \n" +
	"\fquery_set_id\x18\x01 \x01(\tR\n" +
	"querySetId\x12+\n" +
	"\x04data\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04data\x12+\n" +
	"\x04unit\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04unit\x12\x1a\n" +
	"\bprovider\x18\x04 \x01(\tR\bprovider\x12.\n" +
	"\x13cloud_service_group\x18\x05 \x01(\tR\x11cloudServiceGroup\x12,\n" +
	"\x12cloud_service_type\x18\x06 \x01(\tR\x10cloudServiceType\x12\x1f\n" +
	"\vregion_code\x18\a \x01(\tR\n" +
	"regionCode\x12\x18\n" +
	"\aaccount\x18\b \x01(\tR\aaccount\x12@\n" +
	"\x0fadditional_info\x18\t \x01(\v2\x17.google.protobuf.StructR\x0eadditionalInfo\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\x12!\n" +
	"\fcreated_date\x18\x1f \x01(\tR\vcreatedDate\"\x83\x01\n" +
	"\x15CloudServiceStatsInfo\x12I\n" +
	"\aresults\x18\x01 \x03(\v2/.spaceone.api.inventory.v1.CloudServiceStatInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"\x85\x01\n" +
	"\x1dCloudServiceStatsAnalyzeQuery\x12B\n" +
	"\x05query\x18\x01 \x01(\v2,.spaceone.api.core.v2.TimeSeriesAnalyzeQueryR\x05query\x12 \n" +
	"\fquery_set_id\x18\x02 \x01(\tR\n" +
	"querySetId\"{\n" +
	"\x1aCloudServiceStatsStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query\x12 \n" +
	"\fquery_set_id\x18\x02 \x01(\tR\n" +
	"querySetId2\xd5\x03\n" +
	"\x11CloudServiceStats\x12\x9e\x01\n" +
	"\x04list\x121.spaceone.api.inventory.v1.CloudServiceStatsQuery\x1a0.spaceone.api.inventory.v1.CloudServiceStatsInfo\"1\x82\xd3\xe4\x93\x02+:\x01*\"&/inventory/v1/cloud-service-stats/list\x12\x92\x01\n" +
	"\aanalyze\x128.spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery\x1a\x17.google.protobuf.Struct\"4\x82\xd3\xe4\x93\x02.:\x01*\")/inventory/v1/cloud-service-stats/analyze\x12\x89\x01\n" +
	"\x04stat\x125.spaceone.api.inventory.v1.CloudServiceStatsStatQuery\x1a\x17.google.protobuf.Struct\"1\x82\xd3\xe4\x93\x02+:\x01*\"&/inventory/v1/cloud-service-stats/statB@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3"

var (
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData []byte
)

func file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc), len(file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc)))
	})
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData
}

var file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_inventory_v1_cloud_service_stats_proto_goTypes = []any{
	(*CloudServiceStatsQuery)(nil),        // 0: spaceone.api.inventory.v1.CloudServiceStatsQuery
	(*CloudServiceStatInfo)(nil),          // 1: spaceone.api.inventory.v1.CloudServiceStatInfo
	(*CloudServiceStatsInfo)(nil),         // 2: spaceone.api.inventory.v1.CloudServiceStatsInfo
	(*CloudServiceStatsAnalyzeQuery)(nil), // 3: spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery
	(*CloudServiceStatsStatQuery)(nil),    // 4: spaceone.api.inventory.v1.CloudServiceStatsStatQuery
	(*v2.Query)(nil),                      // 5: spaceone.api.core.v2.Query
	(*_struct.Struct)(nil),                // 6: google.protobuf.Struct
	(*v2.TimeSeriesAnalyzeQuery)(nil),     // 7: spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	(*v2.StatisticsQuery)(nil),            // 8: spaceone.api.core.v2.StatisticsQuery
}
var file_spaceone_api_inventory_v1_cloud_service_stats_proto_depIdxs = []int32{
	5,  // 0: spaceone.api.inventory.v1.CloudServiceStatsQuery.query:type_name -> spaceone.api.core.v2.Query
	6,  // 1: spaceone.api.inventory.v1.CloudServiceStatInfo.data:type_name -> google.protobuf.Struct
	6,  // 2: spaceone.api.inventory.v1.CloudServiceStatInfo.unit:type_name -> google.protobuf.Struct
	6,  // 3: spaceone.api.inventory.v1.CloudServiceStatInfo.additional_info:type_name -> google.protobuf.Struct
	1,  // 4: spaceone.api.inventory.v1.CloudServiceStatsInfo.results:type_name -> spaceone.api.inventory.v1.CloudServiceStatInfo
	7,  // 5: spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery.query:type_name -> spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	8,  // 6: spaceone.api.inventory.v1.CloudServiceStatsStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 7: spaceone.api.inventory.v1.CloudServiceStats.list:input_type -> spaceone.api.inventory.v1.CloudServiceStatsQuery
	3,  // 8: spaceone.api.inventory.v1.CloudServiceStats.analyze:input_type -> spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery
	4,  // 9: spaceone.api.inventory.v1.CloudServiceStats.stat:input_type -> spaceone.api.inventory.v1.CloudServiceStatsStatQuery
	2,  // 10: spaceone.api.inventory.v1.CloudServiceStats.list:output_type -> spaceone.api.inventory.v1.CloudServiceStatsInfo
	6,  // 11: spaceone.api.inventory.v1.CloudServiceStats.analyze:output_type -> google.protobuf.Struct
	6,  // 12: spaceone.api.inventory.v1.CloudServiceStats.stat:output_type -> google.protobuf.Struct
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v1_cloud_service_stats_proto_init() }
func file_spaceone_api_inventory_v1_cloud_service_stats_proto_init() {
	if File_spaceone_api_inventory_v1_cloud_service_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc), len(file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v1_cloud_service_stats_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v1_cloud_service_stats_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v1_cloud_service_stats_proto = out.File
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_goTypes = nil
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_depIdxs = nil
}
