package main

import (
	"context"
	"log"
	"net"

	"github.com/postmannen/go-learning/grpc/01-greet/greetpb"
	"google.golang.org/grpc"
)

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

	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("error: failed to open listener: %v\n", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatalf("error: failed to serve: %v\n", err)
	}

}
