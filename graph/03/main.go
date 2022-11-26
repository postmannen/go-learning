// Create svg:
//   dot -Tsvg -O mygraph.gv

package main

import (
	"fmt"
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	g := graph.New(graph.IntHash, graph.Directed())

	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 3)
	_ = g.AddEdge(3, 4)

	_ = graph.DFS(g, 1, func(value int) bool {
		fmt.Println(value)
		return false
	})

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)
}
