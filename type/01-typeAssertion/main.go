/*
	Example who shows how to do type checking with a switch statement.
*/
package main

import "fmt"

func start(d map[int]interface{}) {
	for k, v := range d {
		switch v.(type) {
		case int:
			fmt.Println(k, "The type of the interface value was an int")
		case string:
			fmt.Println(k, "The type of the interface value was a string")
		default:
			fmt.Println(k, "type not detected")
		}

	}
}

func main() {
	db := map[int]interface{}{
		1: "apekatt",                           //a string
		2: 1024,                                //an int
		3: struct{ name string }{name: "hest"}, //does not match either int or string
	}

	start(db)
}
