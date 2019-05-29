package main

import (
	"github.com/gorilla/websocket"
)

//client represents a single chatting user.
type client struct {
	//socket is the websocket for this client.
	socket *websocket.Conn
	//send is a channel where messages are sent to be forwarded to the users browser
	send chan []byte
	//room is the room this client is chatting in
	room *room
}

//read will read messages comming in on the socket,
// and send them to the room so all the clients get them.
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

//write will write messages to to the client,
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
