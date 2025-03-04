package handlers

import (
	"github.com/tucnak/telebot"
)

const (
	OutputHelp = "Введите команду, чтобы узнать, что я могу."
)

func HandleHelp(b Bot) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		b.LoggerInfo(m)
		b.TBot().Send(m.Chat, OutputHelp)
	}
}
