package handler

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"gopkg.in/telebot.v4"
)

func HandleLanguage(bm BotManager, buttons ButtonRenderer) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		printer := i18n.Printer(c.Get("lang").(string))

		return c.Send(printer.Sprintf("command.language"), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		}, buttons.Display())
	}
}
