package menu

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/commands"
	"gopkg.in/telebot.v4"
)

func Create(bot *telebot.Bot) {
	bot.SetCommands(commands.ToTelebotCommands())
}
