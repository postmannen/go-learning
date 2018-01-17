package main

import "fmt"

type struct1 struct {
	verdi11 int
	verdi12 int
}

type struct2 struct {
	verdi21 int
	verdi22 int
	slice23 []int
}

type allStructs struct {
	allVar1 struct1
	allVar2 struct2
}

func main() {
	a := allStructs{}

	a.allVar1.verdi11 = 111
	a.allVar1.verdi12 = 112
	a.allVar2.verdi21 = 221
	a.allVar2.verdi22 = 222
	a.allVar2.slice23 = []int{10, 20, 30}
	fmt.Println(a)

	a.allVar2.slice23[0] = 100
	fmt.Println(a)

	for i := range a.allVar2.slice23 {
		fmt.Println(i)
	}
}
