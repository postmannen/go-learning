/*
The idea here is to create buffer, and if the buffer
is saturated, then drop the request.
*/
package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	const readers = 10
	const cap = 2
	const requests = 100

	ch := make(chan string, cap)

	//Read the pretended requests comming on the channel
	for i := 1; i <= readers; i++ {
		go func() {
			for req := range ch {
				log.Println("Server: Received : ", req)
			}
		}()
	}

	//Create request:
	//The for loop simulates the clients creating the requests,
	//and the select statement is the logic on the server to
	//either accept or drop the request based on the buffer is
	//full or not.
	for i := 0; i < requests; i++ {
		select {
		case ch <- "new Request": //If there is room on the buffer ?
			log.Println("Client: sending req, buffer got room")
		default:
			log.Println("Client: Dropping request, buffers full")

		}
	}

}
