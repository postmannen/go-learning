package main

import (
	"fmt"

	"github.com/postmannen/go-learning/package/01-test/db"
	"github.com/postmannen/go-learning/package/01-test/house"
	"github.com/postmannen/go-learning/package/01-test/web"
)

func main() {
	web.Print()
	db.Print()
	d := db.DStruct{}
	fmt.Println(d)

	house.MakeWindow()

}
