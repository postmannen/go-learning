package main

import "fmt"

func main() {
	x := []string{"Donald", "Duck", "Andeby", "Norge"}
	y := []string{"Fantomet", "Walker", "Hodeskallegrotten", "Afrika"}
	fmt.Println(x)
	fmt.Println(y)

	x = append(x[:1], x[2:]...) //:1 tar med seg index 0, men dropper index 1. [:1 går til 1 man tar ikke med 1]
	//,og fortsetter så fra Index 2 med [2:]
	//Resultatet er at du slicer bort index nr. 1 siden alt igjen blir appended til x.
	fmt.Println("Fjerner index nr.1 som er etternavnet = ", x)

	x = append(x, y...) //Husk ... hvis man appender en annen slice
	fmt.Println("Etter at y er appended til x, så er x = ", x)

	fmt.Println("´n")

	for i, v := range x { //range gir samme funksjonalitet som foreach i perl
		fmt.Printf("%v\t%v\n", i, v)
	}

}
