# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: remote.proto for package 'gitaly'

require 'grpc'
require 'remote_pb'

module Gitaly
  module RemoteService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.RemoteService'

      rpc :AddRemote, AddRemoteRequest, AddRemoteResponse
      rpc :FetchInternalRemote, FetchInternalRemoteRequest, FetchInternalRemoteResponse
      rpc :RemoveRemote, RemoveRemoteRequest, RemoveRemoteResponse
      rpc :UpdateRemoteMirror, stream(UpdateRemoteMirrorRequest), UpdateRemoteMirrorResponse
      rpc :FindRemoteRepository, FindRemoteRepositoryRequest, FindRemoteRepositoryResponse
      rpc :FindRemoteRootRef, FindRemoteRootRefRequest, FindRemoteRootRefResponse
    end

    Stub = Service.rpc_stub_class
  end
end
