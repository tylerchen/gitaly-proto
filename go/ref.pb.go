// Code generated by protoc-gen-go.
// source: ref.proto
// DO NOT EDIT!

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FindLocalBranchesRequest_SortBy int32

const (
	FindLocalBranchesRequest_NAME         FindLocalBranchesRequest_SortBy = 0
	FindLocalBranchesRequest_UPDATED_ASC  FindLocalBranchesRequest_SortBy = 1
	FindLocalBranchesRequest_UPDATED_DESC FindLocalBranchesRequest_SortBy = 2
)

var FindLocalBranchesRequest_SortBy_name = map[int32]string{
	0: "NAME",
	1: "UPDATED_ASC",
	2: "UPDATED_DESC",
}
var FindLocalBranchesRequest_SortBy_value = map[string]int32{
	"NAME":         0,
	"UPDATED_ASC":  1,
	"UPDATED_DESC": 2,
}

func (x FindLocalBranchesRequest_SortBy) String() string {
	return proto.EnumName(FindLocalBranchesRequest_SortBy_name, int32(x))
}
func (FindLocalBranchesRequest_SortBy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{8, 0}
}

type FindDefaultBranchNameRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *FindDefaultBranchNameRequest) Reset()                    { *m = FindDefaultBranchNameRequest{} }
func (m *FindDefaultBranchNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindDefaultBranchNameRequest) ProtoMessage()               {}
func (*FindDefaultBranchNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *FindDefaultBranchNameRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type FindDefaultBranchNameResponse struct {
	Name []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *FindDefaultBranchNameResponse) Reset()                    { *m = FindDefaultBranchNameResponse{} }
func (m *FindDefaultBranchNameResponse) String() string            { return proto.CompactTextString(m) }
func (*FindDefaultBranchNameResponse) ProtoMessage()               {}
func (*FindDefaultBranchNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FindDefaultBranchNameResponse) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type FindAllBranchNamesRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *FindAllBranchNamesRequest) Reset()                    { *m = FindAllBranchNamesRequest{} }
func (m *FindAllBranchNamesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAllBranchNamesRequest) ProtoMessage()               {}
func (*FindAllBranchNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *FindAllBranchNamesRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type FindAllBranchNamesResponse struct {
	Names [][]byte `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (m *FindAllBranchNamesResponse) Reset()                    { *m = FindAllBranchNamesResponse{} }
func (m *FindAllBranchNamesResponse) String() string            { return proto.CompactTextString(m) }
func (*FindAllBranchNamesResponse) ProtoMessage()               {}
func (*FindAllBranchNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *FindAllBranchNamesResponse) GetNames() [][]byte {
	if m != nil {
		return m.Names
	}
	return nil
}

type FindAllTagNamesRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *FindAllTagNamesRequest) Reset()                    { *m = FindAllTagNamesRequest{} }
func (m *FindAllTagNamesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAllTagNamesRequest) ProtoMessage()               {}
func (*FindAllTagNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *FindAllTagNamesRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type FindAllTagNamesResponse struct {
	Names [][]byte `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (m *FindAllTagNamesResponse) Reset()                    { *m = FindAllTagNamesResponse{} }
func (m *FindAllTagNamesResponse) String() string            { return proto.CompactTextString(m) }
func (*FindAllTagNamesResponse) ProtoMessage()               {}
func (*FindAllTagNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *FindAllTagNamesResponse) GetNames() [][]byte {
	if m != nil {
		return m.Names
	}
	return nil
}

type FindRefNameRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Require that the resulting ref contains this commit as an ancestor
	CommitId string `protobuf:"bytes,2,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	// Example prefix: "refs/heads/". Type bytes because that is the type of ref names.
	Prefix []byte `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (m *FindRefNameRequest) Reset()                    { *m = FindRefNameRequest{} }
func (m *FindRefNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRefNameRequest) ProtoMessage()               {}
func (*FindRefNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *FindRefNameRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FindRefNameRequest) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

func (m *FindRefNameRequest) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

type FindRefNameResponse struct {
	// Example name: "refs/heads/master". Cannot assume UTF8, so the type is bytes.
	Name []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *FindRefNameResponse) Reset()                    { *m = FindRefNameResponse{} }
func (m *FindRefNameResponse) String() string            { return proto.CompactTextString(m) }
func (*FindRefNameResponse) ProtoMessage()               {}
func (*FindRefNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *FindRefNameResponse) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type FindLocalBranchesRequest struct {
	Repository *Repository                     `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	SortBy     FindLocalBranchesRequest_SortBy `protobuf:"varint,2,opt,name=sort_by,json=sortBy,enum=gitaly.FindLocalBranchesRequest_SortBy" json:"sort_by,omitempty"`
}

func (m *FindLocalBranchesRequest) Reset()                    { *m = FindLocalBranchesRequest{} }
func (m *FindLocalBranchesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindLocalBranchesRequest) ProtoMessage()               {}
func (*FindLocalBranchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *FindLocalBranchesRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FindLocalBranchesRequest) GetSortBy() FindLocalBranchesRequest_SortBy {
	if m != nil {
		return m.SortBy
	}
	return FindLocalBranchesRequest_NAME
}

type FindLocalBranchesResponse struct {
	Branches []*FindLocalBranchResponse `protobuf:"bytes,1,rep,name=branches" json:"branches,omitempty"`
}

func (m *FindLocalBranchesResponse) Reset()                    { *m = FindLocalBranchesResponse{} }
func (m *FindLocalBranchesResponse) String() string            { return proto.CompactTextString(m) }
func (*FindLocalBranchesResponse) ProtoMessage()               {}
func (*FindLocalBranchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *FindLocalBranchesResponse) GetBranches() []*FindLocalBranchResponse {
	if m != nil {
		return m.Branches
	}
	return nil
}

type FindLocalBranchResponse struct {
	Name            []byte                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CommitId        string                       `protobuf:"bytes,2,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	CommitSubject   []byte                       `protobuf:"bytes,3,opt,name=commit_subject,json=commitSubject,proto3" json:"commit_subject,omitempty"`
	CommitAuthor    *FindLocalBranchCommitAuthor `protobuf:"bytes,4,opt,name=commit_author,json=commitAuthor" json:"commit_author,omitempty"`
	CommitCommitter *FindLocalBranchCommitAuthor `protobuf:"bytes,5,opt,name=commit_committer,json=commitCommitter" json:"commit_committer,omitempty"`
}

func (m *FindLocalBranchResponse) Reset()                    { *m = FindLocalBranchResponse{} }
func (m *FindLocalBranchResponse) String() string            { return proto.CompactTextString(m) }
func (*FindLocalBranchResponse) ProtoMessage()               {}
func (*FindLocalBranchResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *FindLocalBranchResponse) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *FindLocalBranchResponse) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

func (m *FindLocalBranchResponse) GetCommitSubject() []byte {
	if m != nil {
		return m.CommitSubject
	}
	return nil
}

func (m *FindLocalBranchResponse) GetCommitAuthor() *FindLocalBranchCommitAuthor {
	if m != nil {
		return m.CommitAuthor
	}
	return nil
}

func (m *FindLocalBranchResponse) GetCommitCommitter() *FindLocalBranchCommitAuthor {
	if m != nil {
		return m.CommitCommitter
	}
	return nil
}

type FindLocalBranchCommitAuthor struct {
	Name  []byte                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email []byte                      `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date  *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
}

func (m *FindLocalBranchCommitAuthor) Reset()                    { *m = FindLocalBranchCommitAuthor{} }
func (m *FindLocalBranchCommitAuthor) String() string            { return proto.CompactTextString(m) }
func (*FindLocalBranchCommitAuthor) ProtoMessage()               {}
func (*FindLocalBranchCommitAuthor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *FindLocalBranchCommitAuthor) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *FindLocalBranchCommitAuthor) GetEmail() []byte {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *FindLocalBranchCommitAuthor) GetDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterType((*FindDefaultBranchNameRequest)(nil), "gitaly.FindDefaultBranchNameRequest")
	proto.RegisterType((*FindDefaultBranchNameResponse)(nil), "gitaly.FindDefaultBranchNameResponse")
	proto.RegisterType((*FindAllBranchNamesRequest)(nil), "gitaly.FindAllBranchNamesRequest")
	proto.RegisterType((*FindAllBranchNamesResponse)(nil), "gitaly.FindAllBranchNamesResponse")
	proto.RegisterType((*FindAllTagNamesRequest)(nil), "gitaly.FindAllTagNamesRequest")
	proto.RegisterType((*FindAllTagNamesResponse)(nil), "gitaly.FindAllTagNamesResponse")
	proto.RegisterType((*FindRefNameRequest)(nil), "gitaly.FindRefNameRequest")
	proto.RegisterType((*FindRefNameResponse)(nil), "gitaly.FindRefNameResponse")
	proto.RegisterType((*FindLocalBranchesRequest)(nil), "gitaly.FindLocalBranchesRequest")
	proto.RegisterType((*FindLocalBranchesResponse)(nil), "gitaly.FindLocalBranchesResponse")
	proto.RegisterType((*FindLocalBranchResponse)(nil), "gitaly.FindLocalBranchResponse")
	proto.RegisterType((*FindLocalBranchCommitAuthor)(nil), "gitaly.FindLocalBranchCommitAuthor")
	proto.RegisterEnum("gitaly.FindLocalBranchesRequest_SortBy", FindLocalBranchesRequest_SortBy_name, FindLocalBranchesRequest_SortBy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ref service

type RefClient interface {
	FindDefaultBranchName(ctx context.Context, in *FindDefaultBranchNameRequest, opts ...grpc.CallOption) (*FindDefaultBranchNameResponse, error)
	FindAllBranchNames(ctx context.Context, in *FindAllBranchNamesRequest, opts ...grpc.CallOption) (Ref_FindAllBranchNamesClient, error)
	FindAllTagNames(ctx context.Context, in *FindAllTagNamesRequest, opts ...grpc.CallOption) (Ref_FindAllTagNamesClient, error)
	// Find a Ref matching the given constraints. Response may be empty.
	FindRefName(ctx context.Context, in *FindRefNameRequest, opts ...grpc.CallOption) (*FindRefNameResponse, error)
	// Return a stream so we can divide the response in chunks of branches
	FindLocalBranches(ctx context.Context, in *FindLocalBranchesRequest, opts ...grpc.CallOption) (Ref_FindLocalBranchesClient, error)
}

type refClient struct {
	cc *grpc.ClientConn
}

func NewRefClient(cc *grpc.ClientConn) RefClient {
	return &refClient{cc}
}

func (c *refClient) FindDefaultBranchName(ctx context.Context, in *FindDefaultBranchNameRequest, opts ...grpc.CallOption) (*FindDefaultBranchNameResponse, error) {
	out := new(FindDefaultBranchNameResponse)
	err := grpc.Invoke(ctx, "/gitaly.Ref/FindDefaultBranchName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refClient) FindAllBranchNames(ctx context.Context, in *FindAllBranchNamesRequest, opts ...grpc.CallOption) (Ref_FindAllBranchNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ref_serviceDesc.Streams[0], c.cc, "/gitaly.Ref/FindAllBranchNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &refFindAllBranchNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ref_FindAllBranchNamesClient interface {
	Recv() (*FindAllBranchNamesResponse, error)
	grpc.ClientStream
}

type refFindAllBranchNamesClient struct {
	grpc.ClientStream
}

func (x *refFindAllBranchNamesClient) Recv() (*FindAllBranchNamesResponse, error) {
	m := new(FindAllBranchNamesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *refClient) FindAllTagNames(ctx context.Context, in *FindAllTagNamesRequest, opts ...grpc.CallOption) (Ref_FindAllTagNamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ref_serviceDesc.Streams[1], c.cc, "/gitaly.Ref/FindAllTagNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &refFindAllTagNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ref_FindAllTagNamesClient interface {
	Recv() (*FindAllTagNamesResponse, error)
	grpc.ClientStream
}

type refFindAllTagNamesClient struct {
	grpc.ClientStream
}

func (x *refFindAllTagNamesClient) Recv() (*FindAllTagNamesResponse, error) {
	m := new(FindAllTagNamesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *refClient) FindRefName(ctx context.Context, in *FindRefNameRequest, opts ...grpc.CallOption) (*FindRefNameResponse, error) {
	out := new(FindRefNameResponse)
	err := grpc.Invoke(ctx, "/gitaly.Ref/FindRefName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refClient) FindLocalBranches(ctx context.Context, in *FindLocalBranchesRequest, opts ...grpc.CallOption) (Ref_FindLocalBranchesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ref_serviceDesc.Streams[2], c.cc, "/gitaly.Ref/FindLocalBranches", opts...)
	if err != nil {
		return nil, err
	}
	x := &refFindLocalBranchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ref_FindLocalBranchesClient interface {
	Recv() (*FindLocalBranchesResponse, error)
	grpc.ClientStream
}

type refFindLocalBranchesClient struct {
	grpc.ClientStream
}

func (x *refFindLocalBranchesClient) Recv() (*FindLocalBranchesResponse, error) {
	m := new(FindLocalBranchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ref service

type RefServer interface {
	FindDefaultBranchName(context.Context, *FindDefaultBranchNameRequest) (*FindDefaultBranchNameResponse, error)
	FindAllBranchNames(*FindAllBranchNamesRequest, Ref_FindAllBranchNamesServer) error
	FindAllTagNames(*FindAllTagNamesRequest, Ref_FindAllTagNamesServer) error
	// Find a Ref matching the given constraints. Response may be empty.
	FindRefName(context.Context, *FindRefNameRequest) (*FindRefNameResponse, error)
	// Return a stream so we can divide the response in chunks of branches
	FindLocalBranches(*FindLocalBranchesRequest, Ref_FindLocalBranchesServer) error
}

func RegisterRefServer(s *grpc.Server, srv RefServer) {
	s.RegisterService(&_Ref_serviceDesc, srv)
}

func _Ref_FindDefaultBranchName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDefaultBranchNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefServer).FindDefaultBranchName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Ref/FindDefaultBranchName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefServer).FindDefaultBranchName(ctx, req.(*FindDefaultBranchNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ref_FindAllBranchNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindAllBranchNamesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefServer).FindAllBranchNames(m, &refFindAllBranchNamesServer{stream})
}

type Ref_FindAllBranchNamesServer interface {
	Send(*FindAllBranchNamesResponse) error
	grpc.ServerStream
}

type refFindAllBranchNamesServer struct {
	grpc.ServerStream
}

func (x *refFindAllBranchNamesServer) Send(m *FindAllBranchNamesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Ref_FindAllTagNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindAllTagNamesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefServer).FindAllTagNames(m, &refFindAllTagNamesServer{stream})
}

type Ref_FindAllTagNamesServer interface {
	Send(*FindAllTagNamesResponse) error
	grpc.ServerStream
}

type refFindAllTagNamesServer struct {
	grpc.ServerStream
}

func (x *refFindAllTagNamesServer) Send(m *FindAllTagNamesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Ref_FindRefName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRefNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefServer).FindRefName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.Ref/FindRefName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefServer).FindRefName(ctx, req.(*FindRefNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ref_FindLocalBranches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindLocalBranchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefServer).FindLocalBranches(m, &refFindLocalBranchesServer{stream})
}

type Ref_FindLocalBranchesServer interface {
	Send(*FindLocalBranchesResponse) error
	grpc.ServerStream
}

type refFindLocalBranchesServer struct {
	grpc.ServerStream
}

func (x *refFindLocalBranchesServer) Send(m *FindLocalBranchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Ref_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.Ref",
	HandlerType: (*RefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindDefaultBranchName",
			Handler:    _Ref_FindDefaultBranchName_Handler,
		},
		{
			MethodName: "FindRefName",
			Handler:    _Ref_FindRefName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindAllBranchNames",
			Handler:       _Ref_FindAllBranchNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindAllTagNames",
			Handler:       _Ref_FindAllTagNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindLocalBranches",
			Handler:       _Ref_FindLocalBranches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ref.proto",
}

func init() { proto.RegisterFile("ref.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x6f, 0x12, 0x41,
	0x10, 0x77, 0x5b, 0xc0, 0xeb, 0x40, 0x5b, 0x5c, 0x6b, 0x3d, 0xaf, 0xff, 0xe8, 0xb5, 0xb5, 0x18,
	0x93, 0xab, 0xa1, 0xf1, 0xc9, 0x17, 0xaf, 0xfc, 0x69, 0x9b, 0x20, 0x35, 0x0b, 0x4d, 0xf4, 0xe9,
	0x72, 0xc0, 0x02, 0x67, 0x80, 0xc3, 0xbb, 0x25, 0x4a, 0x8c, 0x2f, 0xbc, 0xf8, 0xe0, 0x83, 0xf1,
	0xc3, 0xf0, 0x19, 0xfc, 0x5a, 0xe6, 0x76, 0x97, 0x4a, 0xe9, 0x41, 0x4d, 0x78, 0xba, 0x9b, 0xd9,
	0xdf, 0xcc, 0x6f, 0x76, 0x66, 0xf6, 0x07, 0x2b, 0x1e, 0x6d, 0x18, 0x3d, 0xcf, 0x65, 0x2e, 0x8e,
	0x35, 0x1d, 0x66, 0xb7, 0x07, 0x5a, 0xc2, 0x6f, 0xd9, 0x1e, 0xad, 0x0b, 0xaf, 0xb6, 0xd7, 0x74,
	0xdd, 0x66, 0x9b, 0x9e, 0x70, 0xab, 0xda, 0x6f, 0x9c, 0x30, 0xa7, 0x43, 0x7d, 0x66, 0x77, 0x7a,
	0x12, 0xb0, 0xea, 0xf6, 0x98, 0xe3, 0x76, 0x7d, 0x61, 0xea, 0x04, 0xb6, 0x0b, 0x4e, 0xb7, 0x9e,
	0xa3, 0x0d, 0xbb, 0xdf, 0x66, 0x67, 0x9e, 0xdd, 0xad, 0xb5, 0x4a, 0x76, 0x87, 0x12, 0xfa, 0xb9,
	0x4f, 0x7d, 0x86, 0x33, 0x00, 0x1e, 0xed, 0xb9, 0xbe, 0xc3, 0x5c, 0x6f, 0xa0, 0xa2, 0x14, 0x4a,
	0xc7, 0x33, 0xd8, 0x10, 0xd4, 0x06, 0xb9, 0x39, 0x21, 0x13, 0x28, 0xfd, 0x14, 0x76, 0x66, 0xe4,
	0xf4, 0x7b, 0x6e, 0xd7, 0xa7, 0x18, 0x43, 0xa4, 0x6b, 0x77, 0x28, 0x4f, 0x97, 0x20, 0xfc, 0x5f,
	0xbf, 0x82, 0x67, 0x41, 0x90, 0xd9, 0x6e, 0xff, 0x0b, 0xf0, 0x17, 0xa9, 0x22, 0x03, 0x5a, 0x58,
	0x42, 0x59, 0xc2, 0x06, 0x44, 0x03, 0x5a, 0x5f, 0x45, 0xa9, 0xe5, 0x74, 0x82, 0x08, 0x43, 0x2f,
	0xc2, 0xa6, 0x8c, 0xa9, 0xd8, 0xcd, 0x85, 0x2b, 0x38, 0x81, 0xa7, 0x77, 0xb2, 0xcd, 0xa5, 0xff,
	0x0e, 0x38, 0x08, 0x20, 0xb4, 0xb1, 0xe0, 0x08, 0xf0, 0x16, 0xac, 0xd4, 0xdc, 0x4e, 0xc7, 0x61,
	0x96, 0x53, 0x57, 0x97, 0x52, 0x28, 0xbd, 0x42, 0x14, 0xe1, 0xb8, 0xac, 0xe3, 0x4d, 0x88, 0xf5,
	0x3c, 0xda, 0x70, 0xbe, 0xaa, 0xcb, 0x7c, 0x00, 0xd2, 0xd2, 0x5f, 0xc0, 0xe3, 0x5b, 0xf4, 0x73,
	0xa6, 0xf5, 0x07, 0x81, 0x1a, 0x60, 0x8b, 0x6e, 0xcd, 0x96, 0xfd, 0x5d, 0xa8, 0x57, 0xf8, 0x2d,
	0x3c, 0xf4, 0x5d, 0x8f, 0x59, 0xd5, 0x01, 0x2f, 0x77, 0x2d, 0x73, 0x3c, 0x0e, 0x98, 0x45, 0x63,
	0x94, 0x5d, 0x8f, 0x9d, 0x0d, 0x48, 0xcc, 0xe7, 0x5f, 0xfd, 0x35, 0xc4, 0x84, 0x07, 0x2b, 0x10,
	0x29, 0x99, 0xef, 0xf2, 0xc9, 0x07, 0x78, 0x1d, 0xe2, 0xd7, 0xef, 0x73, 0x66, 0x25, 0x9f, 0xb3,
	0xcc, 0x72, 0x36, 0x89, 0x70, 0x12, 0x12, 0x63, 0x47, 0x2e, 0x5f, 0xce, 0x26, 0x97, 0xf4, 0x0f,
	0x62, 0xef, 0xa6, 0x18, 0xe4, 0xd5, 0xdf, 0x80, 0x52, 0x95, 0x3e, 0x3e, 0xa9, 0x78, 0x66, 0x6f,
	0x46, 0x59, 0xe3, 0x10, 0x72, 0x13, 0xa0, 0xff, 0x5c, 0x12, 0xf3, 0x0f, 0x41, 0x85, 0xf5, 0x74,
	0xfe, 0xcc, 0x8e, 0x60, 0x4d, 0x1e, 0xfa, 0xfd, 0xea, 0x27, 0x5a, 0x63, 0x72, 0x76, 0xab, 0xc2,
	0x5b, 0x16, 0x4e, 0x7c, 0x01, 0xd2, 0x61, 0xd9, 0x7d, 0xd6, 0x72, 0x3d, 0x35, 0xc2, 0xbb, 0x7f,
	0x30, 0xa3, 0xea, 0x2c, 0xc7, 0x9a, 0x1c, 0x4a, 0x12, 0xb5, 0x09, 0x0b, 0x97, 0x20, 0x29, 0x33,
	0x89, 0x0f, 0xa3, 0x9e, 0x1a, 0xfd, 0xff, 0x64, 0xeb, 0x22, 0x2a, 0x3b, 0x8e, 0xd5, 0xbf, 0xc0,
	0xd6, 0x1c, 0x7c, 0x68, 0x43, 0x36, 0x20, 0x4a, 0x3b, 0xb6, 0xd3, 0xe6, 0xcd, 0x48, 0x10, 0x61,
	0x60, 0x03, 0x22, 0x75, 0x9b, 0x51, 0x7e, 0xff, 0x78, 0x46, 0x33, 0x84, 0xe0, 0x19, 0x63, 0xc1,
	0x33, 0x2a, 0x63, 0xc1, 0x23, 0x1c, 0x97, 0xf9, 0x15, 0x85, 0x65, 0x42, 0x1b, 0xf8, 0x37, 0x82,
	0x27, 0xa1, 0xb2, 0x84, 0x0f, 0x27, 0x2f, 0x34, 0x4b, 0x09, 0xb5, 0xa3, 0x7b, 0x50, 0x62, 0xb2,
	0xfa, 0xcb, 0xe1, 0x48, 0x3d, 0x56, 0x10, 0x3e, 0x38, 0xbf, 0xac, 0x98, 0xc5, 0x8f, 0x16, 0xc9,
	0x17, 0xac, 0xc2, 0x65, 0x29, 0xd8, 0xb7, 0x82, 0x79, 0x5d, 0xac, 0x58, 0x67, 0xc4, 0x2c, 0x65,
	0x2f, 0xac, 0x60, 0x43, 0xf1, 0x0f, 0x24, 0x5e, 0xfc, 0x6d, 0x91, 0xc2, 0xfb, 0x93, 0x54, 0xa1,
	0x8a, 0xa8, 0xe9, 0xf3, 0x20, 0xb2, 0x94, 0xf4, 0x70, 0xa4, 0x1e, 0x2a, 0x08, 0xa7, 0xa6, 0x4b,
	0x31, 0x8b, 0xc5, 0xc9, 0x32, 0xca, 0xaf, 0x10, 0xfe, 0x06, 0xeb, 0x53, 0x5a, 0x85, 0x77, 0xa7,
	0x28, 0xa6, 0x24, 0x51, 0xdb, 0x9b, 0x79, 0x2e, 0xf9, 0x8f, 0x86, 0x23, 0x75, 0x5f, 0x41, 0x78,
	0x27, 0x8c, 0xbf, 0x62, 0x9e, 0xdf, 0x90, 0xb7, 0x20, 0x3e, 0x21, 0x3c, 0x58, 0x9b, 0x4c, 0x7c,
	0x5b, 0x0c, 0xb5, 0xad, 0xd0, 0x33, 0x49, 0x98, 0x1a, 0x8e, 0xd4, 0x6d, 0x05, 0x61, 0x75, 0x9a,
	0x30, 0xf8, 0xe1, 0x0d, 0x1f, 0x22, 0x78, 0x74, 0xe7, 0xb9, 0xe3, 0xd4, 0x7d, 0x5a, 0xa3, 0xed,
	0xcf, 0x41, 0x48, 0xf2, 0xe7, 0xc3, 0x91, 0xaa, 0x2b, 0x08, 0xef, 0x4e, 0x93, 0x17, 0xaf, 0xb2,
	0xe6, 0xb8, 0xdf, 0xc1, 0x75, 0xab, 0x31, 0xbe, 0xab, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x5b, 0x6f, 0x58, 0xcd, 0x07, 0x00, 0x00,
}
