// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

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

type UserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (m *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(m, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

func (m *UserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResp struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Title                []string `protobuf:"bytes,4,rep,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResp) Reset()         { *m = UserResp{} }
func (m *UserResp) String() string { return proto.CompactTextString(m) }
func (*UserResp) ProtoMessage()    {}
func (*UserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResp.Unmarshal(m, b)
}
func (m *UserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResp.Marshal(b, m, deterministic)
}
func (m *UserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResp.Merge(m, src)
}
func (m *UserResp) XXX_Size() int {
	return xxx_messageInfo_UserResp.Size(m)
}
func (m *UserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserResp proto.InternalMessageInfo

func (m *UserResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResp) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserResp) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

func init() {
	proto.RegisterType((*UserReq)(nil), "proto.UserReq")
	proto.RegisterType((*UserResp)(nil), "proto.UserResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xb2, 0x5c, 0xec, 0xa1, 0xc5,
	0xa9, 0x45, 0x41, 0xa9, 0x85, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x52, 0x18, 0x17, 0x07, 0x44, 0xba, 0xb8, 0x40, 0x88, 0x8f,
	0x8b, 0x29, 0x33, 0x05, 0x2c, 0xcb, 0x1a, 0xc4, 0x94, 0x99, 0x02, 0x57, 0xcf, 0x84, 0x50, 0x2f,
	0x24, 0xc0, 0xc5, 0x9c, 0x98, 0x9e, 0x2a, 0xc1, 0x0c, 0x56, 0x04, 0x62, 0x0a, 0x89, 0x70, 0xb1,
	0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0xb0, 0x28, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x38, 0x46, 0xce,
	0x5c, 0xfc, 0x20, 0x73, 0x3d, 0xf3, 0xd2, 0xf2, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85,
	0x0c, 0xb8, 0xb8, 0xdd, 0x53, 0x4b, 0x60, 0xa2, 0x42, 0x7c, 0x10, 0x77, 0xea, 0x41, 0x5d, 0x27,
	0xc5, 0x8f, 0xc2, 0x2f, 0x2e, 0x50, 0x62, 0x48, 0x62, 0x03, 0x8b, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xa6, 0xaf, 0xaa, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInfoServiceClient is the client API for UserInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInfoServiceClient interface {
	GetUserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error)
}

type userInfoServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoServiceClient(cc *grpc.ClientConn) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) GetUserInfo(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error) {
	out := new(UserResp)
	err := c.cc.Invoke(ctx, "/proto.UserInfoService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServiceServer is the server API for UserInfoService service.
type UserInfoServiceServer interface {
	GetUserInfo(context.Context, *UserReq) (*UserResp, error)
}

// UnimplementedUserInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserInfoServiceServer struct {
}

func (*UnimplementedUserInfoServiceServer) GetUserInfo(ctx context.Context, req *UserReq) (*UserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterUserInfoServiceServer(s *grpc.Server, srv UserInfoServiceServer) {
	s.RegisterService(&_UserInfoService_serviceDesc, srv)
}

func _UserInfoService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserInfoService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInfoService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
