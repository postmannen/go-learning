package main

import "fmt"

func main() {

	x := map[string][]string{
		"Donald_Duck": []string{"Hus", "bil"},
		"Mikke_Mus":   []string{"Bøker", "fly"}, //slice blir 'value' av map (definert etter :)
	}

	//for å legge til en ekstra i map

	x["langbein"] = []string{"Helikopter", "Båter"}
	fmt.Println(x)

	for key1, value1 := range x {
		for index, value2 := range value1 {
			fmt.Println("key1 : ", key1, ", og value1 : ", value1)
			fmt.Println("	", index, value2)
		}

	}
	//To delete a key from the map
	// delete(x, "Donald_Duck")

}
