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
	"fmt"
	"testing"
)

func BenchmarkWithoutReaderBenchmark(b *testing.B) {
	by := []byte{60, 70}
	for n := 0; n < b.N; n++ {
		s := withoutReader(by)
		_ = fmt.Sprint(s)
	}
}

func BenchmarkWithReader(b *testing.B) {
	by := []byte{60, 70}
	for n := 0; n < b.N; n++ {
		s := withReader(by)
		_ = fmt.Sprint(s)
	}
}
