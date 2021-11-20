package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	svg "github.com/ajstarks/svgo"
)

func main() {
	http.Handle("/", http.HandlerFunc(draw))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type rect struct {
	xRect int
	yRect int
	wRect int
	hRect int

	distanceRect int
	textPaddingX int
	textPaddingY int

	fontSize int
	maxChr   int

	amount int
}

func firstRect() *rect {
	r := rect{}

	r.xRect = 50
	r.yRect = 50
	r.wRect = 100
	r.hRect = 100

	r.distanceRect = 50
	r.textPaddingX = 10
	r.textPaddingY = 20

	r.fontSize = 15
	r.maxChr = 10

	return &r
}

func (r *rect) draw(s *svg.SVG, text string) {
	// Draw a rectangle.
	s.Roundrect(r.xRect, r.yRect, r.wRect, r.hRect, 10, 10, "fill:none;stroke:black")

	// Put text in the rectangle

	//maxChr := r.wRect / r.fontSize
	maxChr := 10

	textSlice := sliceNString(text, maxChr)
	fSize := strconv.Itoa(r.fontSize)
	textStyle := fmt.Sprintf("text-anchor:left;font-size:%spx;fill:black", fSize)
	lineSize := 0

	for _, v := range textSlice {
		s.Text(r.xRect+r.textPaddingX, r.yRect+r.textPaddingY+lineSize, v, textStyle)
		lineSize += 15
	}

	if r.amount != 0 {
		s.Line(r.xRect, r.yRect+(r.hRect/2), r.xRect-r.distanceRect, r.yRect+(r.hRect/2), "fill:none;stroke:black")
	}

	r.amount++

	// Move the position for the next rectangle.
	r.xRect = r.xRect + r.wRect + r.distanceRect
}

func sliceNString(text string, n int) []string {
	textSlice := []string{}
	counter := 1

	var str string
	for i, v := range text {
		str = str + string(v)

		switch {
		case i >= len(text)-1:
			textSlice = append(textSlice, str)
		case v == ' ':
			textSlice = append(textSlice, str)
			str = ""
			counter = 1
		case counter >= n:
			textSlice = append(textSlice, str)
			str = ""
			counter = 1
		case counter < n:
			counter++
		}
	}

	return textSlice
}

func draw(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(1000, 1000)
	defer s.End()

	r := firstRect()
	r.draw(s, "abcdefghijklmnopqrstuvwxyz")
	r.draw(s, "123456789012345678901234567890")
	r.draw(s, "this is an odd sentence")
}
