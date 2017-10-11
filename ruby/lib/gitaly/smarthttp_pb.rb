# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: smarthttp.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.InfoRefsRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    repeated :git_config_options, :string, 2
  end
  add_message "gitaly.InfoRefsResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.PostUploadPackRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :data, :bytes, 2
    repeated :git_config_options, :string, 3
  end
  add_message "gitaly.PostUploadPackResponse" do
    optional :data, :bytes, 1
  end
  add_message "gitaly.PostReceivePackRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :data, :bytes, 2
    optional :gl_id, :string, 3
    optional :gl_repository, :string, 4
    optional :gl_username, :string, 5
  end
  add_message "gitaly.PostReceivePackResponse" do
    optional :data, :bytes, 1
  end
end

module Gitaly
  InfoRefsRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.InfoRefsRequest").msgclass
  InfoRefsResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.InfoRefsResponse").msgclass
  PostUploadPackRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.PostUploadPackRequest").msgclass
  PostUploadPackResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.PostUploadPackResponse").msgclass
  PostReceivePackRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.PostReceivePackRequest").msgclass
  PostReceivePackResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.PostReceivePackResponse").msgclass
end
