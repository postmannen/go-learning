package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type person struct {
	FirstName string
	LastName  string
	HairColor string
	Age       int
}

func main() {
	p1 := person{
		FirstName: "Ole",
		LastName:  "Olsen",
		HairColor: "Brown",
		Age:       30,
	}

	var f *os.File

	_, err := os.Stat("test.txt")
	//check if file exist
	if os.IsNotExist(err) {
		//if not exist, create file
		f, err = os.Create("test.txt")
		if err != nil {
			fmt.Println("Error: Create File : ", err)
		}
	} else {
		//if exist, open file
		f, err = os.OpenFile("test.txt", os.O_RDWR, 0777)
		if err != nil {
			fmt.Println("Error: open File : ", err)
		}
	}

	marshalledText, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error: marshal", err)
	}

	//write the marshaled text to file
	_, err = f.Write(marshalledText)
	if err != nil {
		fmt.Println("error: file write", err)
	}
	f.Close()

	//read from file into a []byte
	fr, err := os.OpenFile("test.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("error: open file in read mode: ", err)
	}
	defer fr.Close()

	var b []byte
	b, err = ioutil.ReadAll(fr)
	if err != nil {
		fmt.Println("error: file read", err)
	}
	//fmt.Println("n = ", n)
	fmt.Println("Printing what's Read from file", string(b))

	var p2 person

	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println("error: Unmarshal : ", err)
	}

	fmt.Println("Printing the unmarshaled data : ", p2)
}
