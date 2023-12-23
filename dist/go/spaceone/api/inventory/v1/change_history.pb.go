// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: spaceone/api/inventory/v1/change_history.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChangeHistoryQuery_RecordAction int32

const (
	ChangeHistoryQuery_NONE   ChangeHistoryQuery_RecordAction = 0
	ChangeHistoryQuery_CREATE ChangeHistoryQuery_RecordAction = 1
	ChangeHistoryQuery_UPDATE ChangeHistoryQuery_RecordAction = 2
	ChangeHistoryQuery_DELETE ChangeHistoryQuery_RecordAction = 3
)

// Enum value maps for ChangeHistoryQuery_RecordAction.
var (
	ChangeHistoryQuery_RecordAction_name = map[int32]string{
		0: "NONE",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	ChangeHistoryQuery_RecordAction_value = map[string]int32{
		"NONE":   0,
		"CREATE": 1,
		"UPDATE": 2,
		"DELETE": 3,
	}
)

func (x ChangeHistoryQuery_RecordAction) Enum() *ChangeHistoryQuery_RecordAction {
	p := new(ChangeHistoryQuery_RecordAction)
	*p = x
	return p
}

func (x ChangeHistoryQuery_RecordAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeHistoryQuery_RecordAction) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_inventory_v1_change_history_proto_enumTypes[0].Descriptor()
}

func (ChangeHistoryQuery_RecordAction) Type() protoreflect.EnumType {
	return &file_spaceone_api_inventory_v1_change_history_proto_enumTypes[0]
}

func (x ChangeHistoryQuery_RecordAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeHistoryQuery_RecordAction.Descriptor instead.
func (ChangeHistoryQuery_RecordAction) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP(), []int{0, 0}
}

type RecordInfo_RecordAction int32

const (
	RecordInfo_NONE   RecordInfo_RecordAction = 0
	RecordInfo_CREATE RecordInfo_RecordAction = 1
	RecordInfo_UPDATE RecordInfo_RecordAction = 2
	RecordInfo_DELETE RecordInfo_RecordAction = 3
)

// Enum value maps for RecordInfo_RecordAction.
var (
	RecordInfo_RecordAction_name = map[int32]string{
		0: "NONE",
		1: "CREATE",
		2: "UPDATE",
		3: "DELETE",
	}
	RecordInfo_RecordAction_value = map[string]int32{
		"NONE":   0,
		"CREATE": 1,
		"UPDATE": 2,
		"DELETE": 3,
	}
)

func (x RecordInfo_RecordAction) Enum() *RecordInfo_RecordAction {
	p := new(RecordInfo_RecordAction)
	*p = x
	return p
}

func (x RecordInfo_RecordAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordInfo_RecordAction) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_inventory_v1_change_history_proto_enumTypes[1].Descriptor()
}

func (RecordInfo_RecordAction) Type() protoreflect.EnumType {
	return &file_spaceone_api_inventory_v1_change_history_proto_enumTypes[1]
}

func (x RecordInfo_RecordAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordInfo_RecordAction.Descriptor instead.
func (RecordInfo_RecordAction) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP(), []int{1, 0}
}

type ChangeHistoryQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query          *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	CloudServiceId string    `protobuf:"bytes,2,opt,name=cloud_service_id,json=cloudServiceId,proto3" json:"cloud_service_id,omitempty"`
	// +optional
	Action ChangeHistoryQuery_RecordAction `protobuf:"varint,3,opt,name=action,proto3,enum=spaceone.api.inventory.v1.ChangeHistoryQuery_RecordAction" json:"action,omitempty"`
	// +optional
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// +optional
	CollectorId string `protobuf:"bytes,5,opt,name=collector_id,json=collectorId,proto3" json:"collector_id,omitempty"`
	// +optional
	JobId string `protobuf:"bytes,6,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// +optional
	UpdatedBy string `protobuf:"bytes,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,22,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ChangeHistoryQuery) Reset() {
	*x = ChangeHistoryQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHistoryQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHistoryQuery) ProtoMessage() {}

