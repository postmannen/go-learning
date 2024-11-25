/*
Benchmarking the difference between using LittleEndian conversion
with a reader, and without a reader specifying the type directly.

The result is that without the reader is twice as fast, but the
drawback in relation to the code complexity used in the parser
is that we need to find the correct method to use with the binary
package, as opposed to using binary.Read which will convert the
output of the read into the type of the argument we want to read
into without specifying an specific method for type to convert into.
*/
package main

import (
	"testing"
)

func BenchmarkConvWithReader(b *testing.B) {
	by := []byte{60, 70}
	for n := 0; n < b.N; n++ {
		_ = convWithReader(by)
		//_ = fmt.Sprint(s)
	}
}

func BenchmarkConvLittleEndianUint16(b *testing.B) {
	by := []byte{60, 70}

	for n := 0; n < b.N; n++ {
		var o uint16

		convLittleEndian(by, &o)
	}
}

func BenchmarkConvLittleEndianInt32(b *testing.B) {
	by := []byte{0, 0, 0, 1}

	for n := 0; n < b.N; n++ {
		var o int32
		convLittleEndian(by, &o)
	}
}
