package main

import (
	"testing"
)

func BenchmarkMix1(t *testing.B){
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 0}

	for i:=0;i<t.N;i++ {
	_ = mix1(s1, s2)		
	}
}

func BenchmarkMix2(t *testing.B){
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 0}

	for i:=0;i<t.N;i++ {
	_ = mix1(s1, s2)		
	}
}