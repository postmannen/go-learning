package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

/*
u8 1 unsigned 8bit value
i8 1 signed 8bit value
u16 2 unsigned 16bit value
i16 2 signed 16bit value
u32 4 unsigned 32bit value
i32 4 signed 32bit value
u64 8 unsigned 64bit value
i64 8 signed 64bit value
float 4 IEEE-754 single precision
double 8 IEEE-754 double precision
string * Null terminated string (C-String)
(Variable size)
enum 4 Per command defined enum
*/

// int8Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type float32Type struct {
	value  float32
	length int
}

// i8 makes a type for int8 data, and tells the length of bytes it
// is made of.
var float = float32Type{
	length: 4,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *float32Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *float32Type) argDecode(b []byte) (err error) {
	fmt.Printf("running the float32.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val float32

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return err
}

// ------------------------------------------------------------------------------------

// int8Type makes a type for int8 data.
// The value is for storing the parsed value, length tells the length of bytes it
// is made of.
type int8Type struct {
	value  int8
	length int
}

var i8 = int8Type{
	length: 1,
}

// getLength will get the length value of the type. We need this method
// since we're working with interface type's in the final iteration, and
// interface types is all about methods, and not concrete values stored
// in a struct. Thats why we use a method to get that value from the struct.
func (f *int8Type) getLength() int {
	return f.length
}

// argDecode will decode the []byte given as input, and store it
// into f.
func (f *int8Type) argDecode(b []byte) (err error) {
	fmt.Printf("running the int8Type.argDecode method, b = %v\n", b)
	bReader := bytes.NewReader(b)
	var val int8

	err = binary.Read(bReader, binary.LittleEndian, &val)
	if err != nil {
		log.Println("error: failed binary.Read: ", err)
		return err
	}

	f.value = val

	fmt.Printf("Content of f = %#v\n", *f)

	return err
}

// ------------------------------------------------------------------------------------

// argDecoder is an interface type which tells that any type that
// have the methods argDecode([]byte) error, and getLength() int
// is of the interface type argDecoder.
type argDecoder interface {
	argDecode([]byte) error
	getLength() int
}

// argumentState is a type for keeping track of the start position of the
// data to read in a slice.
type argumentsState struct {
	position int
}

// argumentsToDecode takes a []byte and any number of the interface type argDecoder
// is input.
// It will loop through the argDecoder methods, and run the concrete types method,
// one by one until all methods are done.
// The method will use the getLength() method to know the size of the portion of data
// to work with, and increase the the position with the size of the last data read to
// know where the next piece of data starts.
// TODO: Make logic check if there are given the correct amount of argDecoders to
// handle the length of the data slice given as input, and return error if they don't
// match.
func (as *argumentsState) argumentsToDecode(d []byte, a ...argDecoder) {
	as.position = 0

	for _, v := range a {
		fmt.Println("------------------Decoding byte or bytes----------------------")
		err := v.argDecode(d[as.position : as.position+v.getLength()])
		if err != nil {
			fmt.Println("error: argumentsToDecode: failed looping over v ", err)
		}

		l := v.getLength()
		as.position += l
	}
}

func main() {
	//The data below should decode
	//bytes 1->4 = a float,
	//byte 5 = an int8
	tmpData := []byte{154, 221, 45, 61, 83}

	a := argumentsState{}
	a.argumentsToDecode(tmpData, &float, &i8)

	//fmt.Printf("%+v\n", a)
}

