package main

import (
	"fmt"
)

func main() {
	myString := "{{{([])}}"

	myStack := []byte{}

	for i := range myString {
		fmt.Println(myString[i])

		switch string(myString[i]) {
		case "{", "(", "[":
			fmt.Println("found start = ", string(myString[i]))
			fmt.Println("putting ", string(myString[i]), " on the stack")
			myStack = append(myStack, myString[i])
			fmt.Println("The stack content = ", myStack)
		case "}", ")", "]":
			myStack = myStack[:len(myStack)-1]
			fmt.Println("found closing ", string(myString[i]), "deleting from stack")
			fmt.Println("Stack now contains: ", myStack)
		}
	}

	if len(myStack) == 0 {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("The string was consistent with starting and closing")
	} else {
		fmt.Println("-------------------------------------------------------")
		fmt.Println("The string was not consistent with starting and closing")

		for _, v := range myStack {
			fmt.Println("The values that are missing clousure are : ", string(v))
		}
	}
}
