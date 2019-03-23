package main

import "fmt"

func main() {
	a := []struct {
		aa string
	}{
		{aa: "some aa"},
		{aa: "some ab"},
		{aa: "some ac"},
	}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%v\n\n", a)

	for i, v := range a {
		fmt.Printf("i = %v, v = %#v\n", i, v)
	}
}
