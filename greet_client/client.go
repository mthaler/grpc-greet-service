package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-greet-service/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	fmt.Println("Hello, I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	//doUnary(c)

	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("Starting to do unary RPC...")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Richard",
			LastName: "Feynman",
		},
	}

	res, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting to do server streaming RPC...")
	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Richard",
			LastName: "Feynman",
		},
	}

	stream, err := c.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg)
	}
}

