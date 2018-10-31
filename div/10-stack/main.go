//The idea is to make a push/pop stack, where you push
// and pop elements to the end of a stack with append.
package main

import (
	"fmt"
)

//stack will keep track of where we are working in the iteration,
type stack struct {
	data []string
}

func newStack() *stack {
	return &stack{}
}

//push will add another item to the end of the stack with a normal append
func (s *stack) push(d string) {
	s.data = append(s.data, d)
}

//pop will remove the last element of the stack
func (s *stack) pop() {
	last := len(s.data)
	s.data = append(s.data[0:0], s.data[:last-1]...)
}

func main() {
	s := newStack()
	s.push("ape")
	s.push("bever")
	s.push("cat")
	s.push("dog")
	s.push("elephant")
	s.push("frog")

	fmt.Println("Elements of stack : ", s)
	fmt.Println("And the length of the stack : ", len(s.data))
	fmt.Println("==========================================================")

	s.pop()
	fmt.Println("Elements of stack : ", s)
	fmt.Println("And the length of the stack : ", len(s.data))
	fmt.Println("==========================================================")

}
