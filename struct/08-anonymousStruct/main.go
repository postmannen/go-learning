package main

import "fmt"

func main() {
	data := struct {
		var1 int
		var2 int
	}{
		var1: 10,
		var2: 20,
	}

	fmt.Println(data)
}
