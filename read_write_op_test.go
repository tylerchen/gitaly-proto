package main

import (
	"testing"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"gitlab.com/gitlab-org/gitaly-proto/go/gitalypb"
)

func TestReadWriteOps(t *testing.T) {
	for _, req := range []descriptor.Message{&gitalypb.RepositoryExistsRequest{}, &gitalypb.RepackIncrementalRequest{}} {
		_, md := descriptor.ForMessage(req)
		ex, err := proto.GetExtension(md.GetOptions(), gitalypb.E_OpType)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(ex)
	}

	t.Fail()
}
