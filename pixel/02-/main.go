package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"time"

	_ "image/png"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	//pixelgl.WindowConfig is a struct that lets us do the initial config of a window
	//setting Vsync=true will enable the win.Update to follow the displays refresh interval,
	//so it does not consume all the cpu by running constantly and as fast as it can.
	cfg := pixelgl.WindowConfig{
		Title:  "test window",
		Bounds: pixel.R(0, 0, 500, 500),
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

	//load the picture from disk
	pic, err := loadPicture("sailor.png")
	if err != nil {
		fmt.Println("error: Failed loading picture : ", err)
	}

	//Create a sprite from the picture.
	//the second argument is for what part of the picture to use. Here we use the whole picture with pic.Bounds()
	sprite := pixel.NewSprite(pic, pic.Bounds())

	//angle := 0.0
	var angle []float64
	last := time.Now()

	//slice to remember all the different sprite positions when key was pressed
	spritePosition := []pixel.Vec{}
	x := []float64{}
	y := []float64{}

	//create a loop that keeps the window open, unless the close button in the corner is pushed
	//The function win.Update fetches new events (key presses, mouse moves and clicks, etc.) and redraws the window.
	for !win.Closed() {

		if win.JustPressed(pixelgl.KeySpace) {
			fmt.Println("Space pressed")
			spritePosition = append(spritePosition, pixel.V(250, 250))
			angle = append(angle, 0.0)
			x = append(x, 0.0)
			y = append(y, 0.0)
		}

		deltaTime := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)
		for i := range spritePosition {
			angle[i] += 0.5 * float64(deltaTime)
			var mat pixel.Matrix //this one is not needed, just added for clarity
			mat = pixel.IM
			mat = mat.Scaled(pixel.ZV, 0.3)
			mat = mat.Moved(pixel.V(math.Cos(angle[i])*spritePosition[i].X, math.Cos(angle[i])*spritePosition[i].Y))
			mat = mat.Rotated(spritePosition[i], angle[i])

			sprite.Draw(win, mat)

			if spritePosition[i].X > 500 || spritePosition[i].Y > 500 {
				spritePosition[i].X = 0.0
				spritePosition[i].Y = 0.0
			} else {

				spritePosition[i].X += 1
				spritePosition[i].Y += 1
			}

		}

		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		win.Update()

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
