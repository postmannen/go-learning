package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	a := uint16(0xAABB)
	fmt.Printf("The initial value of a: %x\n", a)
	b := *(*[2]byte)(unsafe.Pointer(&a))
	fmt.Printf("How the value of a looks layed out in memory: %x\n", b)

	var h, l uint8 = uint8(a >> 8), uint8(a & 0xff)
	ca := []byte{h, l}
	c := binary.LittleEndian.Uint16(ca)
	fmt.Printf("The initial value a converted to little endian: %x\n", c)

}
