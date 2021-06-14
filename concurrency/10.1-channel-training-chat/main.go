package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type room struct {
	ID      int
	msg     chan []byte      //the active chat messages for all in room
	buf     bytes.Buffer     //buffer of recent messages
	clients map[*client]bool //true if client is in room
	joining chan *client
	leaving chan *client
	command chan []byte //user for sending command action messages to server
}

//create a new room
func newRoom(id int) *room {
	return &room{
		ID:      id,
		clients: make(map[*client]bool),
		msg:     make(chan []byte, 10), //initialize channel, or...deadlock
		joining: make(chan *client),    //initialize channel, or...deadlock
		leaving: make(chan *client),    //initialize channel, or...deadlock
	}
}

//The room will allways be up, and do things based on what is received on its channels
func (ro *room) run() {
	log.Println("Starting up the room with ID = ", ro.ID)

	for {
		select {
		//any new incomming messages to the room ? if so...send them out on each client.msg channel to
		//be handlet by the client methods
		case msg := <-ro.msg:
			log.Printf("room%v: %v\n", ro.ID, string(msg))
			for k := range ro.clients {
				log.Println("Active clients to get message = ", k)
				k.msg <- msg
			}

		//create a reference of the client in the room struct, and set its value to true
		//to indicate that this client is in the room
		case c := <-ro.joining:
			ro.clients[c] = true
			c.msg <- []byte("Welcome to the room !\n")
			//TODO: make the client tell the room it has left, so client is removed from the room
		case l := <-ro.leaving:
			log.Printf("room: client%v leaving room\n", l.ID)
			ro.msg <- []byte(fmt.Sprintf("room: client%v leaving room\n", l.ID))
			delete(ro.clients, l)
		}
	}

}

type client struct {
	ID   int         //unique ID of client
	room *room       //room that client belongs to
	msg  chan []byte //the channel to send a message directly to a client
	exit chan bool   //for telling if a client has left
	conn net.Conn
}

//create a new client with unique ID
func newClient(id int, con net.Conn) *client {
	return &client{
		ID:   id,
		msg:  make(chan []byte),
		exit: make(chan bool),
		conn: con,
	}
}

//attach a room given as input to the client
func (c *client) joinRoom(ro *room) {
	c.room = ro
	log.Printf("joinRoom: client1.ID =%v, is now in the room client.room.ID = %v\n", c.ID, c.room.ID)
	c.room.msg <- []byte(fmt.Sprintf("Hello, I'm client%v, and entering the room\n", c.ID))

	//set the client active in the room
	c.room.joining <- c
}

//check the client channels
func (c *client) checkChannels() {
	for {
		select {
		//if message received, write it to the telnet session
		case msg := <-c.msg:
			//fmt.Println("content of msg", msg)
			log.Printf("client%v direct message: %v\n", c.ID, string(msg))
			//write the client message to the telnet session
			c.conn.Write(msg)
		case <-c.exit:
			log.Printf("leaving checkChannels for client goRoutine, client%v disconnected\n", c.ID)
			break
		}
	}
}

//Read the the telnet session, and exit go-routine if error.
//Check for ascii value 4 (EOT=end of transmission) since it will indicate that the client connection was lost.
func (c *client) handleTelnet() {
	defer c.conn.Close()
	for {
		b := make([]byte, 256)
		_, err := c.conn.Read(b)
		if err != nil {
			log.Println("error: handleTelnet read:", err)
			c.room.leaving <- c
			break
		} else if b[0] == 4 {
			//above we check for ascii value 4 (EOT), since it will tell if the client session is lost
			c.exit <- true
			//print error message to room, and leave handleTelnet go routine
			c.room.msg <- []byte(fmt.Sprintf("client%v unexpectedly lost connection\n", c.ID))
			c.room.leaving <- c
			break
		} else {
			//if all is OK above, send the message to the room, and then the room will send out to all clients
			b := []byte(fmt.Sprintf("client%v: %v", c.ID, string(b)))
			c.room.msg <- b
		}

	}
	log.Printf("leaving handleTelnet for client%v\n", c.ID)
	//return err
}

var clientID = 1

func main() {
	addr := flag.String("addr", ":8000", "<address:port>")
	flag.Parse()
	fmt.Println(*addr)

	room1 := newRoom(1)
	go room1.run()
	//time.Sleep(time.Millisecond * 50) //let the room fully start before starting clients, will be removed later.

	//start telnet server
	server, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Println("Failed starting net listen:", err)
		os.Exit(1)
	}
	defer server.Close()

	//wait for telnet connection, create new client, join default room, and start the client.
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Failed net accept:", err)
		}

		client1 := newClient(clientID, conn)
		client1.joinRoom(room1)
		go client1.checkChannels()
		go client1.handleTelnet()

		clientID++
	}

}
