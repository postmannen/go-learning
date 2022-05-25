package main

import "fmt"

type some[T any] []T

func newSome[T any](s some[T]) some[T] {
	return s
}

func (t some[T]) getConn() {
	fmt.Println("127.0.0.1:3128")
}

// ---

func concatenateSlice(s []string) string {
	str := ""

	for _, v := range s {
		str = str + fmt.Sprintf(" %v", v)
	}

	return str
}

func main() {
	a := []string{"large", "some", "what a", "gigantic"}
	b := newSome(a)
	b.getConn()
	fmt.Printf("%v\n", b[0])

	fmt.Printf("%v\n", concatenateSlice(b))
}
