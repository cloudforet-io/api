// A ProjectAlertConfig is a resource to set the alert policies for a Project.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/monitoring/v1/project_alert_config.proto

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

type AlertOptions_UrgencyOption int32

const (
	AlertOptions_URGENCY_NONE AlertOptions_UrgencyOption = 0
	AlertOptions_ALL          AlertOptions_UrgencyOption = 1
	AlertOptions_HIGH_ONLY    AlertOptions_UrgencyOption = 2
)

// Enum value maps for AlertOptions_UrgencyOption.
var (
	AlertOptions_UrgencyOption_name = map[int32]string{
		0: "URGENCY_NONE",
		1: "ALL",
		2: "HIGH_ONLY",
	}
	AlertOptions_UrgencyOption_value = map[string]int32{
		"URGENCY_NONE": 0,
		"ALL":          1,
		"HIGH_ONLY":    2,
	}
)

func (x AlertOptions_UrgencyOption) Enum() *AlertOptions_UrgencyOption {
	p := new(AlertOptions_UrgencyOption)
	*p = x
	return p
}

func (x AlertOptions_UrgencyOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertOptions_UrgencyOption) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_enumTypes[0].Descriptor()
}

func (AlertOptions_UrgencyOption) Type() protoreflect.EnumType {
	return &file_spaceone_api_monitoring_v1_project_alert_config_proto_enumTypes[0]
}

func (x AlertOptions_UrgencyOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertOptions_UrgencyOption.Descriptor instead.
func (AlertOptions_UrgencyOption) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{0, 0}
}

type AlertOptions_RecoveryOption int32

const (
	AlertOptions_RECOVERY_NONE AlertOptions_RecoveryOption = 0
	AlertOptions_AUTO          AlertOptions_RecoveryOption = 1
	AlertOptions_MANUAL        AlertOptions_RecoveryOption = 2
)

// Enum value maps for AlertOptions_RecoveryOption.
var (
	AlertOptions_RecoveryOption_name = map[int32]string{
		0: "RECOVERY_NONE",
		1: "AUTO",
		2: "MANUAL",
	}
	AlertOptions_RecoveryOption_value = map[string]int32{
		"RECOVERY_NONE": 0,
		"AUTO":          1,
		"MANUAL":        2,
	}
)

func (x AlertOptions_RecoveryOption) Enum() *AlertOptions_RecoveryOption {
	p := new(AlertOptions_RecoveryOption)
	*p = x
	return p
}

func (x AlertOptions_RecoveryOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertOptions_RecoveryOption) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_enumTypes[1].Descriptor()
}

func (AlertOptions_RecoveryOption) Type() protoreflect.EnumType {
	return &file_spaceone_api_monitoring_v1_project_alert_config_proto_enumTypes[1]
}

func (x AlertOptions_RecoveryOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertOptions_RecoveryOption.Descriptor instead.
func (AlertOptions_RecoveryOption) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{0, 1}
}

