// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/system.proto

package v2

import (
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type SystemInitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Admin         *Admin                 `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Force         bool                   `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SystemInitRequest) Reset() {
	*x = SystemInitRequest{}
	mi := &file_spaceone_api_identity_v2_system_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInitRequest) ProtoMessage() {}

func (x *SystemInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_system_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInitRequest.ProtoReflect.Descriptor instead.
func (*SystemInitRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_system_proto_rawDescGZIP(), []int{0}
}

func (x *SystemInitRequest) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

func (x *SystemInitRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type SystemInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DomainId      string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SystemToken   string                 `protobuf:"bytes,3,opt,name=system_token,json=systemToken,proto3" json:"system_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	mi := &file_spaceone_api_identity_v2_system_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_system_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_system_proto_rawDescGZIP(), []int{1}
}

func (x *SystemInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *SystemInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemInfo) GetSystemToken() string {
	if x != nil {
		return x.SystemToken
	}
	return ""
}

var File_spaceone_api_identity_v2_system_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_system_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x22, 0x60, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x65, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x5b, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_system_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_system_proto_rawDescData = file_spaceone_api_identity_v2_system_proto_rawDesc
)

func file_spaceone_api_identity_v2_system_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_system_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_system_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_system_proto_rawDescData
}

var file_spaceone_api_identity_v2_system_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spaceone_api_identity_v2_system_proto_goTypes = []any{
	(*SystemInitRequest)(nil), // 0: spaceone.api.identity.v2.SystemInitRequest
	(*SystemInfo)(nil),        // 1: spaceone.api.identity.v2.SystemInfo
	(*Admin)(nil),             // 2: spaceone.api.identity.v2.Admin
}
var file_spaceone_api_identity_v2_system_proto_depIdxs = []int32{
	2, // 0: spaceone.api.identity.v2.SystemInitRequest.admin:type_name -> spaceone.api.identity.v2.Admin
	0, // 1: spaceone.api.identity.v2.System.init:input_type -> spaceone.api.identity.v2.SystemInitRequest
	1, // 2: spaceone.api.identity.v2.System.init:output_type -> spaceone.api.identity.v2.SystemInfo
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_system_proto_init() }
func file_spaceone_api_identity_v2_system_proto_init() {
	if File_spaceone_api_identity_v2_system_proto != nil {
		return
	}
	file_spaceone_api_identity_v2_domain_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_identity_v2_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_system_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_system_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v2_system_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_system_proto = out.File
	file_spaceone_api_identity_v2_system_proto_rawDesc = nil
	file_spaceone_api_identity_v2_system_proto_goTypes = nil
	file_spaceone_api_identity_v2_system_proto_depIdxs = nil
}
