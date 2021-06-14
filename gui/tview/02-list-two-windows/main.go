package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	nodes1 := tview.NewList().ShowSecondaryText(false)
	nodes1.SetBorder(true).SetTitle("nodes1")
	nodes2 := tview.NewList().ShowSecondaryText(false)
	nodes2.SetBorder(true).SetTitle("nodes1")

	flex := tview.NewFlex().AddItem(nodes1, 0, 2, true).AddItem(nodes2, 0, 2, false)

	ships := []string{"ship1", "ship2", "ship3"}

	for _, v := range ships {
		nodes1.AddItem(v, "", rune(0), nodesSelected(nodes1, nodes2, app))
	}

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func nodesSelected(nodes1 *tview.List, nodes2 *tview.List, a *tview.Application) func() {
	shipsdb := map[string]string{
		"ship1": "ship one",
		"ship2": "ship two",
		"ship3": "ship three",
	}

	f := func() {
		index := nodes1.GetCurrentItem()
		text, _ := nodes1.GetItemText(index)
		nodes2.Clear()
		nodes2.AddItem(fmt.Sprintf("%v", shipsdb[text]), "", '0', nil)
	}

	return f
}
