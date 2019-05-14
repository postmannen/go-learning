package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//Person is just a type for testing
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

	err = client.Call("PersonRPC.Print", &p, &reply)
	if err != nil {
		log.Println("error: failed to client.Call: ", err)
	}
	fmt.Println("* The response :", reply)

}
