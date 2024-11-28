// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: spaceone/api/opsflow/v1/task_type.proto

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskTypeCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Fields []*TaskField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// +optional
	AssigneePool []string `protobuf:"bytes,4,rep,name=assignee_pool,json=assigneePool,proto3" json:"assignee_pool,omitempty"`
	// +optional
	Tags       *_struct.Struct `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	CategoryId string          `protobuf:"bytes,21,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *TaskTypeCreateRequest) Reset() {
	*x = TaskTypeCreateRequest{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeCreateRequest) ProtoMessage() {}

func (x *TaskTypeCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskTypeCreateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{0}
}

func (x *TaskTypeCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskTypeCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskTypeCreateRequest) GetFields() []*TaskField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *TaskTypeCreateRequest) GetAssigneePool() []string {
	if x != nil {
		return x.AssigneePool
	}
	return nil
}

func (x *TaskTypeCreateRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TaskTypeCreateRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type TaskTypeUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskTypeId string `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	AssigneePool []string `protobuf:"bytes,4,rep,name=assignee_pool,json=assigneePool,proto3" json:"assignee_pool,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	CategoryId string `protobuf:"bytes,21,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *TaskTypeUpdateRequest) Reset() {
	*x = TaskTypeUpdateRequest{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeUpdateRequest) ProtoMessage() {}

func (x *TaskTypeUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeUpdateRequest.ProtoReflect.Descriptor instead.
func (*TaskTypeUpdateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{1}
}

func (x *TaskTypeUpdateRequest) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

func (x *TaskTypeUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskTypeUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskTypeUpdateRequest) GetAssigneePool() []string {
	if x != nil {
		return x.AssigneePool
	}
	return nil
}

func (x *TaskTypeUpdateRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TaskTypeUpdateRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type TaskTypeUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskTypeId string       `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	Fields     []*TaskField `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Force      bool         `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *TaskTypeUpdateFieldsRequest) Reset() {
	*x = TaskTypeUpdateFieldsRequest{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeUpdateFieldsRequest) ProtoMessage() {}

func (x *TaskTypeUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*TaskTypeUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{2}
}

func (x *TaskTypeUpdateFieldsRequest) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

func (x *TaskTypeUpdateFieldsRequest) GetFields() []*TaskField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *TaskTypeUpdateFieldsRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type TaskTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskTypeId string `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
}

func (x *TaskTypeRequest) Reset() {
	*x = TaskTypeRequest{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeRequest) ProtoMessage() {}

func (x *TaskTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeRequest.ProtoReflect.Descriptor instead.
func (*TaskTypeRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{3}
}

func (x *TaskTypeRequest) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

type TaskTypeSearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	TaskTypeId string `protobuf:"bytes,2,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TaskTypeSearchQuery) Reset() {
	*x = TaskTypeSearchQuery{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeSearchQuery) ProtoMessage() {}

func (x *TaskTypeSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeSearchQuery.ProtoReflect.Descriptor instead.
func (*TaskTypeSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{4}
}

func (x *TaskTypeSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *TaskTypeSearchQuery) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

func (x *TaskTypeSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TaskTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskTypeId   string          `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	Name         string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Fields       []*TaskField    `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	AssigneePool []string        `protobuf:"bytes,5,rep,name=assignee_pool,json=assigneePool,proto3" json:"assignee_pool,omitempty"`
	Tags         *_struct.Struct `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId     string          `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CategoryId   string          `protobuf:"bytes,22,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CreatedAt    string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string          `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TaskTypeInfo) Reset() {
	*x = TaskTypeInfo{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeInfo) ProtoMessage() {}

func (x *TaskTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeInfo.ProtoReflect.Descriptor instead.
func (*TaskTypeInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{5}
}

func (x *TaskTypeInfo) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

func (x *TaskTypeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskTypeInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskTypeInfo) GetFields() []*TaskField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *TaskTypeInfo) GetAssigneePool() []string {
	if x != nil {
		return x.AssigneePool
	}
	return nil
}

func (x *TaskTypeInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *TaskTypeInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *TaskTypeInfo) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *TaskTypeInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TaskTypeInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type TaskTypesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*TaskTypeInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32           `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *TaskTypesInfo) Reset() {
	*x = TaskTypesInfo{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypesInfo) ProtoMessage() {}

func (x *TaskTypesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypesInfo.ProtoReflect.Descriptor instead.
func (*TaskTypesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{6}
}

func (x *TaskTypesInfo) GetResults() []*TaskTypeInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *TaskTypesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type TaskTypeStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *v2.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *TaskTypeStatQuery) Reset() {
	*x = TaskTypeStatQuery{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeStatQuery) ProtoMessage() {}

func (x *TaskTypeStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTypeStatQuery.ProtoReflect.Descriptor instead.
func (*TaskTypeStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{7}
}

func (x *TaskTypeStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_opsflow_v1_task_type_proto protoreflect.FileDescriptor

var file_spaceone_api_opsflow_v1_task_type_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x73, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x15, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x2b, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x15, 0x54, 0x61,
	0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x50, 0x6f, 0x6f, 0x6c,
	0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x91,
	0x01, 0x0a, 0x1b, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x22, 0x33, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x13, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf0, 0x02, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x71, 0x0a, 0x0d, 0x54, 0x61,
	0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a,
	0x11, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x32,
	0xac, 0x07, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x88, 0x01, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x6f, 0x70, 0x73, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x6f,
	0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x73, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x6f, 0x70, 0x73, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7c, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x28, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x6f, 0x70, 0x73, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x67, 0x65, 0x74, 0x12, 0x83, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x26, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a,
	0x2f, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x72, 0x0a, 0x04, 0x73, 0x74,
	0x61, 0x74, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a,
	0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_opsflow_v1_task_type_proto_rawDescOnce sync.Once
	file_spaceone_api_opsflow_v1_task_type_proto_rawDescData = file_spaceone_api_opsflow_v1_task_type_proto_rawDesc
)

func file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP() []byte {
	file_spaceone_api_opsflow_v1_task_type_proto_rawDescOnce.Do(func() {
		file_spaceone_api_opsflow_v1_task_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_opsflow_v1_task_type_proto_rawDescData)
	})
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescData
}

var file_spaceone_api_opsflow_v1_task_type_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_opsflow_v1_task_type_proto_goTypes = []any{
	(*TaskTypeCreateRequest)(nil),       // 0: spaceone.api.opsflow.v1.TaskTypeCreateRequest
	(*TaskTypeUpdateRequest)(nil),       // 1: spaceone.api.opsflow.v1.TaskTypeUpdateRequest
	(*TaskTypeUpdateFieldsRequest)(nil), // 2: spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest
	(*TaskTypeRequest)(nil),             // 3: spaceone.api.opsflow.v1.TaskTypeRequest
	(*TaskTypeSearchQuery)(nil),         // 4: spaceone.api.opsflow.v1.TaskTypeSearchQuery
	(*TaskTypeInfo)(nil),                // 5: spaceone.api.opsflow.v1.TaskTypeInfo
	(*TaskTypesInfo)(nil),               // 6: spaceone.api.opsflow.v1.TaskTypesInfo
	(*TaskTypeStatQuery)(nil),           // 7: spaceone.api.opsflow.v1.TaskTypeStatQuery
	(*TaskField)(nil),                   // 8: spaceone.api.opsflow.v1.TaskField
	(*_struct.Struct)(nil),              // 9: google.protobuf.Struct
	(*v2.Query)(nil),                    // 10: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),          // 11: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                 // 12: google.protobuf.Empty
}
var file_spaceone_api_opsflow_v1_task_type_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.opsflow.v1.TaskTypeCreateRequest.fields:type_name -> spaceone.api.opsflow.v1.TaskField
	9,  // 1: spaceone.api.opsflow.v1.TaskTypeCreateRequest.tags:type_name -> google.protobuf.Struct
	9,  // 2: spaceone.api.opsflow.v1.TaskTypeUpdateRequest.tags:type_name -> google.protobuf.Struct
	8,  // 3: spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest.fields:type_name -> spaceone.api.opsflow.v1.TaskField
	10, // 4: spaceone.api.opsflow.v1.TaskTypeSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	8,  // 5: spaceone.api.opsflow.v1.TaskTypeInfo.fields:type_name -> spaceone.api.opsflow.v1.TaskField
	9,  // 6: spaceone.api.opsflow.v1.TaskTypeInfo.tags:type_name -> google.protobuf.Struct
	5,  // 7: spaceone.api.opsflow.v1.TaskTypesInfo.results:type_name -> spaceone.api.opsflow.v1.TaskTypeInfo
	11, // 8: spaceone.api.opsflow.v1.TaskTypeStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 9: spaceone.api.opsflow.v1.TaskType.create:input_type -> spaceone.api.opsflow.v1.TaskTypeCreateRequest
	1,  // 10: spaceone.api.opsflow.v1.TaskType.update:input_type -> spaceone.api.opsflow.v1.TaskTypeUpdateRequest
	2,  // 11: spaceone.api.opsflow.v1.TaskType.update_fields:input_type -> spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest
	3,  // 12: spaceone.api.opsflow.v1.TaskType.delete:input_type -> spaceone.api.opsflow.v1.TaskTypeRequest
	3,  // 13: spaceone.api.opsflow.v1.TaskType.get:input_type -> spaceone.api.opsflow.v1.TaskTypeRequest
	4,  // 14: spaceone.api.opsflow.v1.TaskType.list:input_type -> spaceone.api.opsflow.v1.TaskTypeSearchQuery
	7,  // 15: spaceone.api.opsflow.v1.TaskType.stat:input_type -> spaceone.api.opsflow.v1.TaskTypeStatQuery
	5,  // 16: spaceone.api.opsflow.v1.TaskType.create:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	5,  // 17: spaceone.api.opsflow.v1.TaskType.update:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	5,  // 18: spaceone.api.opsflow.v1.TaskType.update_fields:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	12, // 19: spaceone.api.opsflow.v1.TaskType.delete:output_type -> google.protobuf.Empty
	5,  // 20: spaceone.api.opsflow.v1.TaskType.get:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	6,  // 21: spaceone.api.opsflow.v1.TaskType.list:output_type -> spaceone.api.opsflow.v1.TaskTypesInfo
	9,  // 22: spaceone.api.opsflow.v1.TaskType.stat:output_type -> google.protobuf.Struct
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_opsflow_v1_task_type_proto_init() }
func file_spaceone_api_opsflow_v1_task_type_proto_init() {
	if File_spaceone_api_opsflow_v1_task_type_proto != nil {
		return
	}
	file_spaceone_api_opsflow_v1_task_category_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_opsflow_v1_task_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_opsflow_v1_task_type_proto_goTypes,
		DependencyIndexes: file_spaceone_api_opsflow_v1_task_type_proto_depIdxs,
		MessageInfos:      file_spaceone_api_opsflow_v1_task_type_proto_msgTypes,
	}.Build()
	File_spaceone_api_opsflow_v1_task_type_proto = out.File
	file_spaceone_api_opsflow_v1_task_type_proto_rawDesc = nil
	file_spaceone_api_opsflow_v1_task_type_proto_goTypes = nil
	file_spaceone_api_opsflow_v1_task_type_proto_depIdxs = nil
}
