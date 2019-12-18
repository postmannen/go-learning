package main

import "google.golang.org/grpc"

import "log"

import "github.com/postmannen/go-learning/grpc/02-messaging/messagingpb"

import "context"

import "fmt"

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to do grpc.Dial... : ", err)
	}
	defer conn.Close()

	c := messagingpb.NewChatterClient(conn)

	msg := messagingpb.ChatRequest{
		Text: "This is a chat message",
	}

	res, err := c.Chat(context.Background(), &msg)
	if err != nil {
		log.Println("Chat from client failed : ", err)
	}

	fmt.Printf("return value for OK received : %v\n", res.Ok)
}
