package main

import (
	"fmt"
)

type driver interface {
	drive()
}

type tractor struct {
	name string
}

func (t tractor) drive() {
	fmt.Println("I'm driving a ", t.name, "tractor !")
}

type car struct {
	name string
}

func (c car) drive() {
	fmt.Println("I'm driving a ", c.name, "car !")
}

func driveSome(t driver) {
	t.drive()
}

func main() {
	tractor1 := tractor{
		name: "John Deere",
	}

	car1 := car{
		name: "Nissan",
	}

	driveSome(tractor1)
	driveSome(car1)

}
