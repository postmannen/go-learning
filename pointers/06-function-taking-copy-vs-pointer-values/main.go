package main

import (
	"fmt"
)

func changeOne(t string) {
	t = "pig"
}

func changeTwo(t *string) {
	*t = "pig"
}

func main() {
	animal := "horse"

	//This one will not change the value of animal, since Go is just passing a copy of the value to the function,
	//so the original animal variable is never changed

	changeOne(animal)
	fmt.Println(animal)

	//This will change the value of animal, since we're passing the adress of the variable containing the value of animal to the function,
	//so the value will be changed directly at its memory location from the function

	changeTwo(&animal)
	fmt.Println(animal)
}
