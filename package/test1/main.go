package main

import (
	"fmt"

	"github.com/postmannen/training/package/test1/house"

	"github.com/postmannen/training/package/test1/db"
	"github.com/postmannen/training/package/test1/web"
)

func main() {
	web.Print()
	db.Print()
	d := db.DStruct{}
	fmt.Println(d)

	house.MakeWindow()

}
