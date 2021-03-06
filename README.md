# Protobuf specifications and client libraries for Gitaly

Gitaly is part of GitLab. It is a [server
application](https://gitlab.com/gitlab-org/gitaly) that uses its own
gRPC protocol to communicate with its clients. This repository
contains the protocol definition and automatically generated wrapper
code for Go and Ruby.

The .proto files define the remote procedure calls for interacting
with Gitaly. We keep auto-generated client libraries for Ruby and Go
in their respective subdirectories.

Run `make` from the root of the repository to regenerate the client
libraries after updating .proto files.

See
[developers.google.com](https://developers.google.com/protocol-buffers/docs/proto3)
for documentation of the 'proto3' Protocol buffer specification
language.

## gRPC/Protobuf concepts

The core Protobuf concepts we use are rpc, service and message. We use
these to define the Gitaly **protocol**.

-   **rpc** a function that can be called from the client and that gets
    executed on the server. Belongs to a service. Can have one of four
    request/response signatures: message/message (example: get metadata for
    commit xxx), message/stream (example: get contents of blob xxx),
    stream/message (example: create new blob with contents xxx),
    stream/stream (example: git SSH session).
-   **service** a logical group of RPC's.
-   **message** like a JSON object except it has pre-defined types.
-   **stream** an unbounded sequence of messages. In the Ruby clients
    this looks like an Enumerator.

gRPC provides an implementation framework based on these Protobuf concepts.

-   A gRPC **server** implements one or more services behind a network
    listener. Example: the Gitaly server application.
-   The gRPC toolchain automatically generates **client libraries** that
    handle serialization and connection management. Example: the Go
    client package and Ruby gem in this repository.
-   gRPC **clients** use the client libraries to make remote procedure
    calls. These clients must decide what network address to reach their
    gRPC servers on and handle connection reuse: it is possible to
    spread different gRPC services over multiple connections to the same
    gRPC server.
-   Officially a gRPC connection is called a **channel**. In the Go gRPC
    library these channels are called **client connections** because
    'channel' is already a concept in Go itself. In Ruby a gRPC channel
    is an instance of GRPC::Core::Channel. We use the word 'connection'
    in this document. The underlying transport of gRPC, HTTP/2, allows
    multiple remote procedure calls to happen at the same time on a
    single connection to a gRPC server. In principle, a multi-threaded
    gRPC client needs only one connection to a gRPC server.

## Design decisions

1.  In Gitaly's case there is one server application
    https://gitlab.com/gitlab-org/gitaly which implements all services
    in the protocol.
1.  In default GitLab installations each Gitaly client interacts with
    exactly 1 Gitaly server, on the same host, via a Unix domain socket.
    In a larger installation each Gitaly client will interact with many
    different Gitaly servers (one per GitLab storage shard) via TCP
    connections.
1.  Gitaly uses
    [grpc.Errorf](https://godoc.org/google.golang.org/grpc#Errorf) to
    return meaningful
    [errors](https://godoc.org/google.golang.org/grpc/codes#Code) to its
    clients.
1.  Each RPC `FooBar` has its own `FooBarRequest` and `FooBarResponse`
    message types. Try to keep the structure of these messages as flat as
    possible. Only add abstractions when they have a practical benefit.
1.  We never make backwards incompatible changes to an RPC that is
    already implemented on either the client side or server side.
    Instead we just create a new RPC call and start a deprecation
    procedure (see below) for the old one.
1.  It is encouraged to put comments (starting with `//`) in .proto files.
    Please put comments on their own lines. This will cause them to be
    treated as documentation by the protoc compiler.
1.  When choosing an RPC name don't use the service name as context.
    Good: `service CommitService { rpc CommitExists }`. Bad:
    `service CommitService { rpc Exists }`.

## Contributing

The CI at https://gitlab.com/gitlab-org/gitaly-proto regenerates the
client libraries to guard against the mistake of updating the .proto
files but not the client libraries. This check uses `git diff` to look
for changes. Some of the code in the Go client libraries is sensitive
to implementation details of the Go standard library (specifically,
the output of gzip). **Use the same Go version as .gitlab-ci.yml (Go
1.8)** when generating new client libraries for a merge request.

[DCO + License](CONTRIBUTING.md)

### Build process

After you change or add a .proto file you need to re-generate the Go
and Ruby libraries before committing your change.

```
# Re-generate Go and Ruby libraries
make generate
```

## How to deprecate an RPC call

See [DEPRECATION.md](DEPRECATION.md).

## Release

This will tag and release the gitaly-proto library, including
pushing the gem to rubygems.org

```
make release version=X.Y.Z
```


## How to manually push the gem

If the release script fails the gem may not be pushed. This is how you can do that after the fact:

```shell
# Use a sub-shell to limit scope of 'set -e'
(
  set -e

  # Replace X.Y.Z with the version you are pushing
  GEM_VERSION=X.Y.Z

  git checkout v$GEM_VERSION
  gem build gitaly.gemspec
  gem push gitaly-$GEM_VERSION.gem
)
```
