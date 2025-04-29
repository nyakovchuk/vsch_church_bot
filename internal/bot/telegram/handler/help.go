package handler

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"gopkg.in/telebot.v4"
)

func HandleHelp(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		printer := i18n.Printer(c.Get("lang").(string))

		return c.Send(printer.Sprintf("command.help"), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}
