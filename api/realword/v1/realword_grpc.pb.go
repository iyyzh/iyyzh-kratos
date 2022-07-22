// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: realword/v1/realword.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RealwordClient is the client API for Realword service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealwordClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type realwordClient struct {
	cc grpc.ClientConnInterface
}

func NewRealwordClient(cc grpc.ClientConnInterface) RealwordClient {
	return &realwordClient{cc}
}

func (c *realwordClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/realword.v1.Realword/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realwordClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, "/realword.v1.Realword/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realwordClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, "/realword.v1.Realword/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realwordClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, "/realword.v1.Realword/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realwordClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error) {
	out := new(GetUserListReply)
	err := c.cc.Invoke(ctx, "/realword.v1.Realword/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realwordClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/realword.v1.Realword/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealwordServer is the server API for Realword service.
// All implementations must embed UnimplementedRealwordServer
// for forward compatibility
type RealwordServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedRealwordServer()
}

// UnimplementedRealwordServer must be embedded to have forward compatible implementations.
type UnimplementedRealwordServer struct {
}

func (UnimplementedRealwordServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedRealwordServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedRealwordServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedRealwordServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedRealwordServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedRealwordServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRealwordServer) mustEmbedUnimplementedRealwordServer() {}

// UnsafeRealwordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealwordServer will
// result in compilation errors.
type UnsafeRealwordServer interface {
	mustEmbedUnimplementedRealwordServer()
}

func RegisterRealwordServer(s grpc.ServiceRegistrar, srv RealwordServer) {
	s.RegisterService(&Realword_ServiceDesc, srv)
}

func _Realword_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealwordServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realword.v1.Realword/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealwordServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realword_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealwordServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realword.v1.Realword/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealwordServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realword_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealwordServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realword.v1.Realword/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealwordServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realword_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealwordServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realword.v1.Realword/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealwordServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realword_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealwordServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realword.v1.Realword/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealwordServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realword_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealwordServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realword.v1.Realword/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealwordServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Realword_ServiceDesc is the grpc.ServiceDesc for Realword service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Realword_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realword.v1.Realword",
	HandlerType: (*RealwordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Realword_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Realword_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Realword_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Realword_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _Realword_GetUserList_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Realword_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realword/v1/realword.proto",
}