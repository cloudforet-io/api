// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/project_group.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProjectGroup_Create_FullMethodName            = "/spaceone.api.identity.v2.ProjectGroup/create"
	ProjectGroup_Update_FullMethodName            = "/spaceone.api.identity.v2.ProjectGroup/update"
	ProjectGroup_ChangeParentGroup_FullMethodName = "/spaceone.api.identity.v2.ProjectGroup/change_parent_group"
	ProjectGroup_Delete_FullMethodName            = "/spaceone.api.identity.v2.ProjectGroup/delete"
	ProjectGroup_AddUsers_FullMethodName          = "/spaceone.api.identity.v2.ProjectGroup/add_users"
	ProjectGroup_RemoveUsers_FullMethodName       = "/spaceone.api.identity.v2.ProjectGroup/remove_users"
	ProjectGroup_Get_FullMethodName               = "/spaceone.api.identity.v2.ProjectGroup/get"
	ProjectGroup_List_FullMethodName              = "/spaceone.api.identity.v2.ProjectGroup/list"
	ProjectGroup_Stat_FullMethodName              = "/spaceone.api.identity.v2.ProjectGroup/stat"
)

// ProjectGroupClient is the client API for ProjectGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectGroupClient interface {
	Create(ctx context.Context, in *CreateProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error)
	Update(ctx context.Context, in *UpdateProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error)
	ChangeParentGroup(ctx context.Context, in *ChangeParentGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error)
	Delete(ctx context.Context, in *ProjectGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddUsers(ctx context.Context, in *UsersProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error)
	RemoveUsers(ctx context.Context, in *UsersProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error)
	Get(ctx context.Context, in *ProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error)
	List(ctx context.Context, in *ProjectGroupSearchQuery, opts ...grpc.CallOption) (*ProjectGroupsInfo, error)
	Stat(ctx context.Context, in *ProjectGroupStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type projectGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectGroupClient(cc grpc.ClientConnInterface) ProjectGroupClient {
	return &projectGroupClient{cc}
}

func (c *projectGroupClient) Create(ctx context.Context, in *CreateProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) Update(ctx context.Context, in *UpdateProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) ChangeParentGroup(ctx context.Context, in *ChangeParentGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_ChangeParentGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) Delete(ctx context.Context, in *ProjectGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectGroup_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) AddUsers(ctx context.Context, in *UsersProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_AddUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) RemoveUsers(ctx context.Context, in *UsersProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_RemoveUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) Get(ctx context.Context, in *ProjectGroupRequest, opts ...grpc.CallOption) (*ProjectGroupInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) List(ctx context.Context, in *ProjectGroupSearchQuery, opts ...grpc.CallOption) (*ProjectGroupsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProjectGroupsInfo)
	err := c.cc.Invoke(ctx, ProjectGroup_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectGroupClient) Stat(ctx context.Context, in *ProjectGroupStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, ProjectGroup_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectGroupServer is the server API for ProjectGroup service.
// All implementations must embed UnimplementedProjectGroupServer
// for forward compatibility.
type ProjectGroupServer interface {
	Create(context.Context, *CreateProjectGroupRequest) (*ProjectGroupInfo, error)
	Update(context.Context, *UpdateProjectGroupRequest) (*ProjectGroupInfo, error)
	ChangeParentGroup(context.Context, *ChangeParentGroupRequest) (*ProjectGroupInfo, error)
	Delete(context.Context, *ProjectGroupRequest) (*empty.Empty, error)
	AddUsers(context.Context, *UsersProjectGroupRequest) (*ProjectGroupInfo, error)
	RemoveUsers(context.Context, *UsersProjectGroupRequest) (*ProjectGroupInfo, error)
	Get(context.Context, *ProjectGroupRequest) (*ProjectGroupInfo, error)
	List(context.Context, *ProjectGroupSearchQuery) (*ProjectGroupsInfo, error)
	Stat(context.Context, *ProjectGroupStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedProjectGroupServer()
}

// UnimplementedProjectGroupServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectGroupServer struct{}

func (UnimplementedProjectGroupServer) Create(context.Context, *CreateProjectGroupRequest) (*ProjectGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProjectGroupServer) Update(context.Context, *UpdateProjectGroupRequest) (*ProjectGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProjectGroupServer) ChangeParentGroup(context.Context, *ChangeParentGroupRequest) (*ProjectGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeParentGroup not implemented")
}
func (UnimplementedProjectGroupServer) Delete(context.Context, *ProjectGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProjectGroupServer) AddUsers(context.Context, *UsersProjectGroupRequest) (*ProjectGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsers not implemented")
}
func (UnimplementedProjectGroupServer) RemoveUsers(context.Context, *UsersProjectGroupRequest) (*ProjectGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUsers not implemented")
}
func (UnimplementedProjectGroupServer) Get(context.Context, *ProjectGroupRequest) (*ProjectGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProjectGroupServer) List(context.Context, *ProjectGroupSearchQuery) (*ProjectGroupsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectGroupServer) Stat(context.Context, *ProjectGroupStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedProjectGroupServer) mustEmbedUnimplementedProjectGroupServer() {}
func (UnimplementedProjectGroupServer) testEmbeddedByValue()                      {}

// UnsafeProjectGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectGroupServer will
// result in compilation errors.
type UnsafeProjectGroupServer interface {
	mustEmbedUnimplementedProjectGroupServer()
}

func RegisterProjectGroupServer(s grpc.ServiceRegistrar, srv ProjectGroupServer) {
	// If the following call pancis, it indicates UnimplementedProjectGroupServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProjectGroup_ServiceDesc, srv)
}

func _ProjectGroup_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).Create(ctx, req.(*CreateProjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).Update(ctx, req.(*UpdateProjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_ChangeParentGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeParentGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).ChangeParentGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_ChangeParentGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).ChangeParentGroup(ctx, req.(*ChangeParentGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).Delete(ctx, req.(*ProjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_AddUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersProjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).AddUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_AddUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).AddUsers(ctx, req.(*UsersProjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_RemoveUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersProjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).RemoveUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_RemoveUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).RemoveUsers(ctx, req.(*UsersProjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).Get(ctx, req.(*ProjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectGroupSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).List(ctx, req.(*ProjectGroupSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectGroup_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectGroupStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectGroupServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectGroup_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectGroupServer).Stat(ctx, req.(*ProjectGroupStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectGroup_ServiceDesc is the grpc.ServiceDesc for ProjectGroup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectGroup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.ProjectGroup",
	HandlerType: (*ProjectGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ProjectGroup_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ProjectGroup_Update_Handler,
		},
		{
			MethodName: "change_parent_group",
			Handler:    _ProjectGroup_ChangeParentGroup_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ProjectGroup_Delete_Handler,
		},
		{
			MethodName: "add_users",
			Handler:    _ProjectGroup_AddUsers_Handler,
		},
		{
			MethodName: "remove_users",
			Handler:    _ProjectGroup_RemoveUsers_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ProjectGroup_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _ProjectGroup_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _ProjectGroup_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/project_group.proto",
}
