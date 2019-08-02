package main

import "fmt"

// An interface for all types that have a speak() method.
type thing interface {
	speak()
}

type person struct {
	firstName string
	lastName  string
}

type animal struct {
	firstName string
	age       int
}

// person will have a speak method.
func (s person) speak() {
	fmt.Println("Personens navn er : ", s.firstName, s.lastName)
}

// animal will have a speak method.
func (a animal) speak() {
	fmt.Println("The characteristics of the animal is : ", a.firstName, ", and age : ", a.age)
}

func main() {
	person1 := person{
		firstName: "Donald",
		lastName:  "Duck",
	}

	person2 := person{
		firstName: "Mikke",
		lastName:  "Mus",
	}

	person1.speak()
	person2.speak()

	var ix thing = person{
		"anne",
		"annesen",
	}

	ix.speak()
	fmt.Printf("The type of ix = %T\n", ix)
}
