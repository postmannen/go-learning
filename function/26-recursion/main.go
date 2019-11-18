package main

import "fmt"

func checkID(ID int) bool {
	fmt.Println("Checking if ID = 10, is now : ", ID)
	if ID == 10 {
		fmt.Println("ID is 10")
		return true
	}

	ID++
	checkID(ID)
	fmt.Println("After calling checkID recursively, when ID = ", ID)

	return false
}

func main() {
	fmt.Println(checkID(1))
}
