package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	ice       []string
}

func main() {
	p1 := person{
		firstName: "Donald",
		lastName:  "Duck",
		ice:       []string{"båtis", "kroneis"},
	}

	p2 := person{
		firstName: "Mikke",
		lastName:  "Mus",
		ice:       []string{"lollipop", "løveis"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.ice {
		fmt.Println("is nr. ", i, " er ", v)
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p2.ice {
		fmt.Println("is nr. ", i, " er ", v)
	}
}
