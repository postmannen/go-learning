package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	// Put text for the rectangle

	textSlice, longest := sliceNString(text, " ")
	fSize := strconv.Itoa(r.fontSize)
	textStyle := fmt.Sprintf("text-anchor:left;font-size:%spx;fill:black", fSize)
	lineSize := 0

	//set the width of the rectangle according to longest slice element.

	for i := 0; i <= len(textSlice)-1; i++ {
		s.Text(r.xRect+r.textPaddingX, r.yRect+r.textPaddingY+lineSize, textSlice[i], textStyle)
		lineSize += 15
	}

	r.wRect = longest * r.fontSize

	// Draw a rectangle.
	s.Roundrect(r.xRect, r.yRect, r.wRect, r.hRect, 10, 10, "fill:none;stroke:black")

	if r.amount != 0 {
		s.Line(r.xRect, r.yRect+(r.hRect/2), r.xRect-r.distanceRect, r.yRect+(r.hRect/2), "fill:none;stroke:black")
	}

	r.amount++

	// Move the position for the next rectangle.
	r.xRect = r.xRect + r.wRect + r.distanceRect
}

// Will split string at every n character or space.
// Will return a []string, and the length of the longest
// slice element.
func sliceNString(text string, splitStr string) ([]string, int) {
	textSlice := strings.Split(text, splitStr)

	var longest int

	for _, v := range textSlice {
		var runeSlice []rune
		for _, vv := range v {
			runeSlice = append(runeSlice, vv)
		}

		if len(v) > longest {
			longest = len(runeSlice)
		}
	}

	return textSlice, longest
}

func draw(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(1000, 1000)
	defer s.End()

	r := firstRect()
	r.draw(s, "abcdefghi jklmnopqr stuvwxyz")
	r.draw(s, "1234567890 1234567890 1234567890")
	r.draw(s, "this is an odd sentence")
	r.draw(s, "世a界世bcd界efg 世h世i世j世k世l世")
}
