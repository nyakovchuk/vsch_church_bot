package handlers

import (
	"github.com/tucnak/telebot"
)

const (
	OutputStart = "Привет! Я ваш Telegram-бот!!!"
)

func HandleStart(b Bot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		b.LoggerInfo(m)
		b.TBot().Send(m.Chat, OutputStart)
	}
}
