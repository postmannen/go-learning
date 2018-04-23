package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	//forward is a channel that holds the incomming messages
	//that should be forwarded to other clients
	forward chan []byte
	//join is a channel for client who wants to join the room
	join chan *client
	//leave is a channel for client who want to leave the room
	leave chan *client
	//client is a map holds all current clients in this room
	clients map[*client]bool
}

func (r *room) run() {
	//This for loop will run forever
	//The select statements will watch the 3 channels, and do the case
	//defined if anything is received on that channel.
	//Only one block of case code will run at a time, and this is how
	//we make sure that only one client is modifying the map at a time.
	for {
		select {
		case client := <-r.join:
			//joining. Sets the map value for specific client to true.
			r.clients[client] = true
		case client := <-r.leave:
			//leaving. Deleting the specific client index from map the type room clients map.
			delete(r.clients, client)
			//closing the send channel for the specific client.
			//Here we are receiving a pointer to a client, so we can pick up
			//the client from the channel and close its channel directly based
			//only on the pointer to client we received on the channel.
			close(client.send)
		case msg := <-r.forward:
			//if we receive a message on the room forward channel, we will
			//iterate over all the clients, and add the message on all the
			//clients send channel.
			//Then the write method of the client will pick it up and send
			//it down the socket to the browser.
			//Send msg to all clients in the room.
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

//Creating variable of a pointer to struct &websocket.Upgrader
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

/*
creating a ServeHTTP method means the room now can act as a handler.
The ServeHTTP method will be called once each time a new client opens the web page,
and it will keep the page until the client leaves.
*/
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("The serveHTTP method for room was called")
	//upgrading the http connection to a websocket
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r, //here we assign a pointer to the existing room to the client variable.
	}

	//let the client join the room by passing it to the join channel of the room
	r.join <- client //defenition of r.join is 'join chan *client'
	defer func() { r.leave <- client }()
	go client.writeToSocket()
	//client.read() has a for loop, so it will read the socket continously.
	//since client.read has a for loop it will block the rest of the operations in this
	//function, and keep the connection alive for the client.
	client.readFromSocket()
}

// newRoom makes a new room.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}
