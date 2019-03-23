package main

import "fmt"

func main() {
	a := struct {
		aa func()
	}{
		aa: func() {
			fmt.Println("testing func in aa")
		},
	}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%v\n\n", a)

	a.aa()
}
