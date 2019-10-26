package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var data = `
{
	"window": 
		[
			{
				"addButton": "buttonTemplate1",
				"addHeader": "socketTemplate1",
				"addParagraph": "paragraphTemplate1"
			},
			{
				"addButton": "buttonTemplate2",
				"addHeader": "socketTemplate2",
				"addParagraph": "paragraphTemplate2"
			}
		]
}
`

// WindowContent is describing a window
type WindowContent struct {
	Addbutton    string `json:"addButton"`
	AddHeader    string `json:"addHeader"`
	AddParagraph string `json:"addParagraph"`
}

// Window are a collection of windows
type Window struct {
	WindowContent []WindowContent `json:"window"`
}

func unmarshalWindow(data string) (Window, error) {
	var w Window
	err := json.Unmarshal([]byte(data), &w)
	if err != nil {
		return w, err
	}

	return w, nil
}

func main() {
	w, err := unmarshalWindow(data)
	if err != nil {
		log.Println("unmarshaling failed : ", err)
	}

	fmt.Println(w)
}
