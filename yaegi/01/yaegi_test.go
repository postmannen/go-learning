package main

import (
	"fmt"
	"testing"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func TestGetFunc(t *testing.T) {
	i := interp.New(interp.Options{GoPath: "./_gopath/"})
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	if _, err := i.Eval(`import "github.com/mypkg"`); err != nil {
		t.Fatal(err)
	}

	val, err := i.Eval(`mypkg.func1`)
	if err != nil {
		t.Fatal(err)
	}

	result := val.Call(nil)
	fmt.Printf("%#v\n", result)
}
