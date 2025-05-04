package inline

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"gopkg.in/telebot.v4"
)

type SingleButton struct {
	IButtons   *button.TgBtn
	ButtonData *button.ButtonConfig
}

func (sb *SingleButton) Display() *telebot.ReplyMarkup {

	btnSingleButton := sb.Сreate()

	sb.IButtons.Reply.Inline(
		sb.IButtons.Reply.Row(btnSingleButton),
	)

	return sb.IButtons.Reply
}

func (sb *SingleButton) Сreate() telebot.Btn {
	return sb.IButtons.Reply.Data(
		sb.ButtonData.Label,
		sb.ButtonData.Data,
	)
}
