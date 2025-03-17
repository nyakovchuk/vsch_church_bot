package messages

import (
	tc "github.com/nyakovchuk/vsch_church_bot/internal/bot/messages/type_commands"
	"gopkg.in/telebot.v4"
)

func GetTypeCommand(bot telebot.Context) CommandInfo {
	switch {
	case bot.Message().Location != nil:
		le := tc.NewLocationEvent(bot)
		return le
	default:
		cm := tc.NewCommandMessage(bot)
		return cm
	}
}
