// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

/*
Package gitaly is a generated protocol buffer package.

It is generated from these files:
	blob.proto
	commit.proto
	deprecated-services.proto
	diff.proto
	notifications.proto
	ref.proto
	repository-service.proto
	shared.proto
	smarthttp.proto
	ssh.proto

It has these top-level messages:
	GetBlobRequest
	GetBlobResponse
	CommitIsAncestorRequest
	CommitIsAncestorResponse
	TreeEntryRequest
	TreeEntryResponse
	CommitsBetweenRequest
	CommitsBetweenResponse
	CommitDiffRequest
	CommitDiffResponse
	CommitDeltaRequest
	CommitDelta
	CommitDeltaResponse
	PostReceiveRequest
	PostReceiveResponse
	FindDefaultBranchNameRequest
	FindDefaultBranchNameResponse
	FindAllBranchNamesRequest
	FindAllBranchNamesResponse
	FindAllTagNamesRequest
	FindAllTagNamesResponse
	FindRefNameRequest
	FindRefNameResponse
	FindLocalBranchesRequest
	FindLocalBranchesResponse
	FindLocalBranchResponse
	FindLocalBranchCommitAuthor
	RepositoryExistsRequest
	RepositoryExistsResponse
	Repository
	GitCommit
	CommitAuthor
	ExitStatus
	InfoRefsRequest
	InfoRefsResponse
	PostUploadPackRequest
	PostUploadPackResponse
	PostReceivePackRequest
	PostReceivePackResponse
	SSHUploadPackRequest
	SSHUploadPackResponse
	SSHReceivePackRequest
	SSHReceivePackResponse
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetBlobRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Oid        string      `protobuf:"bytes,2,opt,name=oid" json:"oid,omitempty"`
}

func (m *GetBlobRequest) Reset()                    { *m = GetBlobRequest{} }
func (m *GetBlobRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlobRequest) ProtoMessage()               {}
func (*GetBlobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetBlobRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetBlobRequest) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

type GetBlobResponse struct {
	Size int64  `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Oid  string `protobuf:"bytes,3,opt,name=oid" json:"oid,omitempty"`
}

func (m *GetBlobResponse) Reset()                    { *m = GetBlobResponse{} }
func (m *GetBlobResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlobResponse) ProtoMessage()               {}
func (*GetBlobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetBlobResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlobResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlobResponse) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBlobRequest)(nil), "gitaly.GetBlobRequest")
	proto.RegisterType((*GetBlobResponse)(nil), "gitaly.GetBlobResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlobService service

type BlobServiceClient interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error)
}

type blobServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlobServiceClient(cc *grpc.ClientConn) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[0], c.cc, "/gitaly.BlobService/GetBlob", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetBlobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetBlobClient interface {
	Recv() (*GetBlobResponse, error)
	grpc.ClientStream
}

type blobServiceGetBlobClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetBlobClient) Recv() (*GetBlobResponse, error) {
	m := new(GetBlobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BlobService service

type BlobServiceServer interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(*GetBlobRequest, BlobService_GetBlobServer) error
}

func RegisterBlobServiceServer(s *grpc.Server, srv BlobServiceServer) {
	s.RegisterService(&_BlobService_serviceDesc, srv)
}

func _BlobService_GetBlob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetBlob(m, &blobServiceGetBlobServer{stream})
}

type BlobService_GetBlobServer interface {
	Send(*GetBlobResponse) error
	grpc.ServerStream
}

type blobServiceGetBlobServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetBlobServer) Send(m *GetBlobResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlob",
			Handler:       _BlobService_GetBlob_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blob.proto",
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xc9, 0x4f,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x2c, 0x49, 0xcc, 0xa9, 0x94, 0xe2,
	0x29, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0x81, 0x88, 0x2a, 0x85, 0x71, 0xf1, 0xb9, 0xa7, 0x96, 0x38,
	0xe5, 0xe4, 0x27, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x19, 0x71, 0x71, 0x15, 0xa5,
	0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09,
	0xe9, 0x41, 0x34, 0xeb, 0x05, 0xc1, 0x65, 0x82, 0x90, 0x54, 0x09, 0x09, 0x70, 0x31, 0xe7, 0x67,
	0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0xde, 0x5c, 0xfc, 0x70, 0x73,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x8a, 0x33, 0xab, 0x52, 0xc1, 0x46,
	0x32, 0x07, 0x81, 0xd9, 0x20, 0xb1, 0x94, 0xc4, 0x92, 0x44, 0xb0, 0x4e, 0x9e, 0x20, 0x30, 0x1b,
	0x66, 0x18, 0x33, 0xdc, 0x30, 0x23, 0x5f, 0x2e, 0x6e, 0x90, 0x49, 0xc1, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x42, 0x76, 0x5c, 0xec, 0x50, 0xb3, 0x85, 0xc4, 0x60, 0x0e, 0x43, 0xf5, 0x84, 0x94,
	0x38, 0x86, 0x38, 0xc4, 0x11, 0x4a, 0x0c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0xaf, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0x00, 0x89, 0x74, 0x1e, 0x01, 0x00, 0x00,
}
