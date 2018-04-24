/*
Initially it's working, but if a client don't do a graceful exit then the exit message spawns alot of blank lines in
the other clients terminals.
If the client exits by closing the connection it works ok.
*/

package main

import (
	"fmt"
	"log"
	"net"
)

type room struct {
	ID       int
	messages chan []byte      //the active chat messages for all in room
	clients  map[*client]bool //true if client is in room
	joining  chan *client
	leaving  chan *client
}

//create a new room
func newRoom(id int) *room {
	return &room{
		ID:       id,
		clients:  make(map[*client]bool),
		messages: make(chan []byte),  //initialize channel, or...deadlock
		joining:  make(chan *client), //initialize channel, or...deadlock
		leaving:  make(chan *client), //initialize channel, or...deadlock
	}
}

//The room will allways be up, and do things based on what is received on its channels
func (r *room) run() {
	fmt.Println("-----------Starting up the room with ID = ", r.ID, "-----------")

	for {
		select {
		//any new incomming messages to the room ?
		case msg := <-r.messages:
			fmt.Printf("room%v received msg: %v\n", r.ID, string(msg))
			//send the room message out to all the active clients in the room
			for cl := range r.clients {
				cl.msg <- msg
			}
		//create a reference of the client in the room struct, and set its value to true
		//to indicate that this client is in the room
		case c := <-r.joining:
			r.clients[c] = true
			//sends message directly to the client, and not in room for all to see
			c.msg <- []byte("Welcome to the room !")
		case c := <-r.leaving:
			fmt.Printf("----------- client%v are leaving room%v-----------\n", c.ID, r.ID)
			delete(r.clients, c)
		}
	}

}

type client struct {
	ID   int         //unique ID of client
	room *room       //room that client belongs to
	msg  chan []byte //the channel to send a message directly to a client
}

//attach a room given as input to the client
func (c *client) joinRoom(r *room) {
	c.room = r
	fmt.Printf("-----------joinRoom: client1.ID =%v, is now in the room client.room.ID = %v-----------\n", c.ID, c.room.ID)
	//Since Hello message goes to the room all clients will se it
	myString := fmt.Sprintf("Hello, I'm client%v, and entering the room", c.ID)
	c.room.messages <- []byte(myString)

	//set the client active in the room
	c.room.joining <- c
}

func (c *client) leaveRoom(r *room) {
	r.leaving <- c
}

//read from client network connection, and put the data on the messages channel to the room
func (c *client) readFromNet(conn net.Conn, r *room) {
	defer conn.Close()
	for {
		data := make([]byte, 1024) //Will let the Read method read 1024

		_, err := conn.Read(data)
		if err != nil {
			log.Println("Error client.readFromNet:", err)
			break
		}

		r.messages <- []byte(data)

	}

	s := fmt.Sprintf("client%v left the room !", c.ID)
	r.messages <- []byte(s)
}

func (c *client) writeToNet(conn net.Conn, r *room) {
	//check if something is received on the client.msg channel, and write it to the telnet session
	defer conn.Close()
	for {
		for m := range c.msg {
			m2 := fmt.Sprintf("client%v received msg: %v\n", c.ID, string(m))
			_, err := conn.Write([]byte(m2))
			if err != nil {
				fmt.Printf("---> error: client%v.writeToNet: %v\n", c.ID, err)
				//c.leaveRoom(r)
			}
		}
	}
}

func handleClient(conn net.Conn, cID int, r *room) {
	client := newClient(cID)
	client.joinRoom(r)

	go client.readFromNet(conn, r)

	go client.writeToNet(conn, r)
}

//create a new client with unique ID
func newClient(id int) *client {
	return &client{
		ID:  id,
		msg: make(chan []byte),
	}
}

func main() {
	//create a default room for clients to connect to
	room1 := newRoom(1)
	go room1.run()

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
