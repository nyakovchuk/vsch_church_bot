package message

import (
	ct "github.com/nyakovchuk/vsch_church_bot/internal/bot/message/commandType"
	"gopkg.in/telebot.v4"
)

func GetTypeCommand(bot telebot.Context) CommandInfo {
	switch {
	case bot.Message().Location != nil:
		le := ct.NewLocationEvent(bot)
		return le
	case isTextMessage(bot.Message().Text):
		te := ct.NewTextEvent(bot)
		return te
	default:
		cm := ct.NewCommandMessage(bot)
		return cm
	}
}

func isTextMessage(text string) bool {
	if len(text) > 0 && text[0] == '/' {
		return false
	}

	return true
}
