package main

import "fmt"

const myString string = "<start>This is a test</test>"

func main() {
	for i, v := range myString {
		if v == '>' {
			fmt.Println("found >")

			if checkForChrAfter(i, myString) {
				fmt.Println("found start bracket later")
			}

		}
		fmt.Println(i, string(v))
	}

}

func checkForChrAfter(curPos int, s string) (found bool) {
	for i := curPos; i < len(s)-1; i++ {
		if s[i] == '<' {
			return true
		}
	}

	return false
}
