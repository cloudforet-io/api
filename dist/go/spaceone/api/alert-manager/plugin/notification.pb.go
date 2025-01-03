// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: spaceone/api/alert_manager/plugin/notification.proto

package plugin

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type DispatchRequest_NotificationType int32

const (
	DispatchRequest_NOTIFICATION_TYPE_NONE DispatchRequest_NotificationType = 0
	DispatchRequest_ERROR                  DispatchRequest_NotificationType = 1
	DispatchRequest_WARNING                DispatchRequest_NotificationType = 2
	DispatchRequest_SUCCESS                DispatchRequest_NotificationType = 3
	DispatchRequest_INFO                   DispatchRequest_NotificationType = 4
)

// Enum value maps for DispatchRequest_NotificationType.
var (
	DispatchRequest_NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_NONE",
		1: "ERROR",
		2: "WARNING",
		3: "SUCCESS",
		4: "INFO",
	}
	DispatchRequest_NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_NONE": 0,
		"ERROR":                  1,
		"WARNING":                2,
		"SUCCESS":                3,
		"INFO":                   4,
	}
)

func (x DispatchRequest_NotificationType) Enum() *DispatchRequest_NotificationType {
	p := new(DispatchRequest_NotificationType)
	*p = x
	return p
}

func (x DispatchRequest_NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DispatchRequest_NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_alert_manager_plugin_notification_proto_enumTypes[0].Descriptor()
}

func (DispatchRequest_NotificationType) Type() protoreflect.EnumType {
	return &file_spaceone_api_alert_manager_plugin_notification_proto_enumTypes[0]
}

func (x DispatchRequest_NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DispatchRequest_NotificationType.Descriptor instead.
func (DispatchRequest_NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_notification_proto_rawDescGZIP(), []int{3, 0}
}

type NotificationTag struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Options       *_struct.Struct        `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationTag) Reset() {
	*x = NotificationTag{}
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTag) ProtoMessage() {}

func (x *NotificationTag) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTag.ProtoReflect.Descriptor instead.
func (*NotificationTag) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationTag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *NotificationTag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *NotificationTag) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type NotificationCallback struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Options       *_struct.Struct        `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationCallback) Reset() {
	*x = NotificationCallback{}
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationCallback) ProtoMessage() {}

func (x *NotificationCallback) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationCallback.ProtoReflect.Descriptor instead.
func (*NotificationCallback) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationCallback) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *NotificationCallback) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NotificationCallback) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type NotificationMessage struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Title         string                  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Link          string                  `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Description   string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl      string                  `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Tags          []*NotificationTag      `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	Callbacks     []*NotificationCallback `protobuf:"bytes,12,rep,name=callbacks,proto3" json:"callbacks,omitempty"`
	OccurredAt    string                  `protobuf:"bytes,31,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationMessage) Reset() {
	*x = NotificationMessage{}
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationMessage) ProtoMessage() {}

func (x *NotificationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationMessage.ProtoReflect.Descriptor instead.
func (*NotificationMessage) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotificationMessage) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *NotificationMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotificationMessage) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *NotificationMessage) GetTags() []*NotificationTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *NotificationMessage) GetCallbacks() []*NotificationCallback {
	if x != nil {
		return x.Callbacks
	}
	return nil
}

func (x *NotificationMessage) GetOccurredAt() string {
	if x != nil {
		return x.OccurredAt
	}
	return ""
}

type DispatchRequest struct {
	state            protoimpl.MessageState           `protogen:"open.v1"`
	Options          *_struct.Struct                  `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData       *_struct.Struct                  `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	ChannelData      *_struct.Struct                  `protobuf:"bytes,3,opt,name=channel_data,json=channelData,proto3" json:"channel_data,omitempty"`
	Message          *NotificationMessage             `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	NotificationType DispatchRequest_NotificationType `protobuf:"varint,5,opt,name=notification_type,json=notificationType,proto3,enum=spaceone.api.alert_manager.plugin.DispatchRequest_NotificationType" json:"notification_type,omitempty"`
	DomainId         string                           `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_plugin_notification_proto_rawDescGZIP(), []int{3}
}

func (x *DispatchRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *DispatchRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *DispatchRequest) GetChannelData() *_struct.Struct {
	if x != nil {
		return x.ChannelData
	}
	return nil
}

func (x *DispatchRequest) GetMessage() *NotificationMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *DispatchRequest) GetNotificationType() DispatchRequest_NotificationType {
	if x != nil {
		return x.NotificationType
	}
	return DispatchRequest_NOTIFICATION_TYPE_NONE
}

