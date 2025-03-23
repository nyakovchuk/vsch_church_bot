package handler

import (
	"gopkg.in/telebot.v4"
)

const (
	OutputStart = "Привет! Я ваш Telegram-бот!!!"
)

func HandleStart(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		return c.Send(OutputStart)
	}
}
