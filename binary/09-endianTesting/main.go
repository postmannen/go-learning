package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func withReader(in []byte) uint16 {
	var ui16 uint16
	binary.Read(bytes.NewReader(in), binary.LittleEndian, &ui16)
	return ui16
}

func withoutReader(in []byte) uint16 {
	var ui16 uint16
	ui16 = binary.LittleEndian.Uint16(in)
	return ui16
}

func newIdea(in []byte, out interface{}) {
	//fmt.Printf("Before switch: %+v, type=%T\n", out, out)

	switch out := out.(type) {
	case *uint16:
		//fmt.Printf("Inside case: %+v, type = %T\n", out, out)
		//tmp := binary.LittleEndian.Uint16(in)
		//out = &tmp
		*out = binary.LittleEndian.Uint16(in)
	}
	//fmt.Printf("out after switch: %+v, type = %T\n", out, out)

}

func main() {
	data := []byte{60, 70}
	var o uint16

	newIdea(data, &o)
	fmt.Printf("In main after calling function o = %+v\n", o)
}
