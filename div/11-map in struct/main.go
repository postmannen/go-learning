package main

import (
	"fmt"
)

type tree struct {
	tagName string
	values
	comment string
}

type values map[string]string

func newTree(s string) *tree {
	return &tree{
		tagName: s,
		values:  make(map[string]string),
	}
}

func setID(t *tree) {
	t.values["id"] = "1"
}

func setComment(t *tree) {
	t.values["comment"] = "this is some comment that belongs to the tree node"
}

func main() {
	t := newTree("project")

	setID(t)
	setComment(t)

	fmt.Println("RAW:", t)

	for k, v := range t.values {
		fmt.Printf("Key: %v, Value: %v \n", k, v)
	}
}
