package commandType

import "gopkg.in/telebot.v4"

type TextEvent struct {
	tg telebot.Context
}

func NewTextEvent(tg telebot.Context) *TextEvent {
	return &TextEvent{tg: tg}
}

func (m *TextEvent) Command() string {
	return "Text"
}

func (m *TextEvent) Data() string {
	return m.tg.Message().Text
}
