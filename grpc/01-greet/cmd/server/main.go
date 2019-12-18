package main

import (
	"context"
	"log"
	"net"

	"github.com/postmannen/go-learning/grpc/01-greet/greetpb"
	"google.golang.org/grpc"
)

// server is the type where we implement all the methods of the GreetService.
// This type must implement all the methods of the GreetService to satisfy
// that interfave type
//
// The type is later down in the code registered with the grpc.Server by calling
// greetpb.RegisterGreetServiceServer(s, &server{}).
type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	r := req.GetGreeting()
	fName := r.GetFirstName()
	lName := r.GetLastName()
	result := "Hello " + fName + " " + lName
	response := greetpb.GreetResponse{
		Result: result,
	}
	return &response, nil
}

func main() {
	log.Printf("Preparing to start gRPC listener...\n")

	// Create a normal TCP Listener.
	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("error: failed to open listener: %v\n", err)
	}

	// Create a gRPC server. Will return a *grpc.Server
	s := grpc.NewServer()

	// RegsiterService registers the gRPC server s with the
	// server{} object we have created.
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatalf("error: failed to serve: %v\n", err)
	}

}
