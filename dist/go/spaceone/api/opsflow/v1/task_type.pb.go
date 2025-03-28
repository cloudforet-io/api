// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskTypeCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Fields []*TaskField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// +optional
	AssigneePool []string `protobuf:"bytes,5,rep,name=assignee_pool,json=assigneePool,proto3" json:"assignee_pool,omitempty"`
	// +optional
	Vars *_struct.Struct `protobuf:"bytes,6,opt,name=vars,proto3" json:"vars,omitempty"`
	// +optional
	RequireProject bool `protobuf:"varint,7,opt,name=require_project,json=requireProject,proto3" json:"require_project,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	CategoryId    string          `protobuf:"bytes,21,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *TaskTypeCreateRequest) GetVars() *_struct.Struct {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *TaskTypeCreateRequest) GetRequireProject() bool {
	if x != nil {
		return x.RequireProject
	}
	return false
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
	state      protoimpl.MessageState `protogen:"open.v1"`
	TaskTypeId string                 `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	AssigneePool []string `protobuf:"bytes,4,rep,name=assignee_pool,json=assigneePool,proto3" json:"assignee_pool,omitempty"`
	// +optional
	Vars *_struct.Struct `protobuf:"bytes,5,opt,name=vars,proto3" json:"vars,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	CategoryId    string `protobuf:"bytes,21,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *TaskTypeUpdateRequest) GetVars() *_struct.Struct {
	if x != nil {
		return x.Vars
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskTypeId    string                 `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	Fields        []*TaskField           `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Force         bool                   `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

type TaskTypeDeleteRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	TaskTypeId string                 `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	// +optional
	NewTaskTypeId string `protobuf:"bytes,2,opt,name=new_task_type_id,json=newTaskTypeId,proto3" json:"new_task_type_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskTypeDeleteRequest) Reset() {
	*x = TaskTypeDeleteRequest{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeDeleteRequest) ProtoMessage() {}

func (x *TaskTypeDeleteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskTypeDeleteRequest.ProtoReflect.Descriptor instead.
func (*TaskTypeDeleteRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{3}
}

func (x *TaskTypeDeleteRequest) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

func (x *TaskTypeDeleteRequest) GetNewTaskTypeId() string {
	if x != nil {
		return x.NewTaskTypeId
	}
	return ""
}

type TaskTypeRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	TaskTypeId string                 `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	// +optional
	IncludeCategoryFields bool `protobuf:"varint,2,opt,name=include_category_fields,json=includeCategoryFields,proto3" json:"include_category_fields,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *TaskTypeRequest) Reset() {
	*x = TaskTypeRequest{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeRequest) ProtoMessage() {}

func (x *TaskTypeRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskTypeRequest.ProtoReflect.Descriptor instead.
func (*TaskTypeRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{4}
}

func (x *TaskTypeRequest) GetTaskTypeId() string {
	if x != nil {
		return x.TaskTypeId
	}
	return ""
}

func (x *TaskTypeRequest) GetIncludeCategoryFields() bool {
	if x != nil {
		return x.IncludeCategoryFields
	}
	return false
}

type TaskTypeSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	TaskTypeId string `protobuf:"bytes,2,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	// +optional
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskTypeSearchQuery) Reset() {
	*x = TaskTypeSearchQuery{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeSearchQuery) ProtoMessage() {}

func (x *TaskTypeSearchQuery) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskTypeSearchQuery.ProtoReflect.Descriptor instead.
func (*TaskTypeSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{5}
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
	state          protoimpl.MessageState `protogen:"open.v1"`
	TaskTypeId     string                 `protobuf:"bytes,1,opt,name=task_type_id,json=taskTypeId,proto3" json:"task_type_id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Fields         []*TaskField           `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	AssigneePool   []string               `protobuf:"bytes,6,rep,name=assignee_pool,json=assigneePool,proto3" json:"assignee_pool,omitempty"`
	Vars           *_struct.Struct        `protobuf:"bytes,7,opt,name=vars,proto3" json:"vars,omitempty"`
	RequireProject bool                   `protobuf:"varint,8,opt,name=require_project,json=requireProject,proto3" json:"require_project,omitempty"`
	Tags           *_struct.Struct        `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId       string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CategoryId     string                 `protobuf:"bytes,22,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CreatedAt      string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TaskTypeInfo) Reset() {
	*x = TaskTypeInfo{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeInfo) ProtoMessage() {}

func (x *TaskTypeInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskTypeInfo.ProtoReflect.Descriptor instead.
func (*TaskTypeInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{6}
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

func (x *TaskTypeInfo) GetVars() *_struct.Struct {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *TaskTypeInfo) GetRequireProject() bool {
	if x != nil {
		return x.RequireProject
	}
	return false
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*TaskTypeInfo        `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskTypesInfo) Reset() {
	*x = TaskTypesInfo{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypesInfo) ProtoMessage() {}

func (x *TaskTypesInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskTypesInfo.ProtoReflect.Descriptor instead.
func (*TaskTypesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{7}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskTypeStatQuery) Reset() {
	*x = TaskTypeStatQuery{}
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskTypeStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTypeStatQuery) ProtoMessage() {}

func (x *TaskTypeStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_opsflow_v1_task_type_proto_msgTypes[8]
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
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP(), []int{8}
}

func (x *TaskTypeStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_opsflow_v1_task_type_proto protoreflect.FileDescriptor

const file_spaceone_api_opsflow_v1_task_type_proto_rawDesc = "" +
	"\n" +
	"'spaceone/api/opsflow/v1/task_type.proto\x12\x17spaceone.api.opsflow.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a+spaceone/api/opsflow/v1/task_category.proto\"\xd2\x02\n" +
	"\x15TaskTypeCreateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12:\n" +
	"\x06fields\x18\x03 \x03(\v2\".spaceone.api.opsflow.v1.TaskFieldR\x06fields\x12#\n" +
	"\rassignee_pool\x18\x05 \x03(\tR\fassigneePool\x12+\n" +
	"\x04vars\x18\x06 \x01(\v2\x17.google.protobuf.StructR\x04vars\x12'\n" +
	"\x0frequire_project\x18\a \x01(\bR\x0erequireProject\x12+\n" +
	"\x04tags\x18\v \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1f\n" +
	"\vcategory_id\x18\x15 \x01(\tR\n" +
	"categoryId\"\x8f\x02\n" +
	"\x15TaskTypeUpdateRequest\x12 \n" +
	"\ftask_type_id\x18\x01 \x01(\tR\n" +
	"taskTypeId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12#\n" +
	"\rassignee_pool\x18\x04 \x03(\tR\fassigneePool\x12+\n" +
	"\x04vars\x18\x05 \x01(\v2\x17.google.protobuf.StructR\x04vars\x12+\n" +
	"\x04tags\x18\v \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1f\n" +
	"\vcategory_id\x18\x15 \x01(\tR\n" +
	"categoryId\"\x91\x01\n" +
	"\x1bTaskTypeUpdateFieldsRequest\x12 \n" +
	"\ftask_type_id\x18\x01 \x01(\tR\n" +
	"taskTypeId\x12:\n" +
	"\x06fields\x18\x02 \x03(\v2\".spaceone.api.opsflow.v1.TaskFieldR\x06fields\x12\x14\n" +
	"\x05force\x18\x03 \x01(\bR\x05force\"b\n" +
	"\x15TaskTypeDeleteRequest\x12 \n" +
	"\ftask_type_id\x18\x01 \x01(\tR\n" +
	"taskTypeId\x12'\n" +
	"\x10new_task_type_id\x18\x02 \x01(\tR\rnewTaskTypeId\"k\n" +
	"\x0fTaskTypeRequest\x12 \n" +
	"\ftask_type_id\x18\x01 \x01(\tR\n" +
	"taskTypeId\x126\n" +
	"\x17include_category_fields\x18\x02 \x01(\bR\x15includeCategoryFields\"~\n" +
	"\x13TaskTypeSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12 \n" +
	"\ftask_type_id\x18\x02 \x01(\tR\n" +
	"taskTypeId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\"\xc6\x03\n" +
	"\fTaskTypeInfo\x12 \n" +
	"\ftask_type_id\x18\x01 \x01(\tR\n" +
	"taskTypeId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12:\n" +
	"\x06fields\x18\x04 \x03(\v2\".spaceone.api.opsflow.v1.TaskFieldR\x06fields\x12#\n" +
	"\rassignee_pool\x18\x06 \x03(\tR\fassigneePool\x12+\n" +
	"\x04vars\x18\a \x01(\v2\x17.google.protobuf.StructR\x04vars\x12'\n" +
	"\x0frequire_project\x18\b \x01(\bR\x0erequireProject\x12+\n" +
	"\x04tags\x18\v \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12\x1f\n" +
	"\vcategory_id\x18\x16 \x01(\tR\n" +
	"categoryId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18  \x01(\tR\tupdatedAt\"q\n" +
	"\rTaskTypesInfo\x12?\n" +
	"\aresults\x18\x01 \x03(\v2%.spaceone.api.opsflow.v1.TaskTypeInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"P\n" +
	"\x11TaskTypeStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\xb2\a\n" +
	"\bTaskType\x12\x88\x01\n" +
	"\x06create\x12..spaceone.api.opsflow.v1.TaskTypeCreateRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/opsflow/v1/task-type/create\x12\x88\x01\n" +
	"\x06update\x12..spaceone.api.opsflow.v1.TaskTypeUpdateRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/opsflow/v1/task-type/update\x12\x9c\x01\n" +
	"\rupdate_fields\x124.spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\".\x82\xd3\xe4\x93\x02(:\x01*\"#/opsflow/v1/task-type/update_fields\x12y\n" +
	"\x06delete\x12..spaceone.api.opsflow.v1.TaskTypeDeleteRequest\x1a\x16.google.protobuf.Empty\"'\x82\xd3\xe4\x93\x02!:\x01*\"\x1c/opsflow/v1/task-type/delete\x12|\n" +
	"\x03get\x12(.spaceone.api.opsflow.v1.TaskTypeRequest\x1a%.spaceone.api.opsflow.v1.TaskTypeInfo\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/opsflow/v1/task-type/get\x12\x83\x01\n" +
	"\x04list\x12,.spaceone.api.opsflow.v1.TaskTypeSearchQuery\x1a&.spaceone.api.opsflow.v1.TaskTypesInfo\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/opsflow/v1/task-type/list\x12r\n" +
	"\x04stat\x12*.spaceone.api.opsflow.v1.TaskTypeStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/opsflow/v1/task-type/statB>Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1b\x06proto3"

var (
	file_spaceone_api_opsflow_v1_task_type_proto_rawDescOnce sync.Once
	file_spaceone_api_opsflow_v1_task_type_proto_rawDescData []byte
)

func file_spaceone_api_opsflow_v1_task_type_proto_rawDescGZIP() []byte {
	file_spaceone_api_opsflow_v1_task_type_proto_rawDescOnce.Do(func() {
		file_spaceone_api_opsflow_v1_task_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_opsflow_v1_task_type_proto_rawDesc), len(file_spaceone_api_opsflow_v1_task_type_proto_rawDesc)))
	})
	return file_spaceone_api_opsflow_v1_task_type_proto_rawDescData
}

var file_spaceone_api_opsflow_v1_task_type_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spaceone_api_opsflow_v1_task_type_proto_goTypes = []any{
	(*TaskTypeCreateRequest)(nil),       // 0: spaceone.api.opsflow.v1.TaskTypeCreateRequest
	(*TaskTypeUpdateRequest)(nil),       // 1: spaceone.api.opsflow.v1.TaskTypeUpdateRequest
	(*TaskTypeUpdateFieldsRequest)(nil), // 2: spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest
	(*TaskTypeDeleteRequest)(nil),       // 3: spaceone.api.opsflow.v1.TaskTypeDeleteRequest
	(*TaskTypeRequest)(nil),             // 4: spaceone.api.opsflow.v1.TaskTypeRequest
	(*TaskTypeSearchQuery)(nil),         // 5: spaceone.api.opsflow.v1.TaskTypeSearchQuery
	(*TaskTypeInfo)(nil),                // 6: spaceone.api.opsflow.v1.TaskTypeInfo
	(*TaskTypesInfo)(nil),               // 7: spaceone.api.opsflow.v1.TaskTypesInfo
	(*TaskTypeStatQuery)(nil),           // 8: spaceone.api.opsflow.v1.TaskTypeStatQuery
	(*TaskField)(nil),                   // 9: spaceone.api.opsflow.v1.TaskField
	(*_struct.Struct)(nil),              // 10: google.protobuf.Struct
	(*v2.Query)(nil),                    // 11: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),          // 12: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),                 // 13: google.protobuf.Empty
}
var file_spaceone_api_opsflow_v1_task_type_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.opsflow.v1.TaskTypeCreateRequest.fields:type_name -> spaceone.api.opsflow.v1.TaskField
	10, // 1: spaceone.api.opsflow.v1.TaskTypeCreateRequest.vars:type_name -> google.protobuf.Struct
	10, // 2: spaceone.api.opsflow.v1.TaskTypeCreateRequest.tags:type_name -> google.protobuf.Struct
	10, // 3: spaceone.api.opsflow.v1.TaskTypeUpdateRequest.vars:type_name -> google.protobuf.Struct
	10, // 4: spaceone.api.opsflow.v1.TaskTypeUpdateRequest.tags:type_name -> google.protobuf.Struct
	9,  // 5: spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest.fields:type_name -> spaceone.api.opsflow.v1.TaskField
	11, // 6: spaceone.api.opsflow.v1.TaskTypeSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	9,  // 7: spaceone.api.opsflow.v1.TaskTypeInfo.fields:type_name -> spaceone.api.opsflow.v1.TaskField
	10, // 8: spaceone.api.opsflow.v1.TaskTypeInfo.vars:type_name -> google.protobuf.Struct
	10, // 9: spaceone.api.opsflow.v1.TaskTypeInfo.tags:type_name -> google.protobuf.Struct
	6,  // 10: spaceone.api.opsflow.v1.TaskTypesInfo.results:type_name -> spaceone.api.opsflow.v1.TaskTypeInfo
	12, // 11: spaceone.api.opsflow.v1.TaskTypeStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	0,  // 12: spaceone.api.opsflow.v1.TaskType.create:input_type -> spaceone.api.opsflow.v1.TaskTypeCreateRequest
	1,  // 13: spaceone.api.opsflow.v1.TaskType.update:input_type -> spaceone.api.opsflow.v1.TaskTypeUpdateRequest
	2,  // 14: spaceone.api.opsflow.v1.TaskType.update_fields:input_type -> spaceone.api.opsflow.v1.TaskTypeUpdateFieldsRequest
	3,  // 15: spaceone.api.opsflow.v1.TaskType.delete:input_type -> spaceone.api.opsflow.v1.TaskTypeDeleteRequest
	4,  // 16: spaceone.api.opsflow.v1.TaskType.get:input_type -> spaceone.api.opsflow.v1.TaskTypeRequest
	5,  // 17: spaceone.api.opsflow.v1.TaskType.list:input_type -> spaceone.api.opsflow.v1.TaskTypeSearchQuery
	8,  // 18: spaceone.api.opsflow.v1.TaskType.stat:input_type -> spaceone.api.opsflow.v1.TaskTypeStatQuery
	6,  // 19: spaceone.api.opsflow.v1.TaskType.create:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	6,  // 20: spaceone.api.opsflow.v1.TaskType.update:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	6,  // 21: spaceone.api.opsflow.v1.TaskType.update_fields:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	13, // 22: spaceone.api.opsflow.v1.TaskType.delete:output_type -> google.protobuf.Empty
	6,  // 23: spaceone.api.opsflow.v1.TaskType.get:output_type -> spaceone.api.opsflow.v1.TaskTypeInfo
	7,  // 24: spaceone.api.opsflow.v1.TaskType.list:output_type -> spaceone.api.opsflow.v1.TaskTypesInfo
	10, // 25: spaceone.api.opsflow.v1.TaskType.stat:output_type -> google.protobuf.Struct
	19, // [19:26] is the sub-list for method output_type
	12, // [12:19] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_opsflow_v1_task_type_proto_rawDesc), len(file_spaceone_api_opsflow_v1_task_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_opsflow_v1_task_type_proto_goTypes,
		DependencyIndexes: file_spaceone_api_opsflow_v1_task_type_proto_depIdxs,
		MessageInfos:      file_spaceone_api_opsflow_v1_task_type_proto_msgTypes,
	}.Build()
	File_spaceone_api_opsflow_v1_task_type_proto = out.File
	file_spaceone_api_opsflow_v1_task_type_proto_goTypes = nil
	file_spaceone_api_opsflow_v1_task_type_proto_depIdxs = nil
}
