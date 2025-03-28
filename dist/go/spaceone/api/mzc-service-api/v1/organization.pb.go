// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/mzc_service_api/v1/organization.proto

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

type OrganizationState int32

const (
	OrganizationState_ORGANIZATION_STATE_NONE OrganizationState = 0
	OrganizationState_ENABLED                 OrganizationState = 1
	OrganizationState_DISABLED                OrganizationState = 2
)

// Enum value maps for OrganizationState.
var (
	OrganizationState_name = map[int32]string{
		0: "ORGANIZATION_STATE_NONE",
		1: "ENABLED",
		2: "DISABLED",
	}
	OrganizationState_value = map[string]int32{
		"ORGANIZATION_STATE_NONE": 0,
		"ENABLED":                 1,
		"DISABLED":                2,
	}
)

func (x OrganizationState) Enum() *OrganizationState {
	p := new(OrganizationState)
	*p = x
	return p
}

func (x OrganizationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationState) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_enumTypes[0].Descriptor()
}

func (OrganizationState) Type() protoreflect.EnumType {
	return &file_spaceone_api_mzc_service_api_v1_organization_proto_enumTypes[0]
}

func (x OrganizationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrganizationState.Descriptor instead.
func (OrganizationState) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{0}
}

type OrganizationCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Industry string `protobuf:"bytes,3,opt,name=industry,proto3" json:"industry,omitempty"`
	// +optional
	RegistrationNumber string `protobuf:"bytes,4,opt,name=registration_number,json=registrationNumber,proto3" json:"registration_number,omitempty"`
	// +optional
	Country string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	// +optional
	Address string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	// +optional
	ContactPerson string `protobuf:"bytes,7,opt,name=contact_person,json=contactPerson,proto3" json:"contact_person,omitempty"`
	// +optional
	ContactEmail string `protobuf:"bytes,8,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	// +optional
	ContactPhone string `protobuf:"bytes,9,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationCreateRequest) Reset() {
	*x = OrganizationCreateRequest{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationCreateRequest) ProtoMessage() {}

func (x *OrganizationCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationCreateRequest.ProtoReflect.Descriptor instead.
func (*OrganizationCreateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrganizationCreateRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *OrganizationCreateRequest) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *OrganizationCreateRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *OrganizationCreateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationCreateRequest) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *OrganizationCreateRequest) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *OrganizationCreateRequest) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *OrganizationCreateRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type OrganizationUpdateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	OrgId string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Name  string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Industry string `protobuf:"bytes,4,opt,name=industry,proto3" json:"industry,omitempty"`
	// +optional
	RegistrationNumber string `protobuf:"bytes,5,opt,name=registration_number,json=registrationNumber,proto3" json:"registration_number,omitempty"`
	// +optional
	Country string `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	// +optional
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// +optional
	ContactPerson string `protobuf:"bytes,8,opt,name=contact_person,json=contactPerson,proto3" json:"contact_person,omitempty"`
	// +optional
	ContactEmail string `protobuf:"bytes,9,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	// +optional
	ContactPhone string `protobuf:"bytes,10,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationUpdateRequest) Reset() {
	*x = OrganizationUpdateRequest{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationUpdateRequest) ProtoMessage() {}

func (x *OrganizationUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrganizationUpdateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationUpdateRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *OrganizationUpdateRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type OrganizationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationRequest) Reset() {
	*x = OrganizationRequest{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationRequest) ProtoMessage() {}

func (x *OrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationRequest.ProtoReflect.Descriptor instead.
func (*OrganizationRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

type OrganizationSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	State         OrganizationState `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.mzc_service_api.v1.OrganizationState" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationSearchQuery) Reset() {
	*x = OrganizationSearchQuery{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSearchQuery) ProtoMessage() {}

func (x *OrganizationSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSearchQuery.ProtoReflect.Descriptor instead.
func (*OrganizationSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OrganizationSearchQuery) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OrganizationSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationSearchQuery) GetState() OrganizationState {
	if x != nil {
		return x.State
	}
	return OrganizationState_ORGANIZATION_STATE_NONE
}

