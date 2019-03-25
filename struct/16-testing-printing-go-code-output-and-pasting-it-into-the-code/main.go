package main

import "fmt"

func main() {
	a := struct {
		name string
		id   uint8
	}{name: "drone", id: 1}
	//The code below is just a from the output I got from printing the the 'a' struct out
	// with %#v,then typing 'b :=' and pasting the output into the code.
	b := struct {
		name string
		id   uint8
	}{name: "drone", id: 0x1}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
}
