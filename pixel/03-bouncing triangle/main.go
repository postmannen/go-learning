package main

import (
	"fmt"
	"image"
	"os"

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

	x1 := []float64{100, 200, 100, 100}
	y1 := []float64{100, 200, 200, 100}
	x1Direction := []bool{true, true, true, true}
	y1Direction := []bool{true, true, true, true}
	angle := []float64{1, 2, 3, 4}

	//last := time.Now()

	//create a loop that keeps the window open, unless the close button in the corner is pushed
	//The function win.Update fetches new events (key presses, mouse moves and clicks, etc.) and redraws the window.
	for !win.Closed() {

		//deltaTime := time.Since(last).Seconds()
		//last = time.Now()

		win.Clear(colornames.Skyblue)

		something := imdraw.New(nil)

		for i := 0; i < len(x1); i++ {
			something.Push(pixel.V(x1[i], y1[i]))
		}

		something.Line(2)
		something.Draw(win)

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Update()

		for i := 0; i < len(x1); i++ {
			if x1[i] >= windowX2 && x1Direction[i] == true {
				x1Direction[i] = false
			}
			if x1[i] <= windowX1 && x1Direction[i] == false {
				x1Direction[i] = true
			}

			if y1[i] >= windowY2 && y1Direction[i] == true {
				y1Direction[i] = false
			}
			if y1[i] <= windowY1 && y1Direction[i] == false {
				y1Direction[i] = true
			}
		}

		for i := 0; i < len(x1); i++ {
			if x1Direction[i] == true {
				x1[i] = x1[i] + 4
			} else {
				x1[i] = x1[i] - 4
			}

			if y1Direction[i] == true {
				y1[i] = y1[i] + 5
			} else {
				y1[i] = y1[i] - 5
			}

			angle[i] = angle[i] + 0.01
		}
		x1[len(x1)-1] = x1[0]
		y1[len(x1)-1] = y1[0]

	}
}

//loadPicture loads the picture from the path given as input, and returns a pixel.Picture
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil

}

func main() {
	pixelgl.Run(run)
}