type OrganizationInfo struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	OrgId              string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Name               string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State              OrganizationState      `protobuf:"varint,3,opt,name=state,proto3,enum=spaceone.api.mzc_service_api.v1.OrganizationState" json:"state,omitempty"`
	Description        string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Industry           string                 `protobuf:"bytes,5,opt,name=industry,proto3" json:"industry,omitempty"`
	RegistrationNumber string                 `protobuf:"bytes,6,opt,name=registration_number,json=registrationNumber,proto3" json:"registration_number,omitempty"`
	Country            string                 `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Address            string                 `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	ContactPerson      string                 `protobuf:"bytes,9,opt,name=contact_person,json=contactPerson,proto3" json:"contact_person,omitempty"`
	ContactEmail       string                 `protobuf:"bytes,10,opt,name=contact_email,json=contactEmail,proto3" json:"contact_email,omitempty"`
	ContactPhone       string                 `protobuf:"bytes,11,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty"`
	Tags               *_struct.Struct        `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId           string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt          string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *OrganizationInfo) Reset() {
	*x = OrganizationInfo{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationInfo) ProtoMessage() {}

func (x *OrganizationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationInfo.ProtoReflect.Descriptor instead.
func (*OrganizationInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{4}
}

func (x *OrganizationInfo) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *OrganizationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationInfo) GetState() OrganizationState {
	if x != nil {
		return x.State
	}
	return OrganizationState_ORGANIZATION_STATE_NONE
}

func (x *OrganizationInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OrganizationInfo) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

func (x *OrganizationInfo) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *OrganizationInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *OrganizationInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationInfo) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *OrganizationInfo) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *OrganizationInfo) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *OrganizationInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *OrganizationInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *OrganizationInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrganizationInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type OrganizationsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*OrganizationInfo    `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationsInfo) Reset() {
	*x = OrganizationsInfo{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationsInfo) ProtoMessage() {}

func (x *OrganizationsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationsInfo.ProtoReflect.Descriptor instead.
func (*OrganizationsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{5}
}

func (x *OrganizationsInfo) GetResults() []*OrganizationInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *OrganizationsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type OrganizationStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationStatQuery) Reset() {
	*x = OrganizationStatQuery{}
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationStatQuery) ProtoMessage() {}

func (x *OrganizationStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationStatQuery.ProtoReflect.Descriptor instead.
func (*OrganizationStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP(), []int{6}
}

func (x *OrganizationStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_mzc_service_api_v1_organization_proto protoreflect.FileDescriptor

const file_spaceone_api_mzc_service_api_v1_organization_proto_rawDesc = "" +
	"\n" +
	"2spaceone/api/mzc_service_api/v1/organization.proto\x12\x1fspaceone.api.mzc_service_api.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xf0\x02\n" +
	"\x19OrganizationCreateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x1a\n" +
	"\bindustry\x18\x03 \x01(\tR\bindustry\x12/\n" +
	"\x13registration_number\x18\x04 \x01(\tR\x12registrationNumber\x12\x18\n" +
	"\acountry\x18\x05 \x01(\tR\acountry\x12\x18\n" +
	"\aaddress\x18\x06 \x01(\tR\aaddress\x12%\n" +
	"\x0econtact_person\x18\a \x01(\tR\rcontactPerson\x12#\n" +
	"\rcontact_email\x18\b \x01(\tR\fcontactEmail\x12#\n" +
	"\rcontact_phone\x18\t \x01(\tR\fcontactPhone\x12+\n" +
	"\x04tags\x18\x0f \x01(\v2\x17.google.protobuf.StructR\x04tags\"\x87\x03\n" +
	"\x19OrganizationUpdateRequest\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1a\n" +
	"\bindustry\x18\x04 \x01(\tR\bindustry\x12/\n" +
	"\x13registration_number\x18\x05 \x01(\tR\x12registrationNumber\x12\x18\n" +
	"\acountry\x18\x06 \x01(\tR\acountry\x12\x18\n" +
	"\aaddress\x18\a \x01(\tR\aaddress\x12%\n" +
	"\x0econtact_person\x18\b \x01(\tR\rcontactPerson\x12#\n" +
	"\rcontact_email\x18\t \x01(\tR\fcontactEmail\x12#\n" +
	"\rcontact_phone\x18\n" +
	" \x01(\tR\fcontactPhone\x12+\n" +
	"\x04tags\x18\x0f \x01(\v2\x17.google.protobuf.StructR\x04tags\",\n" +
	"\x13OrganizationRequest\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\"\xc1\x01\n" +
	"\x17OrganizationSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x15\n" +
	"\x06org_id\x18\x02 \x01(\tR\x05orgId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12H\n" +
	"\x05state\x18\x04 \x01(\x0e22.spaceone.api.mzc_service_api.v1.OrganizationStateR\x05state\"\xa3\x04\n" +
	"\x10OrganizationInfo\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12H\n" +
	"\x05state\x18\x03 \x01(\x0e22.spaceone.api.mzc_service_api.v1.OrganizationStateR\x05state\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x1a\n" +
	"\bindustry\x18\x05 \x01(\tR\bindustry\x12/\n" +
	"\x13registration_number\x18\x06 \x01(\tR\x12registrationNumber\x12\x18\n" +
	"\acountry\x18\a \x01(\tR\acountry\x12\x18\n" +
	"\aaddress\x18\b \x01(\tR\aaddress\x12%\n" +
	"\x0econtact_person\x18\t \x01(\tR\rcontactPerson\x12#\n" +
	"\rcontact_email\x18\n" +
	" \x01(\tR\fcontactEmail\x12#\n" +
	"\rcontact_phone\x18\v \x01(\tR\fcontactPhone\x12+\n" +
	"\x04tags\x18\x0f \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18  \x01(\tR\tupdatedAt\"\x81\x01\n" +
	"\x11OrganizationsInfo\x12K\n" +
	"\aresults\x18\x01 \x03(\v21.spaceone.api.mzc_service_api.v1.OrganizationInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"T\n" +
	"\x15OrganizationStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query*K\n" +
	"\x11OrganizationState\x12\x1b\n" +
	"\x17ORGANIZATION_STATE_NONE\x10\x00\x12\v\n" +
	"\aENABLED\x10\x01\x12\f\n" +
	"\bDISABLED\x10\x022\xa0\n" +
	"\n" +
	"\fOrganization\x12\xab\x01\n" +
	"\x06create\x12:.spaceone.api.mzc_service_api.v1.OrganizationCreateRequest\x1a1.spaceone.api.mzc_service_api.v1.OrganizationInfo\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/mzc-service-api/v1/organization/create\x12\xab\x01\n" +
	"\x06update\x12:.spaceone.api.mzc_service_api.v1.OrganizationUpdateRequest\x1a1.spaceone.api.mzc_service_api.v1.OrganizationInfo\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/mzc-service-api/v1/organization/update\x12\xa5\x01\n" +
	"\x06enable\x124.spaceone.api.mzc_service_api.v1.OrganizationRequest\x1a1.spaceone.api.mzc_service_api.v1.OrganizationInfo\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/mzc-service-api/v1/organization/enable\x12\xa7\x01\n" +
	"\adisable\x124.spaceone.api.mzc_service_api.v1.OrganizationRequest\x1a1.spaceone.api.mzc_service_api.v1.OrganizationInfo\"3\x82\xd3\xe4\x93\x02-:\x01*\"(/mzc-service-api/v1/organization/disable\x12\x8a\x01\n" +
	"\x06delete\x124.spaceone.api.mzc_service_api.v1.OrganizationRequest\x1a\x16.google.protobuf.Empty\"2\x82\xd3\xe4\x93\x02,:\x01*\"'/mzc-service-api/v1/organization/delete\x12\x9f\x01\n" +
	"\x03get\x124.spaceone.api.mzc_service_api.v1.OrganizationRequest\x1a1.spaceone.api.mzc_service_api.v1.OrganizationInfo\"/\x82\xd3\xe4\x93\x02):\x01*\"$/mzc-service-api/v1/organization/get\x12\xa6\x01\n" +
	"\x04list\x128.spaceone.api.mzc_service_api.v1.OrganizationSearchQuery\x1a2.spaceone.api.mzc_service_api.v1.OrganizationsInfo\"0\x82\xd3\xe4\x93\x02*:\x01*\"%/mzc-service-api/v1/organization/list\x12\x89\x01\n" +
	"\x04stat\x126.spaceone.api.mzc_service_api.v1.OrganizationStatQuery\x1a\x17.google.protobuf.Struct\"0\x82\xd3\xe4\x93\x02*:\x01*\"%/mzc-service-api/v1/organization/statBFZDgithub.com/cloudforet-io/api/dist/go/spaceone/api/mzc-service-api/v1b\x06proto3"

var (
	file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescOnce sync.Once
	file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescData []byte
)

func file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescGZIP() []byte {
	file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescOnce.Do(func() {
		file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_mzc_service_api_v1_organization_proto_rawDesc), len(file_spaceone_api_mzc_service_api_v1_organization_proto_rawDesc)))
	})
	return file_spaceone_api_mzc_service_api_v1_organization_proto_rawDescData
}

