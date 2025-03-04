package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/handlers"
)

type routs map[string]string

// Commands for the bot
var (
	commands = routs{
		"start": "/start",
		"help":  "/help",
	}
)

func (b *Bot) Handlers() {
	b.bot.Handle(commands["start"], handlers.HandleStart(b))
	b.bot.Handle(commands["help"], handlers.HandleHelp(b))
}
