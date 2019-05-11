/*
	encoding/gob testing
	Encode a struct into gob, put it on a buffer of type bytes.Buffer, then decode it.

*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

//Person describes a person
type Person struct {
	FirstName string
	LastName  string
}

func main() {
	persons := []Person{
		{
			FirstName: "Arne",
			LastName:  "Arnesen",
		},
		{
			FirstName: "Knut",
			LastName:  "Knutsen",
		},
	}

	// A buffer to encode into
	var b bytes.Buffer
	//NewEncoder. Since bytes.Buffer's Write method got a pointer receiver we have to pass a pointer
	// to the bytes buffer with &b
	e := gob.NewEncoder(&b)
	//Encode what is in persons into &b
	err := e.Encode(persons)
	if err != nil {
		fmt.Printf("encoding failed : %v\n", err)
	}

	fmt.Println(b.String())

	animals := []Person{}
	//NewDecoder. Since bytes.Buffer's read method got a pointer receiver we have to pass a pointer
	// to the bytes buffer with &b
	d := gob.NewDecoder(&b)
	//Decode what is in the buffer b into &animals
	err = d.Decode(&animals)
	if err != nil {
		fmt.Printf("decoding failed : %v\n", err)
	}
	fmt.Printf("%#v\n", animals)
}
