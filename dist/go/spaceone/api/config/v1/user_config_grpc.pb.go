// UserConfig API which configure environments for user

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/config/v1/user_config.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserConfig_Create_FullMethodName = "/spaceone.api.config.v1.UserConfig/create"
	UserConfig_Update_FullMethodName = "/spaceone.api.config.v1.UserConfig/update"
	UserConfig_Set_FullMethodName    = "/spaceone.api.config.v1.UserConfig/set"
	UserConfig_Delete_FullMethodName = "/spaceone.api.config.v1.UserConfig/delete"
	UserConfig_Get_FullMethodName    = "/spaceone.api.config.v1.UserConfig/get"
	UserConfig_List_FullMethodName   = "/spaceone.api.config.v1.UserConfig/list"
	UserConfig_Stat_FullMethodName   = "/spaceone.api.config.v1.UserConfig/stat"
)

// UserConfigClient is the client API for UserConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserConfigClient interface {
	Create(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error)
	Update(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error)
	Set(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error)
	Delete(ctx context.Context, in *UserConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *UserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error)
	List(ctx context.Context, in *UserConfigQuery, opts ...grpc.CallOption) (*UserConfigsInfo, error)
	Stat(ctx context.Context, in *UserConfigStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type userConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewUserConfigClient(cc grpc.ClientConnInterface) UserConfigClient {
	return &userConfigClient{cc}
}

func (c *userConfigClient) Create(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserConfigInfo)
	err := c.cc.Invoke(ctx, UserConfig_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) Update(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserConfigInfo)
	err := c.cc.Invoke(ctx, UserConfig_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) Set(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserConfigInfo)
	err := c.cc.Invoke(ctx, UserConfig_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) Delete(ctx context.Context, in *UserConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserConfig_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) Get(ctx context.Context, in *UserConfigRequest, opts ...grpc.CallOption) (*UserConfigInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserConfigInfo)
	err := c.cc.Invoke(ctx, UserConfig_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) List(ctx context.Context, in *UserConfigQuery, opts ...grpc.CallOption) (*UserConfigsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserConfigsInfo)
	err := c.cc.Invoke(ctx, UserConfig_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) Stat(ctx context.Context, in *UserConfigStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, UserConfig_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserConfigServer is the server API for UserConfig service.
// All implementations must embed UnimplementedUserConfigServer
// for forward compatibility.
type UserConfigServer interface {
	Create(context.Context, *SetUserConfigRequest) (*UserConfigInfo, error)
	Update(context.Context, *SetUserConfigRequest) (*UserConfigInfo, error)
	Set(context.Context, *SetUserConfigRequest) (*UserConfigInfo, error)
	Delete(context.Context, *UserConfigRequest) (*empty.Empty, error)
	Get(context.Context, *UserConfigRequest) (*UserConfigInfo, error)
	List(context.Context, *UserConfigQuery) (*UserConfigsInfo, error)
	Stat(context.Context, *UserConfigStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedUserConfigServer()
}

// UnimplementedUserConfigServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserConfigServer struct{}

func (UnimplementedUserConfigServer) Create(context.Context, *SetUserConfigRequest) (*UserConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserConfigServer) Update(context.Context, *SetUserConfigRequest) (*UserConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserConfigServer) Set(context.Context, *SetUserConfigRequest) (*UserConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedUserConfigServer) Delete(context.Context, *UserConfigRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserConfigServer) Get(context.Context, *UserConfigRequest) (*UserConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserConfigServer) List(context.Context, *UserConfigQuery) (*UserConfigsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserConfigServer) Stat(context.Context, *UserConfigStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedUserConfigServer) mustEmbedUnimplementedUserConfigServer() {}
func (UnimplementedUserConfigServer) testEmbeddedByValue()                    {}

// UnsafeUserConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserConfigServer will
// result in compilation errors.
type UnsafeUserConfigServer interface {
	mustEmbedUnimplementedUserConfigServer()
}

func RegisterUserConfigServer(s grpc.ServiceRegistrar, srv UserConfigServer) {
	// If the following call pancis, it indicates UnimplementedUserConfigServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserConfig_ServiceDesc, srv)
}

func _UserConfig_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).Create(ctx, req.(*SetUserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).Update(ctx, req.(*SetUserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).Set(ctx, req.(*SetUserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).Delete(ctx, req.(*UserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).Get(ctx, req.(*UserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).List(ctx, req.(*UserConfigQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConfigStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).Stat(ctx, req.(*UserConfigStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// UserConfig_ServiceDesc is the grpc.ServiceDesc for UserConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.config.v1.UserConfig",
	HandlerType: (*UserConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _UserConfig_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _UserConfig_Update_Handler,
		},
		{
			MethodName: "set",
			Handler:    _UserConfig_Set_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _UserConfig_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _UserConfig_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _UserConfig_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _UserConfig_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/config/v1/user_config.proto",
}
