// A CostQuerySet is a set of saved queries a User frequently uses as a setting.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/cost_analysis/v1/cost_query_set.proto

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
	CostQuerySet_Create_FullMethodName = "/spaceone.api.cost_analysis.v1.CostQuerySet/create"
	CostQuerySet_Update_FullMethodName = "/spaceone.api.cost_analysis.v1.CostQuerySet/update"
	CostQuerySet_Delete_FullMethodName = "/spaceone.api.cost_analysis.v1.CostQuerySet/delete"
	CostQuerySet_Get_FullMethodName    = "/spaceone.api.cost_analysis.v1.CostQuerySet/get"
	CostQuerySet_List_FullMethodName   = "/spaceone.api.cost_analysis.v1.CostQuerySet/list"
	CostQuerySet_Stat_FullMethodName   = "/spaceone.api.cost_analysis.v1.CostQuerySet/stat"
)

// CostQuerySetClient is the client API for CostQuerySet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostQuerySetClient interface {
	// Creates a new CostQuerySet. You can make your own custom query that meets your needs, and input it as an `option` parameter of the resource. Queries such as `group_by` and `granularity` are provided by default.
	Create(ctx context.Context, in *CreateCostQuerySetRequest, opts ...grpc.CallOption) (*CostQuerySetInfo, error)
	// Updates a specific CostQuerySet. You can make changes in the details of queries.
	Update(ctx context.Context, in *UpdateCostQuerySetRequest, opts ...grpc.CallOption) (*CostQuerySetInfo, error)
	// Deletes a specific CostQuerySet. You must specify the `cost_query_set_id` of the CostQuerySet to delete.
	Delete(ctx context.Context, in *CostQuerySetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific CostQuerySet. Prints detailed information about the CostQuerySet, including the details of queries.
	Get(ctx context.Context, in *CostQuerySetRequest, opts ...grpc.CallOption) (*CostQuerySetInfo, error)
	// Gets a list of all CostQuerySets. You can use a query to get a filtered list of CostQuerySets.
	List(ctx context.Context, in *CostQuerySetQuery, opts ...grpc.CallOption) (*CostQuerySetsInfo, error)
	Stat(ctx context.Context, in *CostQuerySetStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type costQuerySetClient struct {
	cc grpc.ClientConnInterface
}

func NewCostQuerySetClient(cc grpc.ClientConnInterface) CostQuerySetClient {
	return &costQuerySetClient{cc}
}

func (c *costQuerySetClient) Create(ctx context.Context, in *CreateCostQuerySetRequest, opts ...grpc.CallOption) (*CostQuerySetInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostQuerySetInfo)
	err := c.cc.Invoke(ctx, CostQuerySet_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costQuerySetClient) Update(ctx context.Context, in *UpdateCostQuerySetRequest, opts ...grpc.CallOption) (*CostQuerySetInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostQuerySetInfo)
	err := c.cc.Invoke(ctx, CostQuerySet_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costQuerySetClient) Delete(ctx context.Context, in *CostQuerySetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CostQuerySet_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costQuerySetClient) Get(ctx context.Context, in *CostQuerySetRequest, opts ...grpc.CallOption) (*CostQuerySetInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostQuerySetInfo)
	err := c.cc.Invoke(ctx, CostQuerySet_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costQuerySetClient) List(ctx context.Context, in *CostQuerySetQuery, opts ...grpc.CallOption) (*CostQuerySetsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostQuerySetsInfo)
	err := c.cc.Invoke(ctx, CostQuerySet_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costQuerySetClient) Stat(ctx context.Context, in *CostQuerySetStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CostQuerySet_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CostQuerySetServer is the server API for CostQuerySet service.
// All implementations must embed UnimplementedCostQuerySetServer
// for forward compatibility.
type CostQuerySetServer interface {
	// Creates a new CostQuerySet. You can make your own custom query that meets your needs, and input it as an `option` parameter of the resource. Queries such as `group_by` and `granularity` are provided by default.
	Create(context.Context, *CreateCostQuerySetRequest) (*CostQuerySetInfo, error)
	// Updates a specific CostQuerySet. You can make changes in the details of queries.
	Update(context.Context, *UpdateCostQuerySetRequest) (*CostQuerySetInfo, error)
	// Deletes a specific CostQuerySet. You must specify the `cost_query_set_id` of the CostQuerySet to delete.
	Delete(context.Context, *CostQuerySetRequest) (*empty.Empty, error)
	// Gets a specific CostQuerySet. Prints detailed information about the CostQuerySet, including the details of queries.
	Get(context.Context, *CostQuerySetRequest) (*CostQuerySetInfo, error)
	// Gets a list of all CostQuerySets. You can use a query to get a filtered list of CostQuerySets.
	List(context.Context, *CostQuerySetQuery) (*CostQuerySetsInfo, error)
	Stat(context.Context, *CostQuerySetStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCostQuerySetServer()
}

// UnimplementedCostQuerySetServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCostQuerySetServer struct{}

func (UnimplementedCostQuerySetServer) Create(context.Context, *CreateCostQuerySetRequest) (*CostQuerySetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCostQuerySetServer) Update(context.Context, *UpdateCostQuerySetRequest) (*CostQuerySetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCostQuerySetServer) Delete(context.Context, *CostQuerySetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCostQuerySetServer) Get(context.Context, *CostQuerySetRequest) (*CostQuerySetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCostQuerySetServer) List(context.Context, *CostQuerySetQuery) (*CostQuerySetsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCostQuerySetServer) Stat(context.Context, *CostQuerySetStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCostQuerySetServer) mustEmbedUnimplementedCostQuerySetServer() {}
func (UnimplementedCostQuerySetServer) testEmbeddedByValue()                      {}

// UnsafeCostQuerySetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostQuerySetServer will
// result in compilation errors.
type UnsafeCostQuerySetServer interface {
	mustEmbedUnimplementedCostQuerySetServer()
}

func RegisterCostQuerySetServer(s grpc.ServiceRegistrar, srv CostQuerySetServer) {
	// If the following call pancis, it indicates UnimplementedCostQuerySetServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CostQuerySet_ServiceDesc, srv)
}

func _CostQuerySet_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCostQuerySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostQuerySetServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostQuerySet_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostQuerySetServer).Create(ctx, req.(*CreateCostQuerySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostQuerySet_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCostQuerySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostQuerySetServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostQuerySet_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostQuerySetServer).Update(ctx, req.(*UpdateCostQuerySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostQuerySet_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostQuerySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostQuerySetServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostQuerySet_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostQuerySetServer).Delete(ctx, req.(*CostQuerySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostQuerySet_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostQuerySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostQuerySetServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostQuerySet_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostQuerySetServer).Get(ctx, req.(*CostQuerySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostQuerySet_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostQuerySetQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostQuerySetServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostQuerySet_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostQuerySetServer).List(ctx, req.(*CostQuerySetQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostQuerySet_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostQuerySetStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostQuerySetServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostQuerySet_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostQuerySetServer).Stat(ctx, req.(*CostQuerySetStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CostQuerySet_ServiceDesc is the grpc.ServiceDesc for CostQuerySet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CostQuerySet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.cost_analysis.v1.CostQuerySet",
	HandlerType: (*CostQuerySetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CostQuerySet_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CostQuerySet_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CostQuerySet_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _CostQuerySet_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CostQuerySet_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CostQuerySet_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/cost_analysis/v1/cost_query_set.proto",
}
