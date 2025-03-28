// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/core/v2/handler.proto

package v2

import (
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

type AuthenticationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DomainId      string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DomainId      string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	PublicKey     string                 `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationResponse) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AuthenticationResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type EventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Resource      string                 `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Verb          string                 `protobuf:"bytes,3,opt,name=verb,proto3" json:"verb,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Message       *_struct.Struct        `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v2_handler_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v2_handler_proto_rawDescGZIP(), []int{2}
}

func (x *EventRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EventRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *EventRequest) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *EventRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EventRequest) GetMessage() *_struct.Struct {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_spaceone_api_core_v2_handler_proto protoreflect.FileDescriptor

const file_spaceone_api_core_v2_handler_proto_rawDesc = "" +
	"\n" +
	"\"spaceone/api/core/v2/handler.proto\x12\x14spaceone.api.core.v2\x1a\x1cgoogle/protobuf/struct.proto\"4\n" +
	"\x15AuthenticationRequest\x12\x1b\n" +
	"\tdomain_id\x18\x01 \x01(\tR\bdomainId\"T\n" +
	"\x16AuthenticationResponse\x12\x1b\n" +
	"\tdomain_id\x18\x01 \x01(\tR\bdomainId\x12\x1d\n" +
	"\n" +
	"public_key\x18\x02 \x01(\tR\tpublicKey\"\xa3\x01\n" +
	"\fEventRequest\x12\x18\n" +
	"\aservice\x18\x01 \x01(\tR\aservice\x12\x1a\n" +
	"\bresource\x18\x02 \x01(\tR\bresource\x12\x12\n" +
	"\x04verb\x18\x03 \x01(\tR\x04verb\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\x121\n" +
	"\amessage\x18\x05 \x01(\v2\x17.google.protobuf.StructR\amessageB;Z9github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2b\x06proto3"

var (
	file_spaceone_api_core_v2_handler_proto_rawDescOnce sync.Once
	file_spaceone_api_core_v2_handler_proto_rawDescData []byte
)

func file_spaceone_api_core_v2_handler_proto_rawDescGZIP() []byte {
	file_spaceone_api_core_v2_handler_proto_rawDescOnce.Do(func() {
		file_spaceone_api_core_v2_handler_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_core_v2_handler_proto_rawDesc), len(file_spaceone_api_core_v2_handler_proto_rawDesc)))
	})
	return file_spaceone_api_core_v2_handler_proto_rawDescData
}

var file_spaceone_api_core_v2_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_core_v2_handler_proto_goTypes = []any{
	(*AuthenticationRequest)(nil),  // 0: spaceone.api.core.v2.AuthenticationRequest
	(*AuthenticationResponse)(nil), // 1: spaceone.api.core.v2.AuthenticationResponse
	(*EventRequest)(nil),           // 2: spaceone.api.core.v2.EventRequest
	(*_struct.Struct)(nil),         // 3: google.protobuf.Struct
}
var file_spaceone_api_core_v2_handler_proto_depIdxs = []int32{
	3, // 0: spaceone.api.core.v2.EventRequest.message:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spaceone_api_core_v2_handler_proto_init() }
func file_spaceone_api_core_v2_handler_proto_init() {
	if File_spaceone_api_core_v2_handler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_core_v2_handler_proto_rawDesc), len(file_spaceone_api_core_v2_handler_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spaceone_api_core_v2_handler_proto_goTypes,
		DependencyIndexes: file_spaceone_api_core_v2_handler_proto_depIdxs,
		MessageInfos:      file_spaceone_api_core_v2_handler_proto_msgTypes,
	}.Build()
	File_spaceone_api_core_v2_handler_proto = out.File
	file_spaceone_api_core_v2_handler_proto_goTypes = nil
	file_spaceone_api_core_v2_handler_proto_depIdxs = nil
}
