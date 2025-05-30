package handler

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/reply"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"gopkg.in/telebot.v4"
)

func HandleStart(bm BotManager) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		langCode := c.Get("lang").(string)

		btnLocation := reply.BtnLocation(langCode)

		printer := i18n.Printer(langCode)

		return c.Send(printer.Sprintf("command.start"), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		}, btnLocation)
	}
}
