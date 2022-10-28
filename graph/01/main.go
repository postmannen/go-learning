package main

import (
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	type City struct {
		Name string
	}

	cityHash := func(c City) string {
		return c.Name
	}

	g := graph.New(cityHash)

	london := City{Name: "london city"}

	g.AddVertex(london)

	file, _ := os.Create("my-graph.gv")
	draw.DOT(g, file)

}
