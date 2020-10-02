package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func NewVehicle(m map[string]interface{}) vehicle {
	switch m["type"].(string) {
	case "car":
		return NewCar(m)
	case "plane":
		return NewPlane(m)
	}
	return nil
}

func NewCar(m map[string]interface{}) *car {
	return &car{
		Type:  m["type"].(string),
		Color: m["color"].(string),
		HP:    int(m["hp"].(float64)),
		Doors: int(m["doors"].(float64)),
	}
}

func NewPlane(m map[string]interface{}) *plane {
	return &plane{
		Type:    m["type"].(string),
		Color:   m["color"].(string),
		Engines: int(m["engines"].(float64)),
	}
}

func main() {
	var vehicles []vehicle

	objs := []map[string]interface{}{}
	err := json.Unmarshal(js, &objs)
	if err != nil {
		log.Fatal(err)
	}

	for _, obj := range objs {
		vehicles = append(vehicles, NewVehicle(obj))
	}

	fmt.Printf("%#v\n", vehicles)
}

var js = []byte(`[{
    "type": "car",
    "color": "red",
    "hp": 85,
    "doors": 4
}, {
    "type": "plane",
    "color": "blue",
    "engines": 3
}]`)

type vehicle interface {
	vehicle()
}

type car struct {
	Type  string
	Color string
	HP    int
	Doors int
}

func (car) vehicle() { return }

type plane struct {
	Type    string
	Color   string
	Engines int
}

func (plane) vehicle() { return }

