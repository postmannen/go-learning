package main

import "fmt"

const (
	// iota will start an uint8 with the first value of 1, then 2 and so on.
	farmMonkey uint8 = iota // 1
	farmHorse               // 2
	farmPig                 // 3
	farmSheep               // 4
)

func main() {
	a := farmHorse
	fmt.Println(a)
}