var file_spaceone_api_mzc_service_api_v1_organization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_mzc_service_api_v1_organization_proto_goTypes = []any{
	(OrganizationState)(0),            // 0: spaceone.api.mzc_service_api.v1.OrganizationState
	(*OrganizationCreateRequest)(nil), // 1: spaceone.api.mzc_service_api.v1.OrganizationCreateRequest
	(*OrganizationUpdateRequest)(nil), // 2: spaceone.api.mzc_service_api.v1.OrganizationUpdateRequest
	(*OrganizationRequest)(nil),       // 3: spaceone.api.mzc_service_api.v1.OrganizationRequest
	(*OrganizationSearchQuery)(nil),   // 4: spaceone.api.mzc_service_api.v1.OrganizationSearchQuery
	(*OrganizationInfo)(nil),          // 5: spaceone.api.mzc_service_api.v1.OrganizationInfo
	(*OrganizationsInfo)(nil),         // 6: spaceone.api.mzc_service_api.v1.OrganizationsInfo
	(*OrganizationStatQuery)(nil),     // 7: spaceone.api.mzc_service_api.v1.OrganizationStatQuery
	(*_struct.Struct)(nil),            // 8: google.protobuf.Struct
	(*v2.Query)(nil),                  // 9: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),        // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_spaceone_api_mzc_service_api_v1_organization_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.mzc_service_api.v1.OrganizationCreateRequest.tags:type_name -> google.protobuf.Struct
	8,  // 1: spaceone.api.mzc_service_api.v1.OrganizationUpdateRequest.tags:type_name -> google.protobuf.Struct
	9,  // 2: spaceone.api.mzc_service_api.v1.OrganizationSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	0,  // 3: spaceone.api.mzc_service_api.v1.OrganizationSearchQuery.state:type_name -> spaceone.api.mzc_service_api.v1.OrganizationState
	0,  // 4: spaceone.api.mzc_service_api.v1.OrganizationInfo.state:type_name -> spaceone.api.mzc_service_api.v1.OrganizationState
	8,  // 5: spaceone.api.mzc_service_api.v1.OrganizationInfo.tags:type_name -> google.protobuf.Struct
	5,  // 6: spaceone.api.mzc_service_api.v1.OrganizationsInfo.results:type_name -> spaceone.api.mzc_service_api.v1.OrganizationInfo
	10, // 7: spaceone.api.mzc_service_api.v1.OrganizationStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	1,  // 8: spaceone.api.mzc_service_api.v1.Organization.create:input_type -> spaceone.api.mzc_service_api.v1.OrganizationCreateRequest
	2,  // 9: spaceone.api.mzc_service_api.v1.Organization.update:input_type -> spaceone.api.mzc_service_api.v1.OrganizationUpdateRequest
	3,  // 10: spaceone.api.mzc_service_api.v1.Organization.enable:input_type -> spaceone.api.mzc_service_api.v1.OrganizationRequest
	3,  // 11: spaceone.api.mzc_service_api.v1.Organization.disable:input_type -> spaceone.api.mzc_service_api.v1.OrganizationRequest
	3,  // 12: spaceone.api.mzc_service_api.v1.Organization.delete:input_type -> spaceone.api.mzc_service_api.v1.OrganizationRequest
	3,  // 13: spaceone.api.mzc_service_api.v1.Organization.get:input_type -> spaceone.api.mzc_service_api.v1.OrganizationRequest
	4,  // 14: spaceone.api.mzc_service_api.v1.Organization.list:input_type -> spaceone.api.mzc_service_api.v1.OrganizationSearchQuery
	7,  // 15: spaceone.api.mzc_service_api.v1.Organization.stat:input_type -> spaceone.api.mzc_service_api.v1.OrganizationStatQuery
	5,  // 16: spaceone.api.mzc_service_api.v1.Organization.create:output_type -> spaceone.api.mzc_service_api.v1.OrganizationInfo
	5,  // 17: spaceone.api.mzc_service_api.v1.Organization.update:output_type -> spaceone.api.mzc_service_api.v1.OrganizationInfo
	5,  // 18: spaceone.api.mzc_service_api.v1.Organization.enable:output_type -> spaceone.api.mzc_service_api.v1.OrganizationInfo
	5,  // 19: spaceone.api.mzc_service_api.v1.Organization.disable:output_type -> spaceone.api.mzc_service_api.v1.OrganizationInfo
	11, // 20: spaceone.api.mzc_service_api.v1.Organization.delete:output_type -> google.protobuf.Empty
	5,  // 21: spaceone.api.mzc_service_api.v1.Organization.get:output_type -> spaceone.api.mzc_service_api.v1.OrganizationInfo
	6,  // 22: spaceone.api.mzc_service_api.v1.Organization.list:output_type -> spaceone.api.mzc_service_api.v1.OrganizationsInfo
	8,  // 23: spaceone.api.mzc_service_api.v1.Organization.stat:output_type -> google.protobuf.Struct
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_spaceone_api_mzc_service_api_v1_organization_proto_init() }
func file_spaceone_api_mzc_service_api_v1_organization_proto_init() {
	if File_spaceone_api_mzc_service_api_v1_organization_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_mzc_service_api_v1_organization_proto_rawDesc), len(file_spaceone_api_mzc_service_api_v1_organization_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_mzc_service_api_v1_organization_proto_goTypes,
		DependencyIndexes: file_spaceone_api_mzc_service_api_v1_organization_proto_depIdxs,
		EnumInfos:         file_spaceone_api_mzc_service_api_v1_organization_proto_enumTypes,
		MessageInfos:      file_spaceone_api_mzc_service_api_v1_organization_proto_msgTypes,
	}.Build()
	File_spaceone_api_mzc_service_api_v1_organization_proto = out.File
	file_spaceone_api_mzc_service_api_v1_organization_proto_goTypes = nil
	file_spaceone_api_mzc_service_api_v1_organization_proto_depIdxs = nil
}