func (x *ChangeHistoryQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHistoryQuery.ProtoReflect.Descriptor instead.
func (*ChangeHistoryQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeHistoryQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ChangeHistoryQuery) GetCloudServiceId() string {
	if x != nil {
		return x.CloudServiceId
	}
	return ""
}

func (x *ChangeHistoryQuery) GetAction() ChangeHistoryQuery_RecordAction {
	if x != nil {
		return x.Action
	}
	return ChangeHistoryQuery_NONE
}

func (x *ChangeHistoryQuery) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChangeHistoryQuery) GetCollectorId() string {
	if x != nil {
		return x.CollectorId
	}
	return ""
}

func (x *ChangeHistoryQuery) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *ChangeHistoryQuery) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *ChangeHistoryQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *ChangeHistoryQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type RecordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId       string                  `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	CloudServiceId string                  `protobuf:"bytes,2,opt,name=cloud_service_id,json=cloudServiceId,proto3" json:"cloud_service_id,omitempty"`
	Action         RecordInfo_RecordAction `protobuf:"varint,3,opt,name=action,proto3,enum=spaceone.api.inventory.v1.RecordInfo_RecordAction" json:"action,omitempty"`
	Diff           *_struct.ListValue      `protobuf:"bytes,4,opt,name=diff,proto3" json:"diff,omitempty"`
	DiffCount      int32                   `protobuf:"varint,5,opt,name=diff_count,json=diffCount,proto3" json:"diff_count,omitempty"`
	UserId         string                  `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CollectorId    string                  `protobuf:"bytes,7,opt,name=collector_id,json=collectorId,proto3" json:"collector_id,omitempty"`
	JobId          string                  `protobuf:"bytes,8,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	UpdatedBy      string                  `protobuf:"bytes,9,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	DomainId       string                  `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId    string                  `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId      string                  `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CreatedAt      string                  `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *RecordInfo) Reset() {
	*x = RecordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordInfo) ProtoMessage() {}

func (x *RecordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordInfo.ProtoReflect.Descriptor instead.
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP(), []int{1}
}

func (x *RecordInfo) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *RecordInfo) GetCloudServiceId() string {
	if x != nil {
		return x.CloudServiceId
	}
	return ""
}

func (x *RecordInfo) GetAction() RecordInfo_RecordAction {
	if x != nil {
		return x.Action
	}
	return RecordInfo_NONE
}

func (x *RecordInfo) GetDiff() *_struct.ListValue {
	if x != nil {
		return x.Diff
	}
	return nil
}

func (x *RecordInfo) GetDiffCount() int32 {
	if x != nil {
		return x.DiffCount
	}
	return 0
}

func (x *RecordInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RecordInfo) GetCollectorId() string {
	if x != nil {
		return x.CollectorId
	}
	return ""
}

func (x *RecordInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *RecordInfo) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *RecordInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *RecordInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *RecordInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *RecordInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ChangeHistoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*RecordInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ChangeHistoryInfo) Reset() {
	*x = ChangeHistoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHistoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHistoryInfo) ProtoMessage() {}

func (x *ChangeHistoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHistoryInfo.ProtoReflect.Descriptor instead.
func (*ChangeHistoryInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeHistoryInfo) GetResults() []*RecordInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ChangeHistoryInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ChangeHistoryStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query          *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	CloudServiceId string              `protobuf:"bytes,3,opt,name=cloud_service_id,json=cloudServiceId,proto3" json:"cloud_service_id,omitempty"`
}

func (x *ChangeHistoryStatQuery) Reset() {
	*x = ChangeHistoryStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHistoryStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHistoryStatQuery) ProtoMessage() {}

func (x *ChangeHistoryStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_change_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHistoryStatQuery.ProtoReflect.Descriptor instead.
func (*ChangeHistoryStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeHistoryStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ChangeHistoryStatQuery) GetCloudServiceId() string {
	if x != nil {
		return x.CloudServiceId
	}
	return ""
}

var File_spaceone_api_inventory_v1_change_history_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_v1_change_history_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x03, 0x0a, 0x12, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x52, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x03, 0x22, 0x9c, 0x04, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x64, 0x69, 0x66, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x69, 0x66, 0x66, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x03, 0x22, 0x75, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7f, 0x0a, 0x16, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x28, 0x0a, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x32, 0xa6, 0x02, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x91, 0x01, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x80, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a,
	0x22, 0x21, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_inventory_v1_change_history_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v1_change_history_proto_rawDescData = file_spaceone_api_inventory_v1_change_history_proto_rawDesc
)

