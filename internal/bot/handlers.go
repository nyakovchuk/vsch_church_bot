package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/handlers"
)

func (b *Bot) Handlers() {
	commandStart := b.Commands().GetByName("start")
	commandHelp := b.Commands().GetByName("help")

	b.bot.Handle(commandStart.Route, handlers.HandleStart(b))
	b.bot.Handle(commandHelp.Route, handlers.HandleHelp(b))
}
