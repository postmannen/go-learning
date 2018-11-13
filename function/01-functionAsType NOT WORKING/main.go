package main

import (
	"fmt"
)

type deleter interface {
	deleteLine()
}

type db struct{}

func (d db) deleteLine() {
	fmt.Println("Deleting line")
}

func main() {
	var myDB db
	myDB.deleteLine()

	var a deleter

	fmt.Printf("a is of type %T\n", a)

}
