package main

import "fmt"

const (
	farmMonkey uint8 = iota
	farmHorse
	farmPig
	farmSheep
)

func main() {
	a := farmHorse
	fmt.Println(a)
}
