// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/plugin/data_source.proto

package plugin

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// {
//
// }
type InitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Options       *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	DomainId      string                 `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes[0]
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
	return file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *InitRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type PluginVerifyRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Options    *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData *_struct.Struct        `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// +optional
	Schema        string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	DomainId      string `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginVerifyRequest) Reset() {
	*x = PluginVerifyRequest{}
	mi := &file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginVerifyRequest) ProtoMessage() {}

func (x *PluginVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes[1]
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
	return file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescGZIP(), []int{1}
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

func (x *PluginVerifyRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *PluginVerifyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

// {
//
// }
type PluginInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *_struct.Struct        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	mi := &file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes[2]
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
	return file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescGZIP(), []int{2}
}

func (x *PluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spaceone_api_cost_analysis_plugin_data_source_proto protoreflect.FileDescriptor

const file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDesc = "" +
	"\n" +
	"3spaceone/api/cost_analysis/plugin/data_source.proto\x12!spaceone.api.cost_analysis.plugin\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"]\n" +
	"\vInitRequest\x121\n" +
	"\aoptions\x18\x01 \x01(\v2\x17.google.protobuf.StructR\aoptions\x12\x1b\n" +
	"\tdomain_id\x18\x02 \x01(\tR\bdomainId\"\xb7\x01\n" +
	"\x13PluginVerifyRequest\x121\n" +
	"\aoptions\x18\x01 \x01(\v2\x17.google.protobuf.StructR\aoptions\x128\n" +
	"\vsecret_data\x18\x02 \x01(\v2\x17.google.protobuf.StructR\n" +
	"secretData\x12\x16\n" +
	"\x06schema\x18\x03 \x01(\tR\x06schema\x12\x1b\n" +
	"\tdomain_id\x18\x04 \x01(\tR\bdomainId\"A\n" +
	"\n" +
	"PluginInfo\x123\n" +
	"\bmetadata\x18\x01 \x01(\v2\x17.google.protobuf.StructR\bmetadata2\xd1\x01\n" +
	"\n" +
	"DataSource\x12g\n" +
	"\x04init\x12..spaceone.api.cost_analysis.plugin.InitRequest\x1a-.spaceone.api.cost_analysis.plugin.PluginInfo\"\x00\x12Z\n" +
	"\x06verify\x126.spaceone.api.cost_analysis.plugin.PluginVerifyRequest\x1a\x16.google.protobuf.Empty\"\x00BHZFgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/pluginb\x06proto3"

var (
	file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescData []byte
)

func file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDesc), len(file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDesc)))
	})
	return file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDescData
}

var file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_cost_analysis_plugin_data_source_proto_goTypes = []any{
	(*InitRequest)(nil),         // 0: spaceone.api.cost_analysis.plugin.InitRequest
	(*PluginVerifyRequest)(nil), // 1: spaceone.api.cost_analysis.plugin.PluginVerifyRequest
	(*PluginInfo)(nil),          // 2: spaceone.api.cost_analysis.plugin.PluginInfo
	(*_struct.Struct)(nil),      // 3: google.protobuf.Struct
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_spaceone_api_cost_analysis_plugin_data_source_proto_depIdxs = []int32{
	3, // 0: spaceone.api.cost_analysis.plugin.InitRequest.options:type_name -> google.protobuf.Struct
	3, // 1: spaceone.api.cost_analysis.plugin.PluginVerifyRequest.options:type_name -> google.protobuf.Struct
	3, // 2: spaceone.api.cost_analysis.plugin.PluginVerifyRequest.secret_data:type_name -> google.protobuf.Struct
	3, // 3: spaceone.api.cost_analysis.plugin.PluginInfo.metadata:type_name -> google.protobuf.Struct
	0, // 4: spaceone.api.cost_analysis.plugin.DataSource.init:input_type -> spaceone.api.cost_analysis.plugin.InitRequest
	1, // 5: spaceone.api.cost_analysis.plugin.DataSource.verify:input_type -> spaceone.api.cost_analysis.plugin.PluginVerifyRequest
	2, // 6: spaceone.api.cost_analysis.plugin.DataSource.init:output_type -> spaceone.api.cost_analysis.plugin.PluginInfo
	4, // 7: spaceone.api.cost_analysis.plugin.DataSource.verify:output_type -> google.protobuf.Empty
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_plugin_data_source_proto_init() }
func file_spaceone_api_cost_analysis_plugin_data_source_proto_init() {
	if File_spaceone_api_cost_analysis_plugin_data_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDesc), len(file_spaceone_api_cost_analysis_plugin_data_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_plugin_data_source_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_plugin_data_source_proto_depIdxs,
		MessageInfos:      file_spaceone_api_cost_analysis_plugin_data_source_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_plugin_data_source_proto = out.File
	file_spaceone_api_cost_analysis_plugin_data_source_proto_goTypes = nil
	file_spaceone_api_cost_analysis_plugin_data_source_proto_depIdxs = nil
}
