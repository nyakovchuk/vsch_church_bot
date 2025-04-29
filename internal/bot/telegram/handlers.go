package telegram

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/handler"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/language"
)

func (b *Bot) Handlers() {
	commandStart := b.Commands().GetByName("start")
	commandHelp := b.Commands().GetByName("help")
	commandLanguage := b.Commands().GetByName("language")

	languageBtns := language.NewButtons()

	b.bot.Handle(commandStart.Route, handler.HandleStart(b))
	b.bot.Handle(commandHelp.Route, handler.HandleHelp(b))
	b.bot.Handle(commandLanguage.Route, handler.HandleLanguage(b, languageBtns))
	b.bot.Handle("/getdb", handler.HandleGetDB(b))
}
