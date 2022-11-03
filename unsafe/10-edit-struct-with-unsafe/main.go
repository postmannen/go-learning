package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func main() {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)
}
