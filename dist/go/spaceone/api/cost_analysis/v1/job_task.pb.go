// A JobTask is a task unit subdividing a Job. The division criteria are specified in each DataSource plugin.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/job_task.proto

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

type JobTaskQuery_Status int32

const (
	JobTaskQuery_SCOPE_NONE  JobTaskQuery_Status = 0
	JobTaskQuery_PENDING     JobTaskQuery_Status = 1
	JobTaskQuery_IN_PROGRESS JobTaskQuery_Status = 2
	JobTaskQuery_SUCCESS     JobTaskQuery_Status = 3
	JobTaskQuery_FAILURE     JobTaskQuery_Status = 4
	JobTaskQuery_TIMEOUT     JobTaskQuery_Status = 5
	JobTaskQuery_CANCELED    JobTaskQuery_Status = 6
)

// Enum value maps for JobTaskQuery_Status.
var (
	JobTaskQuery_Status_name = map[int32]string{
		0: "SCOPE_NONE",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "SUCCESS",
		4: "FAILURE",
		5: "TIMEOUT",
		6: "CANCELED",
	}
	JobTaskQuery_Status_value = map[string]int32{
		"SCOPE_NONE":  0,
		"PENDING":     1,
		"IN_PROGRESS": 2,
		"SUCCESS":     3,
		"FAILURE":     4,
		"TIMEOUT":     5,
		"CANCELED":    6,
	}
)

func (x JobTaskQuery_Status) Enum() *JobTaskQuery_Status {
	p := new(JobTaskQuery_Status)
	*p = x
	return p
}

func (x JobTaskQuery_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobTaskQuery_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes[0].Descriptor()
}

func (JobTaskQuery_Status) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes[0]
}

func (x JobTaskQuery_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobTaskQuery_Status.Descriptor instead.
func (JobTaskQuery_Status) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{1, 0}
}

type JobTaskInfo_Status int32

const (
	JobTaskInfo_SCOPE_NONE  JobTaskInfo_Status = 0
	JobTaskInfo_PENDING     JobTaskInfo_Status = 1
	JobTaskInfo_IN_PROGRESS JobTaskInfo_Status = 2
	JobTaskInfo_SUCCESS     JobTaskInfo_Status = 3
	JobTaskInfo_FAILURE     JobTaskInfo_Status = 4
	JobTaskInfo_TIMEOUT     JobTaskInfo_Status = 5
	JobTaskInfo_CANCELED    JobTaskInfo_Status = 6
)

// Enum value maps for JobTaskInfo_Status.
var (
	JobTaskInfo_Status_name = map[int32]string{
		0: "SCOPE_NONE",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "SUCCESS",
		4: "FAILURE",
		5: "TIMEOUT",
		6: "CANCELED",
	}
	JobTaskInfo_Status_value = map[string]int32{
		"SCOPE_NONE":  0,
		"PENDING":     1,
		"IN_PROGRESS": 2,
		"SUCCESS":     3,
		"FAILURE":     4,
		"TIMEOUT":     5,
		"CANCELED":    6,
	}
)

func (x JobTaskInfo_Status) Enum() *JobTaskInfo_Status {
	p := new(JobTaskInfo_Status)
	*p = x
	return p
}

func (x JobTaskInfo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobTaskInfo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes[1].Descriptor()
}

func (JobTaskInfo_Status) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes[1]
}

func (x JobTaskInfo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobTaskInfo_Status.Descriptor instead.
func (JobTaskInfo_Status) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{2, 0}
}

type JobTaskInfo_ResourceGroup int32

const (
	JobTaskInfo_RESOURCE_GROUP_NONE JobTaskInfo_ResourceGroup = 0
	JobTaskInfo_DOMAIN              JobTaskInfo_ResourceGroup = 1
	JobTaskInfo_WORKSPACE           JobTaskInfo_ResourceGroup = 2
)

// Enum value maps for JobTaskInfo_ResourceGroup.
var (
	JobTaskInfo_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "DOMAIN",
		2: "WORKSPACE",
	}
	JobTaskInfo_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"DOMAIN":              1,
		"WORKSPACE":           2,
	}
)

