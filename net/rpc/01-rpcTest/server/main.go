/*
Test for learning RPC.
RPC uses GOB as the binary format by default when sending the data over the network.
*/
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

//PersonRPC is a type to tie the RPC exposed methods to.
// Not sure if it is correct to make it's own type for the RPC exposed methods, or if they should be
// placed on the data structure Person instead. Have to investigate further on that in another example.
type PersonRPC int

//Print
// An RPC method needs to meet to accept two arguments, and return an error.
// Argument 1, should be the data to receive
// Argument 2, will be the data to return back, in this example a reply code in the form of "ok".
// If the method call fails, an error will be returned, and the argument 2 will be set to zero/0.
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
