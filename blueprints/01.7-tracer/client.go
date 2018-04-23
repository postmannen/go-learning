package main

import (
	"fmt"
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
The read method read's from the socket,
and continually sending any received messages to the room's forward channel
*/
func (c *client) readFromSocket() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Println("Error : client.Read()) : ", err)
			return
		}
		c.room.forward <- msg
	}
}

//write will write messages back to the websocket from the send channel.
//** The messages on the send channel have been put there by run() function **
func (c *client) writeToSocket() {
	defer c.socket.Close()
	//range over the bytes in the c.send channel
	for msg := range c.send {
		fmt.Print(string(msg))
		//put the read byte on the socket
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error : client.Write() : ", err)
			return
		}
	}
	fmt.Println()
}
