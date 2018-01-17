package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

func changeMe(p *name) {
	p.firstName = "Scrooge"
}

func main() {
	p1 := &name{
		firstName: "Donald",
		lastName:  "Duck",
	}

	p2 := name{
		firstName: "Dolly",
		lastName:  "Duck",
	}

	fmt.Printf("%v , %T\n", p1, p1)
	fmt.Printf("%v , %T\n", p1.firstName, p1.firstName)
	fmt.Printf("%v , %T\n", p2, p2)
	fmt.Printf("%v , %T\n", p2.firstName, p2.firstName)

	changeMe(p1)
	fmt.Println("After changeme the firsName = ", p1.firstName)
}
