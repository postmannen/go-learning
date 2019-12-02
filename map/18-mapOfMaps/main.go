package main

import (
	"fmt"
)

func main() {
	data1 := map[string]map[string]string{"cat": map[string]string{"ears": "two"}}
	fmt.Println(data1)

	var data2 map[string]map[string]string
	data2 = map[string]map[string]string{"cat": map[string]string{"ears": "two"}}
	fmt.Println(data2)

	var data3 = make(map[string]map[string]string)
	innerMap := map[string]string{"ears": "two"}
	data3["cat"] = innerMap
	fmt.Println(data3)
}
