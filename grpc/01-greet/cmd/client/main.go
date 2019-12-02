package main

import "google.golang.org/grpc"

import "log"

import "github.com/postmannen/go-learning/grpc/01-greet/greetpb"

import "context"

import "fmt"

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial from client failed %v\n", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "King",
			LastName:  "Olav",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Println("Failed to greet from client ", err)
	}
	fmt.Println(res.GetResult())

}
