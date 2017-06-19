// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ref.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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
	return fileDescriptor3, []int{8, 0}
}

type FindDefaultBranchNameRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *FindDefaultBranchNameRequest) Reset()                    { *m = FindDefaultBranchNameRequest{} }
func (m *FindDefaultBranchNameRequest) String() string            { return proto.CompactTextString(m) }
func (*FindDefaultBranchNameRequest) ProtoMessage()               {}
func (*FindDefaultBranchNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
func (*FindDefaultBranchNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
func (*FindAllBranchNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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
func (*FindAllBranchNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

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
func (*FindAllTagNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

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
func (*FindAllTagNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

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
func (*FindRefNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

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
func (*FindRefNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

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
func (*FindLocalBranchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

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
func (*FindLocalBranchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

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
func (*FindLocalBranchResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

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
	Name  []byte                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email []byte                     `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Date  *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
}

func (m *FindLocalBranchCommitAuthor) Reset()                    { *m = FindLocalBranchCommitAuthor{} }
func (m *FindLocalBranchCommitAuthor) String() string            { return proto.CompactTextString(m) }
func (*FindLocalBranchCommitAuthor) ProtoMessage()               {}
func (*FindLocalBranchCommitAuthor) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

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

func (m *FindLocalBranchCommitAuthor) GetDate() *google_protobuf.Timestamp {
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

func init() { proto.RegisterFile("ref.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0xdd, 0x2e, 0x7f, 0x7e, 0x70, 0xe1, 0xb7, 0xd4, 0x71, 0x5d, 0x6b, 0x51, 0x61, 0xab, 0x1b,
	0xf1, 0xa5, 0x98, 0x6e, 0x7c, 0xf2, 0x45, 0x16, 0x30, 0x6b, 0xb2, 0xa2, 0x19, 0xd0, 0xf8, 0x60,
	0x42, 0x0a, 0x4c, 0xd9, 0x9a, 0x96, 0xc1, 0x76, 0x88, 0xf2, 0xe0, 0x27, 0xf0, 0x83, 0xf9, 0xe6,
	0x67, 0x32, 0xcc, 0xb4, 0x50, 0xd8, 0xb6, 0x6b, 0xc2, 0x53, 0x99, 0x7b, 0xcf, 0xb9, 0xf7, 0xce,
	0x3d, 0xcc, 0x81, 0xa2, 0x47, 0x2c, 0x7d, 0xee, 0x51, 0x46, 0x51, 0x7e, 0x6a, 0x33, 0xd3, 0x59,
	0xaa, 0xb2, 0x47, 0xe6, 0xd4, 0xb7, 0x19, 0xf5, 0x96, 0x22, 0xa3, 0xd6, 0xa6, 0x94, 0x4e, 0x1d,
	0xd2, 0xe4, 0xa7, 0xd1, 0xc2, 0x6a, 0x32, 0xdb, 0x25, 0x3e, 0x33, 0xdd, 0xb9, 0x00, 0x68, 0x18,
	0x1e, 0xbe, 0xb1, 0x67, 0x93, 0x0e, 0xb1, 0xcc, 0x85, 0xc3, 0x2e, 0x3c, 0x73, 0x36, 0xbe, 0xee,
	0x99, 0x2e, 0xc1, 0xe4, 0xdb, 0x82, 0xf8, 0x0c, 0x19, 0x00, 0x9b, 0xa2, 0x8a, 0x54, 0x97, 0x1a,
	0x25, 0x03, 0xe9, 0xa2, 0x9f, 0x8e, 0xd7, 0x19, 0x1c, 0x41, 0x69, 0xe7, 0xf0, 0x28, 0xa1, 0xa6,
	0x3f, 0xa7, 0x33, 0x9f, 0x20, 0x04, 0xd9, 0x99, 0xe9, 0x12, 0x5e, 0xae, 0x8c, 0xf9, 0x6f, 0xed,
	0x3d, 0x3c, 0x58, 0x91, 0x5a, 0x8e, 0xb3, 0x21, 0xf8, 0xfb, 0x4c, 0x61, 0x80, 0x1a, 0x57, 0x30,
	0x18, 0xe1, 0x18, 0x72, 0xab, 0xb6, 0xbe, 0x22, 0xd5, 0x33, 0x8d, 0x32, 0x16, 0x07, 0xed, 0x0a,
	0x4e, 0x02, 0xce, 0xc0, 0x9c, 0xee, 0x3d, 0x41, 0x13, 0xee, 0xdf, 0xa8, 0x96, 0xda, 0xfe, 0x27,
	0xa0, 0x15, 0x01, 0x13, 0x6b, 0x4f, 0x09, 0x50, 0x15, 0x8a, 0x63, 0xea, 0xba, 0x36, 0x1b, 0xda,
	0x13, 0xe5, 0xb0, 0x2e, 0x35, 0x8a, 0xb8, 0x20, 0x02, 0x6f, 0x27, 0xe8, 0x04, 0xf2, 0x73, 0x8f,
	0x58, 0xf6, 0x0f, 0x25, 0xc3, 0x05, 0x08, 0x4e, 0xda, 0x73, 0xb8, 0xbb, 0xd5, 0x3e, 0x45, 0xad,
	0xdf, 0x12, 0x28, 0x2b, 0xec, 0x15, 0x1d, 0x9b, 0xc1, 0x7e, 0xf7, 0xda, 0x15, 0x7a, 0x0d, 0xff,
	0xf9, 0xd4, 0x63, 0xc3, 0xd1, 0x92, 0x8f, 0x7b, 0x64, 0x3c, 0x0b, 0x09, 0x49, 0x6d, 0xf4, 0x3e,
	0xf5, 0xd8, 0xc5, 0x12, 0xe7, 0x7d, 0xfe, 0xd5, 0x5e, 0x42, 0x5e, 0x44, 0x50, 0x01, 0xb2, 0xbd,
	0xd6, 0xbb, 0xae, 0x7c, 0x80, 0x2a, 0x50, 0xfa, 0xf8, 0xa1, 0xd3, 0x1a, 0x74, 0x3b, 0xc3, 0x56,
	0xbf, 0x2d, 0x4b, 0x48, 0x86, 0x72, 0x18, 0xe8, 0x74, 0xfb, 0x6d, 0xf9, 0x50, 0xfb, 0x2c, 0xfe,
	0x77, 0x3b, 0x1d, 0x82, 0xab, 0xbf, 0x82, 0xc2, 0x28, 0x88, 0x71, 0xa5, 0x4a, 0x46, 0x2d, 0x61,
	0xac, 0x90, 0x82, 0xd7, 0x04, 0xed, 0xd7, 0xa1, 0xd0, 0x3f, 0x06, 0x15, 0xb7, 0xd3, 0x74, 0xcd,
	0xce, 0xe0, 0x28, 0x48, 0xfa, 0x8b, 0xd1, 0x57, 0x32, 0x66, 0x81, 0x76, 0xff, 0x8b, 0x68, 0x5f,
	0x04, 0xd1, 0x25, 0x04, 0x81, 0xa1, 0xb9, 0x60, 0xd7, 0xd4, 0x53, 0xb2, 0x7c, 0xfb, 0x4f, 0x12,
	0xa6, 0x6e, 0x73, 0x6c, 0x8b, 0x43, 0x71, 0x79, 0x1c, 0x39, 0xa1, 0x1e, 0xc8, 0x41, 0x25, 0xf1,
	0x61, 0xc4, 0x53, 0x72, 0xff, 0x5e, 0xac, 0x22, 0x58, 0xed, 0x90, 0xab, 0x7d, 0x87, 0x6a, 0x0a,
	0x3e, 0x76, 0x21, 0xc7, 0x90, 0x23, 0xae, 0x69, 0x3b, 0x7c, 0x19, 0x65, 0x2c, 0x0e, 0x48, 0x87,
	0xec, 0xc4, 0x64, 0x84, 0xdf, 0xbf, 0x64, 0xa8, 0xba, 0x70, 0x38, 0x3d, 0x74, 0x38, 0x7d, 0x10,
	0x3a, 0x1c, 0xe6, 0x38, 0xe3, 0x4f, 0x06, 0x32, 0x98, 0x58, 0xc8, 0x82, 0x7b, 0xb1, 0xae, 0x84,
	0x9e, 0x46, 0xef, 0x93, 0x64, 0x84, 0xea, 0xd9, 0x2d, 0x28, 0x21, 0xac, 0x76, 0x80, 0x86, 0xe2,
	0x11, 0x6f, 0xfb, 0x0e, 0x3a, 0x8d, 0xd2, 0x63, 0x4d, 0x4e, 0xd5, 0xd2, 0x20, 0x61, 0xf9, 0x17,
	0x12, 0xfa, 0x04, 0x95, 0x1d, 0x5b, 0x41, 0x8f, 0x77, 0xa8, 0x3b, 0xee, 0xa5, 0xd6, 0x12, 0xf3,
	0x91, 0xba, 0x97, 0x50, 0x8a, 0x3c, 0x7f, 0xa4, 0x46, 0x39, 0xdb, 0x96, 0xa4, 0x56, 0x63, 0x73,
	0xeb, 0x15, 0x7c, 0x81, 0x3b, 0x37, 0xde, 0x14, 0xaa, 0xdf, 0xf6, 0xa0, 0xd5, 0xd3, 0x14, 0xc4,
	0x66, 0xce, 0x51, 0x9e, 0x4b, 0x7d, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xeb, 0x20, 0x3d,
	0x01, 0x07, 0x00, 0x00,
}
