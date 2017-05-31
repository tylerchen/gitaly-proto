// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smarthttp.proto

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

type InfoRefsRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *InfoRefsRequest) Reset()                    { *m = InfoRefsRequest{} }
func (m *InfoRefsRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRefsRequest) ProtoMessage()               {}
func (*InfoRefsRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *InfoRefsRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type InfoRefsResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *InfoRefsResponse) Reset()                    { *m = InfoRefsResponse{} }
func (m *InfoRefsResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoRefsResponse) ProtoMessage()               {}
func (*InfoRefsResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *InfoRefsResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PostUploadPackRequest struct {
	// repository should only be present in the first message of the stream
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Raw data to be copied to stdin of 'git upload-pack'
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PostUploadPackRequest) Reset()                    { *m = PostUploadPackRequest{} }
func (m *PostUploadPackRequest) String() string            { return proto.CompactTextString(m) }
func (*PostUploadPackRequest) ProtoMessage()               {}
func (*PostUploadPackRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *PostUploadPackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PostUploadPackRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PostUploadPackResponse struct {
	// Raw data from stdout of 'git upload-pack'
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PostUploadPackResponse) Reset()                    { *m = PostUploadPackResponse{} }
func (m *PostUploadPackResponse) String() string            { return proto.CompactTextString(m) }
func (*PostUploadPackResponse) ProtoMessage()               {}
func (*PostUploadPackResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *PostUploadPackResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PostReceivePackRequest struct {
	// repository should only be present in the first message of the stream
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Raw data to be copied to stdin of 'git receive-pack'
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// gl_id and gl_repository becomes env variables, used by the Git {pre,post}-receive
	// hooks. They should only be present in the first message of the stream.
	GlId         string `protobuf:"bytes,3,opt,name=gl_id,json=glId" json:"gl_id,omitempty"`
	GlRepository string `protobuf:"bytes,4,opt,name=gl_repository,json=glRepository" json:"gl_repository,omitempty"`
}

func (m *PostReceivePackRequest) Reset()                    { *m = PostReceivePackRequest{} }
func (m *PostReceivePackRequest) String() string            { return proto.CompactTextString(m) }
func (*PostReceivePackRequest) ProtoMessage()               {}
func (*PostReceivePackRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *PostReceivePackRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PostReceivePackRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PostReceivePackRequest) GetGlId() string {
	if m != nil {
		return m.GlId
	}
	return ""
}

func (m *PostReceivePackRequest) GetGlRepository() string {
	if m != nil {
		return m.GlRepository
	}
	return ""
}

type PostReceivePackResponse struct {
	// Raw data from stdout of 'git receive-pack'
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PostReceivePackResponse) Reset()                    { *m = PostReceivePackResponse{} }
func (m *PostReceivePackResponse) String() string            { return proto.CompactTextString(m) }
func (*PostReceivePackResponse) ProtoMessage()               {}
func (*PostReceivePackResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *PostReceivePackResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InfoRefsRequest)(nil), "gitaly.InfoRefsRequest")
	proto.RegisterType((*InfoRefsResponse)(nil), "gitaly.InfoRefsResponse")
	proto.RegisterType((*PostUploadPackRequest)(nil), "gitaly.PostUploadPackRequest")
	proto.RegisterType((*PostUploadPackResponse)(nil), "gitaly.PostUploadPackResponse")
	proto.RegisterType((*PostReceivePackRequest)(nil), "gitaly.PostReceivePackRequest")
	proto.RegisterType((*PostReceivePackResponse)(nil), "gitaly.PostReceivePackResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SmartHTTP service

type SmartHTTPClient interface {
	// The response body for GET /info/refs?service=git-upload-pack
	InfoRefsUploadPack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsUploadPackClient, error)
	// The response body for GET /info/refs?service=git-receive-pack
	InfoRefsReceivePack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsReceivePackClient, error)
	// Request and response body for POST /upload-pack
	PostUploadPack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostUploadPackClient, error)
	// Request and response body for POST /receive-pack
	PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostReceivePackClient, error)
}

type smartHTTPClient struct {
	cc *grpc.ClientConn
}

func NewSmartHTTPClient(cc *grpc.ClientConn) SmartHTTPClient {
	return &smartHTTPClient{cc}
}

func (c *smartHTTPClient) InfoRefsUploadPack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[0], c.cc, "/gitaly.SmartHTTP/InfoRefsUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPInfoRefsUploadPackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartHTTP_InfoRefsUploadPackClient interface {
	Recv() (*InfoRefsResponse, error)
	grpc.ClientStream
}

type smartHTTPInfoRefsUploadPackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPInfoRefsUploadPackClient) Recv() (*InfoRefsResponse, error) {
	m := new(InfoRefsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartHTTPClient) InfoRefsReceivePack(ctx context.Context, in *InfoRefsRequest, opts ...grpc.CallOption) (SmartHTTP_InfoRefsReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[1], c.cc, "/gitaly.SmartHTTP/InfoRefsReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPInfoRefsReceivePackClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartHTTP_InfoRefsReceivePackClient interface {
	Recv() (*InfoRefsResponse, error)
	grpc.ClientStream
}

type smartHTTPInfoRefsReceivePackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPInfoRefsReceivePackClient) Recv() (*InfoRefsResponse, error) {
	m := new(InfoRefsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartHTTPClient) PostUploadPack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostUploadPackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[2], c.cc, "/gitaly.SmartHTTP/PostUploadPack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPPostUploadPackClient{stream}
	return x, nil
}

type SmartHTTP_PostUploadPackClient interface {
	Send(*PostUploadPackRequest) error
	Recv() (*PostUploadPackResponse, error)
	grpc.ClientStream
}

type smartHTTPPostUploadPackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPPostUploadPackClient) Send(m *PostUploadPackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *smartHTTPPostUploadPackClient) Recv() (*PostUploadPackResponse, error) {
	m := new(PostUploadPackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *smartHTTPClient) PostReceivePack(ctx context.Context, opts ...grpc.CallOption) (SmartHTTP_PostReceivePackClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SmartHTTP_serviceDesc.Streams[3], c.cc, "/gitaly.SmartHTTP/PostReceivePack", opts...)
	if err != nil {
		return nil, err
	}
	x := &smartHTTPPostReceivePackClient{stream}
	return x, nil
}

type SmartHTTP_PostReceivePackClient interface {
	Send(*PostReceivePackRequest) error
	Recv() (*PostReceivePackResponse, error)
	grpc.ClientStream
}

type smartHTTPPostReceivePackClient struct {
	grpc.ClientStream
}

func (x *smartHTTPPostReceivePackClient) Send(m *PostReceivePackRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *smartHTTPPostReceivePackClient) Recv() (*PostReceivePackResponse, error) {
	m := new(PostReceivePackResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SmartHTTP service

type SmartHTTPServer interface {
	// The response body for GET /info/refs?service=git-upload-pack
	InfoRefsUploadPack(*InfoRefsRequest, SmartHTTP_InfoRefsUploadPackServer) error
	// The response body for GET /info/refs?service=git-receive-pack
	InfoRefsReceivePack(*InfoRefsRequest, SmartHTTP_InfoRefsReceivePackServer) error
	// Request and response body for POST /upload-pack
	PostUploadPack(SmartHTTP_PostUploadPackServer) error
	// Request and response body for POST /receive-pack
	PostReceivePack(SmartHTTP_PostReceivePackServer) error
}

func RegisterSmartHTTPServer(s *grpc.Server, srv SmartHTTPServer) {
	s.RegisterService(&_SmartHTTP_serviceDesc, srv)
}

func _SmartHTTP_InfoRefsUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InfoRefsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartHTTPServer).InfoRefsUploadPack(m, &smartHTTPInfoRefsUploadPackServer{stream})
}

type SmartHTTP_InfoRefsUploadPackServer interface {
	Send(*InfoRefsResponse) error
	grpc.ServerStream
}

type smartHTTPInfoRefsUploadPackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPInfoRefsUploadPackServer) Send(m *InfoRefsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SmartHTTP_InfoRefsReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InfoRefsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartHTTPServer).InfoRefsReceivePack(m, &smartHTTPInfoRefsReceivePackServer{stream})
}

type SmartHTTP_InfoRefsReceivePackServer interface {
	Send(*InfoRefsResponse) error
	grpc.ServerStream
}

type smartHTTPInfoRefsReceivePackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPInfoRefsReceivePackServer) Send(m *InfoRefsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SmartHTTP_PostUploadPack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SmartHTTPServer).PostUploadPack(&smartHTTPPostUploadPackServer{stream})
}

type SmartHTTP_PostUploadPackServer interface {
	Send(*PostUploadPackResponse) error
	Recv() (*PostUploadPackRequest, error)
	grpc.ServerStream
}

type smartHTTPPostUploadPackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPPostUploadPackServer) Send(m *PostUploadPackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *smartHTTPPostUploadPackServer) Recv() (*PostUploadPackRequest, error) {
	m := new(PostUploadPackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SmartHTTP_PostReceivePack_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SmartHTTPServer).PostReceivePack(&smartHTTPPostReceivePackServer{stream})
}

type SmartHTTP_PostReceivePackServer interface {
	Send(*PostReceivePackResponse) error
	Recv() (*PostReceivePackRequest, error)
	grpc.ServerStream
}

type smartHTTPPostReceivePackServer struct {
	grpc.ServerStream
}

