// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/core/v1/server_info.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VersionInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	mi := &file_spaceone_api_core_v1_server_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v1_server_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_server_info_proto_rawDescGZIP(), []int{0}
}

func (x *VersionInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_spaceone_api_core_v1_server_info_proto protoreflect.FileDescriptor

const file_spaceone_api_core_v1_server_info_proto_rawDesc = "" +
	"\n" +
	"&spaceone/api/core/v1/server_info.proto\x12\x14spaceone.api.core.v1\x1a\x1bgoogle/protobuf/empty.proto\"'\n" +
	"\vVersionInfo\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion2X\n" +
	"\n" +
	"ServerInfo\x12J\n" +
	"\vget_version\x12\x16.google.protobuf.Empty\x1a!.spaceone.api.core.v1.VersionInfo\"\x00B;Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1b\x06proto3"

var (
	file_spaceone_api_core_v1_server_info_proto_rawDescOnce sync.Once
	file_spaceone_api_core_v1_server_info_proto_rawDescData []byte
)

func file_spaceone_api_core_v1_server_info_proto_rawDescGZIP() []byte {
	file_spaceone_api_core_v1_server_info_proto_rawDescOnce.Do(func() {
		file_spaceone_api_core_v1_server_info_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_core_v1_server_info_proto_rawDesc), len(file_spaceone_api_core_v1_server_info_proto_rawDesc)))
	})
	return file_spaceone_api_core_v1_server_info_proto_rawDescData
}

var file_spaceone_api_core_v1_server_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spaceone_api_core_v1_server_info_proto_goTypes = []any{
	(*VersionInfo)(nil), // 0: spaceone.api.core.v1.VersionInfo
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_spaceone_api_core_v1_server_info_proto_depIdxs = []int32{
	1, // 0: spaceone.api.core.v1.ServerInfo.get_version:input_type -> google.protobuf.Empty
	0, // 1: spaceone.api.core.v1.ServerInfo.get_version:output_type -> spaceone.api.core.v1.VersionInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spaceone_api_core_v1_server_info_proto_init() }
func file_spaceone_api_core_v1_server_info_proto_init() {
	if File_spaceone_api_core_v1_server_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_core_v1_server_info_proto_rawDesc), len(file_spaceone_api_core_v1_server_info_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_core_v1_server_info_proto_goTypes,
		DependencyIndexes: file_spaceone_api_core_v1_server_info_proto_depIdxs,
		MessageInfos:      file_spaceone_api_core_v1_server_info_proto_msgTypes,
	}.Build()
	File_spaceone_api_core_v1_server_info_proto = out.File
	file_spaceone_api_core_v1_server_info_proto_goTypes = nil
	file_spaceone_api_core_v1_server_info_proto_depIdxs = nil
}
