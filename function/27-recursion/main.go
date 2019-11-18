package main

import (
	"fmt"
	"log"
)

type nodeType int

const (
	firstName nodeType = 10 + iota
	lastName
	address
)

type node struct {
	nodeType nodeType
	value    string
	subNodes []node
}

var class = []node{
	node{
		nodeType: firstName,
		value:    "Ola",
		subNodes: []node{
			node{
				nodeType: lastName,
				value:    "Olausen",
				subNodes: []node{
					node{
						nodeType: address,
						value:    "Olaveien 1",
					},
				},
			},
		},
	},
	node{
		nodeType: firstName,
		value:    "Per",
		subNodes: []node{
			node{
				nodeType: lastName,
				value:    "Person",
				subNodes: []node{
					node{
						nodeType: address,
						value:    "Persveien 1",
					},
					node{
						nodeType: address,
						value:    "Nilsveien 2",
					},
				},
			},
		},
	},
}

func findNameNode(name string, class []node) (node, error) {
	for _, n := range class {
		if n.nodeType != firstName {
			continue
		}

		if n.value == name {
			return n, nil
		}
	}
	return node{}, fmt.Errorf("did not find name being asked for")
}

func findLastNameNodes(lName string, class []node) (node, error) {
	for _, n := range class {
		fmt.Println("*** Checking node : ", n)
		//if n.nodeType == lastName || n.value == lName {
		//	return n, nil
		//}

		findLastNameNodes(lName, n.subNodes)
	}
	return node{}, fmt.Errorf("did not find name being asked for")
}

func main() {
	fmt.Println(class)
	n, err := findNameNode("Per", class)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("found node for name : ", n)

	fmt.Println()
	ln, err := findLastNameNodes("Person", class)
	if err != nil {
		log.Println("error: ", err)
	}
	fmt.Println("ln : ", ln)

}
