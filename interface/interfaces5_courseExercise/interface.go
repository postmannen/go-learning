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
	area() float64 //husk at hvis en method har return value så må det spesifiseres også i interfacet.
}

func info(s shape) { //funksjon basert på interface 'shape' som tar både square og circle
	//Det gjør at den kan ta både square og circle type siden de begge har methoden area i interfacet shape.
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

	info(sirkel)  //funksjon basert på interface shape som tar både square og circle
	info(firkant) //funksjon basert på interface shape som tar både square og circle
}
