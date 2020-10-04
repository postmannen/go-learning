/*
	Maps needs to be initialized before use with make.
	A single map can be initalized with make(map[string]int)

	If a map contains a map the inner map also needs to be initalied
	before use. That means we need to initalize the outer map first
	which shows the whole structure with the inner map like this
	m:= make(map[string]map[string]int)
	Then we can set the outer map  key value, since the value to the key
	is another map map we can initialize it like this:
	m["somevalue"] = make(map[string]int)
	We can then assign key and value to the inner map:
	m["somevalue"]["anothervalue"] = 10

	Complete example:
	m := make(map[string]map[string]int)
	m["ape"] = make(map[string]int)
	m["ape"]["katt"] = 10
	m["ape"]["katt"] = 20
	m["ape"]["land"] = 30
	fmt.Println(m)

	Below is another example using custom types, and helper functions for
	handling of putting values into the map.
*/

package main

import (
	"fmt"
)

type planet string
type continent string
type country string

func newPlanet(p planet) map[planet]map[continent]map[country]int {
	m := make(map[planet]map[continent]map[country]int)
	m[p] = make(map[continent]map[country]int)
	return m
}

func newContinent(c continent) map[continent]map[country]int {
	m := make(map[continent]map[country]int)
	m[c] = make(map[country]int)
	return m
}

func newCountry(c country) map[country]int {
	m := make(map[country]int)
	m[c] = 0
	return m
}

func main() {
	p := newPlanet("earth")
	p["earth"] = newContinent("europe")
	p["earth"]["europe"] = newCountry("norway")

	fmt.Println(p)

}