func (x *smartHTTPPostReceivePackServer) Send(m *PostReceivePackResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *smartHTTPPostReceivePackServer) Recv() (*PostReceivePackRequest, error) {
	m := new(PostReceivePackRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SmartHTTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.SmartHTTP",
	HandlerType: (*SmartHTTPServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InfoRefsUploadPack",
			Handler:       _SmartHTTP_InfoRefsUploadPack_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "InfoRefsReceivePack",
			Handler:       _SmartHTTP_InfoRefsReceivePack_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostUploadPack",
			Handler:       _SmartHTTP_PostUploadPack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PostReceivePack",
			Handler:       _SmartHTTP_PostReceivePack_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "smarthttp.proto",
}

func init() { proto.RegisterFile("smarthttp.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0x6b, 0x2d, 0x74, 0xac, 0x56, 0xa6, 0x68, 0x4b, 0x40, 0x2d, 0x11, 0xa4, 0x07, 0x0d,
	0x25, 0xfe, 0x06, 0xc1, 0xa2, 0x87, 0xb0, 0xb6, 0xe0, 0x2d, 0xac, 0xcd, 0x36, 0x0d, 0xae, 0xdd,
	0x98, 0x5d, 0x85, 0xfe, 0x15, 0xff, 0x98, 0x7f, 0x47, 0x4c, 0xc8, 0x47, 0x13, 0xe3, 0x41, 0xf1,
	0x16, 0xe6, 0xcd, 0xbc, 0xf7, 0x66, 0x5e, 0x16, 0xba, 0xea, 0x99, 0x45, 0x7a, 0xa9, 0x75, 0x68,
	0x85, 0x91, 0xd4, 0x12, 0x5b, 0x7e, 0xa0, 0x99, 0x58, 0x1b, 0x1d, 0xb5, 0x64, 0x11, 0xf7, 0x92,
	0xaa, 0x79, 0x0d, 0xdd, 0xc9, 0x6a, 0x21, 0x29, 0x5f, 0x28, 0xca, 0x5f, 0x5e, 0xb9, 0xd2, 0x68,
	0x03, 0x44, 0x3c, 0x94, 0x2a, 0xd0, 0x32, 0x5a, 0x0f, 0xc8, 0x90, 0x8c, 0x76, 0x6d, 0xb4, 0x92,
	0x69, 0x8b, 0x66, 0x08, 0x2d, 0x74, 0x99, 0xe7, 0x70, 0x90, 0xd3, 0xa8, 0x50, 0xae, 0x14, 0x47,
	0x84, 0xa6, 0xc7, 0x34, 0x8b, 0x19, 0x3a, 0x34, 0xfe, 0x36, 0x5d, 0x38, 0x74, 0xa4, 0xd2, 0xb3,
	0x50, 0x48, 0xe6, 0x39, 0x6c, 0xfe, 0xf4, 0x07, 0xd1, 0x4c, 0xa0, 0x51, 0x10, 0xb8, 0x80, 0xa3,
	0xb2, 0xc0, 0x0f, 0x76, 0xde, 0x49, 0xd2, 0x4e, 0xf9, 0x9c, 0x07, 0x6f, 0xfc, 0x1f, 0x0c, 0x61,
	0x0f, 0x76, 0x7c, 0xe1, 0x06, 0xde, 0x60, 0x7b, 0x48, 0x46, 0x6d, 0xda, 0xf4, 0xc5, 0xc4, 0xc3,
	0x33, 0xd8, 0xf3, 0x85, 0x5b, 0xe0, 0x6f, 0xc6, 0x60, 0xc7, 0x17, 0x39, 0xb3, 0x79, 0x09, 0xfd,
	0x8a, 0xb7, 0xfa, 0x5d, 0xec, 0x8f, 0x06, 0xb4, 0xef, 0xbf, 0x32, 0xbf, 0x99, 0x4e, 0x1d, 0xbc,
	0x05, 0x4c, 0x03, 0xc9, 0x6f, 0x81, 0xfd, 0x74, 0x81, 0x52, 0xe6, 0xc6, 0xa0, 0x0a, 0x24, 0x52,
	0xe6, 0xd6, 0x98, 0xe0, 0x1d, 0xf4, 0xf2, 0x7a, 0xe6, 0xe6, 0xb7, 0x6c, 0x33, 0xd8, 0xdf, 0x8c,
	0x08, 0x8f, 0xd3, 0xfe, 0x6f, 0xff, 0x0d, 0xe3, 0xa4, 0x0e, 0x4e, 0x49, 0x47, 0x64, 0x4c, 0xf0,
	0x01, 0xba, 0xa5, 0x73, 0xe1, 0xc6, 0x60, 0x35, 0x63, 0xe3, 0xb4, 0x16, 0x2f, 0x32, 0x3f, 0xb6,
	0xe2, 0xa7, 0x72, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x33, 0x8a, 0x1a, 0x53, 0x03, 0x00,
	0x00,
}
