package main

import "fmt"

type tractor struct {
	name string
}

func (t tractor) String() string {
	return fmt.Sprintf("You're working on a %v", t.name)
}

func main() {
	myTractor := tractor{
		name: "John Deere",
	}

	fmt.Println(myTractor)
}
