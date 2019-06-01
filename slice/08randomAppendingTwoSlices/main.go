package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 0}

	s3 := mix1(s1, s2)
	fmt.Printf("s3=%v, length of s1=%v, cap of s1=%v\n", s3, len(s1), cap(s1))
	fmt.Println(s1)

	s4 := mix2(s1, s2)
	fmt.Printf("s4=%v, length of s1=%v, cap of s1=%v\n", s4, len(s1), cap(s1))
	fmt.Println(s1)
}

func mix1(s1, s2 []int) []int {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s2); i++ {
		pos := rand.Intn(len(s1))
		lastPart := s1[pos:len(s1)]
		firstPart := s1[0:pos]

		s3 := []int{}
		s3 = append(s3, firstPart...)
		s3 = append(s3, s2[i])
		s3 = append(s3, lastPart...)
		s1 = s3
	}
	return s1
}

func mix2(s1, s2 []int) []int {
	rand.Seed(time.Now().UnixNano())
	s3 := make([]int, len(s1)+len(s2))

	for i := 0; i < len(s2); i++ {
		pos := rand.Intn(len(s1))
		lastPart := s1[pos:len(s1)]
		firstPart := s1[0:pos]

		s3 = append(s3, firstPart...)
		s3 = append(s3, s2[i])
		s3 = append(s3, lastPart...)
		s1 = s3
	}
	return s1
}
