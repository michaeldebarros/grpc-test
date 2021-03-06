// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
//const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type NewUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserRequest) Reset()         { *m = NewUserRequest{} }
func (m *NewUserRequest) String() string { return proto.CompactTextString(m) }
func (*NewUserRequest) ProtoMessage()    {}
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *NewUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserRequest.Unmarshal(m, b)
}
func (m *NewUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserRequest.Marshal(b, m, deterministic)
}
func (m *NewUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserRequest.Merge(m, src)
}
func (m *NewUserRequest) XXX_Size() int {
	return xxx_messageInfo_NewUserRequest.Size(m)
}
func (m *NewUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserRequest proto.InternalMessageInfo

func (m *NewUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type NewUserResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserResponse) Reset()         { *m = NewUserResponse{} }
func (m *NewUserResponse) String() string { return proto.CompactTextString(m) }
func (*NewUserResponse) ProtoMessage()    {}
func (*NewUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *NewUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserResponse.Unmarshal(m, b)
}
func (m *NewUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserResponse.Marshal(b, m, deterministic)
}
func (m *NewUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserResponse.Merge(m, src)
}
func (m *NewUserResponse) XXX_Size() int {
	return xxx_messageInfo_NewUserResponse.Size(m)
}
func (m *NewUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserResponse proto.InternalMessageInfo

func (m *NewUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*NewUserRequest)(nil), "user.NewUserRequest")
	proto.RegisterType((*NewUserResponse)(nil), "user.NewUserResponse")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2d, 0x4e, 0x2d,
	0xd2, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x01, 0x17,
	0x4b, 0x68, 0x71, 0x6a, 0x91, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x9a, 0x9b, 0x98, 0x99, 0x23, 0xc1,
	0x04, 0x16, 0x84, 0x70, 0x94, 0x0c, 0xb8, 0xf8, 0xfc, 0x52, 0xcb, 0x41, 0x9a, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xe4, 0xb8, 0xc0, 0x66, 0x81, 0xf5, 0x72, 0x1b, 0x71, 0xe9, 0x81,
	0x2d, 0x01, 0x2b, 0x80, 0xd8, 0xa1, 0xcd, 0xc5, 0x0f, 0xd7, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x2a, 0x24, 0xc1, 0xc5, 0x5e, 0x5c, 0x9a, 0x9c, 0x9c, 0x5a, 0x5c, 0x0c, 0xd6, 0xc5, 0x11, 0x04,
	0xe3, 0x1a, 0xb9, 0x73, 0x71, 0x83, 0x54, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59,
	0x70, 0xb1, 0x43, 0xf5, 0x0a, 0x89, 0x40, 0x0c, 0x46, 0xb5, 0x5c, 0x4a, 0x14, 0x4d, 0x14, 0x62,
	0x81, 0x12, 0x83, 0x13, 0x5b, 0x14, 0xd8, 0xf6, 0x24, 0x36, 0xb0, 0x77, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x56, 0x52, 0x6a, 0xa3, 0x01, 0x01, 0x00, 0x00,
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
	// Unary
	NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error) {
	out := new(NewUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/NewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Unary
	NewUser(context.Context, *NewUserRequest) (*NewUserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).NewUser(ctx, req.(*NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUser",
			Handler:    _UserService_NewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
