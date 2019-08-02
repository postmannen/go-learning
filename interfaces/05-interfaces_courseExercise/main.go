package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	// If a method got a return value that must
	// also be put into the definition of the interface.
	area() float64
}

// info, will takes a type shape as input. Both circle and square are a shape.
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	firkant := square{
		length: 10,
		width:  20,
	}

	sirkel := circle{
		radius: 10,
	}

	info(sirkel)
	info(firkant)
}
