package main

import (
	"fmt"
)

func main() {
	dyr := []string{"ape", "hest", "ku"}

	dyrKommune := map[string][]string{
		"spydeberg": dyr,
		"skiptvet":  dyr,
	}

	dyrKommune["ski"] = []string{"høne", "hane", "gås"}

	fmt.Printf("%v\n", dyrKommune)

	for key, value := range dyrKommune {
		fmt.Printf("Kommune : %v, har dyrene %v\n", key, value)
	}

}
