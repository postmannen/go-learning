package main

import (
	"log"

	"github.com/gorilla/websocket"
)

//client represents a single chatting user
type client struct {
	//socket is the websocket for this client
	socket *websocket.Conn
	//send is a channel which messages are sent.
	send chan []byte
	//room is the room this client is chatting in.
	room *room
}

/*
The read method allows our client to read from the socket via the ReadMessage method,
and continually sending any received messages to the forward channel on the room type
*/
func (c *client) read() {
	defer c.socket.Close()
	for {
		//Continuosly read from the client socket, and put whats read on the
		//rooms forward channel. That is the reason *room is defined in the client
		//struct, so we can directly alter the channel field defined in the room struct.
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Println("Error : client.Read()) : ", err)
			return
		}
		c.room.forward <- msg
	}
}

//Here we will write messages back to the websocket
func (c *client) write() {
	defer c.socket.Close()
	//range over the bytes in the c.send channel
	for msg := range c.send {
		//put the read byte on the socket
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error : client.Write() : ", err)
			return
		}
	}
}
