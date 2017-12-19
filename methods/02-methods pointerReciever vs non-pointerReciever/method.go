package main

import "fmt"

type struct1 struct {
	a int
	b int
}

func (s struct1) stayTheSame() {
	s.a = 10
	s.b = 20
}

func (s *struct1) mutate() {
	s.a = 30
	s.b = 40
}

func main() {
	myVar := struct1{
		a: 100,
		b: 200,
	}

	fmt.Println(myVar)

	myVar.stayTheSame()
	fmt.Println(myVar)

	myVar.mutate()
	fmt.Println(myVar)

}
