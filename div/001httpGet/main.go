package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	u := "https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=pizza"
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println("error:failed http.get: ", err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var bodyText string
	for {

		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println("error:reader: ", err)
			break
		}

		fmt.Println("isPrefix=", isPrefix)
		//fmt.Println(string(line))
		bodyText += string(line)
	}

	fmt.Println(bodyText)
}
