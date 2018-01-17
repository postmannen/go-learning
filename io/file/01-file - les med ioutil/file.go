package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")

	if err != nil {
		fmt.Println("ERROR", err)
	}

	for i := range data {
		fmt.Print(string(data[i]))
	}
}
