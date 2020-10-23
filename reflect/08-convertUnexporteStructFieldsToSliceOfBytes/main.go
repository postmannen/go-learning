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
	b uint16
	c uint32
}

func main() {
	var buf bytes.Buffer

	d := data{5, 6, 7}

	rv := reflect.ValueOf(d)

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		v := (*value)(unsafe.Pointer(&f))
		v.flag &^= flagRO
		binary.Write(&buf, binary.LittleEndian, f.Interface())
	}

	fmt.Println(buf.Bytes())
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
