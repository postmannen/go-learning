package main

import "fmt"

type kind int

const (
	general kind = iota
	critical
	other
)

func (k kind) String() string {
	switch k {
	case general:
		return "general error"
	case critical:
		return "critical error"
	case other:
		return "other error"
	}

	return "unknown error kind"
}

func main() {
	e := critical
	fmt.Printf("%v\n", e)
}
