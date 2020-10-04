/*
	Higher order functions.
	Using functions as input parameters to choose the
	filter we want to apply.
*/
package main

import "fmt"

type vehicle struct {
	name   string
	kind   string
	energy string
}

// checkVehicle will execute some function taken as an
// input parameter to do the filtering on the slice of
// vehicles also given as input.
// The checkVehicle do not know anything about how to
// do the filtering, all it know is that is shall execute
// a function `func(vehicle) bool` that takes a single
// vehicle and returns true/false which acts as the filter
// to check if a specified criteria was met or not.
func checkVehicle(v []vehicle, f func(vehicle) bool) []vehicle {
	// prepare a variable to store all the vehicles where
	// the filter criteria was met.
	var result []vehicle

	// range over all the vehicles given as input.
	for _, v := range v {
		// execute the filter function, and if the return
		// value is true, append the vehicle to the result
		// slice
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	vehicles := []vehicle{
		{name: "Toyota", kind: "car", energy: "diesel"},
		{name: "Nissan", kind: "car", energy: "diesel"},
		{name: "Valtra", kind: "tractor", energy: "diesel"},
		{name: "jaguar", kind: "car", energy: "electric"},
	}

	// prepare a filter function with the signature as
	// specified in the checkVehicle function,
	// `func(vehicle) bool`, where we check the energy
	// type of the vehicle, and return true if found.
	checkEnergy := func(v vehicle) bool {
		if v.energy == "diesel" {
			return true
		}
		return false
	}

	{
		// execute the check function with the filter function
		// as the last argument.
		vs := checkVehicle(vehicles, checkEnergy)
		fmt.Printf("%v\n", vs)
	}

	// prepare another filter using the same function
	// signature, to check if there are any tractors.
	// Since the `==` returns true if match we can return
	// the `v.kind == "tractor"` immediately since it will
	// return true if found and false otherwise, and drop
	// the not needed if statement.
	checkTractor := func(v vehicle) bool {
		return v.kind == "tractor"
	}

	{
		// execute the check function with the filter function
		// as the last argument.
		vs := checkVehicle(vehicles, checkTractor)
		fmt.Printf("%v\n", vs)
	}
}
