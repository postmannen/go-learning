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
	radius   float64 = 100
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

	xShapePos := []float64{100, 200, 100, 100}
	yShapePos := []float64{100, 200, 200, 100}
	x := []float64{0, 0, 0, 0}
	y := []float64{0, 0, 0, 0}

	vDirection := []bool{false, false, false, false}

	angle := []float64{0, 0, 0, 0}

	//last := time.Now()

	//create a loop that keeps the window open, unless the close button in the corner is pushed
	//The function win.Update fetches new events (key presses, mouse moves and clicks, etc.) and redraws the window.
	for !win.Closed() {

		//check if one of the corners hit the wall. If so, change direction.
		for i := 0; i < len(x); i++ {
			if x[i] < (windowX1+10) && vDirection[i] == false {
				vDirection[i] = true
			}
			if x[i] > (windowX2-10) && vDirection[i] == true {
				vDirection[i] = false
			}

			if y[i] < (windowY1+10) && vDirection[i] == false {
				vDirection[i] = true
			}
			if y[i] > (windowY2-10) && vDirection[i] == true {
				vDirection[i] = false
			}
		}

		for i := 0; i < len(x); i++ {
			if vDirection[i] == true {
				angle[i] += 0.05
				xShapePos[i]++
				yShapePos[i]++
				fmt.Println("direction er true")
			} else {
				angle[i] -= 0.05
				xShapePos[i]--
				yShapePos[i]--
				fmt.Println("direction er false")
			}

			fmt.Println(x)
			x[i] = (radius * math.Cos(angle[i])) + xShapePos[i]
			y[i] = (radius * math.Sin(angle[i])) + yShapePos[i]

		}

		//set the last vector to the same as the first to make a triangle
		x[len(x)-1] = x[0]
		y[len(x)-1] = y[0]

		//deltaTime := time.Since(last).Seconds()
		//last = time.Now()

		win.Clear(colornames.Skyblue)

		something := imdraw.New(nil)

		for i := 0; i < len(x); i++ {
			something.Push(pixel.V(x[i], y[i]))
		}

		something.Line(2)
		something.Draw(win)

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Update()

	}
}

func main() {
	pixelgl.Run(run)
}
