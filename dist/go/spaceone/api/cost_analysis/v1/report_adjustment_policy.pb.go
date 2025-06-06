// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/report_adjustment_policy.proto

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

type ReportAdjustmentPolicyInfo_Scope int32

const (
	ReportAdjustmentPolicyInfo_SCOPE_NONE      ReportAdjustmentPolicyInfo_Scope = 0
	ReportAdjustmentPolicyInfo_WORKSPACE       ReportAdjustmentPolicyInfo_Scope = 1
	ReportAdjustmentPolicyInfo_PROJECT         ReportAdjustmentPolicyInfo_Scope = 2
	ReportAdjustmentPolicyInfo_SERVICE_ACCOUNT ReportAdjustmentPolicyInfo_Scope = 3
)

// Enum value maps for ReportAdjustmentPolicyInfo_Scope.
var (
	ReportAdjustmentPolicyInfo_Scope_name = map[int32]string{
		0: "SCOPE_NONE",
		1: "WORKSPACE",
		2: "PROJECT",
		3: "SERVICE_ACCOUNT",
	}
	ReportAdjustmentPolicyInfo_Scope_value = map[string]int32{
		"SCOPE_NONE":      0,
		"WORKSPACE":       1,
		"PROJECT":         2,
		"SERVICE_ACCOUNT": 3,
	}
)

func (x ReportAdjustmentPolicyInfo_Scope) Enum() *ReportAdjustmentPolicyInfo_Scope {
	p := new(ReportAdjustmentPolicyInfo_Scope)
	*p = x
	return p
}

func (x ReportAdjustmentPolicyInfo_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportAdjustmentPolicyInfo_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_enumTypes[0].Descriptor()
}

func (ReportAdjustmentPolicyInfo_Scope) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_enumTypes[0]
}

func (x ReportAdjustmentPolicyInfo_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportAdjustmentPolicyInfo_Scope.Descriptor instead.
func (ReportAdjustmentPolicyInfo_Scope) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{6, 0}
}

type AdjustmentPolicyFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkspaceIds  []string               `protobuf:"bytes,1,rep,name=workspace_ids,json=workspaceIds,proto3" json:"workspace_ids,omitempty"`
	ProjectIds    []string               `protobuf:"bytes,2,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdjustmentPolicyFilter) Reset() {
	*x = AdjustmentPolicyFilter{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdjustmentPolicyFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustmentPolicyFilter) ProtoMessage() {}

func (x *AdjustmentPolicyFilter) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustmentPolicyFilter.ProtoReflect.Descriptor instead.
func (*AdjustmentPolicyFilter) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{0}
}

func (x *AdjustmentPolicyFilter) GetWorkspaceIds() []string {
	if x != nil {
		return x.WorkspaceIds
	}
	return nil
}

func (x *AdjustmentPolicyFilter) GetProjectIds() []string {
	if x != nil {
		return x.ProjectIds
	}
	return nil
}

type CreateReportAdjustmentPolicyRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	CostReportConfigId string                 `protobuf:"bytes,1,opt,name=cost_report_config_id,json=costReportConfigId,proto3" json:"cost_report_config_id,omitempty"`
	// +optional
	Order int32 `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	// +optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	PolicyFilter  *AdjustmentPolicyFilter `protobuf:"bytes,5,opt,name=policy_filter,json=policyFilter,proto3" json:"policy_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReportAdjustmentPolicyRequest) Reset() {
	*x = CreateReportAdjustmentPolicyRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReportAdjustmentPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportAdjustmentPolicyRequest) ProtoMessage() {}

func (x *CreateReportAdjustmentPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReportAdjustmentPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateReportAdjustmentPolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReportAdjustmentPolicyRequest) GetCostReportConfigId() string {
	if x != nil {
		return x.CostReportConfigId
	}
	return ""
}

func (x *CreateReportAdjustmentPolicyRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *CreateReportAdjustmentPolicyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateReportAdjustmentPolicyRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateReportAdjustmentPolicyRequest) GetPolicyFilter() *AdjustmentPolicyFilter {
	if x != nil {
		return x.PolicyFilter
	}
	return nil
}

