package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName     string
	LastName      string
	myPrivateInfo string //net exported variables are not marshalled
}

func main() {
	//create a slice of struct person, and fill with values
	p := []person{}
	p = append(p, person{FirstName: "arne", LastName: "arnesen", myPrivateInfo: "Something super secret"})
	p = append(p, person{FirstName: "knut", LastName: "knutsen", myPrivateInfo: "Something even more secret"})

	fmt.Println("Printing slice of struct\n ----------------------\n", p, "\n----------------------")

	pM, err := json.Marshal(p)
	if err != nil {
		log.Println("Error: marshall : ", err)
	}

	fmt.Printf("Printing marshalled text %v, which is of type %T \n", pM, pM)
	fmt.Printf("Printing the stringified slice of bytes %v \n", string(pM))

}
