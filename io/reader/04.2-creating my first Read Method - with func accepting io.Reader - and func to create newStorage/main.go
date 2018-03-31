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

/*
This function will take a type string as input.
It will take that string, put it into the 'data' field of the myStorage struct,
and return a pointer to that new struct with the 'data' variable set to the value of the string given as input with sData.
The returning variable will be of type *myStorage, which contains the field data set to the content of sData.
*/
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
		m.currentPosition = 0 //resets position when all reading is done
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
	/*
		TEST1 :
		Create a new variable of the type myStorage manually, and fill it with data.
		Use myStorage's Read method to read the data stored in myStorage into a new variable.
	*/
	fmt.Println("-----------------------------TEST1-----------------------------")
	disk1 := myStorage{}
	disk1.data = "abcdefghjiklmnopqrstuvwxyz"

	newData := []byte{}

	for {
		readData := make([]byte, 4)
		n, err := disk1.Read(readData)
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
	fmt.Printf("The type of the variable 'disk1' = %T\n", disk1)
	fmt.Println("The data read from storage with Read method : ")
	fmt.Println(string(newData))
	fmt.Println("---------------------------------------------------------------")

	/*
		TEST2 :
		Create a new variable of the type myStorage with the newMyStorage function, and fill it with data when calling the function.
		This will return a *myStorage containing the data its filled with. We can then assign this pointer to a variable,
		and use it like any other variable.
		Use myStorage's Read method to read the data stored in myStorage into a new variable.
	*/
	fmt.Println("-----------------------------TEST2-----------------------------")
	disk2 := newMyStorage("AbCdEfGhIjKlMnOpQrStUvWxYz")

	newData2 := []byte{}

	for {
		readData := make([]byte, 4)
		n, err := disk2.Read(readData)
		if err != nil {
			if err == io.EOF {
				fmt.Println("err: Reached end of string, leaving for loop :", err)
				break
			}
			fmt.Println("error: trying to use the read method from main : ", err)
			break
		}
		newData2 = append(newData2, readData...)
		fmt.Println("main info: bytes read = ", n)
	}

	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("The type of the variable 'disk2' = %T\n", disk2)
	fmt.Println("The data read from storage with Read method : ")
	fmt.Println(string(newData2))
	fmt.Println("---------------------------------------------------------------")

	fmt.Println("-----------------------------TEST3-----------------------------")

	/*
		TEST3 :
		Since we've created a Read method for 'myStorage' with the same input argument, and return values as the Read method defined in
		the io.Reader interface 'Read(p []byte)(n int, err error)', then myStorage becomes accepted as a valid type for the io.Reader interface which also
		is a type.
		Since io.Reader now accepts the type myStorage we can create a new readStorage function with the same code used in the Read loops in TEST1 & TEST2,
		but instead of using input type myStorage for the new readStorage function we can use input type io.Reader, since myStorage now is also an io.Reader.

		When we call readStorage which takes an oi.Reader, and pass it a type of 'myStorage' as input, it will use the Read method that is defined
		with 'myStorage'.
		If we for example called readStorage with a 'myBrandNewOtherType' or 'http' as input, that would also work, as long as they have a Read method.
		If we used a type http with readStorage, then http's Read method found in the http package would be used.
	*/
	readStorage(disk2)

}

func readStorage(p io.Reader) {
	newData := []byte{}

	for {
		readData := make([]byte, 4)
		n, err := p.Read(readData)
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
	fmt.Printf("The type of the variable 'p' = %T\n", p)
	fmt.Println("The data read from storage with Read method : ")
	fmt.Println(string(newData))
	fmt.Println("---------------------------------------------------------------")
}
