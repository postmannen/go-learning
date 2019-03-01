package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	a := 0x00ab
	b := 0x00cd
	c := 0x00ef

	w := &bytes.Buffer{}
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
