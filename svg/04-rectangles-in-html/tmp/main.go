package main

import "fmt"

func main() {
	text := "abcdefghijklmn"
	maxChr := 5

	textSlice := []string{}
	counter := 1

	var str string
	for i, v := range text {
		str = str + string(v)
		fmt.Printf("counter : %v\n", counter)

		switch {
		case i >= len(text)-1:
			textSlice = append(textSlice, str)
		case counter >= maxChr:
			textSlice = append(textSlice, str)
			str = ""
			counter = 1
		case counter < maxChr:
			counter++
		}

	}

	fmt.Printf("%#v\n", textSlice)
}
