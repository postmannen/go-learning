package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

// convLittleEndianSlice takes a a value of any of the standard types
// uint8/int8/uint16/int16/uint32/int32/uint64/int64/float32/float64
// and convert to a []byte.
func convLittleEndianSlice(value reflect.Value) []byte {
	var b []byte

	iv := value.Interface()

	switch v := iv.(type) {
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

func Decode(myStruct interface{}) []byte {
	var bs []byte
	valueOf := reflect.ValueOf(myStruct)

	fmt.Printf("Number of fields in the struct: %v\n", valueOf.NumField())
	fmt.Println("--------------Iterating fields-----------------")
	for i := 0; i < valueOf.NumField(); i++ {
		b := convLittleEndianSlice(valueOf.Field(i))
		fmt.Printf("mySlice = %#v\n", b)

		bs = append(bs, b...)
	}

	return bs
}

type numbers1 struct {
	NumberOne uint16
	NumberTwo uint32
}

type numbers2 struct {
	NumberOne uint32
	NumberTwo uint16
	String1   string
}

type numbers3 struct{}

func main() {
	struct1 := numbers1{
		NumberOne: 255,
		NumberTwo: 255,
	}

	{
		myByteSlice := Decode(struct1)
		fmt.Printf("Resultingslice = %#v\n\n", myByteSlice)
	}

	// -------------

	struct2 := numbers2{
		NumberOne: 255,
		NumberTwo: 255,
		String1:   "AABB",
	}

	{
		myByteSlice := Decode(struct2)
		fmt.Printf("Resultingslice = %#v\n", myByteSlice)
	}

	struct3 := numbers3{}

	{
		myByteSlice := Decode(struct3)
		fmt.Printf("Resultingslice = %#v\n", myByteSlice)
	}
}