func (x *DispatchRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_alert_manager_plugin_notification_proto protoreflect.FileDescriptor

var file_spaceone_api_alert_manager_plugin_notification_proto_rawDesc = []byte{
	0x0a, 0x34, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x71, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbe, 0x02, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x55,
	0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0xfa, 0x03, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a,
	0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x50, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x43, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46,
	0x4f, 0x10, 0x04, 0x32, 0x68, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x32, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x48, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_alert_manager_plugin_notification_proto_rawDescOnce sync.Once
	file_spaceone_api_alert_manager_plugin_notification_proto_rawDescData = file_spaceone_api_alert_manager_plugin_notification_proto_rawDesc
)

func file_spaceone_api_alert_manager_plugin_notification_proto_rawDescGZIP() []byte {
	file_spaceone_api_alert_manager_plugin_notification_proto_rawDescOnce.Do(func() {
		file_spaceone_api_alert_manager_plugin_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_alert_manager_plugin_notification_proto_rawDescData)
	})
	return file_spaceone_api_alert_manager_plugin_notification_proto_rawDescData
}

var file_spaceone_api_alert_manager_plugin_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spaceone_api_alert_manager_plugin_notification_proto_goTypes = []any{
	(DispatchRequest_NotificationType)(0), // 0: spaceone.api.alert_manager.plugin.DispatchRequest.NotificationType
	(*NotificationTag)(nil),               // 1: spaceone.api.alert_manager.plugin.NotificationTag
	(*NotificationCallback)(nil),          // 2: spaceone.api.alert_manager.plugin.NotificationCallback
	(*NotificationMessage)(nil),           // 3: spaceone.api.alert_manager.plugin.NotificationMessage
	(*DispatchRequest)(nil),               // 4: spaceone.api.alert_manager.plugin.DispatchRequest
	(*_struct.Struct)(nil),                // 5: google.protobuf.Struct
	(*empty.Empty)(nil),                   // 6: google.protobuf.Empty
}
var file_spaceone_api_alert_manager_plugin_notification_proto_depIdxs = []int32{
	5,  // 0: spaceone.api.alert_manager.plugin.NotificationTag.options:type_name -> google.protobuf.Struct
	5,  // 1: spaceone.api.alert_manager.plugin.NotificationCallback.options:type_name -> google.protobuf.Struct
	1,  // 2: spaceone.api.alert_manager.plugin.NotificationMessage.tags:type_name -> spaceone.api.alert_manager.plugin.NotificationTag
	2,  // 3: spaceone.api.alert_manager.plugin.NotificationMessage.callbacks:type_name -> spaceone.api.alert_manager.plugin.NotificationCallback
	5,  // 4: spaceone.api.alert_manager.plugin.DispatchRequest.options:type_name -> google.protobuf.Struct
	5,  // 5: spaceone.api.alert_manager.plugin.DispatchRequest.secret_data:type_name -> google.protobuf.Struct
	5,  // 6: spaceone.api.alert_manager.plugin.DispatchRequest.channel_data:type_name -> google.protobuf.Struct
	3,  // 7: spaceone.api.alert_manager.plugin.DispatchRequest.message:type_name -> spaceone.api.alert_manager.plugin.NotificationMessage
	0,  // 8: spaceone.api.alert_manager.plugin.DispatchRequest.notification_type:type_name -> spaceone.api.alert_manager.plugin.DispatchRequest.NotificationType
	4,  // 9: spaceone.api.alert_manager.plugin.Notification.dispatch:input_type -> spaceone.api.alert_manager.plugin.DispatchRequest
	6,  // 10: spaceone.api.alert_manager.plugin.Notification.dispatch:output_type -> google.protobuf.Empty
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_alert_manager_plugin_notification_proto_init() }
func file_spaceone_api_alert_manager_plugin_notification_proto_init() {
	if File_spaceone_api_alert_manager_plugin_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_alert_manager_plugin_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_alert_manager_plugin_notification_proto_goTypes,
		DependencyIndexes: file_spaceone_api_alert_manager_plugin_notification_proto_depIdxs,
		EnumInfos:         file_spaceone_api_alert_manager_plugin_notification_proto_enumTypes,
		MessageInfos:      file_spaceone_api_alert_manager_plugin_notification_proto_msgTypes,
	}.Build()
	File_spaceone_api_alert_manager_plugin_notification_proto = out.File
	file_spaceone_api_alert_manager_plugin_notification_proto_rawDesc = nil
	file_spaceone_api_alert_manager_plugin_notification_proto_goTypes = nil
	file_spaceone_api_alert_manager_plugin_notification_proto_depIdxs = nil
}
