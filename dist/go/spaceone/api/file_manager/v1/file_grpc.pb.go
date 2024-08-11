// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/file_manager/v1/file.proto

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
	File_Add_FullMethodName            = "/spaceone.api.file_manager.v1.File/add"
	File_Update_FullMethodName         = "/spaceone.api.file_manager.v1.File/update"
	File_Delete_FullMethodName         = "/spaceone.api.file_manager.v1.File/delete"
	File_GetDownloadUrl_FullMethodName = "/spaceone.api.file_manager.v1.File/get_download_url"
	File_Get_FullMethodName            = "/spaceone.api.file_manager.v1.File/get"
	File_List_FullMethodName           = "/spaceone.api.file_manager.v1.File/list"
	File_Stat_FullMethodName           = "/spaceone.api.file_manager.v1.File/stat"
)

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	Add(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*FileInfo, error)
	Update(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*FileInfo, error)
	Delete(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetDownloadUrl(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileInfo, error)
	Get(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileInfo, error)
	List(ctx context.Context, in *FileSearchQuery, opts ...grpc.CallOption) (*FilesInfo, error)
	Stat(ctx context.Context, in *FileStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) Add(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*FileInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, File_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Update(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*FileInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, File_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Delete(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, File_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) GetDownloadUrl(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, File_GetDownloadUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Get(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, File_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) List(ctx context.Context, in *FileSearchQuery, opts ...grpc.CallOption) (*FilesInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FilesInfo)
	err := c.cc.Invoke(ctx, File_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Stat(ctx context.Context, in *FileStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, File_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility.
type FileServer interface {
	Add(context.Context, *CreateFileRequest) (*FileInfo, error)
	Update(context.Context, *UpdateFileRequest) (*FileInfo, error)
	Delete(context.Context, *FileRequest) (*empty.Empty, error)
	GetDownloadUrl(context.Context, *FileRequest) (*FileInfo, error)
	Get(context.Context, *FileRequest) (*FileInfo, error)
	List(context.Context, *FileSearchQuery) (*FilesInfo, error)
	Stat(context.Context, *FileStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServer struct{}

func (UnimplementedFileServer) Add(context.Context, *CreateFileRequest) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedFileServer) Update(context.Context, *UpdateFileRequest) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFileServer) Delete(context.Context, *FileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileServer) GetDownloadUrl(context.Context, *FileRequest) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadUrl not implemented")
}
func (UnimplementedFileServer) Get(context.Context, *FileRequest) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFileServer) List(context.Context, *FileSearchQuery) (*FilesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFileServer) Stat(context.Context, *FileStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}
func (UnimplementedFileServer) testEmbeddedByValue()              {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	// If the following call pancis, it indicates UnimplementedFileServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Add(ctx, req.(*CreateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Update(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Delete(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_GetDownloadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetDownloadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_GetDownloadUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetDownloadUrl(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Get(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).List(ctx, req.(*FileSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Stat(ctx, req.(*FileStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.file_manager.v1.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _File_Add_Handler,
		},
		{
			MethodName: "update",
			Handler:    _File_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _File_Delete_Handler,
		},
		{
			MethodName: "get_download_url",
			Handler:    _File_GetDownloadUrl_Handler,
		},
		{
			MethodName: "get",
			Handler:    _File_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _File_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _File_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/file_manager/v1/file.proto",
}
