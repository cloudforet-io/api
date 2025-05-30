// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/inventory_v2/v1/metric_data.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	_ "github.com/golang/protobuf/ptypes/empty"
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

// {
//
// }
type MetricDataQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query    *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	MetricId string    `protobuf:"bytes,2,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId     string `protobuf:"bytes,22,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricDataQuery) Reset() {
	*x = MetricDataQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricDataQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDataQuery) ProtoMessage() {}

func (x *MetricDataQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDataQuery.ProtoReflect.Descriptor instead.
func (*MetricDataQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescGZIP(), []int{0}
}

func (x *MetricDataQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MetricDataQuery) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

func (x *MetricDataQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *MetricDataQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

// {
//
// }
type MetricDataInfo struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	MetricId         string                 `protobuf:"bytes,1,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	Value            float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	Unit             string                 `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Labels           *_struct.Struct        `protobuf:"bytes,4,opt,name=labels,proto3" json:"labels,omitempty"`
	DomainId         string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId        string                 `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceAccountId string                 `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	NamespaceId      string                 `protobuf:"bytes,25,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	CreatedYear      string                 `protobuf:"bytes,31,opt,name=created_year,json=createdYear,proto3" json:"created_year,omitempty"`
	CreatedMonth     string                 `protobuf:"bytes,32,opt,name=created_month,json=createdMonth,proto3" json:"created_month,omitempty"`
	CreatedDate      string                 `protobuf:"bytes,33,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MetricDataInfo) Reset() {
	*x = MetricDataInfo{}
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDataInfo) ProtoMessage() {}

func (x *MetricDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDataInfo.ProtoReflect.Descriptor instead.
func (*MetricDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescGZIP(), []int{1}
}

func (x *MetricDataInfo) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

func (x *MetricDataInfo) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *MetricDataInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *MetricDataInfo) GetLabels() *_struct.Struct {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetricDataInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *MetricDataInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *MetricDataInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *MetricDataInfo) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *MetricDataInfo) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *MetricDataInfo) GetCreatedYear() string {
	if x != nil {
		return x.CreatedYear
	}
	return ""
}

func (x *MetricDataInfo) GetCreatedMonth() string {
	if x != nil {
		return x.CreatedMonth
	}
	return ""
}

func (x *MetricDataInfo) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

type MetricDatasInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*MetricDataInfo      `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricDatasInfo) Reset() {
	*x = MetricDatasInfo{}
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricDatasInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDatasInfo) ProtoMessage() {}

func (x *MetricDatasInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDatasInfo.ProtoReflect.Descriptor instead.
func (*MetricDatasInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescGZIP(), []int{2}
}

func (x *MetricDatasInfo) GetResults() []*MetricDataInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *MetricDatasInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type MetricDataAnalyzeQuery struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Query         *v2.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	MetricId      string                     `protobuf:"bytes,2,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricDataAnalyzeQuery) Reset() {
	*x = MetricDataAnalyzeQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricDataAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDataAnalyzeQuery) ProtoMessage() {}

func (x *MetricDataAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDataAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*MetricDataAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescGZIP(), []int{3}
}

func (x *MetricDataAnalyzeQuery) GetQuery() *v2.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MetricDataAnalyzeQuery) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

type MetricDataStatQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Query *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	MetricId      string `protobuf:"bytes,2,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricDataStatQuery) Reset() {
	*x = MetricDataStatQuery{}
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricDataStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDataStatQuery) ProtoMessage() {}

func (x *MetricDataStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDataStatQuery.ProtoReflect.Descriptor instead.
func (*MetricDataStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescGZIP(), []int{4}
}

func (x *MetricDataStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MetricDataStatQuery) GetMetricId() string {
	if x != nil {
		return x.MetricId
	}
	return ""
}

var File_spaceone_api_inventory_v2_v1_metric_data_proto protoreflect.FileDescriptor

const file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDesc = "" +
	"\n" +
	".spaceone/api/inventory_v2/v1/metric_data.proto\x12\x19spaceone.api.inventory.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a)spaceone/api/inventory_v2/v1/metric.proto\"\xa3\x01\n" +
	"\x0fMetricDataQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x1b\n" +
	"\tmetric_id\x18\x02 \x01(\tR\bmetricId\x12!\n" +
	"\fworkspace_id\x18\x15 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x16 \x01(\tR\tprojectId\"\xa3\x03\n" +
	"\x0eMetricDataInfo\x12\x1b\n" +
	"\tmetric_id\x18\x01 \x01(\tR\bmetricId\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x02R\x05value\x12\x12\n" +
	"\x04unit\x18\x03 \x01(\tR\x04unit\x12/\n" +
	"\x06labels\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x06labels\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"project_id\x18\x17 \x01(\tR\tprojectId\x12,\n" +
	"\x12service_account_id\x18\x18 \x01(\tR\x10serviceAccountId\x12!\n" +
	"\fnamespace_id\x18\x19 \x01(\tR\vnamespaceId\x12!\n" +
	"\fcreated_year\x18\x1f \x01(\tR\vcreatedYear\x12#\n" +
	"\rcreated_month\x18  \x01(\tR\fcreatedMonth\x12!\n" +
	"\fcreated_date\x18! \x01(\tR\vcreatedDate\"w\n" +
	"\x0fMetricDatasInfo\x12C\n" +
	"\aresults\x18\x01 \x03(\v2).spaceone.api.inventory.v1.MetricDataInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"y\n" +
	"\x16MetricDataAnalyzeQuery\x12B\n" +
	"\x05query\x18\x01 \x01(\v2,.spaceone.api.core.v2.TimeSeriesAnalyzeQueryR\x05query\x12\x1b\n" +
	"\tmetric_id\x18\x02 \x01(\tR\bmetricId\"o\n" +
	"\x13MetricDataStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query\x12\x1b\n" +
	"\tmetric_id\x18\x02 \x01(\tR\bmetricId2\xa3\x03\n" +
	"\n" +
	"MetricData\x12\x8c\x01\n" +
	"\x04list\x12*.spaceone.api.inventory.v1.MetricDataQuery\x1a*.spaceone.api.inventory.v1.MetricDatasInfo\",\x82\xd3\xe4\x93\x02&:\x01*\"!/inventory-v2/v1/metric-data/list\x12}\n" +
	"\x04stat\x12..spaceone.api.inventory.v1.MetricDataStatQuery\x1a\x17.google.protobuf.Struct\",\x82\xd3\xe4\x93\x02&:\x01*\"!/inventory-v2/v1/metric-data/stat\x12\x86\x01\n" +
	"\aanalyze\x121.spaceone.api.inventory.v1.MetricDataAnalyzeQuery\x1a\x17.google.protobuf.Struct\"/\x82\xd3\xe4\x93\x02):\x01*\"$/inventory-v2/v1/metric-data/analyzeBCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/inventory_v2/v1b\x06proto3"

var (
	file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescData []byte
)

func file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDesc), len(file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDesc)))
	})
	return file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDescData
}

var file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_inventory_v2_v1_metric_data_proto_goTypes = []any{
	(*MetricDataQuery)(nil),           // 0: spaceone.api.inventory.v1.MetricDataQuery
	(*MetricDataInfo)(nil),            // 1: spaceone.api.inventory.v1.MetricDataInfo
	(*MetricDatasInfo)(nil),           // 2: spaceone.api.inventory.v1.MetricDatasInfo
	(*MetricDataAnalyzeQuery)(nil),    // 3: spaceone.api.inventory.v1.MetricDataAnalyzeQuery
	(*MetricDataStatQuery)(nil),       // 4: spaceone.api.inventory.v1.MetricDataStatQuery
	(*v2.Query)(nil),                  // 5: spaceone.api.core.v2.Query
	(*_struct.Struct)(nil),            // 6: google.protobuf.Struct
	(*v2.TimeSeriesAnalyzeQuery)(nil), // 7: spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	(*v2.StatisticsQuery)(nil),        // 8: spaceone.api.core.v2.StatisticsQuery
}
var file_spaceone_api_inventory_v2_v1_metric_data_proto_depIdxs = []int32{
	5, // 0: spaceone.api.inventory.v1.MetricDataQuery.query:type_name -> spaceone.api.core.v2.Query
	6, // 1: spaceone.api.inventory.v1.MetricDataInfo.labels:type_name -> google.protobuf.Struct
	1, // 2: spaceone.api.inventory.v1.MetricDatasInfo.results:type_name -> spaceone.api.inventory.v1.MetricDataInfo
	7, // 3: spaceone.api.inventory.v1.MetricDataAnalyzeQuery.query:type_name -> spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	8, // 4: spaceone.api.inventory.v1.MetricDataStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0, // 5: spaceone.api.inventory.v1.MetricData.list:input_type -> spaceone.api.inventory.v1.MetricDataQuery
	4, // 6: spaceone.api.inventory.v1.MetricData.stat:input_type -> spaceone.api.inventory.v1.MetricDataStatQuery
	3, // 7: spaceone.api.inventory.v1.MetricData.analyze:input_type -> spaceone.api.inventory.v1.MetricDataAnalyzeQuery
	2, // 8: spaceone.api.inventory.v1.MetricData.list:output_type -> spaceone.api.inventory.v1.MetricDatasInfo
	6, // 9: spaceone.api.inventory.v1.MetricData.stat:output_type -> google.protobuf.Struct
	6, // 10: spaceone.api.inventory.v1.MetricData.analyze:output_type -> google.protobuf.Struct
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v2_v1_metric_data_proto_init() }
func file_spaceone_api_inventory_v2_v1_metric_data_proto_init() {
	if File_spaceone_api_inventory_v2_v1_metric_data_proto != nil {
		return
	}
	file_spaceone_api_inventory_v2_v1_metric_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDesc), len(file_spaceone_api_inventory_v2_v1_metric_data_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v2_v1_metric_data_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v2_v1_metric_data_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_v2_v1_metric_data_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v2_v1_metric_data_proto = out.File
	file_spaceone_api_inventory_v2_v1_metric_data_proto_goTypes = nil
	file_spaceone_api_inventory_v2_v1_metric_data_proto_depIdxs = nil
}
