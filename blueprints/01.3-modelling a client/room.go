package main

type room struct {
	//forward is a channel that holds the incomming messages
	//that should be forwarded to other clients
	forward chan []byte
}
