package main

import (
	"fmt"

	"github.com/faiface/pixel/imdraw"

	_ "image/png"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	windowX1   float64
	windowY1   float64
	windowX2   float64 = 500
	windowY2   float64 = 500
	radius     float64 = 100
	mainRadius float64 = 90
)

func run() {
	//pixelgl.WindowConfig is a struct that lets us do the initial config of a window
	//setting Vsync=true will enable the win.Update to follow the displays refresh interval,
	//so it does not consume all the cpu by running constantly and as fast as it can.
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

	//create a loop that keeps the window open, unless the close button in the corner is pushed
	//The function win.Update fetches new events (key presses, mouse moves and clicks, etc.) and redraws the window.
	for !win.Closed() {

		//deltaTime := time.Since(last).Seconds()
		//last = time.Now()

		win.Clear(colornames.Skyblue)

		something := imdraw.New(nil)

		/*		for i := 0; i < len(triPosX); i++ {
				something.Push(pixel.V(triPosX[i], triPosY[i]))
			}*/

		something.Line(2)
		something.Draw(win)

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
