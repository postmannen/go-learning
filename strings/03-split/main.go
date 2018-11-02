package main

import (
	"fmt"
	"log"
)

//chrPositions , finds the positions containing a chr in a string
func findChrPositions(myString string, chr byte) (equalPosition []int) {
	for i := 0; i < len(myString); i++ {
		//find the positions of the "=" character
		if myString[i] == byte(chr) {
			//fmt.Println(myString[i])
			equalPosition = append(equalPosition, i)
			//fmt.Println("DEBUG : equalPosition : ", equalPosition)
		}
	}
	return
}

//findPriorOccurance .
// Searches backwards in a string from a given positions,
// for the first occurence of a character.
func findPriorOccurance(myString string, preChr byte, origChrPositions []int) (preChrPositions []int) {
	for _, v := range origChrPositions {
		vv := v
		//fmt.Println("DEBUG: vv : ", vv)

		//find the first space before the preceding word
		for {
			vv--

			if vv < 0 {
				log.Println("Found no space before the equal sign, reached the beginning of the line")
				break
			}

			if myString[vv] == preChr {
				preChrPositions = append(preChrPositions, vv)
				break
			}
		}
	}
	//will return the preceding chr's positions found
	return
}

//findNextOccurance .
// Searches forward in a string from a given positions,
// for the first occurence of a character after it.
// The function takes multiple positions as input,
// and will also return multiplex positions
func findNextOccurance(myString string, preChr byte, origChrPositions []int) (nextChrPositions []int) {
	for _, v := range origChrPositions {
		vv := v
		//fmt.Println("DEBUG: vv : ", vv)

		//find the first space before the preceding word
		for {
			vv++

			if vv > len(myString)-1 {
				log.Println("Found no space before the equal sign, reached the end of the line")
				break
			}

			if myString[vv] == preChr {
				nextChrPositions = append(nextChrPositions, vv)
				break
			}
		}
	}

	//will return the preceding chr's positions found
	return
}

//findLettersBetween
func findLettersBetween(myString string, firstPositions []int, secondPositions []int) (words []string) {
	for i, v := range firstPositions {
		letters := []byte{}

		//as long as first position is lower than second position....
		for v < secondPositions[i] {
			//fmt.Printf("DEBUG: *** v : %v, secondposition : %v\n", v, secondPositions[i])

			letters = append(letters, myString[v])

			v++
		}
		words = append(words, string(letters))
	}
	return
}

func main() {
	myString := `<arg name="longitude" type="double" ape="apekatt" hest="folafolablakken">`

	//Find the positions where there is an equal sign in the string
	equalPositions := findChrPositions(myString, '=')

	preChrPositions := findPriorOccurance(myString, ' ', equalPositions)

	//==============find the word before the equal sign==============================

	//We need to add 1 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range preChrPositions {
		preChrPositions[i]++
	}

	attributeNames := findLettersBetween(myString, preChrPositions, equalPositions)
	fmt.Printf("DEBUG: Found word >%v<\n", attributeNames)
	//fmt.Printf("Equal position : %v, preChrPositions : %v \n", equalPositions, preChrPositions)

	// =================find the word after the equal and between " "===========================

	fmt.Println("===================================================================")

	nextChrPositions := findNextOccurance(myString, '"', equalPositions)
	//fmt.Println("********nextChrPositions : ", nextChrPositions)

	nextNextChrPositions := findNextOccurance(myString, '"', nextChrPositions)
	//fmt.Println("********nextNextChrPositions : ", nextNextChrPositions)

	//We need to add 2 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range nextChrPositions {
		nextChrPositions[i] = nextChrPositions[i] + 1
	}

	attributeValues := findLettersBetween(myString, nextChrPositions, nextNextChrPositions)
	fmt.Printf("DEBUG: Found word >%v<\n", attributeValues)
}
