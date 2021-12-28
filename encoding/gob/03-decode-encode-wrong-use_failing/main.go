package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type data struct {
	Field1 string
	Field2 int
	Field3 string
}

func main() {
	s1 := &data{
		Field1: "Hello Gob 1",
		Field2: 999,
		Field3: "one",
	}
	s2 := &data{
		Field1: "Hello Gob 2",
		Field2: 888,
		Field3: "two",
	}
	log.Println("Original value s1:", s1)
	log.Println("Original value s2:", s2)

	var buf bytes.Buffer

	encoder := gob.NewEncoder(&buf)
	var err error

	err = encoder.Encode(s1)
	if err != nil {
		log.Println("Encode:", err)
		return
	}
	fmt.Printf(" * len of buf enc s1: %v\n", len(buf.Bytes()))

	err = encoder.Encode(s2)
	if err != nil {
		log.Println("Encode:", err)
		return
	}
	fmt.Printf(" * len of buf enc s2: %v\n", len(buf.Bytes()))

	// ---------- Decode ------------

	fmt.Printf(" * buf content: %v\n", buf.String())

	s3 := &data{}
	s4 := &data{}

	decoder := gob.NewDecoder(&buf)

	err = decoder.Decode(s3)
	if err != nil {
		log.Println("Decode s3:", err)
		return
	}
	log.Println("Decoded value:", s3)
	fmt.Printf(" * len of buf dec s3: %v\n", len(buf.Bytes()))

	decoder2 := gob.NewDecoder(&buf)
	err = decoder2.Decode(s4)
	if err != nil {
		log.Println("Decode s4:", err)
		return
	}
	log.Println("Decoded value:", s4)
	fmt.Printf(" * len of buf enc s4: %v\n", len(buf.Bytes()))

}
