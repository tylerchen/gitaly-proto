// Code generated by protoc-gen-go.
// source: diff.proto
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

type CommitDiffRequest struct {
	Repository             *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	LeftCommitId           string      `protobuf:"bytes,2,opt,name=left_commit_id,json=leftCommitId" json:"left_commit_id,omitempty"`
	RightCommitId          string      `protobuf:"bytes,3,opt,name=right_commit_id,json=rightCommitId" json:"right_commit_id,omitempty"`
	IgnoreWhitespaceChange bool        `protobuf:"varint,4,opt,name=ignore_whitespace_change,json=ignoreWhitespaceChange" json:"ignore_whitespace_change,omitempty"`
	Paths                  [][]byte    `protobuf:"bytes,5,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (m *CommitDiffRequest) Reset()                    { *m = CommitDiffRequest{} }
func (m *CommitDiffRequest) String() string            { return proto.CompactTextString(m) }
func (*CommitDiffRequest) ProtoMessage()               {}
func (*CommitDiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CommitDiffRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *CommitDiffRequest) GetLeftCommitId() string {
	if m != nil {
		return m.LeftCommitId
	}
	return ""
}

func (m *CommitDiffRequest) GetRightCommitId() string {
	if m != nil {
		return m.RightCommitId
	}
	return ""
}

func (m *CommitDiffRequest) GetIgnoreWhitespaceChange() bool {
	if m != nil {
		return m.IgnoreWhitespaceChange
	}
	return false
}

func (m *CommitDiffRequest) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

// A CommitDiffResponse corresponds to a single changed file in a commit.
type CommitDiffResponse struct {
	FromPath []byte `protobuf:"bytes,1,opt,name=from_path,json=fromPath,proto3" json:"from_path,omitempty"`
	ToPath   []byte `protobuf:"bytes,2,opt,name=to_path,json=toPath,proto3" json:"to_path,omitempty"`
	// Blob ID as returned via `git diff --full-index`
	FromId    string   `protobuf:"bytes,3,opt,name=from_id,json=fromId" json:"from_id,omitempty"`
	ToId      string   `protobuf:"bytes,4,opt,name=to_id,json=toId" json:"to_id,omitempty"`
	OldMode   int32    `protobuf:"varint,5,opt,name=old_mode,json=oldMode" json:"old_mode,omitempty"`
	NewMode   int32    `protobuf:"varint,6,opt,name=new_mode,json=newMode" json:"new_mode,omitempty"`
	Binary    bool     `protobuf:"varint,7,opt,name=binary" json:"binary,omitempty"`
	RawChunks [][]byte `protobuf:"bytes,8,rep,name=raw_chunks,json=rawChunks,proto3" json:"raw_chunks,omitempty"`
}

func (m *CommitDiffResponse) Reset()                    { *m = CommitDiffResponse{} }
func (m *CommitDiffResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitDiffResponse) ProtoMessage()               {}
func (*CommitDiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CommitDiffResponse) GetFromPath() []byte {
	if m != nil {
		return m.FromPath
	}
	return nil
}

func (m *CommitDiffResponse) GetToPath() []byte {
	if m != nil {
		return m.ToPath
	}
	return nil
}

func (m *CommitDiffResponse) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *CommitDiffResponse) GetToId() string {
	if m != nil {
		return m.ToId
	}
	return ""
}

func (m *CommitDiffResponse) GetOldMode() int32 {
	if m != nil {
		return m.OldMode
	}
	return 0
}

func (m *CommitDiffResponse) GetNewMode() int32 {
	if m != nil {
		return m.NewMode
	}
	return 0
}

func (m *CommitDiffResponse) GetBinary() bool {
	if m != nil {
		return m.Binary
	}
	return false
}

func (m *CommitDiffResponse) GetRawChunks() [][]byte {
	if m != nil {
		return m.RawChunks
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitDiffRequest)(nil), "gitaly.CommitDiffRequest")
	proto.RegisterType((*CommitDiffResponse)(nil), "gitaly.CommitDiffResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Diff service

type DiffClient interface {
	// Returns stream of CommitDiffResponse: 1 per changed file
	CommitDiff(ctx context.Context, in *CommitDiffRequest, opts ...grpc.CallOption) (Diff_CommitDiffClient, error)
}

type diffClient struct {
	cc *grpc.ClientConn
}

func NewDiffClient(cc *grpc.ClientConn) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) CommitDiff(ctx context.Context, in *CommitDiffRequest, opts ...grpc.CallOption) (Diff_CommitDiffClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Diff_serviceDesc.Streams[0], c.cc, "/gitaly.Diff/CommitDiff", opts...)
	if err != nil {
		return nil, err
	}
	x := &diffCommitDiffClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Diff_CommitDiffClient interface {
	Recv() (*CommitDiffResponse, error)
	grpc.ClientStream
}

type diffCommitDiffClient struct {
	grpc.ClientStream
}

func (x *diffCommitDiffClient) Recv() (*CommitDiffResponse, error) {
	m := new(CommitDiffResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Diff service

type DiffServer interface {
	// Returns stream of CommitDiffResponse: 1 per changed file
	CommitDiff(*CommitDiffRequest, Diff_CommitDiffServer) error
}

func RegisterDiffServer(s *grpc.Server, srv DiffServer) {
	s.RegisterService(&_Diff_serviceDesc, srv)
}

func _Diff_CommitDiff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommitDiffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiffServer).CommitDiff(m, &diffCommitDiffServer{stream})
}

type Diff_CommitDiffServer interface {
	Send(*CommitDiffResponse) error
	grpc.ServerStream
}

type diffCommitDiffServer struct {
	grpc.ServerStream
}

func (x *diffCommitDiffServer) Send(m *CommitDiffResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Diff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CommitDiff",
			Handler:       _Diff_CommitDiff_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "diff.proto",
}

func init() { proto.RegisterFile("diff.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x8b, 0xd3, 0x40,
	0x14, 0x85, 0xc9, 0xb6, 0x49, 0xd3, 0x6b, 0x57, 0x71, 0x94, 0xdd, 0xd9, 0x2e, 0x62, 0x58, 0x44,
	0xf2, 0x54, 0xa4, 0xbe, 0xf8, 0x2a, 0x5d, 0x56, 0x02, 0x16, 0x25, 0x2c, 0x88, 0x4f, 0xc3, 0x6c,
	0x67, 0xd2, 0x0c, 0xb6, 0x73, 0xe3, 0xcc, 0x2c, 0xa1, 0xaf, 0xfe, 0x27, 0x7f, 0x92, 0xff, 0x43,
	0x66, 0xc6, 0x8d, 0x05, 0x7d, 0xcb, 0x39, 0xdf, 0xc9, 0x4d, 0xce, 0xdc, 0x01, 0x10, 0xaa, 0x69,
	0x16, 0x9d, 0x41, 0x87, 0x24, 0xdb, 0x2a, 0xc7, 0x77, 0x87, 0xf9, 0xcc, 0xb6, 0xdc, 0x48, 0x11,
	0xdd, 0xf9, 0x29, 0x76, 0x4e, 0xa1, 0xb6, 0x51, 0x5e, 0xfd, 0x4a, 0xe0, 0xe9, 0x0a, 0xf7, 0x7b,
	0xe5, 0xae, 0x55, 0xd3, 0xd4, 0xf2, 0xfb, 0xbd, 0xb4, 0x8e, 0x2c, 0x01, 0x8c, 0xec, 0xd0, 0x2a,
	0x87, 0xe6, 0x40, 0x93, 0x22, 0x29, 0x1f, 0x2d, 0xc9, 0x22, 0xce, 0x5b, 0xd4, 0x03, 0xa9, 0x8f,
	0x52, 0xe4, 0x15, 0x3c, 0xde, 0xc9, 0xc6, 0xb1, 0x4d, 0x98, 0xc6, 0x94, 0xa0, 0x27, 0x45, 0x52,
	0x4e, 0xeb, 0x99, 0x77, 0xe3, 0x27, 0x2a, 0x41, 0x5e, 0xc3, 0x13, 0xa3, 0xb6, 0xed, 0x71, 0x6c,
	0x14, 0x62, 0xa7, 0xc1, 0x1e, 0x72, 0xef, 0x80, 0xaa, 0xad, 0x46, 0x23, 0x59, 0xdf, 0x2a, 0x27,
	0x6d, 0xc7, 0x37, 0x92, 0x6d, 0x5a, 0xae, 0xb7, 0x92, 0x8e, 0x8b, 0xa4, 0xcc, 0xeb, 0xb3, 0xc8,
	0xbf, 0x0c, 0x78, 0x15, 0x28, 0x79, 0x0e, 0x69, 0xc7, 0x5d, 0x6b, 0x69, 0x5a, 0x8c, 0xca, 0x59,
	0x1d, 0x85, 0xef, 0x49, 0x8e, 0x7b, 0xda, 0x0e, 0xb5, 0x95, 0xe4, 0x12, 0xa6, 0x8d, 0xc1, 0x3d,
	0xf3, 0xa1, 0xd0, 0x73, 0x56, 0xe7, 0xde, 0xf8, 0xcc, 0x5d, 0x4b, 0xce, 0x61, 0xe2, 0x30, 0xa2,
	0x93, 0x80, 0x32, 0x87, 0x0f, 0x20, 0xbc, 0x35, 0xfc, 0x7c, 0xe6, 0x65, 0x25, 0xc8, 0x33, 0x48,
	0x1d, 0x7a, 0x7b, 0x1c, 0xec, 0xb1, 0xc3, 0x4a, 0x90, 0x0b, 0xc8, 0x71, 0x27, 0xd8, 0x1e, 0x85,
	0xa4, 0x69, 0x91, 0x94, 0x69, 0x3d, 0xc1, 0x9d, 0x58, 0xa3, 0x90, 0x1e, 0x69, 0xd9, 0x47, 0x94,
	0x45, 0xa4, 0x65, 0x1f, 0xd0, 0x19, 0x64, 0x77, 0x4a, 0x73, 0x73, 0xa0, 0x93, 0x50, 0xf7, 0x8f,
	0x22, 0x2f, 0x00, 0x0c, 0xef, 0xd9, 0xa6, 0xbd, 0xd7, 0xdf, 0x2c, 0xcd, 0x43, 0xc7, 0xa9, 0xe1,
	0xfd, 0x2a, 0x18, 0x4b, 0x0d, 0x63, 0x5f, 0x90, 0x34, 0x00, 0x7f, 0xeb, 0x92, 0x8b, 0x87, 0xdd,
	0xfd, 0xb3, 0xea, 0xf9, 0xfc, 0x7f, 0x28, 0x9e, 0xce, 0xd5, 0xcb, 0x1f, 0x3f, 0xe9, 0x65, 0x3e,
	0x22, 0xe7, 0x1f, 0xaa, 0xdb, 0xf7, 0x1f, 0xbf, 0xb2, 0xeb, 0xea, 0xe6, 0x86, 0xad, 0x3e, 0xad,
	0xd7, 0xd5, 0x6d, 0x78, 0x7e, 0x93, 0xdc, 0x65, 0xe1, 0x1a, 0xbd, 0xfd, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0xc8, 0x0a, 0xe5, 0xaf, 0x79, 0x02, 0x00, 0x00,
}
