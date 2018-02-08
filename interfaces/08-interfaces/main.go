package main

import "fmt"

type fuelTyper interface {
	fuelTypePrinter()
}

type car struct {
	name string
	fuel string
}

func (f car) fuelTypePrinter() {
	fmt.Printf("%v rund on %v\n", f.name, f.fuel)
}

type tractor struct {
	name string
	fuel string
}

func (t tractor) fuelTypePrinter() {
	fmt.Printf("%v runs on %v\n", t.name, t.fuel)
}

type vehicle struct {
	car
	tractor
}

func runsOn(f fuelTyper) {
	f.fuelTypePrinter()
}

func main() {
	car1 := car{
		name: "opel",
		fuel: "electric",
	}
	runsOn(car1)
	car2 := car{
		name: "nissan",
		fuel: "diesel",
	}
	runsOn(car2)

	tractor1 := tractor{
		name: "john deere",
		fuel: "diesel",
	}
	runsOn(tractor1)

	vehicle1 := vehicle{
		car: car{
			name: "fiat",
			fuel: "gazoline",
		},
		tractor: tractor{
			name: "fendt",
			fuel: "diesel",
		},
	}
	fmt.Println(vehicle1)
	vehicle1.car.fuelTypePrinter()
	runsOn(vehicle1.car)
}