type AlertOptions struct {
	state               protoimpl.MessageState      `protogen:"open.v1"`
	NotificationUrgency AlertOptions_UrgencyOption  `protobuf:"varint,1,opt,name=notification_urgency,json=notificationUrgency,proto3,enum=spaceone.api.monitoring.v1.AlertOptions_UrgencyOption" json:"notification_urgency,omitempty"`
	RecoveryMode        AlertOptions_RecoveryOption `protobuf:"varint,2,opt,name=recovery_mode,json=recoveryMode,proto3,enum=spaceone.api.monitoring.v1.AlertOptions_RecoveryOption" json:"recovery_mode,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AlertOptions) Reset() {
	*x = AlertOptions{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertOptions) ProtoMessage() {}

func (x *AlertOptions) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertOptions.ProtoReflect.Descriptor instead.
func (*AlertOptions) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{0}
}

func (x *AlertOptions) GetNotificationUrgency() AlertOptions_UrgencyOption {
	if x != nil {
		return x.NotificationUrgency
	}
	return AlertOptions_URGENCY_NONE
}

func (x *AlertOptions) GetRecoveryMode() AlertOptions_RecoveryOption {
	if x != nil {
		return x.RecoveryMode
	}
	return AlertOptions_RECOVERY_NONE
}

//	{
//	   "project_id": "project-dee2a81d4859",
//	   "escalation_policy_id": "ep-b441abe04ca9",
//	   "options": {
//	       "notification_urgency": "ALL",
//	       "recovery_mode": "AUTO"
//	   },
//	   "domain_id": "domain-58010aa2e451"
//	}
type CreateProjectAlertConfigRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	ProjectId string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	EscalationPolicyId string `protobuf:"bytes,2,opt,name=escalation_policy_id,json=escalationPolicyId,proto3" json:"escalation_policy_id,omitempty"`
	// +optional
	Options       *AlertOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProjectAlertConfigRequest) Reset() {
	*x = CreateProjectAlertConfigRequest{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProjectAlertConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectAlertConfigRequest) ProtoMessage() {}

func (x *CreateProjectAlertConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectAlertConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectAlertConfigRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectAlertConfigRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateProjectAlertConfigRequest) GetEscalationPolicyId() string {
	if x != nil {
		return x.EscalationPolicyId
	}
	return ""
}

func (x *CreateProjectAlertConfigRequest) GetOptions() *AlertOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

//	{
//	   "project_id": "project-dee2a81d4859",
//	   "escalation_policy_id": "ep-4ee42a9b2d96",
//	   "options": {
//	       "notification_urgency": "ALL",
//	       "recovery_mode": "MANUAL"
//	   },
//	   "domain_id": "domain-58010aa2e451"
//	}
type UpdateProjectAlertConfigRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	ProjectId string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	EscalationPolicyId string `protobuf:"bytes,2,opt,name=escalation_policy_id,json=escalationPolicyId,proto3" json:"escalation_policy_id,omitempty"`
	// +optional
	Options       *AlertOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProjectAlertConfigRequest) Reset() {
	*x = UpdateProjectAlertConfigRequest{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProjectAlertConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectAlertConfigRequest) ProtoMessage() {}

func (x *UpdateProjectAlertConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectAlertConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectAlertConfigRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProjectAlertConfigRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateProjectAlertConfigRequest) GetEscalationPolicyId() string {
	if x != nil {
		return x.EscalationPolicyId
	}
	return ""
}

func (x *UpdateProjectAlertConfigRequest) GetOptions() *AlertOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

//	{
//	   "project_id": "project-dee2a81d4859",
//	   "domain_id": "domain-58010aa2e451"
//	}
type ProjectAlertConfigRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectAlertConfigRequest) Reset() {
	*x = ProjectAlertConfigRequest{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectAlertConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAlertConfigRequest) ProtoMessage() {}

func (x *ProjectAlertConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAlertConfigRequest.ProtoReflect.Descriptor instead.
func (*ProjectAlertConfigRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectAlertConfigRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

//	{
//	   "query": {},
//	   "domain_id": "domain-58010aa2e451"
//	}
type ProjectAlertConfigQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	EscalationPolicyId string `protobuf:"bytes,21,opt,name=escalation_policy_id,json=escalationPolicyId,proto3" json:"escalation_policy_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ProjectAlertConfigQuery) Reset() {
	*x = ProjectAlertConfigQuery{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectAlertConfigQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAlertConfigQuery) ProtoMessage() {}

func (x *ProjectAlertConfigQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAlertConfigQuery.ProtoReflect.Descriptor instead.
func (*ProjectAlertConfigQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectAlertConfigQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ProjectAlertConfigQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectAlertConfigQuery) GetEscalationPolicyId() string {
	if x != nil {
		return x.EscalationPolicyId
	}
	return ""
}

//	{
//	   "project_id": "project-dee2a81d4859",
//	   "options": {
//	       "notification_urgency": "ALL",
//	       "recovery_mode": "AUTO"
//	   },
//	   "escalation_policy_info": {
//	       "escalation_policy_id": "ep-b441abe04ca9",
//	       "name": "Global New Policy"
//	   },
//	   "domain_id": "domain-58010aa2e451",
//	   "created_at": "2022-06-27T05:12:22.998Z"
//	}
type ProjectAlertConfigInfo struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	ProjectId            string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Options              *AlertOptions          `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	EscalationPolicyInfo *EscalationPolicyInfo  `protobuf:"bytes,3,opt,name=escalation_policy_info,json=escalationPolicyInfo,proto3" json:"escalation_policy_info,omitempty"`
	DomainId             string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId          string                 `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	CreatedAt            string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ProjectAlertConfigInfo) Reset() {
	*x = ProjectAlertConfigInfo{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectAlertConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAlertConfigInfo) ProtoMessage() {}

func (x *ProjectAlertConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAlertConfigInfo.ProtoReflect.Descriptor instead.
func (*ProjectAlertConfigInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectAlertConfigInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectAlertConfigInfo) GetOptions() *AlertOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ProjectAlertConfigInfo) GetEscalationPolicyInfo() *EscalationPolicyInfo {
	if x != nil {
		return x.EscalationPolicyInfo
	}
	return nil
}

func (x *ProjectAlertConfigInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *ProjectAlertConfigInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *ProjectAlertConfigInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "project_id": "project-18655561c535",
//	           "options": {
//	               "notification_urgency": "ALL",
//	               "recovery_mode": "MANUAL"
//	           },
//	           "escalation_policy_info": {
//	               "escalation_policy_id": "ep-4ee42a9b2d96",
//	               "name": "HAHA",
//	               "is_default": true
//	           },
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2022-05-17T02:09:19.839Z"
//	       },
//	       {
//	           "project_id": "project-9074eea97d7e",
//	           "options": {
//	               "notification_urgency": "ALL",
//	               "recovery_mode": "MANUAL"
//	           },
//	           "escalation_policy_info": {
//	               "escalation_policy_id": "ep-b441abe04ca9",
//	               "name": "Global New Policy"
//	           },
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2021-06-24T02:50:50.535Z"
//	       }
//	   ],
//	   "total_count": 2
//	}
type ProjectAlertConfigsInfo struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Results       []*ProjectAlertConfigInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                     `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectAlertConfigsInfo) Reset() {
	*x = ProjectAlertConfigsInfo{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectAlertConfigsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAlertConfigsInfo) ProtoMessage() {}

func (x *ProjectAlertConfigsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAlertConfigsInfo.ProtoReflect.Descriptor instead.
func (*ProjectAlertConfigsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectAlertConfigsInfo) GetResults() []*ProjectAlertConfigInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ProjectAlertConfigsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type ProjectAlertConfigStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectAlertConfigStatQuery) Reset() {
	*x = ProjectAlertConfigStatQuery{}
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectAlertConfigStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAlertConfigStatQuery) ProtoMessage() {}

func (x *ProjectAlertConfigStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAlertConfigStatQuery.ProtoReflect.Descriptor instead.
func (*ProjectAlertConfigStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectAlertConfigStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_monitoring_v1_project_alert_config_proto protoreflect.FileDescriptor

const file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDesc = "" +
	"\n" +
	"5spaceone/api/monitoring/v1/project_alert_config.proto\x12\x1aspaceone.api.monitoring.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a2spaceone/api/monitoring/v1/escalation_policy.proto\"\xcd\x02\n" +
	"\fAlertOptions\x12i\n" +
	"\x14notification_urgency\x18\x01 \x01(\x0e26.spaceone.api.monitoring.v1.AlertOptions.UrgencyOptionR\x13notificationUrgency\x12\\\n" +
	"\rrecovery_mode\x18\x02 \x01(\x0e27.spaceone.api.monitoring.v1.AlertOptions.RecoveryOptionR\frecoveryMode\"9\n" +
	"\rUrgencyOption\x12\x10\n" +
	"\fURGENCY_NONE\x10\x00\x12\a\n" +
	"\x03ALL\x10\x01\x12\r\n" +
	"\tHIGH_ONLY\x10\x02\"9\n" +
	"\x0eRecoveryOption\x12\x11\n" +
	"\rRECOVERY_NONE\x10\x00\x12\b\n" +
	"\x04AUTO\x10\x01\x12\n" +
	"\n" +
	"\x06MANUAL\x10\x02\"\xb6\x01\n" +
	"\x1fCreateProjectAlertConfigRequest\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\x120\n" +
	"\x14escalation_policy_id\x18\x02 \x01(\tR\x12escalationPolicyId\x12B\n" +
	"\aoptions\x18\x03 \x01(\v2(.spaceone.api.monitoring.v1.AlertOptionsR\aoptions\"\xb6\x01\n" +
	"\x1fUpdateProjectAlertConfigRequest\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\x120\n" +
	"\x14escalation_policy_id\x18\x02 \x01(\tR\x12escalationPolicyId\x12B\n" +
	"\aoptions\x18\x03 \x01(\v2(.spaceone.api.monitoring.v1.AlertOptionsR\aoptions\":\n" +
	"\x19ProjectAlertConfigRequest\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\"\x9d\x01\n" +
	"\x17ProjectAlertConfigQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x1d\n" +
	"\n" +
	"project_id\x18\x02 \x01(\tR\tprojectId\x120\n" +
	"\x14escalation_policy_id\x18\x15 \x01(\tR\x12escalationPolicyId\"\xc2\x02\n" +
	"\x16ProjectAlertConfigInfo\x12\x1d\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tR\tprojectId\x12B\n" +
	"\aoptions\x18\x02 \x01(\v2(.spaceone.api.monitoring.v1.AlertOptionsR\aoptions\x12f\n" +
	"\x16escalation_policy_info\x18\x03 \x01(\v20.spaceone.api.monitoring.v1.EscalationPolicyInfoR\x14escalationPolicyInfo\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12!\n" +
	"\fworkspace_id\x18\x16 \x01(\tR\vworkspaceId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\"\x88\x01\n" +
	"\x17ProjectAlertConfigsInfo\x12L\n" +
	"\aresults\x18\x01 \x03(\v22.spaceone.api.monitoring.v1.ProjectAlertConfigInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"Z\n" +
	"\x1bProjectAlertConfigStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\xea\a\n" +
	"\x12ProjectAlertConfig\x12\xb0\x01\n" +
	"\x06create\x12;.spaceone.api.monitoring.v1.CreateProjectAlertConfigRequest\x1a2.spaceone.api.monitoring.v1.ProjectAlertConfigInfo\"5\x82\xd3\xe4\x93\x02/:\x01*\"*/monitoring/v1/project-alert-config/create\x12\xb0\x01\n" +
	"\x06update\x12;.spaceone.api.monitoring.v1.UpdateProjectAlertConfigRequest\x1a2.spaceone.api.monitoring.v1.ProjectAlertConfigInfo\"5\x82\xd3\xe4\x93\x02/:\x01*\"*/monitoring/v1/project-alert-config/update\x12\x8e\x01\n" +
	"\x06delete\x125.spaceone.api.monitoring.v1.ProjectAlertConfigRequest\x1a\x16.google.protobuf.Empty\"5\x82\xd3\xe4\x93\x02/:\x01*\"*/monitoring/v1/project-alert-config/delete\x12\xa4\x01\n" +
	"\x03get\x125.spaceone.api.monitoring.v1.ProjectAlertConfigRequest\x1a2.spaceone.api.monitoring.v1.ProjectAlertConfigInfo\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/monitoring/v1/project-alert-config/get\x12\xa5\x01\n" +
	"\x04list\x123.spaceone.api.monitoring.v1.ProjectAlertConfigQuery\x1a3.spaceone.api.monitoring.v1.ProjectAlertConfigsInfo\"3\x82\xd3\xe4\x93\x02-:\x01*\"(/monitoring/v1/project-alert-config/list\x12\x8d\x01\n" +
	"\x04stat\x127.spaceone.api.monitoring.v1.ProjectAlertConfigStatQuery\x1a\x17.google.protobuf.Struct\"3\x82\xd3\xe4\x93\x02-:\x01*\"(/monitoring/v1/project-alert-config/statBAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1b\x06proto3"

var (
	file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescOnce sync.Once
	file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescData []byte
)

func file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescGZIP() []byte {
	file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescOnce.Do(func() {
		file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDesc), len(file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDesc)))
	})
	return file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDescData
}

var file_spaceone_api_monitoring_v1_project_alert_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_monitoring_v1_project_alert_config_proto_goTypes = []any{
	(AlertOptions_UrgencyOption)(0),         // 0: spaceone.api.monitoring.v1.AlertOptions.UrgencyOption
	(AlertOptions_RecoveryOption)(0),        // 1: spaceone.api.monitoring.v1.AlertOptions.RecoveryOption
	(*AlertOptions)(nil),                    // 2: spaceone.api.monitoring.v1.AlertOptions
	(*CreateProjectAlertConfigRequest)(nil), // 3: spaceone.api.monitoring.v1.CreateProjectAlertConfigRequest
	(*UpdateProjectAlertConfigRequest)(nil), // 4: spaceone.api.monitoring.v1.UpdateProjectAlertConfigRequest
	(*ProjectAlertConfigRequest)(nil),       // 5: spaceone.api.monitoring.v1.ProjectAlertConfigRequest
	(*ProjectAlertConfigQuery)(nil),         // 6: spaceone.api.monitoring.v1.ProjectAlertConfigQuery
	(*ProjectAlertConfigInfo)(nil),          // 7: spaceone.api.monitoring.v1.ProjectAlertConfigInfo
	(*ProjectAlertConfigsInfo)(nil),         // 8: spaceone.api.monitoring.v1.ProjectAlertConfigsInfo
	(*ProjectAlertConfigStatQuery)(nil),     // 9: spaceone.api.monitoring.v1.ProjectAlertConfigStatQuery
	(*v2.Query)(nil),                        // 10: spaceone.api.core.v2.Query
	(*EscalationPolicyInfo)(nil),            // 11: spaceone.api.monitoring.v1.EscalationPolicyInfo
	(*v2.StatisticsQuery)(nil),              // 12: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                     // 13: google.protobuf.Empty
	(*_struct.Struct)(nil),                  // 14: google.protobuf.Struct
}
var file_spaceone_api_monitoring_v1_project_alert_config_proto_depIdxs = []int32{
	0,  // 0: spaceone.api.monitoring.v1.AlertOptions.notification_urgency:type_name -> spaceone.api.monitoring.v1.AlertOptions.UrgencyOption
	1,  // 1: spaceone.api.monitoring.v1.AlertOptions.recovery_mode:type_name -> spaceone.api.monitoring.v1.AlertOptions.RecoveryOption
	2,  // 2: spaceone.api.monitoring.v1.CreateProjectAlertConfigRequest.options:type_name -> spaceone.api.monitoring.v1.AlertOptions
	2,  // 3: spaceone.api.monitoring.v1.UpdateProjectAlertConfigRequest.options:type_name -> spaceone.api.monitoring.v1.AlertOptions
	10, // 4: spaceone.api.monitoring.v1.ProjectAlertConfigQuery.query:type_name -> spaceone.api.core.v2.Query
	2,  // 5: spaceone.api.monitoring.v1.ProjectAlertConfigInfo.options:type_name -> spaceone.api.monitoring.v1.AlertOptions
	11, // 6: spaceone.api.monitoring.v1.ProjectAlertConfigInfo.escalation_policy_info:type_name -> spaceone.api.monitoring.v1.EscalationPolicyInfo
	7,  // 7: spaceone.api.monitoring.v1.ProjectAlertConfigsInfo.results:type_name -> spaceone.api.monitoring.v1.ProjectAlertConfigInfo
	12, // 8: spaceone.api.monitoring.v1.ProjectAlertConfigStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	3,  // 9: spaceone.api.monitoring.v1.ProjectAlertConfig.create:input_type -> spaceone.api.monitoring.v1.CreateProjectAlertConfigRequest
	4,  // 10: spaceone.api.monitoring.v1.ProjectAlertConfig.update:input_type -> spaceone.api.monitoring.v1.UpdateProjectAlertConfigRequest
	5,  // 11: spaceone.api.monitoring.v1.ProjectAlertConfig.delete:input_type -> spaceone.api.monitoring.v1.ProjectAlertConfigRequest
	5,  // 12: spaceone.api.monitoring.v1.ProjectAlertConfig.get:input_type -> spaceone.api.monitoring.v1.ProjectAlertConfigRequest
	6,  // 13: spaceone.api.monitoring.v1.ProjectAlertConfig.list:input_type -> spaceone.api.monitoring.v1.ProjectAlertConfigQuery
	9,  // 14: spaceone.api.monitoring.v1.ProjectAlertConfig.stat:input_type -> spaceone.api.monitoring.v1.ProjectAlertConfigStatQuery
	7,  // 15: spaceone.api.monitoring.v1.ProjectAlertConfig.create:output_type -> spaceone.api.monitoring.v1.ProjectAlertConfigInfo
	7,  // 16: spaceone.api.monitoring.v1.ProjectAlertConfig.update:output_type -> spaceone.api.monitoring.v1.ProjectAlertConfigInfo
	13, // 17: spaceone.api.monitoring.v1.ProjectAlertConfig.delete:output_type -> google.protobuf.Empty
	7,  // 18: spaceone.api.monitoring.v1.ProjectAlertConfig.get:output_type -> spaceone.api.monitoring.v1.ProjectAlertConfigInfo
	8,  // 19: spaceone.api.monitoring.v1.ProjectAlertConfig.list:output_type -> spaceone.api.monitoring.v1.ProjectAlertConfigsInfo
	14, // 20: spaceone.api.monitoring.v1.ProjectAlertConfig.stat:output_type -> google.protobuf.Struct
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_monitoring_v1_project_alert_config_proto_init() }
func file_spaceone_api_monitoring_v1_project_alert_config_proto_init() {
	if File_spaceone_api_monitoring_v1_project_alert_config_proto != nil {
		return
	}
	file_spaceone_api_monitoring_v1_escalation_policy_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDesc), len(file_spaceone_api_monitoring_v1_project_alert_config_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_monitoring_v1_project_alert_config_proto_goTypes,
		DependencyIndexes: file_spaceone_api_monitoring_v1_project_alert_config_proto_depIdxs,
		EnumInfos:         file_spaceone_api_monitoring_v1_project_alert_config_proto_enumTypes,
		MessageInfos:      file_spaceone_api_monitoring_v1_project_alert_config_proto_msgTypes,
	}.Build()
	File_spaceone_api_monitoring_v1_project_alert_config_proto = out.File
	file_spaceone_api_monitoring_v1_project_alert_config_proto_goTypes = nil
	file_spaceone_api_monitoring_v1_project_alert_config_proto_depIdxs = nil
}
