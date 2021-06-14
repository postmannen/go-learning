package main

import (
	"encoding/json"
	"fmt"
)

type myS struct {
	Name   string
	Color  string
	Number int
}

func main() {
	cows := []myS{}
	moreCows := []myS{}

	cow1 := myS{
		Name:   "dagros",
		Color:  "brown",
		Number: 100,
	}

	cow2 := myS{
		Name:   "rosalin",
		Color:  "White",
		Number: 101,
	}

	fmt.Println("-----------------Testing marshalling--------------------")
	cow1Mars, err := json.Marshal(cow1)
	if err != nil {
		fmt.Println("Error: marshalling: ", err)
	}

	fmt.Printf("Data: %v, type: %T \n", string(cow1Mars), cow1Mars)

	cows = append(cows, cow1, cow2)
	cowsMars, err := json.Marshal(cows)
	if err != nil {
		fmt.Println("Error: marshalling: ", err)
	}
	fmt.Printf("Data: %v, type: %T \n", string(cowsMars), cowsMars)

	fmt.Println("-----------------Testing unmarshalling--------------------")
	err = json.Unmarshal(cowsMars, &moreCows)
	if err != nil {
		fmt.Println("Error: unmarshalling: ", err)
	}
	fmt.Println("The unmarshalled data: ", moreCows)
}
