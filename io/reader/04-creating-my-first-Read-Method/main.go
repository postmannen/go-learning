/*
The purpose of this program is to create a new type called storage, which can hold some data of type string
The storage type shall have a read method, to read the content of the storage into a variable.
There shall also be a function to create and initialize a new storage (will be done in next version)
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
	/*
		check if there are more to read after current position, if there's nothing more to read we're done.
		Return from function with err value = io.EOF
	*/
	if m.currentPosition >= len(m.data) {
		return 0, io.EOF
	}

	fmt.Println("Read Info: remaining bytes in storage = ", len(m.data)-m.currentPosition)

	/*
		Max position in the slice index it is allowed to read. Will use the capacity if input var 'p' to determine how many bytes to read.
		Check if max position is greater than the length of the string to read. If its greater it will try to read outside the slice, and panic.
		Also, if it is greater it means it will be the last read attempt, and we can set the maxPosition to the length of the string we're reading from.
	*/
	maxPosition := m.currentPosition + cap(p)
	if maxPosition >= len(m.data) {
		maxPosition = len(m.data)
	}
	fmt.Println("Read Info: maxPosition = ", maxPosition)

	/*
		Read the next bytes determined by the capacity of the input slice 'p'.
		Will read on byte at a time, append the read byte to 'tmpRead', over and over again until the capacity of input byte slice is met.
		When all reading is done we will copy the content of the temporary slice into p, which will be returned to main.
	*/
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

	//update the current position variable in the struct to reflect the new value of how many bytes in total that are read.
	m.currentPosition = newPosition

	//return how many bytes are read, and error = nil
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
