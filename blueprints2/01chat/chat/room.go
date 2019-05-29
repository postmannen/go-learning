package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/postmannen/go-learning/blueprints2/01chat/trace"
)

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte
	// join is a channel for clients wishing to join the room.
	join chan *client
	// leave is a channel for clients wishing to leave the room.
	leave chan *client
	// clients holds all current clients in this room.
	clients map[*client]bool
	// tracer will receive trace information of activity
	// in the room.
	tracer trace.Tracer
}

// newRoom makes a new room.
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			r.tracer.Trace("Message received : ", string(msg))
			//forward message to all clients
			for client := range r.clients {
				//since client in the map is a pointer to a *client, we get's all it's methods
				// to use here.
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

const sockeBufferSize = 1024
const messageBufferSize = 256

var upgrader = &websocket.Upgrader{ReadBufferSize: sockeBufferSize, WriteBufferSize: messageBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("error: websocket upgrade: ", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}

	//will add another client to the client map
	r.join <- client
	defer func() {
		r.leave <- client
	}()

	go client.write()
	//The code will block here since client.read() contains a for loop.
	client.read()
}
