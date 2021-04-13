package main

import "fmt"

func main() {
	m := map[string]*string{}
	s := "thing"
	m["some"] = &s

	v, ok := m["some"]
	delete(m, "some")
	if ok {
		// The content of v remains intact even after the map k/v are deleted.
		fmt.Printf("%#v\n", *v)
	}

}
