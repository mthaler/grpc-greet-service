#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=.  greetpb/greet.proto
