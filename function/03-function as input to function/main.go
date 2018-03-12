package main

import (
	"fmt"
)

func two() {
	fmt.Println("calling two")
}

func three() {
	fmt.Println("calling three")
}

func abc(a func()) {
	a()
}

func main() {
	abc(two)

}
