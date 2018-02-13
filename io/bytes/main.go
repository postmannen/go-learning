package main

import (
	"fmt"
	"io"
)

type language struct {
	says string
}

func (m *language) Read(p []byte) (n int, err error) {
	return n, err
}

var someVariable io.Reader

func saySomething(s io.Reader) {
	fmt.Println(s)
}

func main() {

}
