package main

import (
	"fmt"
	"log"
)

// node, some node
type node struct {
	id      int
	subject string
}

// do will just return an error value so we have something
// to do some checking on.
func (n node) do() error {
	return errNode{
		node: node{
			id:      10,
			subject: "some subject",
		},
	}
}

// This is our custom error type.
// errNode will promote the inner fields of a node by just
// specifying the value as we've done here.
// Then when we reference the error values later we can just
// do e.id instead of e.node.id
type errNode struct {
	node
}

// Error, by assigning an error method which returns a string
// we also become a type of the Error interface type.
// This method is also what decides how the output should look
// like when printing the error out.
func (e errNode) Error() string {
	return fmt.Sprintf("problem with node id=%v, subject=%v", e.id, e.subject)
}

func main() {
	// create a node
	n := node{
		id:      10,
		subject: "some subject",
	}

	// do something with the node so we get an error
	err := n.do()
	if err != nil {
		// First we just print the error
		log.Printf("error: %v, and type=%T\n", err, err)

		// Then we do a type check on the error, and if
		// it is we print out the struct values of that
		// error type.
		if e := err; e == err.(errNode) {
			fmt.Printf("the whole structure: %#v\n", e)
		}

	}
}
