package main

import (
	"fmt"
)

type name struct {
	firstName string
	lastName  string
}

func changeMe(p *name) {
	p.firstName = "Scrooge"
}

func main() {
	//with pointer to struct
	p1 := &name{
		firstName: "Donald",
		lastName:  "Duck",
	}

	//standard
	p2 := name{
		firstName: "Dolly",
		lastName:  "Duck",
	}

	changeMe(p1)
	fmt.Println("After changeme the firsName = ", p1.firstName)
	fmt.Println("----Pointer reciever---------------------------------------")
	p1.page1()
	p2.page1()
	fmt.Println("----Non Pointer reciever-----------------------------------")
	p1.page1()
	p2.page2()
	fmt.Println("-----------------------------------------------------------")

}

func (n *name) page1() {
	fmt.Println("Content of 'n' = ", n)
	fmt.Printf("firstName = %v, lastName = %v\n", n.firstName, n.lastName)
}

func (n name) page2() {
	fmt.Println("Content of 'n' = ", n)
	fmt.Printf("firstName = %v, lastName = %v\n", n.firstName, n.lastName)
}
