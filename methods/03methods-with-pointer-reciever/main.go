package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

func (n name) print() {
	fmt.Println("The firstName is = ", n.firstName)
	fmt.Println("The lastName is = ", n.lastName)
}

func (n *name) changeFirst(firstN string) {
	fmt.Println("- Changing name from ", n.firstName, "to", firstN)
	n.firstName = firstN
}

func main() {
	p1 := name{
		firstName: "Donald",
		lastName:  "Duck",
	}

	p1.print()

	p1.changeFirst("Doffen")
	p1.print()

}
