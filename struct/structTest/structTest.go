package main

import "fmt"

type person struct {
	age    int
	name   string
	gender string
}

//Method av typen value reciever da den ikke kan endre noe som helst i structen
func (p person) run() {
	fmt.Println("The person", p.name, "is running")
}

//Method av typen pointer reciever, denne kan endre verdiene som er gitt via
// struct i og med at den er en peker (legg merke til bruk av *)
func (p *person) changeName(name string) string {
	p.name = name
	return "Har byttet navn"
}

func main() {
	var person1 person
	person1.name = "Bob"
	person1.run()

	//change the name of the person1
	fmt.Println(person1.changeName("Roger"))
	fmt.Println("The new name of person1 =", person1.name)
}
