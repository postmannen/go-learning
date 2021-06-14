package main

import "fmt"

type driver interface {
	driving()
}

type fueler interface {
	fuelType()
}

type car struct {
	fuel string
}

func (c car) driving() {
	fmt.Println("Driving car")
}

func (c car) fuelType() {
	fmt.Println("Filling ", c.fuel)
}

type mc struct {
	fuel      string
	twoStroke bool
}

func (m mc) driving() {
	fmt.Println("Driving mc")
}

func (m mc) fuelType() {
	fmt.Println("Filling ", m.fuel)
	if m.twoStroke {
		fmt.Println("	You also have to add 2 percent two stroke oil in the gazoline")
	}
}

func drive(d driver) {
	d.driving()
}

func fillFuel(f fueler) {
	f.fuelType()
}

var opel = car{fuel: "diesel"}
var yamaha = mc{fuel: "gazoline", twoStroke: true}
var honda = mc{fuel: "gazoline", twoStroke: false}

var johnDeere fueler

func main() {
	drive(opel)
	drive(yamaha)

	fillFuel(opel)
	fillFuel(yamaha)
	fillFuel(honda)

}
