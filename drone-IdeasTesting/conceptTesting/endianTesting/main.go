package main

import (
	"bytes"
	"encoding/binary"
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

func main() {

}
