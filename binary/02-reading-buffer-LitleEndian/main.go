package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

/*
 binary read of 16bits a time, from a buffer of 8 bit values.
 This will read the 8 bit slices as they are, but will reverse the first and
 the second 8 bit values in the 16 bit uin16.
 For example. Reading the two values 1000 and 1111, will become 11111000 in the uint16 using
 LittleEndian while reading.
*/
func main() {
	// Create a buffer and put a byte in it
	data := []byte{0xF, 0x8, 0xF, 0x8}

	buf := bytes.Buffer{}
	n, err := buf.Write(data)
	if err != nil {
		log.Println("error: failed writing buffer", err)
	}
	fmt.Println("bytes read = ", n)
	fmt.Printf("Written to buffe : value = %b, type of buf = %T, type of *buf = %T\n", buf, buf, &buf)
	fmt.Println("---------------------------------------------------------------")

	var v2 uint16
	for {
		//We need to use a pointer to &buf to get access to the buf's Read method.
		err = binary.Read(&buf, binary.LittleEndian, &v2)
		if err != nil {
			log.Println("error: failed binary read", err)
			break
		}
		fmt.Printf("Reading 16bits, LittleEndian: value = %b, type = %T\n", v2, v2)
		fmt.Println("---------------------------------------------------------------")
	}

}
