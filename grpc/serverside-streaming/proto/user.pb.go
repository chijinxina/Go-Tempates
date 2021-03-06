// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

//服务端流式grpc

package serverside_stream

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type UserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "serverside.stream.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "serverside.stream.UserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x2a, 0xce,
	0x4c, 0x49, 0xd5, 0x2b, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x55, 0x92, 0xe5, 0xe2, 0x0e, 0x2d, 0x4e,
	0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xf2, 0x74, 0x51, 0x32, 0xe1, 0xe2, 0x81, 0x48, 0x17, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x82, 0x55, 0x70, 0x06,
	0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x4d, 0x20, 0xa6, 0x51,
	0x3c, 0xc4, 0xd0, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xa1, 0x00, 0x2e, 0x6e, 0xf7, 0xd4,
	0x12, 0x90, 0x88, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x9c, 0x1e, 0x86, 0x33, 0xf4, 0x90, 0xdc, 0x20,
	0x25, 0x8f, 0x53, 0x1e, 0xe2, 0x08, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0x7f, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfb, 0xeb, 0x26, 0x10, 0xdd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	//流式响应
	GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (UserService_GetUserInfoClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (UserService_GetUserInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/serverside.stream.UserService/GetUserInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUserInfoClient interface {
	Recv() (*UserResponse, error)
	grpc.ClientStream
}

type userServiceGetUserInfoClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserInfoClient) Recv() (*UserResponse, error) {
	m := new(UserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//流式响应
	GetUserInfo(*UserRequest, UserService_GetUserInfoServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserInfo(req *UserRequest, srv UserService_GetUserInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserInfo(m, &userServiceGetUserInfoServer{stream})
}

type UserService_GetUserInfoServer interface {
	Send(*UserResponse) error
	grpc.ServerStream
}

type userServiceGetUserInfoServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserInfoServer) Send(m *UserResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverside.stream.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserInfo",
			Handler:       _UserService_GetUserInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}
