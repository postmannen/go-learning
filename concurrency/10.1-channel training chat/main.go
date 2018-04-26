package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type room struct {
	ID       int
	messages chan string      //the active chat messages for all in room
	clients  map[*client]bool //true if client is in room
	joining  chan *client     //used detect when to add value to map[*clients
	leaving  chan *client     //used detect when to delete value to map[*clients]
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
		case c := <-ro.leaving:
			log.Printf("client%v disconnecting\n", c.ID)
			delete(ro.clients, c)
		}
	}

}

type client struct {
	ID         int         //unique ID of client
	room       *room       //room that client belongs to
	msg        chan string //the channel to send a message directly to a client
	disconnect chan bool
	conn       *net.Conn
}

//create a new client with unique ID
func newClient(id int, conn *net.Conn) *client {
	return &client{
		ID:         id,
		msg:        make(chan string),
		disconnect: make(chan bool),
		conn:       conn,
	}
}

//attach a room given as input to the client
func (c *client) joinRoom(r *room) {
	c.room = r
	log.Printf("joinRoom: client1.ID =%v, is now in the room client.room.ID = %v\n", c.ID, c.room.ID)
	c.room.messages <- "Hello, I'm client1, and entering the room"

	//set the client active in the room
	c.room.joining <- c
}

func (c *client) run() {

	for c.msg != nil {
		select {
		case m := <-c.msg:
			fmt.Printf("client%v direct message: %v\n", c.ID, m)
		case <-c.disconnect:
			//give message to the room to delete the client from the room
			c.room.leaving <- c
			c.room.messages <- fmt.Sprintf("client%v disconnected\n", c.ID)
			break
		}
	}
}

const number int = 10

var clients = 1

func main() {
	fmt.Println("############################################################")
	room1 := newRoom(1)
	go room1.run()
	time.Sleep(time.Millisecond * 50) //let the room fully start before starting clients, will be removed later.

	nl, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("error: net.Listen :", err)
	}

	//wait for new network connections
	conn, err := nl.Accept()
	if err != nil {
		log.Println("error: net Accept :", err)
	}

	client1 := newClient(1, &conn)
	client1.joinRoom(room1)
	go client1.run()

	time.Sleep(time.Second * 1) //a little delay, will be removed later.
	client1.disconnect <- true

	time.Sleep(time.Second * 2) //a little delay, will be removed later.
	fmt.Println("------------------------------------------------------------")
	fmt.Println("room1 contains : ", *room1)

	panic("show stacks")

}
