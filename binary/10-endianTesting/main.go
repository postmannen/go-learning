package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func convWithReader(in []byte) float32 {
	var f32 float32
	binary.Read(bytes.NewReader(in), binary.LittleEndian, &f32)
	return f32
}

// convLittleEndian takes a []byte, and an *out variable of type
// uint8/int8/uint16/int16/uint32/int32/uint64/int64/float32/float64
// and convert the []byte, and places the result into the *out variable.
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

	fmt.Println("withReader = ", convWithReader(data))
	convLittleEndian(data, &o)
	fmt.Printf("In main after calling function o = %+v\n", o)

	data = []byte{102, 159, 199, 40, 207, 156, 6, 64}
	var o2 float64

	fmt.Println("withReader = ", convWithReader(data))
	convLittleEndian(data, &o2)
	fmt.Printf("In main after calling function o = %+v\n", o2)

	data = []byte{0, 0, 0, 1}
	var o3 int32

	fmt.Println("withReader = ", convWithReader(data))
	convLittleEndian(data, &o3)
	fmt.Printf("In main after calling function o = %+v\n", o3)
}
