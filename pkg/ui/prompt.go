package ui

import (
	"atui/pkg/model"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PromptModel interface {
	Add(r rune)
	InputComplete()
}

type Prompt struct {
	app   *tview.Application
	view  *tview.TextView
	model PromptModel
}

func NewPrompt(app *tview.Application) *Prompt {
	v := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	v.SetBorder(true)
	v.SetText(" >")

	p := &Prompt{
		app:   app,
		view:  v,
		model: model.NewPromptModel(),
	}

	p.view.SetInputCapture(p.handleKeyInput)
	return p
}

func (p *Prompt) View() *tview.TextView {
	return p.view
}

func (p *Prompt) handleKeyInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		p.model.InputComplete()
		p.view.SetText(" >")
	case tcell.KeyRune:
		p.model.Add(event.Rune())
		fmt.Fprintf(p.view, "%c", event.Rune())
	}
	return event
}
