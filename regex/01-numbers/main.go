package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("[0-9]+")
	myVar := re.FindString("hack5432")
	fmt.Println(myVar)
}