func (x JobTaskInfo_ResourceGroup) Enum() *JobTaskInfo_ResourceGroup {
	p := new(JobTaskInfo_ResourceGroup)
	*p = x
	return p
}

func (x JobTaskInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobTaskInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes[2].Descriptor()
}

func (JobTaskInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes[2]
}

func (x JobTaskInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobTaskInfo_ResourceGroup.Descriptor instead.
func (JobTaskInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{2, 1}
}

//	{
//	   "job_task_id": "job-task-3622d860a776"
//	}
type JobTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobTaskId     string                 `protobuf:"bytes,1,opt,name=job_task_id,json=jobTaskId,proto3" json:"job_task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTaskRequest) Reset() {
	*x = JobTaskRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskRequest) ProtoMessage() {}

func (x *JobTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskRequest.ProtoReflect.Descriptor instead.
func (*JobTaskRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{0}
}

func (x *JobTaskRequest) GetJobTaskId() string {
	if x != nil {
		return x.JobTaskId
	}
	return ""
}

//	{
//	   "data_source_id": "ds-12asdsf2a",
//	   "query": {}
//	}
type JobTaskQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	JobTaskId string `protobuf:"bytes,2,opt,name=job_task_id,json=jobTaskId,proto3" json:"job_task_id,omitempty"`
	// +optional
	Status JobTaskQuery_Status `protobuf:"varint,3,opt,name=status,proto3,enum=spaceone.api.cost_analysis.v1.JobTaskQuery_Status" json:"status,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	JobId string `protobuf:"bytes,22,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// +optional
	DataSourceId  string `protobuf:"bytes,23,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTaskQuery) Reset() {
	*x = JobTaskQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTaskQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskQuery) ProtoMessage() {}

func (x *JobTaskQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskQuery.ProtoReflect.Descriptor instead.
func (*JobTaskQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{1}
}

func (x *JobTaskQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *JobTaskQuery) GetJobTaskId() string {
	if x != nil {
		return x.JobTaskId
	}
	return ""
}

func (x *JobTaskQuery) GetStatus() JobTaskQuery_Status {
	if x != nil {
		return x.Status
	}
	return JobTaskQuery_SCOPE_NONE
}

func (x *JobTaskQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *JobTaskQuery) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobTaskQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

//	{
//	   "job_task_id": "job-task-3622d860a776",
//	   "status": "SUCCESS",
//	   "options": {
//	       "month": "202207",
//	       "platform": "gcp"
//	   },
//	   "changed": {
//	       "start": "2024-12",
//	       "end": "2025-01",
//	       "filter": {
//	           "additional_info.Account ID": "3241123123"
//	       }
//	   },
//	   "created_count": 1,
//	   "job_id": "job-85cf2c385252",
//	   "data_source_id": "ds-c96609f5afeb",
//	   "domain_id": "domain-58010aa2e451",
//	   "created_at": "2022-07-17T16:00:08.266Z",
//	   "started_at": "2022-07-17T16:01:28.243Z",
//	   "updated_at": "2022-07-17T16:01:28.939Z",
//	   "finished_at": "2022-07-17T16:01:28.939Z"
//	}
type JobTaskInfo struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	JobTaskId     string                    `protobuf:"bytes,1,opt,name=job_task_id,json=jobTaskId,proto3" json:"job_task_id,omitempty"`
	Status        JobTaskInfo_Status        `protobuf:"varint,2,opt,name=status,proto3,enum=spaceone.api.cost_analysis.v1.JobTaskInfo_Status" json:"status,omitempty"`
	Options       *_struct.Struct           `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	Changed       *ChangedInfo              `protobuf:"bytes,4,opt,name=changed,proto3" json:"changed,omitempty"`
	CreatedCount  int32                     `protobuf:"varint,5,opt,name=created_count,json=createdCount,proto3" json:"created_count,omitempty"`
	ErrorCode     string                    `protobuf:"bytes,6,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage  string                    `protobuf:"bytes,7,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	ResourceGroup JobTaskInfo_ResourceGroup `protobuf:"varint,20,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.cost_analysis.v1.JobTaskInfo_ResourceGroup" json:"resource_group,omitempty"`
	DomainId      string                    `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId   string                    `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	JobId         string                    `protobuf:"bytes,23,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DataSourceId  string                    `protobuf:"bytes,24,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	CreatedAt     string                    `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt     string                    `protobuf:"bytes,32,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	UpdatedAt     string                    `protobuf:"bytes,33,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	FinishedAt    string                    `protobuf:"bytes,34,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTaskInfo) Reset() {
	*x = JobTaskInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskInfo) ProtoMessage() {}

