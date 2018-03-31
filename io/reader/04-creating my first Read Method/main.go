/*
The purpose of this program is to create a new type called storage, which can hold some data of type string
The storage type shall have a read method, to read the content of the storage into a variable.
There shall also be a function to create and initialize a new storage.
*/
package main

import (
	"fmt"
	"io"
)

type myStorage struct {
	data            string //will simulate the data on storage
	currentPosition int    //keeps track of how current bytes been read in storage
}

//NOT IN USE YET, WILL BE USED IN PART 2 :-)
func newMyStorage(sData string) *myStorage {
	return &myStorage{data: sData}
}

//Read takes a []byte as input to store the data that is read into. Returns characters read, and error.
func (m *myStorage) Read(p []byte) (n int, err error) {
	//check if there are more to read after current position
	if m.currentPosition >= len(m.data) {
		return 0, io.EOF
	}

	fmt.Println("Read Info: remaining bytes in storage = ", len(m.data)-m.currentPosition)

	//max long in the slice index it is allowed to read
	maxPosition := m.currentPosition + cap(p)
	if maxPosition >= len(m.data) {
		maxPosition = len(m.data)
	}
	fmt.Println("Read Info: maxPosition = ", maxPosition)

	tmpRead := []byte{} //to store the bytes read
	oldPosition := m.currentPosition
	newPosition := oldPosition
	for newPosition < maxPosition {
		tmpRead = append(tmpRead, m.data[newPosition])
		//fmt.Print(string(m.data[newPosition]))
		newPosition++
	}
	//fmt.Println()
	copy(p, tmpRead)

	m.currentPosition = newPosition

	return newPosition - oldPosition, nil
}

func main() {
	disk := myStorage{}
	disk.data = "abcdefghjiklmnopqrstuvwxyz"

	newData := []byte{}

	for {
		readData := make([]byte, 4)
		n, err := disk.Read(readData)
		if err != nil {
			if err == io.EOF {
				fmt.Println("err: Reached end of string, leaving for loop :", err)
				break
			}
			fmt.Println("error: trying to use the read method from main : ", err)
			break
		}
		newData = append(newData, readData...)
		fmt.Println("main info: bytes read = ", n)
	}

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("The data read from storage with Read method : ")
	fmt.Println(string(newData))
	fmt.Println("---------------------------------------------------------------")

}
