package main

import (
	"testing"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		first := aType{}
		for ii := 1; ii <= 1000000; ii++ {
			first.i = first.aFn(ii)
		}
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		second := &bType{}
		for ii := 1; ii <= 1000000; ii++ {
			second.bFn(ii)
		}
	}
}
