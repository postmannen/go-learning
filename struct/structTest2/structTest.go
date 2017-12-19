/*
type dummy struct {
    a int
}
x := make(map[int]dummy)
x[1] = dummy{a:1}
x[1].a = 2
*/

package main

import "fmt"

type personStruct struct {
	name    string
	surname string
	mail    string
}

func main() {
	person := make(map[int]*personStruct)
	person[1] = &personStruct{}
	person[1].name = "Bob"
	person[1].surname = "Bobson"
	person[1].mail = "bob.bobson@bobson.com"
	fmt.Println("Person[1].name = ", person[1])
	fmt.Println("Slutt....")
}
