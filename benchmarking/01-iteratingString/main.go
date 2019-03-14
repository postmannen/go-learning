/*
 Check for the occurence of a string inside another string. Will start looking from
 the position given as input.
*/
package main

import (
	"fmt"
	"strings"
)

const myString string = "<start>This is a test</test>"

func main() {
	foundEnd := false
	for i, v := range myString {
		if v == '>' && !foundEnd {
			//fmt.Println("found >, we should do a check for </ here !")

			if checkForChrAfter(myString, i, "</") {
				foundEnd = true
				fmt.Println("found end bracket </ later on the same line")
			}

		}
		fmt.Println(i, string(v))
	}

}

//checkForChrAfter takes a string, a position in a string, and a string pattern to look for as input.
// Returns true if the character combination was found after the given position.
// The function will first look for the first character in the given string,
// and if that character is found we get into the inner for loop and loop over the
// character pattern to check if the rest of the characters match.
// If match we return, the value true, and exit the function.
func checkForChrAfter(s string, curPos int, characters string) (found bool) {
	for i := curPos; i < len(s)-1; i++ {
		if s[i] == characters[0] {
			for ii := 1; ii <= len(characters); ii++ {
				if s[i+ii] == characters[ii] {
					return true
				}
			}
		}
	}

	//No match of the complete character string found, return false.
	return false
}

func checkForChrAfter2(s string, characters string) (found bool) {
	if strings.Contains(s, characters) {
		return true
	}

	//No match of the complete character string found, return false.
	return false
}
