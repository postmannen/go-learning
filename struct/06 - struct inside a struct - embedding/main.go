package main

import "fmt"

//People struct
type People struct {
	long bool
}

//Town struct
type Town struct {
	houses      int
	inhabitants int
	People
}

func main() {
	s := Town{}
	s.long = true

	fmt.Println(s)
	fmt.Println(s.long)
}
