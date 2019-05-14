package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//Person is just a type for testing
type Person struct {
	Name string
}

type PersonRPC int

func (m PersonRPC) Print(v Person, status *string) error {
	fmt.Println("** Printing : ", v.Name)
	*status = "ok"
	return nil
}

func main() {
	var err error
	person := new(PersonRPC)
	err = rpc.Register(person)
	if err != nil {
		log.Println("error: failed to register rpc: ", err)
	}
	rpc.HandleHTTP()
	nListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("error: failed to start listener: ", err)
	}

	err = http.Serve(nListener, nil)
	if err != nil {
		log.Println("error: failed to serve http: ", err)
	}
}
