package main

import (
	"encoding/xml"
	"fmt"
)

var myXML = []byte(`
	<sensor>
		<sensorType>Temp Sensor</sensorType>
	</sensor>
	`)

type location struct {
}

//Sensor is for storing the xml values
type Sensor struct {
	XMLName    xml.Name `xml:"sensor"`
	SensorType string   `xml:"sensorType"`
}

func main() {

	sensors := Sensor{}

	err := xml.Unmarshal(myXML, &sensors)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("sensors = ", sensors)
}
