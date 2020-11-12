package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	v := float64(-0.2048)
	b := convLittleEndianSlice(v)
	fmt.Printf("%#v\n", b)

	var v2 float64
	convLittleEndian(b, &v2)
	fmt.Printf("%#v\n", v2)

}

// convLittleEndianSlice takes a a value of any of the standard types
// uint8/int8/uint16/int16/uint32/int32/uint64/int64/float32/float64
// and convert to a []byte.
func convLittleEndianSlice(value interface{}) []byte {
	var b []byte

	switch v := value.(type) {
	case uint8:
		b = []byte{byte(v)}
	case int8:
		b = []byte{byte(v)}
	case uint16:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b, v)
	case int16:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v))
	case uint32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, v)
	case int32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(v))
	case uint64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, v)
	case int64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v))
	case float32:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, math.Float32bits(v))
	case float64:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, math.Float64bits(v))
	case string:
		b = []byte(v)

	}

	return b
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
