package main

import (
	"fmt"
)

func addToString(s string) string {
	return fmt.Sprint(s, " apekatt")

}

//let the function accept another function as a parameter
func stringMe(name string, addS func(string) string) string {
	return addS(name)
}

func main() {
	fmt.Println(stringMe("Anna", addToString))
}
