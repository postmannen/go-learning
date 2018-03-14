package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]string)

	map1["ku"] = "dagros"
	map1["hund"] = "rufus"
	fmt.Println(map1)

	delete(map1, "ku")
	fmt.Println(map1)

	//comma ok, check if a key exist.
	if animal, ok := map1["ku"]; ok {
		fmt.Println("found ku, animal = ", animal)
	} else {
		fmt.Println("did not find ku, animal = ", animal)
	}
}
