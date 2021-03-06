// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRequest
	UserResponse
	UserFilter
*/
package user

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRequest struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserResponse struct {
	Id    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *UserFilter) Reset()                    { *m = UserFilter{} }
func (m *UserFilter) String() string            { return proto.CompactTextString(m) }
func (*UserFilter) ProtoMessage()               {}
func (*UserFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "user.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*UserFilter)(nil), "user.UserFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	// Get all users with filter
	GetUsers(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (User_GetUsersClient, error)
	// Create a new user.
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUsers(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (User_GetUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_User_serviceDesc.Streams[0], c.cc, "/user.User/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type User_GetUsersClient interface {
	Recv() (*UserRequest, error)
	grpc.ClientStream
}

type userGetUsersClient struct {
	grpc.ClientStream
}

func (x *userGetUsersClient) Recv() (*UserRequest, error) {
	m := new(UserRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/user.User/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	// Get all users with filter
	GetUsers(*UserFilter, User_GetUsersServer) error
	// Create a new user.
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServer).GetUsers(m, &userGetUsersServer{stream})
}

type User_GetUsersServer interface {
	Send(*UserRequest) error
	grpc.ServerStream
}

type userGetUsersServer struct {
	grpc.ServerStream
}

func (x *userGetUsersServer) Send(m *UserRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _User_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xbf, 0xca, 0xc2, 0x40,
	0x10, 0xc4, 0x73, 0xf9, 0x92, 0x0f, 0x5d, 0x45, 0x74, 0xb1, 0x08, 0xa9, 0xc2, 0x55, 0xa9, 0x82,
	0x18, 0x7c, 0x02, 0xc1, 0x58, 0x07, 0x7c, 0x80, 0x48, 0xb6, 0x38, 0xc8, 0x3f, 0x6f, 0x2f, 0xef,
	0x2f, 0x77, 0x11, 0x23, 0x58, 0xda, 0xcd, 0x0e, 0x33, 0x3f, 0x98, 0x05, 0x18, 0x99, 0x74, 0x36,
	0xe8, 0xde, 0xf4, 0x18, 0x58, 0x2d, 0x0b, 0x58, 0xdd, 0x98, 0x74, 0x49, 0x8f, 0x91, 0xd8, 0xe0,
	0x06, 0x7c, 0x55, 0x47, 0x22, 0x11, 0x69, 0x58, 0xfa, 0xaa, 0x46, 0x84, 0xa0, 0xab, 0x5a, 0x8a,
	0xfc, 0x44, 0xa4, 0xcb, 0xd2, 0x69, 0xdc, 0x43, 0x48, 0x6d, 0xa5, 0x9a, 0xe8, 0xcf, 0x99, 0xd3,
	0x21, 0xaf, 0xb0, 0x9e, 0x40, 0x3c, 0xf4, 0x1d, 0xd3, 0x0f, 0xa4, 0x04, 0xc0, 0x92, 0x2e, 0xaa,
	0x31, 0xa4, 0xdf, 0x3d, 0x31, 0xf7, 0x8e, 0x1a, 0x02, 0x9b, 0xc0, 0x1c, 0x16, 0x05, 0x19, 0x2b,
	0x19, 0xb7, 0x99, 0xdb, 0x36, 0x37, 0xe3, 0xdd, 0xec, 0xbc, 0xe6, 0x49, 0xef, 0x20, 0xf0, 0x04,
	0x70, 0xd6, 0x54, 0x19, 0x72, 0x88, 0xef, 0x50, 0x8c, 0x9f, 0xd6, 0xb4, 0x46, 0x7a, 0xf7, 0x7f,
	0xf7, 0xb5, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xd8, 0xa9, 0xb4, 0x43, 0x01, 0x00, 0x00,
}
