# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: server.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.ServerInfoRequest" do
  end
  add_message "gitaly.ServerInfoResponse" do
    optional :server_version, :string, 1
    optional :git_version, :string, 2
  end
end

module Gitaly
  ServerInfoRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ServerInfoRequest").msgclass
  ServerInfoResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ServerInfoResponse").msgclass
end