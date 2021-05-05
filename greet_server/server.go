package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-greeting-service/greetpb"
	"log"
	"net"
)

type server struct {}

func main() {
	fmt.Println("Hello, World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if s := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
