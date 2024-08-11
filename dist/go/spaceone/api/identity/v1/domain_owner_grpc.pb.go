// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/identity/v1/domain_owner.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DomainOwner_Create_FullMethodName = "/spaceone.api.identity.v1.DomainOwner/create"
	DomainOwner_Update_FullMethodName = "/spaceone.api.identity.v1.DomainOwner/update"
	DomainOwner_Delete_FullMethodName = "/spaceone.api.identity.v1.DomainOwner/delete"
	DomainOwner_Get_FullMethodName    = "/spaceone.api.identity.v1.DomainOwner/get"
)

// DomainOwnerClient is the client API for DomainOwner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainOwnerClient interface {
	Create(ctx context.Context, in *CreateDomainOwner, opts ...grpc.CallOption) (*DomainOwnerInfo, error)
	Update(ctx context.Context, in *UpdateDomainOwner, opts ...grpc.CallOption) (*DomainOwnerInfo, error)
	Delete(ctx context.Context, in *DomainOwnerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetDomainOwnerRequest, opts ...grpc.CallOption) (*DomainOwnerInfo, error)
}

type domainOwnerClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainOwnerClient(cc grpc.ClientConnInterface) DomainOwnerClient {
	return &domainOwnerClient{cc}
}

func (c *domainOwnerClient) Create(ctx context.Context, in *CreateDomainOwner, opts ...grpc.CallOption) (*DomainOwnerInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DomainOwnerInfo)
	err := c.cc.Invoke(ctx, DomainOwner_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainOwnerClient) Update(ctx context.Context, in *UpdateDomainOwner, opts ...grpc.CallOption) (*DomainOwnerInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DomainOwnerInfo)
	err := c.cc.Invoke(ctx, DomainOwner_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainOwnerClient) Delete(ctx context.Context, in *DomainOwnerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DomainOwner_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainOwnerClient) Get(ctx context.Context, in *GetDomainOwnerRequest, opts ...grpc.CallOption) (*DomainOwnerInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DomainOwnerInfo)
	err := c.cc.Invoke(ctx, DomainOwner_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainOwnerServer is the server API for DomainOwner service.
// All implementations must embed UnimplementedDomainOwnerServer
// for forward compatibility.
type DomainOwnerServer interface {
	Create(context.Context, *CreateDomainOwner) (*DomainOwnerInfo, error)
	Update(context.Context, *UpdateDomainOwner) (*DomainOwnerInfo, error)
	Delete(context.Context, *DomainOwnerRequest) (*empty.Empty, error)
	Get(context.Context, *GetDomainOwnerRequest) (*DomainOwnerInfo, error)
	mustEmbedUnimplementedDomainOwnerServer()
}

// UnimplementedDomainOwnerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDomainOwnerServer struct{}

func (UnimplementedDomainOwnerServer) Create(context.Context, *CreateDomainOwner) (*DomainOwnerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDomainOwnerServer) Update(context.Context, *UpdateDomainOwner) (*DomainOwnerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDomainOwnerServer) Delete(context.Context, *DomainOwnerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDomainOwnerServer) Get(context.Context, *GetDomainOwnerRequest) (*DomainOwnerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDomainOwnerServer) mustEmbedUnimplementedDomainOwnerServer() {}
func (UnimplementedDomainOwnerServer) testEmbeddedByValue()                     {}

// UnsafeDomainOwnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainOwnerServer will
// result in compilation errors.
type UnsafeDomainOwnerServer interface {
	mustEmbedUnimplementedDomainOwnerServer()
}

func RegisterDomainOwnerServer(s grpc.ServiceRegistrar, srv DomainOwnerServer) {
	// If the following call pancis, it indicates UnimplementedDomainOwnerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DomainOwner_ServiceDesc, srv)
}

func _DomainOwner_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainOwnerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainOwner_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainOwnerServer).Create(ctx, req.(*CreateDomainOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainOwner_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainOwnerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainOwner_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainOwnerServer).Update(ctx, req.(*UpdateDomainOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainOwner_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainOwnerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainOwner_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainOwnerServer).Delete(ctx, req.(*DomainOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainOwner_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainOwnerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainOwner_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainOwnerServer).Get(ctx, req.(*GetDomainOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainOwner_ServiceDesc is the grpc.ServiceDesc for DomainOwner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainOwner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v1.DomainOwner",
	HandlerType: (*DomainOwnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _DomainOwner_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _DomainOwner_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _DomainOwner_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _DomainOwner_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v1/domain_owner.proto",
}
