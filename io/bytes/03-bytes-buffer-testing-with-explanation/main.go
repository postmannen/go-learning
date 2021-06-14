package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	byteSlice := []byte{'a', 'b', 'c'}

	/*
		bytes.Buffer is just a struct, that contains a type called 'buf' which is a []byte. (it also contains a few more..),
		so the data is stored as a []byte inside variable of type buffer.
		Since 'buf' is not starting with a capital letter, it is not exported outside the bytes package. It is only the methods of
		Buffer that is exported, so the user can not access buf directly, but have to use Buffer's methods to access the data.
		Buffer have a Read and a Write method (among others), that can be used to Read and Write 'buf'
	*/
	var bb bytes.Buffer

	//Write will take a []byte and write it into the buffer.
	n, err := bb.Write(byteSlice)
	if err != nil {
		fmt.Println("error: bb.Read: ", err)
	}
	fmt.Println("bytes written = ", n)

	//TEST11
	fmt.Println("-----------------------TEST1---------------------------")
	fmt.Printf("bb is of type %T, and contains : \n %v\n", bb, bb)

	/*
		TEST2
		Adding more bytes to the buffer with Write will append the extra bytes to the existing content of the buffer
	*/
	fmt.Println("-----------------------TEST2---------------------------")
	n, err = bb.Write([]byte(",adding 1st dataset to buffer"))
	if err != nil {
		fmt.Println("error: bb.Read: ", err)
	}
	fmt.Println("bytes written = ", n)

	bb.Write([]byte(", adding 2nd dataset to buffer"))
	_, err = bb.WriteTo(os.Stdout)
	fmt.Println()
	if err != nil {
		fmt.Printf("error: WriteTo : %v\n", err)
	}

	/*
		TEST3
		This one fails.
		Reason is that the buffer is now empty after calling the bb.WriteTo method in test2.
	*/
	//create a slice of bytes, with a capacity of 4 bytes to use for reading a chunc of 8 bytes from the buffer
	fmt.Println("-----------------------TEST3---------------------------")
	buf := make([]byte, 4)
	_, err = bb.Read(buf)
	if err != nil {
		fmt.Println("Error: bb.Read :", err)
	}
	fmt.Println("The data read from buffer = ", string(buf))
	fmt.Print("And the buffer bb now contains = ")
	//WriteTo takes an io.Writer as input
	bb.WriteTo(os.Stdout)
	fmt.Println()

	/*
		TEST4
		This one works, since the bytes.Buffer now contains some data to read
		The read method reads 4 bytes from the beginning of the buffer, so it is FIFO.
	*/
	fmt.Println("-----------------------TEST4---------------------------")
	bb.Write([]byte("Putting some new data into the buffer"))
	//again we try to read a chunk of 4 bytes from the buffer
	n, err = bb.Read(buf)
	if err != nil {
		fmt.Println("Error: bb.Read :", err)
	}

	fmt.Println("The data read from buffer = ", string(buf))
	fmt.Print("And the buffer bb now contains = ")
	//WriteTo takes an io.Writer as input
	bb.WriteTo(os.Stdout)
	fmt.Println()

	/*
		TEST5
		Choose the number of bytes to read with Next
	*/
	fmt.Println("-----------------------TEST5---------------------------")
	bb.Write([]byte("Fill the buffer with some new data again"))

	bytesToRead := 7
	bytesRead := bb.Next(bytesToRead)
	fmt.Println("The bytesToRead variable contains = ", string(bytesRead))

	fmt.Print("And the buffer bb now contains = ")
	//WriteTo takes an io.Writer as input
	bb.WriteTo(os.Stdout)
	fmt.Println()

	/*
		This shows that calling any Read or WriteTo method will remove the bytes that the method processes from the buffer.

		In the above code we are using bb.WriteTo(os.StdOut)
		The bytes.WriteTo method accepts an io.Writer as input parameter.
		os.StdOut, StdIn and StdErr are open files pointing to stdout/stdin/stderr. We now know that they are of type File, and if
		we look in pkg/os of the std library, we can see that they have a Read and Write method. Since they have a Read and Write
		method they fullfill the requirement of the io.Reader and io.Writer interface.
		Since the type os.StdOut fullfills the inteface  for io.Writer, it is also an io.Writer, and can be used with the WriteTo method.
	*/
}
