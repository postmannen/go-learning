package main

import "fmt"

type info struct {
	src  string
	dst  string
	port string
}

func main() {
	inf := []info{
		{"10.0.0.1", "10.0.0.2", "80"},
		{"10.0.0.10", "10.0.0.20", "443"},
		{"192.168.0.1", "192.168.0.2", "8080"},
	}

	m1 := map[string]map[string]map[string]info{}

	for _, vinf := range inf {
		var ok bool

		_, ok = m1[vinf.src]
		if !ok {
			m1[vinf.src] = map[string]map[string]info{}
		}

		_, ok = m1[vinf.src][vinf.dst]
		if !ok {
			m1[vinf.src][vinf.dst] = map[string]info{}
		}

		m1[vinf.src][vinf.dst][vinf.port] = vinf
	}

	fmt.Printf("m1: %v\n\n", m1)

	for k1, v1 := range m1 {
		for k2, v2 := range v1 {
			fmt.Printf("k1: %v, k2: %v, v2: %v\n", k1, k2, v2)
		}
	}
}
