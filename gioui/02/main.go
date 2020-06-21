package main

import (
	"fmt"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/key"
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

		// Set the initial position of the rectangle
		var x0, y0, x1, y1 float32 = 100, 100, 150, 150

		// Start an event loop to read all the events happening
		// in window, and maybe do something upon an event.
		for e := range window.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				ops := &op.Ops{}

				paint.ColorOp{Color: color.RGBA{R: 0x80, A: 0xFF}}.Add(ops)
				paint.PaintOp{Rect: f32.Rect(x0, y0, x1, y1)}.Add(ops)
				fmt.Printf("%v %v %v %v", x0, y0, x1, y1)
				e.Frame(ops)
			}

			if k, ok := e.(key.Event); ok {
				fmt.Printf("k=%#v\n", k)
				switch k {
				case key.Event{Name: "↑"}:
					y0--
					y1--
				case key.Event{Name: "↓"}:
					y0++
					y1++
				case key.Event{Name: "←"}:
					x0--
					x1--
				case key.Event{Name: "→"}:
					x0++
					x1++
				}
			}
		}

	}()

	// app deals with the low level parts towards the platform.
	// App also forces us to give up the main() go routine, and
	// that is also why we have to start all work in it's own go
	// routine's before starting app.Main()
	app.Main()
}
