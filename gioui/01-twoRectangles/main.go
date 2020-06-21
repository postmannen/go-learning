package main

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {

	go func() {
		// Overall description:
		// 1. Create a window
		// 2. create an operations buffer.
		// 3. assign an operation (for example draw a rectangle) to the buffer.
		// 4. Then call frame event to redraw the window.
		window := app.NewWindow(app.Size(unit.Dp(400), unit.Dp(400)))

		// Start an event loop to read all the events happening
		// in window, and maybe do something upon an event.
		for e := range window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				ops := &op.Ops{}

				paint.ColorOp{Color: color.RGBA{R: 0x80, A: 0xFF}}.Add(ops)
				paint.PaintOp{Rect: f32.Rect(0, 0, 100, 100)}.Add(ops)
				paint.ColorOp{Color: color.RGBA{R: 0x10, A: 0xFF}}.Add(ops)
				paint.PaintOp{Rect: f32.Rect(100, 100, 200, 200)}.Add(ops)
				e.Frame(ops)
			}
		}

	}()

	// app deals with the low level parts towards the platform.
	// App also forces us to give up the main() go routine, and
	// that is also why we have to start all work in it's own go
	// routine's before starting app.Main()
	app.Main()
}
