package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel/imdraw"

	_ "image/png"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	windowX1 float64
	windowY1 float64
	windowX2 float64 = 500
	windowY2 float64 = 500
)

type mainPoint struct {
	centreX float64
	centreY float64
	radius  float64
	angle   float64
	x       float64
	y       float64
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "test window",
		Bounds: pixel.R(windowX1, windowY1, windowX2, windowY2),
		VSync:  true,
	}

	//create the new window
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Println("error: Creating window : ", err)
		panic(err)
	}

	//will make the lines smooth when sprites are rotated etc. If not set the lines of sprites will become fuzzy.
	win.SetSmooth(true)
	//last := time.Now()

	mp := mainPoint{
		radius: 1,
		angle:  math.Pi / 3,
	}

	mp.centreX = windowX2 / 2
	mp.centreY = windowY2 / 2

	for !win.Closed() {
		//deltaTime := time.Since(last).Seconds()
		//last = time.Now()

		if mp.x >= windowX2 {
			mp.centreX = mp.x
			mp.centreY = mp.y
			mp.radius = 0
			mp.angle += math.Pi / 7
		}
		if mp.x < windowX1 {
			mp.centreX = mp.x
			mp.centreY = mp.y
			mp.radius = 0
			mp.angle += math.Pi / 7
		}
		if mp.y >= windowY2 {
			mp.centreY = mp.y
			mp.centreX = mp.x
			mp.radius = 0
			mp.angle += math.Pi / 7
		}
		if mp.y < windowY1 {
			mp.centreY = mp.y
			mp.centreX = mp.x
			mp.radius = 0
			mp.angle += math.Pi / 7
		}

		mp.radius += 4

		mp.x = mp.radius*math.Cos(mp.angle) + mp.centreX
		mp.y = mp.radius*math.Sin(mp.angle) + mp.centreY

		win.Clear(colornames.Skyblue)
		mainLine := imdraw.New(nil)
		mainLine.Push(pixel.V(mp.centreX, mp.centreY))
		mainLine.Push(pixel.V(mp.x, mp.y))

		mainLine.Line(2)
		mainLine.Draw(win)

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
