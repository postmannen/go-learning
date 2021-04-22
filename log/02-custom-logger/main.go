package main

import (
	"log"
	"reflect"
)

type løgfunc func(interface{})

type løg struct {
	løgfunc
}

func (l *løg) l(s interface{}) {
	l.løgfunc(s)
}

func main() {
	f1 := func(s interface{}) {
		v := reflect.ValueOf(s)
		b := v.Interface().(string)
		log.Printf("%#v, %T", b, b)
	}

	l := løg{løgfunc: f1}

	s := struct {
		a int
		b int
	}{a: 10, b: 20}
	l.l(s)

}
