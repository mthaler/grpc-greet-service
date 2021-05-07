package main

import (
	"context"
	"fmt"
	"github.com/mthaler/grpc-greet-service/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	response := greetpb.GreetResponse{
		Result: result,
	}
	return &response, nil
}

func (*server) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		response := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with %v\n", stream)
	result := ""
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream: %v", err)
		}

		firstName :=  request.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func(*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with %v\n", stream)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream: %v", err)
			return err
		}
		firstName := req.Greeting.GetFirstName()
		result := "Hello " + firstName + "! "
		err = stream.SendMsg(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %w")
			return err
		}
	}
}

func (*server) GreetWithDeadline(ctx context.Context, request *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with %v\n", request)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	response := greetpb.GreetWithDeadlineResponse{
		Result: result,
	}
	return &response, nil
}

func main() {
	fmt.Println("Hello, World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
