package main

import (
	"fmt"
)

type item int

const (
	startTag item = iota
	endTag
	argumentKey
	argumentValue
)

type iMap map[item]string

func main() {
	foundWhileLexing := argumentKey

	if foundWhileLexing == argumentKey {
		fmt.Printf("Value = %v, and the type = %T\n", foundWhileLexing, foundWhileLexing)
	}

	itemMap := make(iMap)
	itemMap[startTag] = "The Start"
	itemMap[endTag] = "The End"
	itemMap[argumentKey] = "Hey, a key argument"
	itemMap[argumentValue] = "Ohh, an argument value"

	fmt.Printf("%v\n", itemMap)

}