type UpdateReportAdjustmentPolicyRequest struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	ReportAdjustmentPolicyId string                 `protobuf:"bytes,1,opt,name=report_adjustment_policy_id,json=reportAdjustmentPolicyId,proto3" json:"report_adjustment_policy_id,omitempty"`
	// +optional
	Description string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tags        *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	PolicyFilter  *AdjustmentPolicyFilter `protobuf:"bytes,4,opt,name=policy_filter,json=policyFilter,proto3" json:"policy_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateReportAdjustmentPolicyRequest) Reset() {
	*x = UpdateReportAdjustmentPolicyRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReportAdjustmentPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReportAdjustmentPolicyRequest) ProtoMessage() {}

func (x *UpdateReportAdjustmentPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReportAdjustmentPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateReportAdjustmentPolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateReportAdjustmentPolicyRequest) GetReportAdjustmentPolicyId() string {
	if x != nil {
		return x.ReportAdjustmentPolicyId
	}
	return ""
}

func (x *UpdateReportAdjustmentPolicyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateReportAdjustmentPolicyRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateReportAdjustmentPolicyRequest) GetPolicyFilter() *AdjustmentPolicyFilter {
	if x != nil {
		return x.PolicyFilter
	}
	return nil
}

type ChangeOrderReportAdjustmentPolicyRequest struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	ReportAdjustmentPolicyId string                 `protobuf:"bytes,1,opt,name=report_adjustment_policy_id,json=reportAdjustmentPolicyId,proto3" json:"report_adjustment_policy_id,omitempty"`
	Order                    int32                  `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ChangeOrderReportAdjustmentPolicyRequest) Reset() {
	*x = ChangeOrderReportAdjustmentPolicyRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeOrderReportAdjustmentPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeOrderReportAdjustmentPolicyRequest) ProtoMessage() {}

func (x *ChangeOrderReportAdjustmentPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeOrderReportAdjustmentPolicyRequest.ProtoReflect.Descriptor instead.
func (*ChangeOrderReportAdjustmentPolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeOrderReportAdjustmentPolicyRequest) GetReportAdjustmentPolicyId() string {
	if x != nil {
		return x.ReportAdjustmentPolicyId
	}
	return ""
}

func (x *ChangeOrderReportAdjustmentPolicyRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type ReportAdjustmentPolicyRequest struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	ReportAdjustmentPolicyId string                 `protobuf:"bytes,1,opt,name=report_adjustment_policy_id,json=reportAdjustmentPolicyId,proto3" json:"report_adjustment_policy_id,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ReportAdjustmentPolicyRequest) Reset() {
	*x = ReportAdjustmentPolicyRequest{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportAdjustmentPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAdjustmentPolicyRequest) ProtoMessage() {}

func (x *ReportAdjustmentPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAdjustmentPolicyRequest.ProtoReflect.Descriptor instead.
func (*ReportAdjustmentPolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{4}
}

func (x *ReportAdjustmentPolicyRequest) GetReportAdjustmentPolicyId() string {
	if x != nil {
		return x.ReportAdjustmentPolicyId
	}
	return ""
}

type ReportAdjustmentPolicyQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	PolicyFilter  *AdjustmentPolicyFilter `protobuf:"bytes,2,opt,name=policy_filter,json=policyFilter,proto3" json:"policy_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportAdjustmentPolicyQuery) Reset() {
	*x = ReportAdjustmentPolicyQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportAdjustmentPolicyQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAdjustmentPolicyQuery) ProtoMessage() {}

func (x *ReportAdjustmentPolicyQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAdjustmentPolicyQuery.ProtoReflect.Descriptor instead.
func (*ReportAdjustmentPolicyQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{5}
}

func (x *ReportAdjustmentPolicyQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ReportAdjustmentPolicyQuery) GetPolicyFilter() *AdjustmentPolicyFilter {
	if x != nil {
		return x.PolicyFilter
	}
	return nil
}

type ReportAdjustmentPolicyInfo struct {
	state                    protoimpl.MessageState           `protogen:"open.v1"`
	ReportAdjustmentPolicyId string                           `protobuf:"bytes,1,opt,name=report_adjustment_policy_id,json=reportAdjustmentPolicyId,proto3" json:"report_adjustment_policy_id,omitempty"`
	Adjustments              []string                         `protobuf:"bytes,2,rep,name=adjustments,proto3" json:"adjustments,omitempty"`
	Scope                    ReportAdjustmentPolicyInfo_Scope `protobuf:"varint,3,opt,name=scope,proto3,enum=spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo_Scope" json:"scope,omitempty"`
	Order                    int32                            `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
	Description              string                           `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Tags                     *_struct.Struct                  `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	PolicyFilter             *AdjustmentPolicyFilter          `protobuf:"bytes,7,opt,name=policy_filter,json=policyFilter,proto3" json:"policy_filter,omitempty"`
	DomainId                 string                           `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CostReportConfigId       string                           `protobuf:"bytes,22,opt,name=cost_report_config_id,json=costReportConfigId,proto3" json:"cost_report_config_id,omitempty"`
	CreatedAt                string                           `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                string                           `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ReportAdjustmentPolicyInfo) Reset() {
	*x = ReportAdjustmentPolicyInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportAdjustmentPolicyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAdjustmentPolicyInfo) ProtoMessage() {}

func (x *ReportAdjustmentPolicyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAdjustmentPolicyInfo.ProtoReflect.Descriptor instead.
func (*ReportAdjustmentPolicyInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{6}
}

func (x *ReportAdjustmentPolicyInfo) GetReportAdjustmentPolicyId() string {
	if x != nil {
		return x.ReportAdjustmentPolicyId
	}
	return ""
}

func (x *ReportAdjustmentPolicyInfo) GetAdjustments() []string {
	if x != nil {
		return x.Adjustments
	}
	return nil
}

func (x *ReportAdjustmentPolicyInfo) GetScope() ReportAdjustmentPolicyInfo_Scope {
	if x != nil {
		return x.Scope
	}
	return ReportAdjustmentPolicyInfo_SCOPE_NONE
}

func (x *ReportAdjustmentPolicyInfo) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *ReportAdjustmentPolicyInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReportAdjustmentPolicyInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ReportAdjustmentPolicyInfo) GetPolicyFilter() *AdjustmentPolicyFilter {
	if x != nil {
		return x.PolicyFilter
	}
	return nil
}

func (x *ReportAdjustmentPolicyInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *ReportAdjustmentPolicyInfo) GetCostReportConfigId() string {
	if x != nil {
		return x.CostReportConfigId
	}
	return ""
}

func (x *ReportAdjustmentPolicyInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ReportAdjustmentPolicyInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ReportAdjustmentPolicesInfo struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	Results       []*ReportAdjustmentPolicyInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportAdjustmentPolicesInfo) Reset() {
	*x = ReportAdjustmentPolicesInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportAdjustmentPolicesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportAdjustmentPolicesInfo) ProtoMessage() {}

func (x *ReportAdjustmentPolicesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportAdjustmentPolicesInfo.ProtoReflect.Descriptor instead.
func (*ReportAdjustmentPolicesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP(), []int{7}
}

func (x *ReportAdjustmentPolicesInfo) GetResults() []*ReportAdjustmentPolicyInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ReportAdjustmentPolicesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto protoreflect.FileDescriptor

const file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDesc = "" +
	"\n" +
	"<spaceone/api/cost_analysis/v1/report_adjustment_policy.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a spaceone/api/core/v2/query.proto\x1a'spaceone/api/cost_analysis/v1/job.proto\"^\n" +
	"\x16AdjustmentPolicyFilter\x12#\n" +
	"\rworkspace_ids\x18\x01 \x03(\tR\fworkspaceIds\x12\x1f\n" +
	"\vproject_ids\x18\x02 \x03(\tR\n" +
	"projectIds\"\x99\x02\n" +
	"#CreateReportAdjustmentPolicyRequest\x121\n" +
	"\x15cost_report_config_id\x18\x01 \x01(\tR\x12costReportConfigId\x12\x14\n" +
	"\x05order\x18\x02 \x01(\x05R\x05order\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12+\n" +
	"\x04tags\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12Z\n" +
	"\rpolicy_filter\x18\x05 \x01(\v25.spaceone.api.cost_analysis.v1.AdjustmentPolicyFilterR\fpolicyFilter\"\x8f\x02\n" +
	"#UpdateReportAdjustmentPolicyRequest\x12=\n" +
	"\x1breport_adjustment_policy_id\x18\x01 \x01(\tR\x18reportAdjustmentPolicyId\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12+\n" +
	"\x04tags\x18\x03 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12Z\n" +
	"\rpolicy_filter\x18\x04 \x01(\v25.spaceone.api.cost_analysis.v1.AdjustmentPolicyFilterR\fpolicyFilter\"\x7f\n" +
	"(ChangeOrderReportAdjustmentPolicyRequest\x12=\n" +
	"\x1breport_adjustment_policy_id\x18\x01 \x01(\tR\x18reportAdjustmentPolicyId\x12\x14\n" +
	"\x05order\x18\x02 \x01(\x05R\x05order\"^\n" +
	"\x1dReportAdjustmentPolicyRequest\x12=\n" +
	"\x1breport_adjustment_policy_id\x18\x01 \x01(\tR\x18reportAdjustmentPolicyId\"\xac\x01\n" +
	"\x1bReportAdjustmentPolicyQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12Z\n" +
	"\rpolicy_filter\x18\x02 \x01(\v25.spaceone.api.cost_analysis.v1.AdjustmentPolicyFilterR\fpolicyFilter\"\xed\x04\n" +
	"\x1aReportAdjustmentPolicyInfo\x12=\n" +
	"\x1breport_adjustment_policy_id\x18\x01 \x01(\tR\x18reportAdjustmentPolicyId\x12 \n" +
	"\vadjustments\x18\x02 \x03(\tR\vadjustments\x12U\n" +
	"\x05scope\x18\x03 \x01(\x0e2?.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo.ScopeR\x05scope\x12\x14\n" +
	"\x05order\x18\x04 \x01(\x05R\x05order\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12+\n" +
	"\x04tags\x18\x06 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12Z\n" +
	"\rpolicy_filter\x18\a \x01(\v25.spaceone.api.cost_analysis.v1.AdjustmentPolicyFilterR\fpolicyFilter\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x121\n" +
	"\x15cost_report_config_id\x18\x16 \x01(\tR\x12costReportConfigId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18  \x01(\tR\tupdatedAt\"H\n" +
	"\x05Scope\x12\x0e\n" +
	"\n" +
	"SCOPE_NONE\x10\x00\x12\r\n" +
	"\tWORKSPACE\x10\x01\x12\v\n" +
	"\aPROJECT\x10\x02\x12\x13\n" +
	"\x0fSERVICE_ACCOUNT\x10\x03\"\x93\x01\n" +
	"\x1bReportAdjustmentPolicesInfo\x12S\n" +
	"\aresults\x18\x01 \x03(\v29.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount2\xe9\n" +
	"\n" +
	"\x16ReportAdjustmentPolicy\x12\xc5\x01\n" +
	"\x06create\x12B.spaceone.api.cost_analysis.v1.CreateReportAdjustmentPolicyRequest\x1a9.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo\"<\x82\xd3\xe4\x93\x026:\x01*\"1/cost-analysis/v1/report-adjustment-policy/create\x12\xc5\x01\n" +
	"\x06update\x12B.spaceone.api.cost_analysis.v1.UpdateReportAdjustmentPolicyRequest\x1a9.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo\"<\x82\xd3\xe4\x93\x026:\x01*\"1/cost-analysis/v1/report-adjustment-policy/update\x12\xd6\x01\n" +
	"\fchange_order\x12G.spaceone.api.cost_analysis.v1.ChangeOrderReportAdjustmentPolicyRequest\x1a9.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo\"B\x82\xd3\xe4\x93\x02<:\x01*\"7/cost-analysis/v1/report-adjustment-policy/change-order\x12\x9c\x01\n" +
	"\x06delete\x12<.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest\x1a\x16.google.protobuf.Empty\"<\x82\xd3\xe4\x93\x026:\x01*\"1/cost-analysis/v1/report-adjustment-policy/delete\x12\xcd\x01\n" +
	"\rsync_currency\x12<.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest\x1a9.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo\"C\x82\xd3\xe4\x93\x02=:\x01*\"8/cost-analysis/v1/report-adjustment-policy/sync-currency\x12\xb9\x01\n" +
	"\x03get\x12<.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest\x1a9.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo\"9\x82\xd3\xe4\x93\x023:\x01*\"./cost-analysis/v1/report-adjustment-policy/get\x12\xba\x01\n" +
	"\x04list\x12:.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyQuery\x1a:.spaceone.api.cost_analysis.v1.ReportAdjustmentPolicesInfo\":\x82\xd3\xe4\x93\x024:\x01*\"//cost-analysis/v1/report-adjustment-policy/listBDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3"

var (
	file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescData []byte
)

func file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDesc)))
	})
	return file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_goTypes = []any{
	(ReportAdjustmentPolicyInfo_Scope)(0),            // 0: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo.Scope
	(*AdjustmentPolicyFilter)(nil),                   // 1: spaceone.api.cost_analysis.v1.AdjustmentPolicyFilter
	(*CreateReportAdjustmentPolicyRequest)(nil),      // 2: spaceone.api.cost_analysis.v1.CreateReportAdjustmentPolicyRequest
	(*UpdateReportAdjustmentPolicyRequest)(nil),      // 3: spaceone.api.cost_analysis.v1.UpdateReportAdjustmentPolicyRequest
	(*ChangeOrderReportAdjustmentPolicyRequest)(nil), // 4: spaceone.api.cost_analysis.v1.ChangeOrderReportAdjustmentPolicyRequest
	(*ReportAdjustmentPolicyRequest)(nil),            // 5: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest
	(*ReportAdjustmentPolicyQuery)(nil),              // 6: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyQuery
	(*ReportAdjustmentPolicyInfo)(nil),               // 7: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	(*ReportAdjustmentPolicesInfo)(nil),              // 8: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicesInfo
	(*_struct.Struct)(nil),                           // 9: google.protobuf.Struct
	(*v2.Query)(nil),                                 // 10: spaceone.api.core.v2.Query
	(*empty.Empty)(nil),                              // 11: google.protobuf.Empty
}
var file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.cost_analysis.v1.CreateReportAdjustmentPolicyRequest.tags:type_name -> google.protobuf.Struct
	1,  // 1: spaceone.api.cost_analysis.v1.CreateReportAdjustmentPolicyRequest.policy_filter:type_name -> spaceone.api.cost_analysis.v1.AdjustmentPolicyFilter
	9,  // 2: spaceone.api.cost_analysis.v1.UpdateReportAdjustmentPolicyRequest.tags:type_name -> google.protobuf.Struct
	1,  // 3: spaceone.api.cost_analysis.v1.UpdateReportAdjustmentPolicyRequest.policy_filter:type_name -> spaceone.api.cost_analysis.v1.AdjustmentPolicyFilter
	10, // 4: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyQuery.query:type_name -> spaceone.api.core.v2.Query
	1,  // 5: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyQuery.policy_filter:type_name -> spaceone.api.cost_analysis.v1.AdjustmentPolicyFilter
	0,  // 6: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo.scope:type_name -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo.Scope
	9,  // 7: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo.tags:type_name -> google.protobuf.Struct
	1,  // 8: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo.policy_filter:type_name -> spaceone.api.cost_analysis.v1.AdjustmentPolicyFilter
	7,  // 9: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicesInfo.results:type_name -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	2,  // 10: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.create:input_type -> spaceone.api.cost_analysis.v1.CreateReportAdjustmentPolicyRequest
	3,  // 11: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.update:input_type -> spaceone.api.cost_analysis.v1.UpdateReportAdjustmentPolicyRequest
	4,  // 12: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.change_order:input_type -> spaceone.api.cost_analysis.v1.ChangeOrderReportAdjustmentPolicyRequest
	5,  // 13: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.delete:input_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest
	5,  // 14: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.sync_currency:input_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest
	5,  // 15: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.get:input_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyRequest
	6,  // 16: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.list:input_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyQuery
	7,  // 17: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.create:output_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	7,  // 18: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.update:output_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	7,  // 19: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.change_order:output_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	11, // 20: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.delete:output_type -> google.protobuf.Empty
	7,  // 21: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.sync_currency:output_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	7,  // 22: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.get:output_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicyInfo
	8,  // 23: spaceone.api.cost_analysis.v1.ReportAdjustmentPolicy.list:output_type -> spaceone.api.cost_analysis.v1.ReportAdjustmentPolicesInfo
	17, // [17:24] is the sub-list for method output_type
	10, // [10:17] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_init() }
func file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_init() {
	if File_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto != nil {
		return
	}
	file_spaceone_api_cost_analysis_v1_job_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_depIdxs,
		EnumInfos:         file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_enumTypes,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto = out.File
	file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_report_adjustment_policy_proto_depIdxs = nil
}
