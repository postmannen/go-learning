// Convert all numeric types to []byte
// Does not work with strings

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

type data struct {
	a uint8
	b uint8
	c uint8
	d uint16
}

func main() {
	d := data{5, 6, 7, 8}

	fmt.Printf("%#v\n", convert(d))
}

func convert(d interface{}) []byte {
	var buf bytes.Buffer

	rv := reflect.ValueOf(d)

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		v := (*value)(unsafe.Pointer(&f))
		v.flag &^= flagRO
		binary.Write(&buf, binary.LittleEndian, f.Interface())
	}

	return buf.Bytes()
}

type value struct {
	_    unsafe.Pointer
	_    unsafe.Pointer
	flag flag
}

type flag uintptr

const (
	flagStickyRO flag = 1 << 5
	flagEmbedRO  flag = 1 << 6
	flagRO       flag = flagStickyRO | flagEmbedRO
)
