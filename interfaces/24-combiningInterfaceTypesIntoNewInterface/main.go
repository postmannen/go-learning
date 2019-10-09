package main

import "fmt"

type human struct {
	name string
}

func (h *human) giveName(s string) {
	h.name = s
}

func (h *human) speak() {
	fmt.Printf("%v are speaking\n", h.name)
}

type animal struct {
	name string
}

func (a *animal) giveName(s string) {
	a.name = s
}

func (a *animal) howl() {
	fmt.Printf("%v are howling\n", a.name)
}

type namer interface {
	giveName(string)
}

type speaker interface {
	speak()
}

type howler interface {
	howl()
}

type speakerNamer interface {
	speaker
	namer
}

func main() {
	var h speakerNamer
	h = &human{name: "knut"}
	h.speak()

	// This will not work, since an animal don't have a speak method.
	//var a speakerNamer
	//a = &animal{name: "monkey"}

}
