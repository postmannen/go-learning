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
	close(c.msg) //TESTING
}

//read from client network connection, and put the data on the messages channel to the room
func (c *client) writeRoomChannel(conn net.Conn, r *room) {
	for {
		data := make([]byte, 1024) //Will let the Read method read 1024

		_, err := conn.Read(data)
		if err != nil {
			log.Println("Error client.wrote:", err)
			break
		}

		r.messages <- bytes.NewBuffer(data)

	}
	conn.Close()

	s := fmt.Sprintf("client%v left the room !", c.ID)
	r.messages <- bytes.NewBufferString(s)
}

func handleClient(conn net.Conn, cID int, r *room) {
	client := newClient(cID)
	client.joinRoom(r)
	//starting detection of client.writeRoomChannel in its own go routine so it not hangs the execution waiting
	go client.writeRoomChannel(conn, r)

	//check if something is received on the client.msg channel, and write it to the telnet session
	for {
		select {
		case m := <-client.msg:
			m2 := fmt.Sprintf("client%v received msg: %v\n", client.ID, m)
			_, err := conn.Write([]byte(m2))
			if err != nil {
				fmt.Println("---> error: conn.Write", err)
				client.leaveRoom(r)
				break
			}

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

func main() {
	//create a default room for clients to connect to
	room1 := newRoom(1)
	go room1.run()
	time.Sleep(time.Millisecond * 50) //let the room fully start before starting clients, will be removed later.

	//send a test message to the room for all clients to receive
	room1.messages <- bytes.NewBufferString("this is the first message for all clients in the room")

	//start a tcp server for accepting clients
	server, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println("error: failed to open tcp listener : ", err)
	}
	defer server.Close()

	clientID := 1
	//wait for new network connection, and start new client sessions
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("error: network connection failed:", err)
		} else {
			go handleClient(conn, clientID, room1)
			clientID++
		}
	}

}
