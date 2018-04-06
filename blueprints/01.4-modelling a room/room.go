package main

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
			//joining
			r.clients[client] = true
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			//if we receive a message on the room forward channel, we will
			//iterate over all the clients, and add the message on all the
			//clients send channel.
			//Then the write method of the client will pick it up and send
			//it down the socket to the browser.
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}
