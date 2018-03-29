package main

import (
	"fmt"
	"image"
	"os"

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

	//load the picture from disk
	pic, err := loadPicture("sailor.png")
	if err != nil {
		fmt.Println("error: Failed loading picture : ", err)
	}

	//Create a sprite from the picture.
	//the second argument is for what part of the picture to use. Here we use the whole picture with pic.Bounds()
	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Skyblue)

	//draw the sprite in the window
	//first parameter is what window, second parameter is where to draw it relative to the windows 0,0 position which
	//is the lower left corner of the window.
	sprite.Draw(win, pixel.IM)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	//create a loop that keeps the window open, unless the close button in the corner is pushed
	//The function win.Update fetches new events (key presses, mouse moves and clicks, etc.) and redraws the window.
	for !win.Closed() {
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
