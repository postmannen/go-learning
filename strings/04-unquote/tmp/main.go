package main

import (
	"fmt"
)

type aType struct{}

func (a aType) read(b []byte) (n int, err error) {
	s := "aType got some data in it"
	n = copy(b, s)

	return 0, nil
}

type bType struct{}

func (bt *bType) read(b []byte) (n int, err error) {
	s := "bType got some data in it"
	n = copy(b, s)

	return n, nil
}

type reader interface {
	read([]byte) (int, error)
}

func retrieve(r reader) {
	data := make([]byte, 7)
	n, _ := r.read(data)
	fmt.Println("Read number of bytes : ", n)
	fmt.Println("data contains : ", string(data))

}

func main() {
	a := &aType{}
	retrieve(a)

	b := &bType{}
	retrieve(b)

}
