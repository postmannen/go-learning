package main

import (
	"log"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

func main() {
	http.Handle("/circle", http.HandlerFunc(circle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(1000, 1000)
	s.Circle(250, 250, 100, "fill:none;stroke:black")
	s.Text(250, 250, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:blue")
	s.Circle(500, 250, 100, "fill:none;stroke:black")
	s.Line(350, 250, 400, 250, "fill:none;stroke:black")
	s.End()
}
