package main

import (
	"log"
)

type løgfunc func(interface{})

type løg struct {
	løgfunc
}

func (l *løg) l(s string) {
	l.løgfunc(s)
}

func main() {
	f := func(s interface{}) {
		log.Printf("%v", s)
	}

	l := løg{løgfunc: f}
	l.l("a\n")

}
