package main

import (
	"fmt"
)

type printer interface {
	print()
}

type tractor struct {
	name string
}

func (t tractor) print() {
	fmt.Printf("name of tractor = %v, and the type = %T\n", t.name, t)
}

type car struct {
	name string
}

func (c car) print() {
	fmt.Printf("name of car = %v, and the type = %T\n", c.name, c)
}

func main() {

	cars := []car{
		car{
			name: "Opel",
		},
		{
			name: "Ford",
		},
	}
	fmt.Println(cars)

	tractors := []tractor{
		tractor{
			name: "John Deere",
		},
		{
			name: "Fendt",
		},
	}
	fmt.Println(tractors)

	//Since both tractor and care have a printer method, they satisfy the printer interface.
	//That means both tractor and car is of type printer, and printer can be used to specify both a tractor and a car
	vehicles1 := []printer{
		tractor{
			name: "Fiat",
		},
		tractor{
			name: "MF",
		},
		car{
			name: "Nissan",
		},
		car{
			name: "Lotus",
		},
	}
	fmt.Printf("var vehicles1 = %v, and the type is %T\n", vehicles1, vehicles1)

	//Since both car and tractor types are defined as a slice of 'printers',
	//we should now be able to range over the slice and run the print method on each of them
	for i, v := range vehicles1 {
		fmt.Println("Ranging over the vehicles slice which is of type []printer, and doing the interface print method on index ", i, " : ")
		v.print()
	}

	//testing if it is possible to append an existing tractor slice to and interface of type []printer
	//	var vehicles2 []printer
	//	vehicles2 = append(vehicles2, tractors...)
	//It does not work to append a slice of concrete type to an interface,
	//since the compiler will not create a loop to convert `[]ConcreteType` to `[]InterfaceType`.
	//The loop to separate out each index of the slice must be created manually,
	//then it should work to append to the interface
	//testing below

	var vehicles2 []printer

	//This works, since we break up the the slice, and append each slice element individually
	//explanation can be found here https://github.com/golang/go/wiki/InterfaceSlice
	for _, v := range tractors {
		vehicles2 = append(vehicles2, v)
	}
	fmt.Println(vehicles2)

}
