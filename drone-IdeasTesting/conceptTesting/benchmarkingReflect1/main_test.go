package main

import (
	"testing"
)

func BenchmarkInsertA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertA(10)
	}
}

func BenchmarkInsertB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertB(int(10))
	}
}
