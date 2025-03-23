package menu

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/command"
	"gopkg.in/telebot.v4"
)

func Create(bot *telebot.Bot) {
	bot.SetCommands(command.ToTelebotCommands())
}
