package main

import (
	"atui/internal/debug"
	"atui/pkg/ui"
	"github.com/rivo/tview"
)

func main() {
	closeFunc := debug.UseDebugMode()
	defer closeFunc()

	app := tview.NewApplication().EnableMouse(true)
	prompt := ui.NewPrompt(app)

	pane := tview.NewFlex().SetDirection(tview.FlexRow)
	pane.AddItem(prompt.View(), 5, 1, false)

	view := tview.NewTextView().
		SetBorder(true)
	pane.AddItem(view, 10, 3, false)
	if err := app.SetRoot(pane, true).Run(); err != nil {
		panic(err)
	}
}
