# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: commit.proto for package 'gitaly'

require 'grpc'
require 'commit_pb'

module Gitaly
  module CommitService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.CommitService'

      rpc :CommitIsAncestor, CommitIsAncestorRequest, CommitIsAncestorResponse
      rpc :TreeEntry, TreeEntryRequest, stream(TreeEntryResponse)
      rpc :CommitsBetween, CommitsBetweenRequest, stream(CommitsBetweenResponse)
      rpc :CountCommits, CountCommitsRequest, CountCommitsResponse
      rpc :GetTreeEntries, GetTreeEntriesRequest, stream(GetTreeEntriesResponse)
      rpc :ListFiles, ListFilesRequest, stream(ListFilesResponse)
      rpc :FindCommit, FindCommitRequest, FindCommitResponse
      rpc :CommitStats, CommitStatsRequest, CommitStatsResponse
      # Use a stream to paginate the result set
      rpc :FindAllCommits, FindAllCommitsRequest, stream(FindAllCommitsResponse)
      rpc :CommitLanguages, CommitLanguagesRequest, CommitLanguagesResponse
      rpc :RawBlame, RawBlameRequest, stream(RawBlameResponse)
      rpc :LastCommitForPath, LastCommitForPathRequest, LastCommitForPathResponse
    end

    Stub = Service.rpc_stub_class
  end
end
