package reply

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"gopkg.in/telebot.v4"
)

func BtnLocation(langCode string) *telebot.ReplyMarkup {

	btn := &telebot.ReplyMarkup{ResizeKeyboard: true}

	printer := i18n.Printer(langCode)

	btnLocation := btn.Location(printer.Sprintf("button.send_location"))

	btn.Reply(btn.Row(btnLocation))

	return btn
}
