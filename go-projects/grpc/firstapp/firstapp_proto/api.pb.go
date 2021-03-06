// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package firstapp_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccountSystem service

type AccountSystemClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	QueryAccount(ctx context.Context, in *AccountQueryRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type accountSystemClient struct {
	cc *grpc.ClientConn
}

func NewAccountSystemClient(cc *grpc.ClientConn) AccountSystemClient {
	return &accountSystemClient{cc}
}

func (c *accountSystemClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := grpc.Invoke(ctx, "/firstapp.proto.AccountSystem/CreateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSystemClient) QueryAccount(ctx context.Context, in *AccountQueryRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := grpc.Invoke(ctx, "/firstapp.proto.AccountSystem/QueryAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSystemClient) Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/firstapp.proto.AccountSystem/ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountSystem service

type AccountSystemServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	QueryAccount(context.Context, *AccountQueryRequest) (*AccountResponse, error)
	Ping(context.Context, *Ping) (*Pong, error)
}

func RegisterAccountSystemServer(s *grpc.Server, srv AccountSystemServer) {
	s.RegisterService(&_AccountSystem_serviceDesc, srv)
}

func _AccountSystem_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountSystemServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firstapp.proto.AccountSystem/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountSystemServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountSystem_QueryAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountSystemServer).QueryAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firstapp.proto.AccountSystem/QueryAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountSystemServer).QueryAccount(ctx, req.(*AccountQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountSystem_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountSystemServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firstapp.proto.AccountSystem/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountSystemServer).Ping(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "firstapp.proto.AccountSystem",
	HandlerType: (*AccountSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountSystem_CreateAccount_Handler,
		},
		{
			MethodName: "QueryAccount",
			Handler:    _AccountSystem_QueryAccount_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _AccountSystem_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0xcb, 0x2c, 0x2a, 0x2e, 0x49, 0x2c, 0x28, 0x80,
	0xf0, 0xa5, 0x78, 0x13, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0xa0, 0x5c, 0x9e, 0xe4, 0xfc, 0xdc,
	0xdc, 0xfc, 0x3c, 0x08, 0xcf, 0xe8, 0x2f, 0x23, 0x17, 0xaf, 0x23, 0x44, 0x3e, 0xb8, 0xb2, 0xb8,
	0x24, 0x35, 0x57, 0x28, 0x8e, 0x8b, 0xd7, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x15, 0x2a, 0x2c, 0xa4,
	0xa2, 0x87, 0x6a, 0xa0, 0x1e, 0x8a, 0x74, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94, 0x2a,
	0x01, 0x55, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x0c, 0x42, 0x61, 0x5c, 0x3c, 0x81, 0xa5,
	0xa9, 0x45, 0x95, 0x30, 0xe3, 0x95, 0xd1, 0x35, 0x42, 0x25, 0xc0, 0x8a, 0x60, 0xa6, 0xcb, 0xe3,
	0x50, 0x84, 0x64, 0xae, 0x09, 0x17, 0x4b, 0x41, 0x66, 0x5e, 0xba, 0x90, 0x08, 0xba, 0xd2, 0x80,
	0xcc, 0xbc, 0x74, 0x29, 0x4c, 0xd1, 0xfc, 0xbc, 0x74, 0x25, 0x06, 0x27, 0xa6, 0x00, 0xc6, 0x24,
	0x36, 0xb0, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xb3, 0xe9, 0x9e, 0x44, 0x01, 0x00,
	0x00,
}
