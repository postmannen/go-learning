package main

import (
	"fmt"
	"log"
)

//chrPositions , finds the positions containing a chr in a string
//
func findChrPositions(myString string, chr byte) (equalPosition []int) {
	for i := 0; i < len(myString); i++ {
		//find the positions of the "=" character
		if myString[i] == byte(chr) {
			equalPosition = append(equalPosition, i)
		}
	}
	return
}

//findPriorOccurance .
// Searches backwards in a string from a given positions,
// for the first occurence of a character.
//
func findPriorOccurance(myString string, preChr byte, origChrPositions []int) (preChrPositions []int) {
	for _, v := range origChrPositions {
		vv := v

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
	return
}

//findNextOccurance .
// Searches forward in a string from a given positions,
// for the first occurence of a character after it.
// The function takes multiple positions as input,
// and will also return multiplex positions
//
func findNextOccurance(myString string, preChr byte, origChrPositions []int) (nextChrPositions []int) {
	for _, v := range origChrPositions {
		vv := v

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
// takes a string, and two positions given as slices as input,
// and returns a slice of string with the words found.
//
func findLettersBetween(myString string, firstPositions []int, secondPositions []int) (words []string) {
	for i, v := range firstPositions {
		letters := []byte{}

		//as long as first position is lower than second position....
		for v < secondPositions[i] {
			letters = append(letters, myString[v])
			v++
		}
		words = append(words, string(letters))
	}
	return
}

//getAttributes
// takes a string as input, and return the attribute names and
// values as two different slices. Reason for using slices and
// not maps are to preserve the order.
//
func getAttributes(s string) (attributeNames []string, attributeValues []string) {
	//Find the positions where there is an equal sign in the string
	equalPositions := findChrPositions(s, '=')

	preChrPositions := findPriorOccurance(s, ' ', equalPositions)

	//==============find the word before the equal sign==============================

	//We need to add 1 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range preChrPositions {
		preChrPositions[i]++
	}

	attributeNames = findLettersBetween(s, preChrPositions, equalPositions)
	//fmt.Printf("DEBUG: Found word >%v<\n", attributeNames)
	//fmt.Printf("Equal position : %v, preChrPositions : %v \n", equalPositions, preChrPositions)

	// =================find the word after the equal and between " "===========================

	fmt.Println("===================================================================")

	nextChrPositions := findNextOccurance(s, '"', equalPositions)
	//fmt.Println("********nextChrPositions : ", nextChrPositions)

	nextNextChrPositions := findNextOccurance(s, '"', nextChrPositions)
	//fmt.Println("********nextNextChrPositions : ", nextNextChrPositions)

	//We need to add 2 to all the pre positions, since the word we're
	// looking for starts after that character.
	for i := range nextChrPositions {
		nextChrPositions[i] = nextChrPositions[i] + 1
	}

	attributeValues = findLettersBetween(s, nextChrPositions, nextNextChrPositions)
	//fmt.Printf("DEBUG: Found word >%v<\n", attributeValues)

	return
}

func main() {
	myString := `<arg name="longitude" type="double" ape="apekatt" hest="folafolablakken">`

	attributeNames, attributeValues := getAttributes(myString)
	fmt.Println("Found the following key/value pairs")
	for i := range attributeNames {
		fmt.Printf("Key = %v, Value = %v \n", attributeNames[i], attributeValues[i])
	}

}
