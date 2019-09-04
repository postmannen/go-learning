package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func withReader(in []byte) float32 {
	var f32 float32
	binary.Read(bytes.NewReader(in), binary.LittleEndian, &f32)
	return f32
}

func withoutReader(in []byte) uint16 {
	var ui16 uint16
	ui16 = binary.LittleEndian.Uint16(in)
	return ui16
}

func convLittleEndian(in []byte, out interface{}) {
	switch out := out.(type) {
	case *uint8:
		*out = uint8(in[0])
	case *int8:
		*out = int8(in[0])
	case *uint16:
		*out = binary.LittleEndian.Uint16(in)
	case *int16:
		*out = int16(binary.LittleEndian.Uint16(in))
	case *uint32:
		*out = binary.LittleEndian.Uint32(in)
	case *int32:
		*out = int32(binary.LittleEndian.Uint32(in))
	case *uint64:
		*out = binary.LittleEndian.Uint64(in)
	case *int64:
		*out = int64(binary.LittleEndian.Uint32(in))
	case *float32:
		bits := binary.LittleEndian.Uint32(in)
		*out = math.Float32frombits(bits)
	case *float64:
		bits := binary.LittleEndian.Uint64(in)
		*out = math.Float64frombits(bits)
	case *string:
		*out = string(in)
	}
}

func main() {
	// 154, 221, 45, 61, 44, 209, 73, 188, 121, 230, 52, 64
	//data := []byte{121, 230, 52, 64}
	data := []byte{121, 230, 52, 64}
	var o float32

	fmt.Println("withReader = ", withReader(data))
	convLittleEndian(data, &o)
	fmt.Printf("In main after calling function o = %+v\n", o)

	data = []byte{102, 159, 199, 40, 207, 156, 6, 64}
	var o2 float64

	fmt.Println("withReader = ", withReader(data))
	convLittleEndian(data, &o2)
	fmt.Printf("In main after calling function o = %+v\n", o2)

	data = []byte{0, 0, 0, 1}
	var o3 int32

	fmt.Println("withReader = ", withReader(data))
	convLittleEndian(data, &o3)
	fmt.Printf("In main after calling function o = %+v\n", o3)
}
