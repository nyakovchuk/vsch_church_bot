package handler

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"gopkg.in/telebot.v4"
)

func HandleChurchesCount(bm BotManager) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		langCode := c.Get("lang").(string)

		printer := i18n.Printer(langCode)

		return c.Send(printer.Sprintf("command.churches_count"), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}
