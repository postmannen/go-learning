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

	xRect := 50
	yRect := 50
	wRect := 100
	hRect := 100

	distance := 150

	textPadding := 20

	s.Roundrect(xRect, yRect, wRect, hRect, 10, 10, "fill:none;stroke:black")
	s.Text(xRect+textPadding, yRect+textPadding, "Hello, SVG", "text-anchor:left;font-size:10px;fill:black")

	s.Roundrect(xRect+wRect+distance, yRect, wRect, hRect, 10, 10, "fill:none;stroke:black")

	s.Line(xRect+wRect, yRect+(hRect/2), xRect+wRect+distance, yRect+(hRect/2), "fill:none;stroke:black")
	s.End()
}
