package main

import (
	"fmt"

	"gioui.org/app"
)

func main() {
	go func() {
		w := app.NewWindow()
		for e := range w.Events() {
			fmt.Printf("%#v\n", e)
		}
	}()
	app.Main()
}
