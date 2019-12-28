package main

import "google.golang.org/grpc"

import "log"

import "github.com/postmannen/go-learning/grpc/01-greet/greetpb"

import "context"

import "fmt"

func main() {
	// Create a GRPC connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial from client failed %v\n", err)
	}
	defer conn.Close()

	// Create a new GreetServiceClient.
	// This will return a GRPC client conn with the connection,
	// and we can then later call the methods registered with that type.
	//
	// The returned *greetServiceClient which is a struct got a field
	// called cc who holds the connection.
	c := greetpb.NewGreetServiceClient(conn)

	// Create a request message to pass into the Greet method
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "King",
			LastName:  "Olav",
		},
	}

	// Do the Greet method on the client.
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Println("Failed to greet from client ", err)
	}
	fmt.Println(res.GetResult())

}
