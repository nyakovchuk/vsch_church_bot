package handlers

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/reply_buttons"
	"gopkg.in/telebot.v4"
)

const (
	commandText = "Поделитесь своим местоположением (только для мобильных)\nС компьютера: отправить ➜ местоположение"
)

func HandleLocation(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		btn := reply_buttons.BtnLocation()

		return c.Send(commandText, btn)
	}
}
