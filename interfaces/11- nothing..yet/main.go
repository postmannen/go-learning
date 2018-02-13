package main

import (
	"fmt"
)

type printer interface {
	print()
}

type vehicle struct {
	VehicleName string
}

func (v vehicle) print() {
	fmt.Println("VehicleName = ", v.VehicleName)
}

//new variable of type interface
var johnDeere printer

func main() {
	johnDeere.print()

}
