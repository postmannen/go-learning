package main

import "fmt"

type some[T any] []T

func newSome[T any](s some[T]) some[T] {
	return s
}

func (t some[T]) test() {
	for i := range t {
		fmt.Printf("%v monkey\n", t[i])
	}
}

func main() {
	a := []string{"a", "some", "what a", "gigantic"}
	b := newSome(a)
	b.test()
}
