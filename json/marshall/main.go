package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   40,
	}
	p2 := person{
		First: "Dr",
		Last:  "no",
		Age:   41,
	}

	pSlice1 := []person{p1, p2}
	fmt.Println(pSlice1)

	bs, err := json.Marshal(pSlice1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("byteSlice: ", string(bs))

	var pSlice2 person

	err = json.Unmarshal(bs, &pSlice2)
	if err != nil {
		log.Println("Error, Unmarshal: ", err)
	}
	fmt.Println("pSlice2 = ", pSlice2)
}
