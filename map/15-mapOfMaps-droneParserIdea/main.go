package main

import "fmt"

type command struct {
	description string
}

type class struct {
	description string
	commands    map[string]command
}

type project struct {
	description string
	classes     map[string]class
}

func main() {
	p := project{}
	p.classes = make(map[string]class)
	p.classes["piloting"] = class{description: "piloting stuff", commands: make(map[string]command)}
	p.classes["piloting"].commands["takeoff"] = command{description: "takeoff stuff"}
	p.classes["piloting"].commands["flattrim"] = command{description: "flattrim stuff"}

	fmt.Println(p)
}
