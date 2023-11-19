// User API which allows member management within project, company, and domain
// note: Administrator must register User first.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/user_group.proto

package v2

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserGroup_Create_FullMethodName      = "/spaceone.api.identity.v2.UserGroup/create"
	UserGroup_Update_FullMethodName      = "/spaceone.api.identity.v2.UserGroup/update"
	UserGroup_Delete_FullMethodName      = "/spaceone.api.identity.v2.UserGroup/delete"
	UserGroup_AddUsers_FullMethodName    = "/spaceone.api.identity.v2.UserGroup/add_users"
	UserGroup_RemoveUsers_FullMethodName = "/spaceone.api.identity.v2.UserGroup/remove_users"
	UserGroup_Get_FullMethodName         = "/spaceone.api.identity.v2.UserGroup/get"
	UserGroup_List_FullMethodName        = "/spaceone.api.identity.v2.UserGroup/list"
	UserGroup_Stat_FullMethodName        = "/spaceone.api.identity.v2.UserGroup/stat"
)

// UserGroupClient is the client API for UserGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGroupClient interface {
	Create(ctx context.Context, in *CreateUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error)
	Update(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error)
	Delete(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddUsers(ctx context.Context, in *UsersUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error)
	RemoveUsers(ctx context.Context, in *UsersUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error)
	Get(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error)
	List(ctx context.Context, in *UserGroupSearchQuery, opts ...grpc.CallOption) (*UserGroupsInfo, error)
	Stat(ctx context.Context, in *UserGroupStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type userGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGroupClient(cc grpc.ClientConnInterface) UserGroupClient {
	return &userGroupClient{cc}
}

func (c *userGroupClient) Create(ctx context.Context, in *CreateUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error) {
	out := new(UserGroupInfo)
	err := c.cc.Invoke(ctx, UserGroup_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) Update(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error) {
	out := new(UserGroupInfo)
	err := c.cc.Invoke(ctx, UserGroup_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) Delete(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserGroup_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) AddUsers(ctx context.Context, in *UsersUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error) {
	out := new(UserGroupInfo)
	err := c.cc.Invoke(ctx, UserGroup_AddUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) RemoveUsers(ctx context.Context, in *UsersUserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error) {
	out := new(UserGroupInfo)
	err := c.cc.Invoke(ctx, UserGroup_RemoveUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) Get(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*UserGroupInfo, error) {
	out := new(UserGroupInfo)
	err := c.cc.Invoke(ctx, UserGroup_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) List(ctx context.Context, in *UserGroupSearchQuery, opts ...grpc.CallOption) (*UserGroupsInfo, error) {
	out := new(UserGroupsInfo)
	err := c.cc.Invoke(ctx, UserGroup_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupClient) Stat(ctx context.Context, in *UserGroupStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, UserGroup_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGroupServer is the server API for UserGroup service.
// All implementations must embed UnimplementedUserGroupServer
// for forward compatibility
type UserGroupServer interface {
	Create(context.Context, *CreateUserGroupRequest) (*UserGroupInfo, error)
	Update(context.Context, *UpdateUserGroupRequest) (*UserGroupInfo, error)
	Delete(context.Context, *UserGroupRequest) (*empty.Empty, error)
	AddUsers(context.Context, *UsersUserGroupRequest) (*UserGroupInfo, error)
	RemoveUsers(context.Context, *UsersUserGroupRequest) (*UserGroupInfo, error)
	Get(context.Context, *UserGroupRequest) (*UserGroupInfo, error)
	List(context.Context, *UserGroupSearchQuery) (*UserGroupsInfo, error)
	Stat(context.Context, *UserGroupStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedUserGroupServer()
}

// UnimplementedUserGroupServer must be embedded to have forward compatible implementations.
type UnimplementedUserGroupServer struct {
}

func (UnimplementedUserGroupServer) Create(context.Context, *CreateUserGroupRequest) (*UserGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserGroupServer) Update(context.Context, *UpdateUserGroupRequest) (*UserGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserGroupServer) Delete(context.Context, *UserGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserGroupServer) AddUsers(context.Context, *UsersUserGroupRequest) (*UserGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsers not implemented")
}
func (UnimplementedUserGroupServer) RemoveUsers(context.Context, *UsersUserGroupRequest) (*UserGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUsers not implemented")
}
func (UnimplementedUserGroupServer) Get(context.Context, *UserGroupRequest) (*UserGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserGroupServer) List(context.Context, *UserGroupSearchQuery) (*UserGroupsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserGroupServer) Stat(context.Context, *UserGroupStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedUserGroupServer) mustEmbedUnimplementedUserGroupServer() {}

// UnsafeUserGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGroupServer will
// result in compilation errors.
type UnsafeUserGroupServer interface {
	mustEmbedUnimplementedUserGroupServer()
}

func RegisterUserGroupServer(s grpc.ServiceRegistrar, srv UserGroupServer) {
	s.RegisterService(&UserGroup_ServiceDesc, srv)
}

func _UserGroup_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).Create(ctx, req.(*CreateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).Update(ctx, req.(*UpdateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).Delete(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_AddUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).AddUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_AddUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).AddUsers(ctx, req.(*UsersUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_RemoveUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).RemoveUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_RemoveUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).RemoveUsers(ctx, req.(*UsersUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).Get(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).List(ctx, req.(*UserGroupSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroup_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGroup_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServer).Stat(ctx, req.(*UserGroupStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGroup_ServiceDesc is the grpc.ServiceDesc for UserGroup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGroup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.UserGroup",
	HandlerType: (*UserGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _UserGroup_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _UserGroup_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _UserGroup_Delete_Handler,
		},
		{
			MethodName: "add_users",
			Handler:    _UserGroup_AddUsers_Handler,
		},
		{
			MethodName: "remove_users",
			Handler:    _UserGroup_RemoveUsers_Handler,
		},
		{
			MethodName: "get",
			Handler:    _UserGroup_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _UserGroup_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _UserGroup_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/user_group.proto",
}
