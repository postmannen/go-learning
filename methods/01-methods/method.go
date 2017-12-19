package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p person) speak() {
	fmt.Println("This persons name is : ", p.first, p.last)
}

func main() {
	person1 := person{
		first: "Donald",
		last:  "Duck",
	}

	person1.speak()

}
