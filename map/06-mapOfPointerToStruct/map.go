package main

import "fmt"

type struct1 struct {
	age int
}

func main() {

	map1 := map[string]*struct1{}
	map1["Donald"] = &struct1{age: 30}

	fmt.Println(map1["Donald"].age)
}
