package main

import (
	"fmt"
	"image"
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

	//create a loop that keeps the window open, unless the close button in the corner is pushed
	//The function win.Update fetches new events (key presses, mouse moves and clicks, etc.) and redraws the window.
	angle := 0.0
	last := time.Now()
	mousePositions := []pixel.Vec{}
	for !win.Closed() {

		if win.JustPressed(pixelgl.KeySpace) {
			fmt.Println("Space pressed")
			mousePositions = append(mousePositions, win.MousePosition())
			fmt.Println(mousePositions)
		}

		deltaTime := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Skyblue)
		for i := range mousePositions {
			angle += 0.5 * float64(deltaTime)
			var mat pixel.Matrix //this one is not needed, just added for clarity
			mat = pixel.IM
			mat = mat.Scaled(pixel.ZV, 0.5)
			mat = mat.Moved(mousePositions[i])          //move sprite to mouse position
			mat = mat.Rotated(mousePositions[i], angle) //rotate it around the current mouse position

			sprite.Draw(win, mat)
		}
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