func file_spaceone_api_inventory_v1_change_history_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v1_change_history_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v1_change_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_inventory_v1_change_history_proto_rawDescData)
	})
	return file_spaceone_api_inventory_v1_change_history_proto_rawDescData
}

var file_spaceone_api_inventory_v1_change_history_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_inventory_v1_change_history_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spaceone_api_inventory_v1_change_history_proto_goTypes = []interface{}{
	(ChangeHistoryQuery_RecordAction)(0), // 0: spaceone.api.inventory.v1.ChangeHistoryQuery.RecordAction
	(RecordInfo_RecordAction)(0),         // 1: spaceone.api.inventory.v1.RecordInfo.RecordAction
	(*ChangeHistoryQuery)(nil),           // 2: spaceone.api.inventory.v1.ChangeHistoryQuery
	(*RecordInfo)(nil),                   // 3: spaceone.api.inventory.v1.RecordInfo
	(*ChangeHistoryInfo)(nil),            // 4: spaceone.api.inventory.v1.ChangeHistoryInfo
	(*ChangeHistoryStatQuery)(nil),       // 5: spaceone.api.inventory.v1.ChangeHistoryStatQuery
	(*v2.Query)(nil),                     // 6: spaceone.api.core.v2.Query
	(*_struct.ListValue)(nil),            // 7: google.protobuf.ListValue
	(*v2.StatisticsQuery)(nil),           // 8: spaceone.api.core.v2.StatisticsQuery
	(*_struct.Struct)(nil),               // 9: google.protobuf.Struct
}
var file_spaceone_api_inventory_v1_change_history_proto_depIdxs = []int32{
	6, // 0: spaceone.api.inventory.v1.ChangeHistoryQuery.query:type_name -> spaceone.api.core.v2.Query
	0, // 1: spaceone.api.inventory.v1.ChangeHistoryQuery.action:type_name -> spaceone.api.inventory.v1.ChangeHistoryQuery.RecordAction
	1, // 2: spaceone.api.inventory.v1.RecordInfo.action:type_name -> spaceone.api.inventory.v1.RecordInfo.RecordAction
	7, // 3: spaceone.api.inventory.v1.RecordInfo.diff:type_name -> google.protobuf.ListValue
	3, // 4: spaceone.api.inventory.v1.ChangeHistoryInfo.results:type_name -> spaceone.api.inventory.v1.RecordInfo
	8, // 5: spaceone.api.inventory.v1.ChangeHistoryStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	2, // 6: spaceone.api.inventory.v1.ChangeHistory.list:input_type -> spaceone.api.inventory.v1.ChangeHistoryQuery
	5, // 7: spaceone.api.inventory.v1.ChangeHistory.stat:input_type -> spaceone.api.inventory.v1.ChangeHistoryStatQuery
	4, // 8: spaceone.api.inventory.v1.ChangeHistory.list:output_type -> spaceone.api.inventory.v1.ChangeHistoryInfo
	9, // 9: spaceone.api.inventory.v1.ChangeHistory.stat:output_type -> google.protobuf.Struct
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v1_change_history_proto_init() }
func file_spaceone_api_inventory_v1_change_history_proto_init() {
	if File_spaceone_api_inventory_v1_change_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_inventory_v1_change_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeHistoryQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_inventory_v1_change_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_inventory_v1_change_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeHistoryInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_inventory_v1_change_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeHistoryStatQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_inventory_v1_change_history_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v1_change_history_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v1_change_history_proto_depIdxs,
		EnumInfos:         file_spaceone_api_inventory_v1_change_history_proto_enumTypes,
		MessageInfos:      file_spaceone_api_inventory_v1_change_history_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v1_change_history_proto = out.File
	file_spaceone_api_inventory_v1_change_history_proto_rawDesc = nil
	file_spaceone_api_inventory_v1_change_history_proto_goTypes = nil
	file_spaceone_api_inventory_v1_change_history_proto_depIdxs = nil
}
