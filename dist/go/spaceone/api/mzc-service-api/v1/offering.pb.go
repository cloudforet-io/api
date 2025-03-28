// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/mzc_service_api/v1/offering.proto

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

type OfferingStatus int32

const (
	OfferingStatus_OFFERING_STATUS_NONE OfferingStatus = 0
	OfferingStatus_AVAILABLE            OfferingStatus = 1
	OfferingStatus_UNAVAILABLE          OfferingStatus = 2
	OfferingStatus_DISCONTINUED         OfferingStatus = 3
)

// Enum value maps for OfferingStatus.
var (
	OfferingStatus_name = map[int32]string{
		0: "OFFERING_STATUS_NONE",
		1: "AVAILABLE",
		2: "UNAVAILABLE",
		3: "DISCONTINUED",
	}
	OfferingStatus_value = map[string]int32{
		"OFFERING_STATUS_NONE": 0,
		"AVAILABLE":            1,
		"UNAVAILABLE":          2,
		"DISCONTINUED":         3,
	}
)

func (x OfferingStatus) Enum() *OfferingStatus {
	p := new(OfferingStatus)
	*p = x
	return p
}

func (x OfferingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfferingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_enumTypes[0].Descriptor()
}

func (OfferingStatus) Type() protoreflect.EnumType {
	return &file_spaceone_api_mzc_service_api_v1_offering_proto_enumTypes[0]
}

func (x OfferingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfferingStatus.Descriptor instead.
func (OfferingStatus) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{0}
}

type OfferingCreateRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category string                 `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// +optional
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	// +optional
	Currency string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	// +optional
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Terms string `protobuf:"bytes,6,opt,name=terms,proto3" json:"terms,omitempty"`
	// +optional
	WebsiteUrl string `protobuf:"bytes,7,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingCreateRequest) Reset() {
	*x = OfferingCreateRequest{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingCreateRequest) ProtoMessage() {}

