package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/postmannen/go-learning/grpc/02-messaging/messagingpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Chat(ctx context.Context, req *messagingpb.ChatRequest) (*messagingpb.ChatResponse, error) {
	text := req.GetText()
	fmt.Println("Server Received : ", text)

	return &messagingpb.ChatResponse{Ok: true}, nil
}

func main() {
	var err error

	fmt.Println("Starting gRPC server process.....")

	// Create a tcp listener
	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to start listener: ", err)
	}

	// Create a new grpc server
	srv := grpc.NewServer()
	// Register the server type with the grpc server.
	// If this is not done the grpc server will not know of the grpc
	// methods to execute.
	messagingpb.RegisterChatterServer(srv, &server{})

	err = srv.Serve(l)
	if err != nil {
		log.Fatal("failed to serve ...", err)
	}

}
