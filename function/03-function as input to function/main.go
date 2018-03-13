package main

import (
	"fmt"
)

func addToString(s string) string {
	return fmt.Sprint(s, " apekatt")

}

func stringMe(name string, addS func(string) string) string {
	return addS(name)
}

func main() {
	fmt.Println(stringMe("Anna", addToString))
}