func (x *OfferingCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingCreateRequest.ProtoReflect.Descriptor instead.
func (*OfferingCreateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{0}
}

func (x *OfferingCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfferingCreateRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *OfferingCreateRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OfferingCreateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *OfferingCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OfferingCreateRequest) GetTerms() string {
	if x != nil {
		return x.Terms
	}
	return ""
}

func (x *OfferingCreateRequest) GetWebsiteUrl() string {
	if x != nil {
		return x.WebsiteUrl
	}
	return ""
}

func (x *OfferingCreateRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type OfferingUpdateRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	OfferingId string                 `protobuf:"bytes,1,opt,name=offering_id,json=offeringId,proto3" json:"offering_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Status OfferingStatus `protobuf:"varint,3,opt,name=status,proto3,enum=spaceone.api.mzc_service_api.v1.OfferingStatus" json:"status,omitempty"`
	// +optional
	Price float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	// +optional
	Currency string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	// +optional
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// +optional
	Terms string `protobuf:"bytes,7,opt,name=terms,proto3" json:"terms,omitempty"`
	// +optional
	WebsiteUrl string `protobuf:"bytes,8,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	// +optional
	Tags          *_struct.Struct `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingUpdateRequest) Reset() {
	*x = OfferingUpdateRequest{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingUpdateRequest) ProtoMessage() {}

func (x *OfferingUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingUpdateRequest.ProtoReflect.Descriptor instead.
func (*OfferingUpdateRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{1}
}

func (x *OfferingUpdateRequest) GetOfferingId() string {
	if x != nil {
		return x.OfferingId
	}
	return ""
}

func (x *OfferingUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfferingUpdateRequest) GetStatus() OfferingStatus {
	if x != nil {
		return x.Status
	}
	return OfferingStatus_OFFERING_STATUS_NONE
}

func (x *OfferingUpdateRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OfferingUpdateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *OfferingUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OfferingUpdateRequest) GetTerms() string {
	if x != nil {
		return x.Terms
	}
	return ""
}

func (x *OfferingUpdateRequest) GetWebsiteUrl() string {
	if x != nil {
		return x.WebsiteUrl
	}
	return ""
}

func (x *OfferingUpdateRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type OfferingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OfferingId    string                 `protobuf:"bytes,1,opt,name=offering_id,json=offeringId,proto3" json:"offering_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingRequest) Reset() {
	*x = OfferingRequest{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingRequest) ProtoMessage() {}

func (x *OfferingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingRequest.ProtoReflect.Descriptor instead.
func (*OfferingRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{2}
}

func (x *OfferingRequest) GetOfferingId() string {
	if x != nil {
		return x.OfferingId
	}
	return ""
}

type OfferingSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	OfferingId string `protobuf:"bytes,2,opt,name=offering_id,json=offeringId,proto3" json:"offering_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Status        OfferingStatus `protobuf:"varint,4,opt,name=status,proto3,enum=spaceone.api.mzc_service_api.v1.OfferingStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingSearchQuery) Reset() {
	*x = OfferingSearchQuery{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingSearchQuery) ProtoMessage() {}

func (x *OfferingSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingSearchQuery.ProtoReflect.Descriptor instead.
func (*OfferingSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{3}
}

func (x *OfferingSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OfferingSearchQuery) GetOfferingId() string {
	if x != nil {
		return x.OfferingId
	}
	return ""
}

func (x *OfferingSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfferingSearchQuery) GetStatus() OfferingStatus {
	if x != nil {
		return x.Status
	}
	return OfferingStatus_OFFERING_STATUS_NONE
}

type OfferingInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OfferingId    string                 `protobuf:"bytes,1,opt,name=offering_id,json=offeringId,proto3" json:"offering_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category      string                 `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Status        OfferingStatus         `protobuf:"varint,4,opt,name=status,proto3,enum=spaceone.api.mzc_service_api.v1.OfferingStatus" json:"status,omitempty"`
	Price         float32                `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Currency      string                 `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Description   string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Terms         string                 `protobuf:"bytes,8,opt,name=terms,proto3" json:"terms,omitempty"`
	WebsiteUrl    string                 `protobuf:"bytes,9,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	Tags          *_struct.Struct        `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingInfo) Reset() {
	*x = OfferingInfo{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingInfo) ProtoMessage() {}

func (x *OfferingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingInfo.ProtoReflect.Descriptor instead.
func (*OfferingInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{4}
}

func (x *OfferingInfo) GetOfferingId() string {
	if x != nil {
		return x.OfferingId
	}
	return ""
}

func (x *OfferingInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfferingInfo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *OfferingInfo) GetStatus() OfferingStatus {
	if x != nil {
		return x.Status
	}
	return OfferingStatus_OFFERING_STATUS_NONE
}

func (x *OfferingInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OfferingInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *OfferingInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OfferingInfo) GetTerms() string {
	if x != nil {
		return x.Terms
	}
	return ""
}

func (x *OfferingInfo) GetWebsiteUrl() string {
	if x != nil {
		return x.WebsiteUrl
	}
	return ""
}

func (x *OfferingInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *OfferingInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *OfferingInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OfferingInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type OfferingsInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*OfferingInfo        `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingsInfo) Reset() {
	*x = OfferingsInfo{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingsInfo) ProtoMessage() {}

func (x *OfferingsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingsInfo.ProtoReflect.Descriptor instead.
func (*OfferingsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{5}
}

func (x *OfferingsInfo) GetResults() []*OfferingInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *OfferingsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type OfferingStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfferingStatQuery) Reset() {
	*x = OfferingStatQuery{}
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfferingStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferingStatQuery) ProtoMessage() {}

func (x *OfferingStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferingStatQuery.ProtoReflect.Descriptor instead.
func (*OfferingStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP(), []int{6}
}

func (x *OfferingStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_mzc_service_api_v1_offering_proto protoreflect.FileDescriptor

const file_spaceone_api_mzc_service_api_v1_offering_proto_rawDesc = "" +
	"\n" +
	".spaceone/api/mzc_service_api/v1/offering.proto\x12\x1fspaceone.api.mzc_service_api.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xff\x01\n" +
	"\x15OfferingCreateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x02 \x01(\tR\bcategory\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12\x14\n" +
	"\x05terms\x18\x06 \x01(\tR\x05terms\x12\x1f\n" +
	"\vwebsite_url\x18\a \x01(\tR\n" +
	"websiteUrl\x12+\n" +
	"\x04tags\x18\x0f \x01(\v2\x17.google.protobuf.StructR\x04tags\"\xcd\x02\n" +
	"\x15OfferingUpdateRequest\x12\x1f\n" +
	"\voffering_id\x18\x01 \x01(\tR\n" +
	"offeringId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12G\n" +
	"\x06status\x18\x03 \x01(\x0e2/.spaceone.api.mzc_service_api.v1.OfferingStatusR\x06status\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x02R\x05price\x12\x1a\n" +
	"\bcurrency\x18\x05 \x01(\tR\bcurrency\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12\x14\n" +
	"\x05terms\x18\a \x01(\tR\x05terms\x12\x1f\n" +
	"\vwebsite_url\x18\b \x01(\tR\n" +
	"websiteUrl\x12+\n" +
	"\x04tags\x18\x0f \x01(\v2\x17.google.protobuf.StructR\x04tags\"2\n" +
	"\x0fOfferingRequest\x12\x1f\n" +
	"\voffering_id\x18\x01 \x01(\tR\n" +
	"offeringId\"\xc6\x01\n" +
	"\x13OfferingSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x1f\n" +
	"\voffering_id\x18\x02 \x01(\tR\n" +
	"offeringId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12G\n" +
	"\x06status\x18\x04 \x01(\x0e2/.spaceone.api.mzc_service_api.v1.OfferingStatusR\x06status\"\xbb\x03\n" +
	"\fOfferingInfo\x12\x1f\n" +
	"\voffering_id\x18\x01 \x01(\tR\n" +
	"offeringId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bcategory\x18\x03 \x01(\tR\bcategory\x12G\n" +
	"\x06status\x18\x04 \x01(\x0e2/.spaceone.api.mzc_service_api.v1.OfferingStatusR\x06status\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x02R\x05price\x12\x1a\n" +
	"\bcurrency\x18\x06 \x01(\tR\bcurrency\x12 \n" +
	"\vdescription\x18\a \x01(\tR\vdescription\x12\x14\n" +
	"\x05terms\x18\b \x01(\tR\x05terms\x12\x1f\n" +
	"\vwebsite_url\x18\t \x01(\tR\n" +
	"websiteUrl\x12+\n" +
	"\x04tags\x18\x0f \x01(\v2\x17.google.protobuf.StructR\x04tags\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18  \x01(\tR\tupdatedAt\"y\n" +
	"\rOfferingsInfo\x12G\n" +
	"\aresults\x18\x01 \x03(\v2-.spaceone.api.mzc_service_api.v1.OfferingInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"P\n" +
	"\x11OfferingStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query*\\\n" +
	"\x0eOfferingStatus\x12\x18\n" +
	"\x14OFFERING_STATUS_NONE\x10\x00\x12\r\n" +
	"\tAVAILABLE\x10\x01\x12\x0f\n" +
	"\vUNAVAILABLE\x10\x02\x12\x10\n" +
	"\fDISCONTINUED\x10\x032\x8a\a\n" +
	"\bOffering\x12\x9f\x01\n" +
	"\x06create\x126.spaceone.api.mzc_service_api.v1.OfferingCreateRequest\x1a-.spaceone.api.mzc_service_api.v1.OfferingInfo\".\x82\xd3\xe4\x93\x02(:\x01*\"#/mzc-service-api/v1/offering/create\x12\x9f\x01\n" +
	"\x06update\x126.spaceone.api.mzc_service_api.v1.OfferingUpdateRequest\x1a-.spaceone.api.mzc_service_api.v1.OfferingInfo\".\x82\xd3\xe4\x93\x02(:\x01*\"#/mzc-service-api/v1/offering/update\x12\x82\x01\n" +
	"\x06delete\x120.spaceone.api.mzc_service_api.v1.OfferingRequest\x1a\x16.google.protobuf.Empty\".\x82\xd3\xe4\x93\x02(:\x01*\"#/mzc-service-api/v1/offering/delete\x12\x93\x01\n" +
	"\x03get\x120.spaceone.api.mzc_service_api.v1.OfferingRequest\x1a-.spaceone.api.mzc_service_api.v1.OfferingInfo\"+\x82\xd3\xe4\x93\x02%:\x01*\" /mzc-service-api/v1/offering/get\x12\x9a\x01\n" +
	"\x04list\x124.spaceone.api.mzc_service_api.v1.OfferingSearchQuery\x1a..spaceone.api.mzc_service_api.v1.OfferingsInfo\",\x82\xd3\xe4\x93\x02&:\x01*\"!/mzc-service-api/v1/offering/list\x12\x81\x01\n" +
	"\x04stat\x122.spaceone.api.mzc_service_api.v1.OfferingStatQuery\x1a\x17.google.protobuf.Struct\",\x82\xd3\xe4\x93\x02&:\x01*\"!/mzc-service-api/v1/offering/statBFZDgithub.com/cloudforet-io/api/dist/go/spaceone/api/mzc-service-api/v1b\x06proto3"

var (
	file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescOnce sync.Once
	file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescData []byte
)

func file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescGZIP() []byte {
	file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescOnce.Do(func() {
		file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_mzc_service_api_v1_offering_proto_rawDesc), len(file_spaceone_api_mzc_service_api_v1_offering_proto_rawDesc)))
	})
	return file_spaceone_api_mzc_service_api_v1_offering_proto_rawDescData
}

var file_spaceone_api_mzc_service_api_v1_offering_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_mzc_service_api_v1_offering_proto_goTypes = []any{
	(OfferingStatus)(0),           // 0: spaceone.api.mzc_service_api.v1.OfferingStatus
	(*OfferingCreateRequest)(nil), // 1: spaceone.api.mzc_service_api.v1.OfferingCreateRequest
	(*OfferingUpdateRequest)(nil), // 2: spaceone.api.mzc_service_api.v1.OfferingUpdateRequest
	(*OfferingRequest)(nil),       // 3: spaceone.api.mzc_service_api.v1.OfferingRequest
	(*OfferingSearchQuery)(nil),   // 4: spaceone.api.mzc_service_api.v1.OfferingSearchQuery
	(*OfferingInfo)(nil),          // 5: spaceone.api.mzc_service_api.v1.OfferingInfo
	(*OfferingsInfo)(nil),         // 6: spaceone.api.mzc_service_api.v1.OfferingsInfo
	(*OfferingStatQuery)(nil),     // 7: spaceone.api.mzc_service_api.v1.OfferingStatQuery
	(*_struct.Struct)(nil),        // 8: google.protobuf.Struct
	(*v2.Query)(nil),              // 9: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),    // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),           // 11: google.protobuf.Empty
}
var file_spaceone_api_mzc_service_api_v1_offering_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.mzc_service_api.v1.OfferingCreateRequest.tags:type_name -> google.protobuf.Struct
	0,  // 1: spaceone.api.mzc_service_api.v1.OfferingUpdateRequest.status:type_name -> spaceone.api.mzc_service_api.v1.OfferingStatus
	8,  // 2: spaceone.api.mzc_service_api.v1.OfferingUpdateRequest.tags:type_name -> google.protobuf.Struct
	9,  // 3: spaceone.api.mzc_service_api.v1.OfferingSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	0,  // 4: spaceone.api.mzc_service_api.v1.OfferingSearchQuery.status:type_name -> spaceone.api.mzc_service_api.v1.OfferingStatus
	0,  // 5: spaceone.api.mzc_service_api.v1.OfferingInfo.status:type_name -> spaceone.api.mzc_service_api.v1.OfferingStatus
	8,  // 6: spaceone.api.mzc_service_api.v1.OfferingInfo.tags:type_name -> google.protobuf.Struct
	5,  // 7: spaceone.api.mzc_service_api.v1.OfferingsInfo.results:type_name -> spaceone.api.mzc_service_api.v1.OfferingInfo
	10, // 8: spaceone.api.mzc_service_api.v1.OfferingStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	1,  // 9: spaceone.api.mzc_service_api.v1.Offering.create:input_type -> spaceone.api.mzc_service_api.v1.OfferingCreateRequest
	2,  // 10: spaceone.api.mzc_service_api.v1.Offering.update:input_type -> spaceone.api.mzc_service_api.v1.OfferingUpdateRequest
	3,  // 11: spaceone.api.mzc_service_api.v1.Offering.delete:input_type -> spaceone.api.mzc_service_api.v1.OfferingRequest
	3,  // 12: spaceone.api.mzc_service_api.v1.Offering.get:input_type -> spaceone.api.mzc_service_api.v1.OfferingRequest
	4,  // 13: spaceone.api.mzc_service_api.v1.Offering.list:input_type -> spaceone.api.mzc_service_api.v1.OfferingSearchQuery
	7,  // 14: spaceone.api.mzc_service_api.v1.Offering.stat:input_type -> spaceone.api.mzc_service_api.v1.OfferingStatQuery
	5,  // 15: spaceone.api.mzc_service_api.v1.Offering.create:output_type -> spaceone.api.mzc_service_api.v1.OfferingInfo
	5,  // 16: spaceone.api.mzc_service_api.v1.Offering.update:output_type -> spaceone.api.mzc_service_api.v1.OfferingInfo
	11, // 17: spaceone.api.mzc_service_api.v1.Offering.delete:output_type -> google.protobuf.Empty
	5,  // 18: spaceone.api.mzc_service_api.v1.Offering.get:output_type -> spaceone.api.mzc_service_api.v1.OfferingInfo
	6,  // 19: spaceone.api.mzc_service_api.v1.Offering.list:output_type -> spaceone.api.mzc_service_api.v1.OfferingsInfo
	8,  // 20: spaceone.api.mzc_service_api.v1.Offering.stat:output_type -> google.protobuf.Struct
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_mzc_service_api_v1_offering_proto_init() }
func file_spaceone_api_mzc_service_api_v1_offering_proto_init() {
	if File_spaceone_api_mzc_service_api_v1_offering_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_mzc_service_api_v1_offering_proto_rawDesc), len(file_spaceone_api_mzc_service_api_v1_offering_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_mzc_service_api_v1_offering_proto_goTypes,
		DependencyIndexes: file_spaceone_api_mzc_service_api_v1_offering_proto_depIdxs,
		EnumInfos:         file_spaceone_api_mzc_service_api_v1_offering_proto_enumTypes,
		MessageInfos:      file_spaceone_api_mzc_service_api_v1_offering_proto_msgTypes,
	}.Build()
	File_spaceone_api_mzc_service_api_v1_offering_proto = out.File
	file_spaceone_api_mzc_service_api_v1_offering_proto_goTypes = nil
	file_spaceone_api_mzc_service_api_v1_offering_proto_depIdxs = nil
}
