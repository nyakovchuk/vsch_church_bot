package handlers

import (
	"github.com/tucnak/telebot"
)

const (
	OutputStart = "Привет! Я ваш Telegram-бот!!!"
)

func HandleStart(bm BotManager) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		bm.LoggerInfo(m)
		bm.TBot().Send(m.Chat, OutputStart)
	}
}
