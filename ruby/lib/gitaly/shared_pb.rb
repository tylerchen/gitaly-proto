# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: shared.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "gitaly.Repository" do
    optional :storage_name, :string, 2
    optional :relative_path, :string, 3
    optional :git_object_directory, :string, 4
    repeated :git_alternate_object_directories, :string, 5
    optional :gl_repository, :string, 6
  end
  add_message "gitaly.GitCommit" do
    optional :id, :string, 1
    optional :subject, :bytes, 2
    optional :body, :bytes, 3
    optional :author, :message, 4, "gitaly.CommitAuthor"
    optional :committer, :message, 5, "gitaly.CommitAuthor"
    repeated :parent_ids, :string, 6
    optional :body_size, :int64, 7
  end
  add_message "gitaly.CommitAuthor" do
    optional :name, :bytes, 1
    optional :email, :bytes, 2
    optional :date, :message, 3, "google.protobuf.Timestamp"
  end
  add_message "gitaly.ExitStatus" do
    optional :value, :int32, 1
  end
  add_message "gitaly.Branch" do
    optional :name, :bytes, 1
    optional :target_commit, :message, 2, "gitaly.GitCommit"
  end
  add_message "gitaly.Tag" do
    optional :name, :bytes, 1
    optional :id, :string, 2
    optional :target_commit, :message, 3, "gitaly.GitCommit"
    optional :message, :bytes, 4
    optional :message_size, :int64, 5
    optional :tagger, :message, 6, "gitaly.CommitAuthor"
  end
  add_message "gitaly.User" do
    optional :gl_id, :string, 1
    optional :name, :bytes, 2
    optional :email, :bytes, 3
    optional :gl_username, :string, 4
  end
end

module Gitaly
  Repository = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.Repository").msgclass
  GitCommit = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.GitCommit").msgclass
  CommitAuthor = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.CommitAuthor").msgclass
  ExitStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.ExitStatus").msgclass
  Branch = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.Branch").msgclass
  Tag = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.Tag").msgclass
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("gitaly.User").msgclass
end
