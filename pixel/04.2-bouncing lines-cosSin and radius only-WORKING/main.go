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

	angleAdd := 1.3
	const points int = 50

	mp := [points]mainPoint{}
	mp[0].angle = math.Pi / 3

	for i := 0; i < len(mp); i++ {
		mp[i].centreX = windowX2 / 2
		mp[i].centreY = windowY2 / 2
		mp[i].angle = math.Pi / float64(points) * float64(i) //check why setting 4 makes the ends to meet
	}
	for !win.Closed() {
		//deltaTime := time.Since(last).Seconds()
		//last = time.Now()
		for i := 0; i < len(mp); i++ {
			if mp[i].x >= windowX2 {
				mp[i].centreX = windowX2
				mp[i].centreY = mp[i].y
				mp[i].radius = 0
				mp[i].angle += angleAdd
			}
			if mp[i].x < windowX1 {
				mp[i].centreX = windowX1
				mp[i].centreY = mp[i].y
				mp[i].radius = 0
				mp[i].angle += angleAdd
			}
			if mp[i].y >= windowY2 {
				mp[i].centreY = windowY2
				mp[i].centreX = mp[i].x
				mp[i].radius = 0
				mp[i].angle += angleAdd
			}
			if mp[i].y < windowY1 {
				mp[i].centreY = windowY1
				mp[i].centreX = mp[i].x
				mp[i].radius = 0
				mp[i].angle += angleAdd
			}

			mp[i].radius += 4

			mp[i].x = mp[i].radius*math.Cos(mp[i].angle) + mp[i].centreX
			mp[i].y = mp[i].radius*math.Sin(mp[i].angle) + mp[i].centreY
			mp[len(mp)-1] = mp[0]
		}

		mainLine := imdraw.New(nil)

		for i := 0; i < len(mp); i++ {

			win.Clear(colornames.Skyblue)
			//mainLine.Push(pixel.V(mp[i].centreX, mp[i].centreY))
			mainLine.Push(pixel.V(mp[i].x, mp[i].y))
		}

		//mainLine.Circle(20, 2)
		mainLine.Line(2)
		//mainLine.Rectangle(2)
		mainLine.Draw(win)

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
