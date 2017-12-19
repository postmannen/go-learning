package main

import "fmt"

func main() {
	mom := map[string]map[string]string{
		"no": map[string]string{},
		"jp": map[string]string{"one": "ichi", "two": "ni"},
		"de": map[string]string{"one": "eins", "three": "drei"}}

	//Add a map
	mom["fr"] = map[string]string{"one": "un", "four": "quatre"}

	//Add an entry to one of the nested maps. The firs key must exist to add more data.
	mom["jp"]["eight"] = "hachi"
	mom["no"]["one"] = "en"

	//Delete an entry in one of the nested maps.
	delete(mom["jp"], "one")

	//Delete one of the nested maps
	delete(mom, "de")

	//Print all elements
	for i := range mom {
		fmt.Printf("(%s) %v\n", i, mom[i])
		for j := range mom[i] {
			fmt.Printf(" (%s) %s => %s\n", i, j, mom[i][j])
		}
	}

	fmt.Printf("Entire DSC: %v\n", mom)

}