func (x *JobTaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskInfo.ProtoReflect.Descriptor instead.
func (*JobTaskInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{2}
}

func (x *JobTaskInfo) GetJobTaskId() string {
	if x != nil {
		return x.JobTaskId
	}
	return ""
}

func (x *JobTaskInfo) GetStatus() JobTaskInfo_Status {
	if x != nil {
		return x.Status
	}
	return JobTaskInfo_SCOPE_NONE
}

func (x *JobTaskInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *JobTaskInfo) GetChanged() *ChangedInfo {
	if x != nil {
		return x.Changed
	}
	return nil
}

func (x *JobTaskInfo) GetCreatedCount() int32 {
	if x != nil {
		return x.CreatedCount
	}
	return 0
}

func (x *JobTaskInfo) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *JobTaskInfo) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *JobTaskInfo) GetResourceGroup() JobTaskInfo_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return JobTaskInfo_RESOURCE_GROUP_NONE
}

func (x *JobTaskInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *JobTaskInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *JobTaskInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobTaskInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *JobTaskInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *JobTaskInfo) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *JobTaskInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *JobTaskInfo) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "job_task_id": "job-task-3622d860a776",
//	           "status": "SUCCESS",
//	           "options": {
//	               "platform": "gcp",
//	               "month": "202207"
//	           },
//	           "created_count": 1,
//	           "job_id": "job-85cf2c385252",
//	           "data_source_id": "ds-c96609f5afeb",
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2022-07-17T16:00:08.266Z",
//	           "started_at": "2022-07-17T16:01:28.243Z",
//	           "updated_at": "2022-07-17T16:01:28.939Z",
//	           "finished_at": "2022-07-17T16:01:28.939Z"
//	       },
//	       {
//	           "job_task_id": "job-task-038c0b076ec5",
//	           "status": "SUCCESS",
//	           "options": {
//	               "account": "257706363616",
//	               "start": "2022-07-01"
//	           },
//	           "created_count": 5756,
//	           "job_id": "job-6b6765f757a9",
//	           "data_source_id": "ds-fcba92ca73b1",
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2022-07-17T16:00:05.099Z",
//	           "started_at": "2022-07-17T16:00:47.356Z",
//	           "updated_at": "2022-07-17T16:01:20.856Z",
//	           "finished_at": "2022-07-17T16:01:20.856Z"
//	       }
//	   ],
//	   "total_count": 720
//	}
type JobTasksInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*JobTaskInfo         `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTasksInfo) Reset() {
	*x = JobTasksInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTasksInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTasksInfo) ProtoMessage() {}

func (x *JobTasksInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTasksInfo.ProtoReflect.Descriptor instead.
func (*JobTasksInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{3}
}

func (x *JobTasksInfo) GetResults() []*JobTaskInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *JobTasksInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type JobTaskStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobTaskStatQuery) Reset() {
	*x = JobTaskStatQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobTaskStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobTaskStatQuery) ProtoMessage() {}

func (x *JobTaskStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobTaskStatQuery.ProtoReflect.Descriptor instead.
func (*JobTaskStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP(), []int{4}
}

func (x *JobTaskStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_cost_analysis_v1_job_task_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_job_task_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62,
	0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0xfa, 0x02, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x6f,
	0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45,
	0x44, 0x10, 0x06, 0x22, 0xe8, 0x06, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x44, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x5f, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6b, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x22, 0x43, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x22, 0x75,
	0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x32, 0xa4, 0x03, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x8b, 0x01, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a,
	0x22, 0x1e, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x65, 0x74,
	0x12, 0x8c, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73,
	0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x7c, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x6f, 0x62, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescData []byte
)

func file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_job_task_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_job_task_proto_rawDesc)))
	})
	return file_spaceone_api_cost_analysis_v1_job_task_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_cost_analysis_v1_job_task_proto_goTypes = []any{
	(JobTaskQuery_Status)(0),       // 0: spaceone.api.cost_analysis.v1.JobTaskQuery.Status
	(JobTaskInfo_Status)(0),        // 1: spaceone.api.cost_analysis.v1.JobTaskInfo.Status
	(JobTaskInfo_ResourceGroup)(0), // 2: spaceone.api.cost_analysis.v1.JobTaskInfo.ResourceGroup
	(*JobTaskRequest)(nil),         // 3: spaceone.api.cost_analysis.v1.JobTaskRequest
	(*JobTaskQuery)(nil),           // 4: spaceone.api.cost_analysis.v1.JobTaskQuery
	(*JobTaskInfo)(nil),            // 5: spaceone.api.cost_analysis.v1.JobTaskInfo
	(*JobTasksInfo)(nil),           // 6: spaceone.api.cost_analysis.v1.JobTasksInfo
	(*JobTaskStatQuery)(nil),       // 7: spaceone.api.cost_analysis.v1.JobTaskStatQuery
	(*v2.Query)(nil),               // 8: spaceone.api.core.v2.Query
	(*_struct.Struct)(nil),         // 9: google.protobuf.Struct
	(*ChangedInfo)(nil),            // 10: spaceone.api.cost_analysis.v1.ChangedInfo
	(*v2.StatisticsQuery)(nil),     // 11: spaceone.api.core.v2.StatisticsQuery
}
var file_spaceone_api_cost_analysis_v1_job_task_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.cost_analysis.v1.JobTaskQuery.query:type_name -> spaceone.api.core.v2.Query
	0,  // 1: spaceone.api.cost_analysis.v1.JobTaskQuery.status:type_name -> spaceone.api.cost_analysis.v1.JobTaskQuery.Status
	1,  // 2: spaceone.api.cost_analysis.v1.JobTaskInfo.status:type_name -> spaceone.api.cost_analysis.v1.JobTaskInfo.Status
	9,  // 3: spaceone.api.cost_analysis.v1.JobTaskInfo.options:type_name -> google.protobuf.Struct
	10, // 4: spaceone.api.cost_analysis.v1.JobTaskInfo.changed:type_name -> spaceone.api.cost_analysis.v1.ChangedInfo
	2,  // 5: spaceone.api.cost_analysis.v1.JobTaskInfo.resource_group:type_name -> spaceone.api.cost_analysis.v1.JobTaskInfo.ResourceGroup
	5,  // 6: spaceone.api.cost_analysis.v1.JobTasksInfo.results:type_name -> spaceone.api.cost_analysis.v1.JobTaskInfo
	11, // 7: spaceone.api.cost_analysis.v1.JobTaskStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	3,  // 8: spaceone.api.cost_analysis.v1.JobTask.get:input_type -> spaceone.api.cost_analysis.v1.JobTaskRequest
	4,  // 9: spaceone.api.cost_analysis.v1.JobTask.list:input_type -> spaceone.api.cost_analysis.v1.JobTaskQuery
	7,  // 10: spaceone.api.cost_analysis.v1.JobTask.stat:input_type -> spaceone.api.cost_analysis.v1.JobTaskStatQuery
	5,  // 11: spaceone.api.cost_analysis.v1.JobTask.get:output_type -> spaceone.api.cost_analysis.v1.JobTaskInfo
	6,  // 12: spaceone.api.cost_analysis.v1.JobTask.list:output_type -> spaceone.api.cost_analysis.v1.JobTasksInfo
	9,  // 13: spaceone.api.cost_analysis.v1.JobTask.stat:output_type -> google.protobuf.Struct
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_job_task_proto_init() }
func file_spaceone_api_cost_analysis_v1_job_task_proto_init() {
	if File_spaceone_api_cost_analysis_v1_job_task_proto != nil {
		return
	}
	file_spaceone_api_cost_analysis_v1_job_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_job_task_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_job_task_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_job_task_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_job_task_proto_depIdxs,
		EnumInfos:         file_spaceone_api_cost_analysis_v1_job_task_proto_enumTypes,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_job_task_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_job_task_proto = out.File
	file_spaceone_api_cost_analysis_v1_job_task_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_job_task_proto_depIdxs = nil
}
