// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/cost_analysis/v1/cost_report.proto

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
	CostReport_Send_FullMethodName   = "/spaceone.api.cost_analysis.v1.CostReport/send"
	CostReport_GetUrl_FullMethodName = "/spaceone.api.cost_analysis.v1.CostReport/get_url"
	CostReport_Get_FullMethodName    = "/spaceone.api.cost_analysis.v1.CostReport/get"
	CostReport_List_FullMethodName   = "/spaceone.api.cost_analysis.v1.CostReport/list"
	CostReport_Stat_FullMethodName   = "/spaceone.api.cost_analysis.v1.CostReport/stat"
)

// CostReportClient is the client API for CostReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CostReportClient interface {
	Send(ctx context.Context, in *CostReportRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUrl(ctx context.Context, in *GetUrlCostReportRequest, opts ...grpc.CallOption) (*_struct.Struct, error)
	Get(ctx context.Context, in *CostReportRequest, opts ...grpc.CallOption) (*CostReportInfo, error)
	List(ctx context.Context, in *CostReportQuery, opts ...grpc.CallOption) (*CostReportsInfo, error)
	Stat(ctx context.Context, in *CostReportStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type costReportClient struct {
	cc grpc.ClientConnInterface
}

func NewCostReportClient(cc grpc.ClientConnInterface) CostReportClient {
	return &costReportClient{cc}
}

func (c *costReportClient) Send(ctx context.Context, in *CostReportRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CostReport_Send_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costReportClient) GetUrl(ctx context.Context, in *GetUrlCostReportRequest, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CostReport_GetUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costReportClient) Get(ctx context.Context, in *CostReportRequest, opts ...grpc.CallOption) (*CostReportInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostReportInfo)
	err := c.cc.Invoke(ctx, CostReport_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costReportClient) List(ctx context.Context, in *CostReportQuery, opts ...grpc.CallOption) (*CostReportsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CostReportsInfo)
	err := c.cc.Invoke(ctx, CostReport_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *costReportClient) Stat(ctx context.Context, in *CostReportStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CostReport_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CostReportServer is the server API for CostReport service.
// All implementations must embed UnimplementedCostReportServer
// for forward compatibility.
type CostReportServer interface {
	Send(context.Context, *CostReportRequest) (*empty.Empty, error)
	GetUrl(context.Context, *GetUrlCostReportRequest) (*_struct.Struct, error)
	Get(context.Context, *CostReportRequest) (*CostReportInfo, error)
	List(context.Context, *CostReportQuery) (*CostReportsInfo, error)
	Stat(context.Context, *CostReportStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCostReportServer()
}

// UnimplementedCostReportServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCostReportServer struct{}

func (UnimplementedCostReportServer) Send(context.Context, *CostReportRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedCostReportServer) GetUrl(context.Context, *GetUrlCostReportRequest) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedCostReportServer) Get(context.Context, *CostReportRequest) (*CostReportInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCostReportServer) List(context.Context, *CostReportQuery) (*CostReportsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCostReportServer) Stat(context.Context, *CostReportStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCostReportServer) mustEmbedUnimplementedCostReportServer() {}
func (UnimplementedCostReportServer) testEmbeddedByValue()                    {}

// UnsafeCostReportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CostReportServer will
// result in compilation errors.
type UnsafeCostReportServer interface {
	mustEmbedUnimplementedCostReportServer()
}

func RegisterCostReportServer(s grpc.ServiceRegistrar, srv CostReportServer) {
	// If the following call pancis, it indicates UnimplementedCostReportServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CostReport_ServiceDesc, srv)
}

func _CostReport_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReport_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportServer).Send(ctx, req.(*CostReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostReport_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlCostReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReport_GetUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportServer).GetUrl(ctx, req.(*GetUrlCostReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostReport_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReport_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportServer).Get(ctx, req.(*CostReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostReport_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReport_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportServer).List(ctx, req.(*CostReportQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CostReport_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CostReportStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CostReportServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CostReport_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CostReportServer).Stat(ctx, req.(*CostReportStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CostReport_ServiceDesc is the grpc.ServiceDesc for CostReport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CostReport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.cost_analysis.v1.CostReport",
	HandlerType: (*CostReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _CostReport_Send_Handler,
		},
		{
			MethodName: "get_url",
			Handler:    _CostReport_GetUrl_Handler,
		},
		{
			MethodName: "get",
			Handler:    _CostReport_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CostReport_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CostReport_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/cost_analysis/v1/cost_report.proto",
}
