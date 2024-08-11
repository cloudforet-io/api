// description of widget

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/dashboard/v1/private_widget.proto

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
	PrivateWidget_Create_FullMethodName = "/spaceone.api.dashboard.v1.PrivateWidget/create"
	PrivateWidget_Update_FullMethodName = "/spaceone.api.dashboard.v1.PrivateWidget/update"
	PrivateWidget_Delete_FullMethodName = "/spaceone.api.dashboard.v1.PrivateWidget/delete"
	PrivateWidget_Load_FullMethodName   = "/spaceone.api.dashboard.v1.PrivateWidget/load"
	PrivateWidget_Get_FullMethodName    = "/spaceone.api.dashboard.v1.PrivateWidget/get"
	PrivateWidget_List_FullMethodName   = "/spaceone.api.dashboard.v1.PrivateWidget/list"
)

// PrivateWidgetClient is the client API for PrivateWidget service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateWidgetClient interface {
	Create(ctx context.Context, in *CreatePrivateWidgetRequest, opts ...grpc.CallOption) (*PrivateWidgetInfo, error)
	Update(ctx context.Context, in *UpdatePrivateWidgetRequest, opts ...grpc.CallOption) (*PrivateWidgetInfo, error)
	Delete(ctx context.Context, in *PrivateWidgetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Load(ctx context.Context, in *LoadPrivateWidgetRequest, opts ...grpc.CallOption) (*_struct.Struct, error)
	Get(ctx context.Context, in *PrivateWidgetRequest, opts ...grpc.CallOption) (*PrivateWidgetInfo, error)
	List(ctx context.Context, in *PrivateWidgetQuery, opts ...grpc.CallOption) (*PrivateWidgetsInfo, error)
}

type privateWidgetClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateWidgetClient(cc grpc.ClientConnInterface) PrivateWidgetClient {
	return &privateWidgetClient{cc}
}

func (c *privateWidgetClient) Create(ctx context.Context, in *CreatePrivateWidgetRequest, opts ...grpc.CallOption) (*PrivateWidgetInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateWidgetInfo)
	err := c.cc.Invoke(ctx, PrivateWidget_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateWidgetClient) Update(ctx context.Context, in *UpdatePrivateWidgetRequest, opts ...grpc.CallOption) (*PrivateWidgetInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateWidgetInfo)
	err := c.cc.Invoke(ctx, PrivateWidget_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateWidgetClient) Delete(ctx context.Context, in *PrivateWidgetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PrivateWidget_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateWidgetClient) Load(ctx context.Context, in *LoadPrivateWidgetRequest, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, PrivateWidget_Load_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateWidgetClient) Get(ctx context.Context, in *PrivateWidgetRequest, opts ...grpc.CallOption) (*PrivateWidgetInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateWidgetInfo)
	err := c.cc.Invoke(ctx, PrivateWidget_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateWidgetClient) List(ctx context.Context, in *PrivateWidgetQuery, opts ...grpc.CallOption) (*PrivateWidgetsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrivateWidgetsInfo)
	err := c.cc.Invoke(ctx, PrivateWidget_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateWidgetServer is the server API for PrivateWidget service.
// All implementations must embed UnimplementedPrivateWidgetServer
// for forward compatibility.
type PrivateWidgetServer interface {
	Create(context.Context, *CreatePrivateWidgetRequest) (*PrivateWidgetInfo, error)
	Update(context.Context, *UpdatePrivateWidgetRequest) (*PrivateWidgetInfo, error)
	Delete(context.Context, *PrivateWidgetRequest) (*empty.Empty, error)
	Load(context.Context, *LoadPrivateWidgetRequest) (*_struct.Struct, error)
	Get(context.Context, *PrivateWidgetRequest) (*PrivateWidgetInfo, error)
	List(context.Context, *PrivateWidgetQuery) (*PrivateWidgetsInfo, error)
	mustEmbedUnimplementedPrivateWidgetServer()
}

// UnimplementedPrivateWidgetServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPrivateWidgetServer struct{}

func (UnimplementedPrivateWidgetServer) Create(context.Context, *CreatePrivateWidgetRequest) (*PrivateWidgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPrivateWidgetServer) Update(context.Context, *UpdatePrivateWidgetRequest) (*PrivateWidgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPrivateWidgetServer) Delete(context.Context, *PrivateWidgetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPrivateWidgetServer) Load(context.Context, *LoadPrivateWidgetRequest) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedPrivateWidgetServer) Get(context.Context, *PrivateWidgetRequest) (*PrivateWidgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPrivateWidgetServer) List(context.Context, *PrivateWidgetQuery) (*PrivateWidgetsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPrivateWidgetServer) mustEmbedUnimplementedPrivateWidgetServer() {}
func (UnimplementedPrivateWidgetServer) testEmbeddedByValue()                       {}

// UnsafePrivateWidgetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateWidgetServer will
// result in compilation errors.
type UnsafePrivateWidgetServer interface {
	mustEmbedUnimplementedPrivateWidgetServer()
}

func RegisterPrivateWidgetServer(s grpc.ServiceRegistrar, srv PrivateWidgetServer) {
	// If the following call pancis, it indicates UnimplementedPrivateWidgetServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PrivateWidget_ServiceDesc, srv)
}

func _PrivateWidget_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateWidgetServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateWidget_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateWidgetServer).Create(ctx, req.(*CreatePrivateWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateWidget_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivateWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateWidgetServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateWidget_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateWidgetServer).Update(ctx, req.(*UpdatePrivateWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateWidget_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateWidgetServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateWidget_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateWidgetServer).Delete(ctx, req.(*PrivateWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateWidget_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadPrivateWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateWidgetServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateWidget_Load_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateWidgetServer).Load(ctx, req.(*LoadPrivateWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateWidget_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateWidgetServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateWidget_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateWidgetServer).Get(ctx, req.(*PrivateWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateWidget_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateWidgetQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateWidgetServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateWidget_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateWidgetServer).List(ctx, req.(*PrivateWidgetQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateWidget_ServiceDesc is the grpc.ServiceDesc for PrivateWidget service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateWidget_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.dashboard.v1.PrivateWidget",
	HandlerType: (*PrivateWidgetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _PrivateWidget_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _PrivateWidget_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _PrivateWidget_Delete_Handler,
		},
		{
			MethodName: "load",
			Handler:    _PrivateWidget_Load_Handler,
		},
		{
			MethodName: "get",
			Handler:    _PrivateWidget_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _PrivateWidget_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/dashboard/v1/private_widget.proto",
}
