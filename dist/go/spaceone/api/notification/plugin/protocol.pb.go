// A Protocol is a plugin instance defining how a User receives data from Cloudforet.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: spaceone/api/notification/plugin/protocol.proto

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

//	{
//	   "options": {}
//	}
type InitRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Option value used when initializing the plugin.
	Options       *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_spaceone_api_notification_plugin_protocol_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_notification_plugin_protocol_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_notification_plugin_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

//	{
//	   "options": {}
//	}
type PluginVerifyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Option values required for the plugin to work.
	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	// The secret value required for the plugin to work.
	// The secret data usually includes the credential information required for the plugin to access the external system.
	SecretData    *_struct.Struct `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginVerifyRequest) Reset() {
	*x = PluginVerifyRequest{}
	mi := &file_spaceone_api_notification_plugin_protocol_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginVerifyRequest) ProtoMessage() {}

func (x *PluginVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_notification_plugin_protocol_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginVerifyRequest.ProtoReflect.Descriptor instead.
func (*PluginVerifyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_notification_plugin_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *PluginVerifyRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *PluginVerifyRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

// {
//
//	   "metadata": {}
//	}
type PluginInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata value required to input various values required for plugin to work.
	// In the case of protocol plugins, when creating a channel, the plugin contains the definition of additional data (channel data) required for channel transmission.
	Metadata      *_struct.Struct `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	mi := &file_spaceone_api_notification_plugin_protocol_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_notification_plugin_protocol_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_notification_plugin_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *PluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spaceone_api_notification_plugin_protocol_proto protoreflect.FileDescriptor

var file_spaceone_api_notification_plugin_protocol_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40,
	0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x82, 0x01, 0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcc, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x65, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x2d, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x06,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74,
	0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_notification_plugin_protocol_proto_rawDescOnce sync.Once
	file_spaceone_api_notification_plugin_protocol_proto_rawDescData = file_spaceone_api_notification_plugin_protocol_proto_rawDesc
)

func file_spaceone_api_notification_plugin_protocol_proto_rawDescGZIP() []byte {
	file_spaceone_api_notification_plugin_protocol_proto_rawDescOnce.Do(func() {
		file_spaceone_api_notification_plugin_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_notification_plugin_protocol_proto_rawDescData)
	})
	return file_spaceone_api_notification_plugin_protocol_proto_rawDescData
}

var file_spaceone_api_notification_plugin_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_notification_plugin_protocol_proto_goTypes = []any{
	(*InitRequest)(nil),         // 0: spaceone.api.notification.plugin.InitRequest
	(*PluginVerifyRequest)(nil), // 1: spaceone.api.notification.plugin.PluginVerifyRequest
	(*PluginInfo)(nil),          // 2: spaceone.api.notification.plugin.PluginInfo
	(*_struct.Struct)(nil),      // 3: google.protobuf.Struct
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_spaceone_api_notification_plugin_protocol_proto_depIdxs = []int32{
	3, // 0: spaceone.api.notification.plugin.InitRequest.options:type_name -> google.protobuf.Struct
	3, // 1: spaceone.api.notification.plugin.PluginVerifyRequest.options:type_name -> google.protobuf.Struct
	3, // 2: spaceone.api.notification.plugin.PluginVerifyRequest.secret_data:type_name -> google.protobuf.Struct
	3, // 3: spaceone.api.notification.plugin.PluginInfo.metadata:type_name -> google.protobuf.Struct
	0, // 4: spaceone.api.notification.plugin.Protocol.init:input_type -> spaceone.api.notification.plugin.InitRequest
	1, // 5: spaceone.api.notification.plugin.Protocol.verify:input_type -> spaceone.api.notification.plugin.PluginVerifyRequest
	2, // 6: spaceone.api.notification.plugin.Protocol.init:output_type -> spaceone.api.notification.plugin.PluginInfo
	4, // 7: spaceone.api.notification.plugin.Protocol.verify:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spaceone_api_notification_plugin_protocol_proto_init() }
func file_spaceone_api_notification_plugin_protocol_proto_init() {
	if File_spaceone_api_notification_plugin_protocol_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_notification_plugin_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_notification_plugin_protocol_proto_goTypes,
		DependencyIndexes: file_spaceone_api_notification_plugin_protocol_proto_depIdxs,
		MessageInfos:      file_spaceone_api_notification_plugin_protocol_proto_msgTypes,
	}.Build()
	File_spaceone_api_notification_plugin_protocol_proto = out.File
	file_spaceone_api_notification_plugin_protocol_proto_rawDesc = nil
	file_spaceone_api_notification_plugin_protocol_proto_goTypes = nil
	file_spaceone_api_notification_plugin_protocol_proto_depIdxs = nil
}
