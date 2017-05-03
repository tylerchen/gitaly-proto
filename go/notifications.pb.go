// Code generated by protoc-gen-go.
// source: notifications.proto
// DO NOT EDIT!

package gitaly

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

type PostReceiveRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *PostReceiveRequest) Reset()                    { *m = PostReceiveRequest{} }
func (m *PostReceiveRequest) String() string            { return proto.CompactTextString(m) }
func (*PostReceiveRequest) ProtoMessage()               {}
func (*PostReceiveRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PostReceiveRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type PostReceiveResponse struct {
}

func (m *PostReceiveResponse) Reset()                    { *m = PostReceiveResponse{} }
func (m *PostReceiveResponse) String() string            { return proto.CompactTextString(m) }
func (*PostReceiveResponse) ProtoMessage()               {}
func (*PostReceiveResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*PostReceiveRequest)(nil), "gitaly.PostReceiveRequest")
	proto.RegisterType((*PostReceiveResponse)(nil), "gitaly.PostReceiveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Notifications service

type NotificationsClient interface {
	PostReceive(ctx context.Context, in *PostReceiveRequest, opts ...grpc.CallOption) (*PostReceiveResponse, error)
}

type notificationsClient struct {
	cc *grpc.ClientConn
}

func NewNotificationsClient(cc *grpc.ClientConn) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) PostReceive(ctx context.Context, in *PostReceiveRequest, opts ...grpc.CallOption) (*PostReceiveResponse, error) {
	out := new(PostReceiveResponse)
	err := grpc.Invoke(ctx, "/gitaly.Notifications/PostReceive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notifications service

type NotificationsServer interface {
	PostReceive(context.Context, *PostReceiveRequest) (*PostReceiveResponse, error)
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_PostReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).PostReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Notifications/PostReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).PostReceive(ctx, req.(*PostReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostReceive",
			Handler:    _Notifications_PostReceive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications.proto",
}

func init() { proto.RegisterFile("notifications.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4b, 0xcf, 0x2c, 0x49, 0xcc, 0xa9, 0x94, 0xe2, 0x29, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0x81,
	0x88, 0x4a, 0xf1, 0xe6, 0x17, 0x20, 0x29, 0x52, 0xf2, 0xe0, 0x12, 0x0a, 0xc8, 0x2f, 0x2e, 0x09,
	0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x32, 0xe2,
	0xe2, 0x2a, 0x4a, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x36, 0x12, 0xd2, 0x83, 0x98, 0xa7, 0x17, 0x04, 0x97, 0x09, 0x42, 0x52, 0xa5, 0x24, 0xca,
	0x25, 0x8c, 0x62, 0x52, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x51, 0x13, 0x23, 0x17, 0xaf, 0x1f,
	0xb2, 0xeb, 0x84, 0x0a, 0xb9, 0xb8, 0x91, 0x14, 0x0a, 0x49, 0xc1, 0xcc, 0xc5, 0x74, 0x87, 0x94,
	0x34, 0x56, 0x39, 0x88, 0xc9, 0x4a, 0x9a, 0x4d, 0x5b, 0x25, 0x54, 0x39, 0x98, 0x85, 0x14, 0xdd,
	0x3d, 0x43, 0x1c, 0x7d, 0x22, 0xe3, 0xfd, 0xfc, 0x43, 0x3c, 0xdd, 0x3c, 0x9d, 0x1d, 0x43, 0x3c,
	0xfd, 0xfd, 0x82, 0xe3, 0x03, 0xfc, 0x83, 0x43, 0xe2, 0x83, 0x5c, 0x9d, 0x5d, 0x3d, 0xc3, 0x5c,
	0x93, 0xd8, 0xc0, 0x9e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x80, 0xbc, 0x85, 0x89, 0x28,
	0x01, 0x00, 0x00,
}
