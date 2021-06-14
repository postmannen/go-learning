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

import "fmt"

// -------------------------------------------------------------------------
// Defining two types of data to receive.
// They are practically identical, except that one uses int as ID,
// and the other uses string as ID.

type obj1 struct {
	id   int
	data []byte
}

func (o obj1) read() (string, []byte) {
	fmt.Printf("Data received of type obj1: The object structure = %#v\n", o)
	return string(o.id), o.data
}

type obj2 struct {
	id   string
	data []byte
}

func (o obj2) read() (string, []byte) {
	fmt.Printf("Data received of type obj2 : The object structure = %#v\n", o)
	return o.id, o.data
}

//Define an interface that reflects the behavior of the obj methods.
type objReader interface {
	read() (string, []byte)
}

// -------------------------------------------------------------------------
// The database storage.

type dbContent struct {
	id   string
	data []byte
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
	myBackendDB := &db{}

	o1 := obj1{id: 101, data: []byte("Some data from obj1")}
	o2 := obj2{id: "1002", data: []byte("Some other data from obj2")}

	writeToDB(o1, myBackendDB)
	writeToDB(o2, myBackendDB)

	fmt.Println("---------------------------------------------------------")
	for i, v := range myBackendDB.records {
		fmt.Printf("The content of backendDB : index = [%v], value = %v\n", i, v)
		fmt.Printf("The string value of v.data at index [%v] = %v\n\n", i, string(v.data))
	}
	fmt.Println("---------------------------------------------------------")
}
