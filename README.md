# grpc-greeting-service
gRPC greeting service from gRPC Master Class: Build Modern API &amp; Microservices on Udemy

## Install protobuf compiler plugins

Install the protobuf compiler plugins for Go using the following commands:

```bash
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Update PATH so that the protoc compiler can find the plugins:

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Compile a gRPC service

```bash
$ protoc --go_out=. --go-grpc_out=.  greetpb/greet.proto
```
## Error handling

The greet service does not implement error handling. A good guide for gRPC error handling can be found here: https://avi.im/grpc-errors/
