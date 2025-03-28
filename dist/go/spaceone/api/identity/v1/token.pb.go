// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/identity/v1/token.proto

package v1

import (
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

type IssueTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	UserId      string          `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Credentials *_struct.Struct `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// LOCAL, EXTERNAL
	// +optional
	UserType string `protobuf:"bytes,3,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	DomainId string `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	Timeout int32 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// +optional
	RefreshCount int32 `protobuf:"varint,6,opt,name=refresh_count,json=refreshCount,proto3" json:"refresh_count,omitempty"`
	// if MFA is enabled, verify_code is required
	// +optional
	VerifyCode    string `protobuf:"bytes,7,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueTokenRequest) Reset() {
	*x = IssueTokenRequest{}
	mi := &file_spaceone_api_identity_v1_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenRequest) ProtoMessage() {}

func (x *IssueTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_token_proto_rawDescGZIP(), []int{0}
}

func (x *IssueTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IssueTokenRequest) GetCredentials() *_struct.Struct {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *IssueTokenRequest) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *IssueTokenRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *IssueTokenRequest) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *IssueTokenRequest) GetRefreshCount() int32 {
	if x != nil {
		return x.RefreshCount
	}
	return 0
}

func (x *IssueTokenRequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

type TokenInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	mi := &file_spaceone_api_identity_v1_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_token_proto_rawDescGZIP(), []int{1}
}

func (x *TokenInfo) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenInfo) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_spaceone_api_identity_v1_token_proto protoreflect.FileDescriptor

const file_spaceone_api_identity_v1_token_proto_rawDesc = "" +
	"\n" +
	"$spaceone/api/identity/v1/token.proto\x12\x18spaceone.api.identity.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"\x81\x02\n" +
	"\x11IssueTokenRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x129\n" +
	"\vcredentials\x18\x02 \x01(\v2\x17.google.protobuf.StructR\vcredentials\x12\x1b\n" +
	"\tuser_type\x18\x03 \x01(\tR\buserType\x12\x1b\n" +
	"\tdomain_id\x18\x04 \x01(\tR\bdomainId\x12\x18\n" +
	"\atimeout\x18\x05 \x01(\x05R\atimeout\x12#\n" +
	"\rrefresh_count\x18\x06 \x01(\x05R\frefreshCount\x12\x1f\n" +
	"\vverify_code\x18\a \x01(\tR\n" +
	"verifyCode\"S\n" +
	"\tTokenInfo\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken2\xf6\x01\n" +
	"\x05Token\x12~\n" +
	"\x05issue\x12+.spaceone.api.identity.v1.IssueTokenRequest\x1a#.spaceone.api.identity.v1.TokenInfo\"#\x82\xd3\xe4\x93\x02\x1d:\x01*\"\x18/identity/v1/token/issue\x12m\n" +
	"\arefresh\x12\x16.google.protobuf.Empty\x1a#.spaceone.api.identity.v1.TokenInfo\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/identity/v1/token/refreshB?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v1b\x06proto3"

var (
	file_spaceone_api_identity_v1_token_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v1_token_proto_rawDescData []byte
)

func file_spaceone_api_identity_v1_token_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v1_token_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v1_token_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v1_token_proto_rawDesc), len(file_spaceone_api_identity_v1_token_proto_rawDesc)))
	})
	return file_spaceone_api_identity_v1_token_proto_rawDescData
}

var file_spaceone_api_identity_v1_token_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spaceone_api_identity_v1_token_proto_goTypes = []any{
	(*IssueTokenRequest)(nil), // 0: spaceone.api.identity.v1.IssueTokenRequest
	(*TokenInfo)(nil),         // 1: spaceone.api.identity.v1.TokenInfo
	(*_struct.Struct)(nil),    // 2: google.protobuf.Struct
	(*empty.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_spaceone_api_identity_v1_token_proto_depIdxs = []int32{
	2, // 0: spaceone.api.identity.v1.IssueTokenRequest.credentials:type_name -> google.protobuf.Struct
	0, // 1: spaceone.api.identity.v1.Token.issue:input_type -> spaceone.api.identity.v1.IssueTokenRequest
	3, // 2: spaceone.api.identity.v1.Token.refresh:input_type -> google.protobuf.Empty
	1, // 3: spaceone.api.identity.v1.Token.issue:output_type -> spaceone.api.identity.v1.TokenInfo
	1, // 4: spaceone.api.identity.v1.Token.refresh:output_type -> spaceone.api.identity.v1.TokenInfo
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v1_token_proto_init() }
func file_spaceone_api_identity_v1_token_proto_init() {
	if File_spaceone_api_identity_v1_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_identity_v1_token_proto_rawDesc), len(file_spaceone_api_identity_v1_token_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v1_token_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v1_token_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v1_token_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v1_token_proto = out.File
	file_spaceone_api_identity_v1_token_proto_goTypes = nil
	file_spaceone_api_identity_v1_token_proto_depIdxs = nil
}
