package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var fasitTall int
	fmt.Println(&fasitTall)

	fmt.Print("Skriv tall: ")

	fmt.Scanln(&fasitTall)
	fmt.Println()
	//fmt.Print(fasitTall)

	//fasitTall = rand.Intn(100) //her finner læreren på ett tall som du skal gjette
	//fasitTall = fasitTall + 3

	//dette er brukerinput
	teller := 0
	for {
		elevTall := rand.Intn(1000) //her skriver brukeren liksom inn ett tall
		teller++
		fmt.Println("Tallet eleven gjettet = ", elevTall, ",")
		if fasitTall == elevTall {
			fmt.Println("eleven fant tallet", elevTall, "etter ", teller, "forsøk")
			break
		}
	}

}
