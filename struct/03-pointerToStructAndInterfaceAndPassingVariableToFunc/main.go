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

	p1.page1("-Pointer Value-")
	p2.page1("-Non Pointer Value-")
	p1.page2("-Pointer Value-")
	p2.page2("-Non Pointer Value-")

}

func (n *name) page1(text string) {
	fmt.Println("-------------------------")
	fmt.Printf("Func: Pointer Receiver: Content of 'n' = %v which was/is %v, and the type is %T\n", n, text, n)
	fmt.Printf("Func: Pointer Receiver: firstName = %v, lastName = %v, and the type is %T\n", n.firstName, n.lastName, n.firstName)
}

func (n name) page2(text string) {
	fmt.Println("-------------------------")
	fmt.Printf("Func: Non-Pointer Receiver: Content of 'n' = %v which was/is %v, and the type is %T\n", n, text, n)
	fmt.Printf("Func: Non-Pointer Receiver: firstName = %v, lastName = %v, and the type is %T\n", n.firstName, n.lastName, n.firstName)
}
