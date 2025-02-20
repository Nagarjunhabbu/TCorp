// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: tcorp.proto

package protocorp

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

const (
	UserInfo_GetUserById_FullMethodName  = "/UserInfo/GetUserById"
	UserInfo_GetUserByIds_FullMethodName = "/UserInfo/GetUserByIds"
	UserInfo_SearchUser_FullMethodName   = "/UserInfo/SearchUser"
)

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdRespnose, error)
	GetUserByIds(ctx context.Context, in *GetUserByIdsRequest, opts ...grpc.CallOption) (*GetUserByIdsResponse, error)
	SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdRespnose, error) {
	out := new(GetUserByIdRespnose)
	err := c.cc.Invoke(ctx, UserInfo_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) GetUserByIds(ctx context.Context, in *GetUserByIdsRequest, opts ...grpc.CallOption) (*GetUserByIdsResponse, error) {
	out := new(GetUserByIdsResponse)
	err := c.cc.Invoke(ctx, UserInfo_GetUserByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, UserInfo_SearchUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations should embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdRespnose, error)
	GetUserByIds(context.Context, *GetUserByIdsRequest) (*GetUserByIdsResponse, error)
	SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
}

// UnimplementedUserInfoServer should be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (UnimplementedUserInfoServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdRespnose, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserInfoServer) GetUserByIds(context.Context, *GetUserByIdsRequest) (*GetUserByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByIds not implemented")
}
func (UnimplementedUserInfoServer) SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_GetUserByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_GetUserByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserByIds(ctx, req.(*GetUserByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_SearchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).SearchUser(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserInfo_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByIds",
			Handler:    _UserInfo_GetUserByIds_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _UserInfo_SearchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tcorp.proto",
}
