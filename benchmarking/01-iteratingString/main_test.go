package main

import (
	"testing"
)

func BenchmarkNormalIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkForChrAfter(myString, 1, "</")
	}
}

func BenchmarkUsingStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkForChrAfter2(myString, "</")
	}
}
