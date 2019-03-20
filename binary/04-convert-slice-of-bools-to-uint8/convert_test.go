package main

import (
	"testing"
)

func BenchmarkConvert(b *testing.B) {
	bol := []bool{true, false, false, false, true, false, false, true}

	for n := 0; n < b.N; n++ {
		_ = convert(bol)
	}
}
