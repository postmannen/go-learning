package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	// Add 0x infront to write a number as hexadecimal
	a := 0x00ab
	b := 0x00cd
	c := 0x00ef

	w := &bytes.Buffer{}

	// Write binary data into w with a byte order of LittleEndian.
	// Endianess works at at the byte level, so the byte slice
	// 0xaabb in will be 0xbbaa written in LittleEndian.
	err := binary.Write(w, binary.LittleEndian, byte(a))
	if err != nil {
		log.Println("error: binary write", err)
	}
	err = binary.Write(w, binary.LittleEndian, byte(b))
	if err != nil {
		log.Println("error: binary write", err)
	}
	err = binary.Write(w, binary.LittleEndian, byte(c))
	if err != nil {
		log.Println("error: binary write", err)
	}

	fmt.Printf("%#v\n", w)
}
