# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gitaly.proto

require 'google/protobuf'

require 'shared_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.CommitIsAncestorResponse" do
    optional :value, :bool, 1
  end
  add_message "gitaly.PostReceiveRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.PostReceiveResponse" do
  end
  add_message "gitaly.FindRefNameRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :commit_id, :string, 2
    optional :prefix, :bytes, 3
  end
  add_message "gitaly.FindRefNameResponse" do
    optional :name, :bytes, 1
  end
  add_message "gitaly.CommitIsAncestorRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :ancestor_id, :string, 2
    optional :child_id, :string, 3
  end
  add_message "gitaly.FindDefaultBranchNameRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindAllBranchNamesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindAllTagNamesRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
  end
  add_message "gitaly.FindDefaultBranchNameResponse" do
    optional :name, :bytes, 1
  end
  add_message "gitaly.FindAllBranchNamesResponse" do
    repeated :names, :bytes, 1
  end
  add_message "gitaly.FindAllTagNamesResponse" do
    repeated :names, :bytes, 1
  end
  add_message "gitaly.CommitDiffRequest" do
    optional :repository, :message, 1, "gitaly.Repository"
    optional :left_commit_id, :string, 2
    optional :right_commit_id, :string, 3
  end
  add_message "gitaly.CommitDiffResponse" do
    optional :from_path, :bytes, 1
    optional :to_path, :bytes, 2
    optional :from_id, :string, 3
    optional :to_id, :string, 4
    optional :old_mode, :int32, 5
    optional :new_mode, :int32, 6
    optional :binary, :bool, 7
    repeated :raw_chunks, :bytes, 8
  end
end

module Gitaly
  CommitIsAncestorResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitIsAncestorResponse").msgclass
  PostReceiveRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.PostReceiveRequest").msgclass
  PostReceiveResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.PostReceiveResponse").msgclass
  FindRefNameRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindRefNameRequest").msgclass
  FindRefNameResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindRefNameResponse").msgclass
  CommitIsAncestorRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitIsAncestorRequest").msgclass
  FindDefaultBranchNameRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindDefaultBranchNameRequest").msgclass
  FindAllBranchNamesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchNamesRequest").msgclass
  FindAllTagNamesRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllTagNamesRequest").msgclass
  FindDefaultBranchNameResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindDefaultBranchNameResponse").msgclass
  FindAllBranchNamesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllBranchNamesResponse").msgclass
  FindAllTagNamesResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.FindAllTagNamesResponse").msgclass
  CommitDiffRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitDiffRequest").msgclass
  CommitDiffResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitDiffResponse").msgclass
end
