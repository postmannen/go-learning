package main

import (
	"encoding/xml"
	"fmt"
)

var myXML = []byte(`
	<sensor>
		<nisse>
			<nissefar>og nissemor</nissefar>
		</nisse>
		<sensorType>Temp Sensor1</sensorType>
		<sensorType>Temp Sensor2</sensorType>
		<sensorType>Temp Sensor3</sensorType>
	</sensor>
	`)

//Sensor is for storing the xml values
type Sensor struct {
	Raw        string `xml:"any"` //have to check more on this one, not doing what I thought it should do :)
	XMLName    xml.Name
	SensorType []string `xml:"sensorType"`
}

func main() {

	sensors := Sensor{}

	err := xml.Unmarshal(myXML, &sensors)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("sensors = ", sensors)
}
