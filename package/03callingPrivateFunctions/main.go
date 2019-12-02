package main

import "fmt"
import "github.com/postmannen/go-learning/package/03callingPrivateFunctions/calc"

func main() {
	sum := calc.Calc(10, 10)
	fmt.Println(sum)
}
