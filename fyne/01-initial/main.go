package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	ap := app.New()
	win := ap.NewWindow("my window")
	wgt := widget.NewVBox(
		widget.NewLabel("my widget"),
		widget.NewButton("button1", func() {
			ap.Quit()
		}),
		widget.NewButton("button2", func() {
			ap.Quit()
		}),
		widget.NewButton("quit", func() {
			ap.Quit()
		}),
	)

	win.SetContent(wgt)

	win.Show()
	ap.Run()
}
