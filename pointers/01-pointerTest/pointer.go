package main

import "fmt"

type Node struct {
	Next     *Node //Lag variabelen Next av samme type Node som den vi definerer
	Value    interface{}
	Previous *Node
}

type list struct {
	First *Node
	Last  *Node
}

func (l *list) insert(val interface{}) { //metode knyttet til *list
	newNode := &Node{
		Next:  nil,
		Value: val,
	}
	if l.First == nil {
		l.First = newNode
	}

	l.Last.Next = newNode
	l.Last = newNode
}

func main() {

	a := list{}
	a.insert("apekatt")
	a.insert("orangutang")
	a.insert("hest")

	//fmt.Printf("%v\t%T\n", first, first)
	//fmt.Println(first)

	visited := make(map[*Node]bool)
	for n := a.First; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

}
