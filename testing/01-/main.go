package main

import (
	"fmt"
	"net/http"
)

func sum(x int, y int) int {
	return x + y
}

func getWeb(u string) (*http.Response, error) {
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println("Error: failed get'ing the web : ", err)
	}
	fmt.Println(resp.Status)

	return resp, err

}

func main() {
	fmt.Println(sum(5, 5))

	_, err := getWeb("https://erter.org")
	fmt.Println("main: err :", err)
}
