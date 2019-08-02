package main

import "fmt"

// walker is an interface type for any type with a walk(int) method.
type walker interface {
	walk(miles int)
}

type camel struct {
	Name string
}

//camel will have walk method.
func (c camel) walk(miles int) {
	fmt.Println(c.Name, "is walking ", miles)
}

// longwalk will accept any walker as input, an camel is a walker since
// it have a walk method.
func longWalk(w walker) {
	w.walk(500)
	w.walk(500)
}

func main() {
	c := camel{"Bill"}
	longWalk(c)
}
