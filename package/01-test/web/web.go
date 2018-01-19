package web

import (
	"fmt"

	"github.com/postmannen/training/package/test1/db"
)

//Print function
func Print() {
	fmt.Println("Printing from web package")

	webVar1 := db.DStruct{}
	webVar1.DVar = "nisse"
	fmt.Printf("Printing webVar1 = %v which is of type %T\n", webVar1, webVar1)
}