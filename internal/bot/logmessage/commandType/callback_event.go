package commandType

import (
	"strings"

	"gopkg.in/telebot.v4"
)

type CallbackEvent struct {
	tg telebot.Context
}

func NewCallbackEvent(tg telebot.Context) *CallbackEvent {
	return &CallbackEvent{tg: tg}
}

func (c *CallbackEvent) Command() string {
	return "OnCallback"
}

func (c *CallbackEvent) Data() string {
	data := strings.TrimSpace(c.tg.Callback().Data)
	return data
}
