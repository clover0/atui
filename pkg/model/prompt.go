package model

import (
	"atui/internal/debug"
)

type Listener interface {
	Complete()
}

type PromptModel struct {
	buff      []rune
	listeners []Listener
}

func NewPromptModel() *PromptModel {
	return &PromptModel{buff: make([]rune, 0)}
}

func (p *PromptModel) Add(r rune) {
	p.buff = append(p.buff, r)
}

func (p *PromptModel) InputComplete() {
	debug.Printf(string(p.buff))
}
