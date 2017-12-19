package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	m1["nisse"] = "far"
	fmt.Println(m1)

	m2 := make(map[string]map[string]string)
	m2["nisse"] = make(map[string]string)
	m2["nisse"]["far"] = "Nordpolen"
	fmt.Println(m2)
	fmt.Print("\n")

	m2["nisse"]["mor"] = "Nordpolen"
	fmt.Println(m2)
	fmt.Print("\n")

}
