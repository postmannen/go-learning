package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

type room struct {
	ID       int
	messages chan string      //the active chat messages for all in room
	buf      bytes.Buffer     //buffer of recent messages
	clients  map[*client]bool //true if client is in room
	joining  chan *client
	leaving  chan *client
}

//create a new room
func newRoom(id int) *room {
	return &room{
		ID:       id,
		clients:  make(map[*client]bool),
		messages: make(chan string),  //initialize channel, or...deadlock
		joining:  make(chan *client), //initialize channel, or...deadlock
		leaving:  make(chan *client), //initialize channel, or...deadlock
	}
}

//The room will allways be up, and do things based on what is received on its channels
func (ro *room) run() {
	log.Println("Starting up the room with ID = ", ro.ID)

	for {
		select {
		//any new incomming messages to the room ?
		case msg := <-ro.messages:
			fmt.Printf("room%v: %v\n", ro.ID, msg)
		//create a reference of the client in the room struct, and set its value to true
		//to indicate that this client is in the room
		case c := <-ro.joining:
			ro.clients[c] = true
			c.msg <- "Welcome to the room !"

		}
	}

}

type client struct {
	ID   int         //unique ID of client
	room *room       //room that client belongs to
	msg  chan string //the channel to send a message directly to a client
}

//attach a room given as input to the client
func (c *client) joinRoom(ro *room) {
	c.room = ro
	log.Printf("joinRoom: client1.ID =%v, is now in the room client.room.ID = %v\n", c.ID, c.room.ID)
	c.room.messages <- "Hello, I'm client1, and entering the room"

	//set the client active in the room
	c.room.joining <- c
}

func (c *client) run() {

	for {
		select {
		case m := <-c.msg:
			fmt.Printf("client%v direct message: %v\n", c.ID, m)
		}
	}
}

//create a new client with unique ID
func newClient(id int) *client {
	return &client{
		ID:  id,
		msg: make(chan string),
	}
}

const number int = 10

func main() {
	room1 := newRoom(1)
	go room1.run()

	client1 := newClient(1)
	client1.joinRoom(room1)
	go client1.run()

	client2 := newClient(2)
	client2.joinRoom(room1)
	go client2.run()

	time.Sleep(time.Second * 2)
	//	fmt.Println("------------------------------------------------------------")
	//	fmt.Println("room1 contains : ", room1)

}
