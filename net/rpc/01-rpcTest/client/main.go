package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//Person is just a type for testing, and is the actual data we send over to the server.
// The type must be the same as used on the server side.
type Person struct {
	Name string
}

func main() {
	p := Person{Name: "Apekatt"}
	var reply string

	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Println("error: failed to DialHTTP: ", err)
	}

	//client.Call will need the method call in it' full length as it is on the server side, including
	// the type the method is attached to.
	// The second argument are the data to send, and the third argument are the data received from
	// the server.
	err = client.Call("PersonRPC.Print", &p, &reply)
	if err != nil {
		log.Println("error: failed to client.Call: ", err)
	}
	fmt.Println("* The response :", reply)

}
