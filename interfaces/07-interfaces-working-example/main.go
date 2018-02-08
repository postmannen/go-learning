package main

import "fmt"

type speaker interface {
	speak()
}

type human struct {
	name string
	say  string
}

type animal struct {
	name string
	say  string
}

func (h human) speak() {
	fmt.Printf("Human %v says %v", h.name, h.say)
}

func (a animal) speak() {
	fmt.Printf("Animal %v says %v", a.name, a.say)
}

//saySomething uses the interface type speaker as input,
//and speaker is interface for the 'speak' method in both human and animal
func saySomething(n speaker) {
	n.speak()
}

func main() {
	human1 := human{
		name: "The King",
		say:  "Good evening",
	}
	animal1 := animal{
		name: "ape",
		say:  "oh oh oh",
	}

	saySomething(human1)
	saySomething(animal1)
}
