package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type room struct {
	ID       int
	messages chan io.Reader   //the active chat messages for all in room
	clients  map[*client]bool //true if client is in room
	joining  chan *client
	leaving  chan *client
}

//create a new room
func newRoom(id int) *room {
	return &room{
		ID:       id,
		clients:  make(map[*client]bool),
		messages: make(chan io.Reader), //initialize channel, or...deadlock
		joining:  make(chan *client),   //initialize channel, or...deadlock
		leaving:  make(chan *client),   //initialize channel, or...deadlock
	}
}

//The room will allways be up, and do things based on what is received on its channels
func (r *room) run() {
	fmt.Println("-----------Starting up the room with ID = ", r.ID, "-----------")

	for {
		select {
		//any new incomming messages to the room ?
		case msg := <-r.messages:
			fmt.Printf("room%v received msg: %v\n", r.ID, msg)
			//send the room message out to all the active clients in the room
			for cl := range r.clients {
				cl.msg <- msg
			}
		//create a reference of the client in the room struct, and set its value to true
		//to indicate that this client is in the room
		case c := <-r.joining:
			r.clients[c] = true
			//sends message directly to the client, and not in room for all to see
			c.msg <- bytes.NewBufferString("Welcome to the room !")
		case c := <-r.leaving:
			fmt.Printf("----------- client%v are leaving room%v-----------\n", c.ID, r.ID)
			delete(r.clients, c)
		}
	}

}

type client struct {
	ID   int            //unique ID of client
	room *room          //room that client belongs to
	msg  chan io.Reader //the channel to send a message directly to a client
}

//attach a room given as input to the client
func (c *client) joinRoom(r *room) {
	c.room = r
	fmt.Printf("-----------joinRoom: client1.ID =%v, is now in the room client.room.ID = %v-----------\n", c.ID, c.room.ID)
	//Since Hello message goes to the room all clients will se it
	myString := fmt.Sprintf("Hello, I'm client%v, and entering the room", c.ID)
	c.room.messages <- bytes.NewBufferString(myString)

	//set the client active in the room
	c.room.joining <- c
}

func (c *client) leaveRoom(r *room) {
	r.leaving <- c
}

func (c *client) run() {
	for {
		select {
		case m := <-c.msg:
			fmt.Printf("client%v received msg: %v\n", c.ID, m)
		}
	}
}

//create a new client with unique ID
func newClient(id int) *client {
	return &client{
		ID:  id,
		msg: make(chan io.Reader),
	}
}

const number int = 10

func main() {
	room1 := newRoom(1)
	go room1.run()
	time.Sleep(time.Millisecond * 50) //let the room fully start before starting clients, will be removed later.

	client1 := newClient(1)
	client1.joinRoom(room1)
	go client1.run()

	//send a test message to the room for all clients to receive
	room1.messages <- bytes.NewBufferString("this is the first message for all clients in the room")

	client2 := newClient(2)
	client2.joinRoom(room1)
	go client2.run()

	//send a test message to the room for all clients to receive
	room1.messages <- bytes.NewBufferString("this is the second message for all clients in the room")

	client1.leaveRoom(room1)

	//send a test message to the room for all clients to receive
	room1.messages <- bytes.NewBufferString("this is the third message for all clients in the room")

	//start a tcp server for clients
	server, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println("error: failed to open tcp listener : ", err)
	}
	defer server.Close()

	clientID := 0
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("error: network connection failed:", err)
		}
		//hvordan skal man starte en ny klient ???????????????
		go 
	}

	time.Sleep(time.Second * 2) //a little delay, will be removed later.
	fmt.Println("------------------------------------------------------------")
	fmt.Println("room1 contains : ", room1)

}
