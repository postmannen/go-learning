/*
 The purpose of this test is to receive data from two different systems
 where the two systems have data in two different formats. The difference
 are the ID field, where one uses the type string, and the other a type int.

 The data from both these systems are read by each type's read method.
 We then create an interface to represent the these two systems, and use that
 interface as input to our Database writer function, which now can handle
 input from them both.

 The datafrom both systems are then stored into the backend database.
*/

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// -------------------------------------------------------------------------
// Defining two types of data to receive.
// They are practically identical, except that one uses int as ID,
// and the other uses string as ID.

//obj1 some type with ID int
type obj1 struct {
	ID   int    `xml:"id"`
	Data []byte `xml:"data"`
}

//read will read the data, and return a ID=string, and Data=[]byte
func (o obj1) read() (string, []byte) {
	fmt.Printf("Data received of type obj1: The object structure = %#v\n", o)
	return string(o.ID), o.Data
}

//obj2 some type with ID string
type obj2 struct {
	ID   string `xml:"id"`
	Data []byte `xml:"data"`
}

//read will read the data, and return a ID=string, and Data=[]byte
func (o obj2) read() (string, []byte) {
	fmt.Printf("Data received of type obj2 : The object structure = %#v\n", o)
	return o.ID, o.Data
}

//obj3 some type with ID UUID
type obj3 struct {
	ID   uuid.UUID `xml:"id"`
	Data []byte    `xml:"data"`
}

//read will read the data, and return a ID=string, and Data=[]byte
func (o obj3) read() (string, []byte) {
	fmt.Printf("Data received of type obj2 : The object structure = %#v\n", o)
	return o.ID.String(), o.Data
}

//Define an interface that reflects the behavior of the obj methods.
type objReader interface {
	read() (string, []byte)
}

// -------------------------------------------------------------------------
// The database storage.

type dbContent struct {
	ID   string `xml:"id"`
	Data []byte `xml:"data"`
}

type db struct {
	records []dbContent
}

//writeToDB takes an objReader interface value, and a myData database as input.
func writeToDB(o objReader, d *db) {
	myID, myData := o.read()
	d.records = append(d.records, dbContent{myID, myData})
}

func main() {
	generateFiles := flag.String("generate", "no", "generate : yes/no")
	flag.Parse()

	if *generateFiles == "yes" {
		err := createFile()
		if err != nil {
			log.Printf("error: failed to create file : %v\n", err)
		}
	}

	myBackendDB := &db{}

	o1 := obj1{ID: 101, Data: []byte("Some data from obj1")}
	o2 := obj2{ID: "1002", Data: []byte("Some other data from obj2")}

	writeToDB(o1, myBackendDB)
	writeToDB(o2, myBackendDB)

	fmt.Println("---------------------------------------------------------")
	for i, v := range myBackendDB.records {
		fmt.Printf("The content of backendDB : index = [%v], value = %v\n", i, v)
		fmt.Printf("The string value of v.data at index [%v] = %v\n\n", i, string(v.Data))
	}
	fmt.Println("---------------------------------------------------------")

}
