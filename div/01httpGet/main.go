package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getTopicCount(topic string) int {
	u := "https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=" + topic
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println("error:failed http.get: ", err)
		return 0
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var respData []byte
	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("error:reader: ", err)
			break
		}

		respData = append(respData, line...)
	}

	jsn := struct {
		Parse struct {
			Title  string `json:"title"`
			PageId int    `json:"pageid"`
			Text   struct {
				Content string `json:"*"`
			} `json:"text"`
		} `json:"parse"`
	}{}

	err = json.Unmarshal(respData, &jsn)
	if err != nil {
		fmt.Println("error:unmarshal: ", err)
		return 0
	}

	// Strip all tags, and leave us with what is between >..here..<.
	//
	// I assumed the test was for occurences outside tags, if it
	// was to also include occurences within tag, replace the rest of
	// the code in the function with just:
	// return strings.Count(jsn.Parse.Text.Content, topic)

	s := strings.Split(jsn.Parse.Text.Content, ">")

	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], "<") {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	var count int

	for i := 0; i < len(s); i++ {
		tmpSlice := strings.Split(s[i], "<")
		tmpItem := strings.ToLower(tmpSlice[0])
		count += strings.Count(tmpItem, topic)
	}
	return count
}

func main() {

	res := getTopicCount("pizza")
	fmt.Println("The number returned from the function = ", res)

}
