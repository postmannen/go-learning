package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = `
{
	"window": {
		"addButton": "buttonTemplate1",
		"addHeader": "socketTemplate1",
		"addParagraph": "paragraphTemplate1"
	}
}
`

// WindowContent is describing a window
type WindowContent struct {
	Addbutton    string `json:"addButton"`
	AddHeader    string `json:"addHeader"`
	AddParagraph string `json:"addParagraph"`
}

type Window struct {
	WindowContent WindowContent `json:"window"`
}

func unmarshalWindow(data string) (Window, error) {
	var w Window
	err := json.Unmarshal([]byte(data), &w)
	if err != nil {
		return w, err
	}

	fmt.Println("1 : ", w)
	return w, nil
}

func main() {
	w, err := unmarshalWindow(data)
	if err != nil {
		log.Println("unmarshaling failed : ", err)
	}

	fmt.Println("2 : ", w)
}
