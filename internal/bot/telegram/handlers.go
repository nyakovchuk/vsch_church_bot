package telegram

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/handler"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/middleware"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/language"
)

func (b *Bot) Handlers() {
	commandStart := b.Commands().GetByName("start")
	commandHelp := b.Commands().GetByName("help")
	commandLanguage := b.Commands().GetByName("language")
	commandChurchesCount := b.Commands().GetByName("churches_count")

	b.bot.Handle(commandStart.Route, handler.HandleStart(b))
	b.bot.Handle(commandHelp.Route, handler.HandleHelp(b))

	languageBtns := language.NewButtons()
	b.bot.Handle(commandLanguage.Route, handler.HandleLanguage(b, languageBtns))

	b.bot.Handle(commandChurchesCount.Route, handler.HandleChurchesCount(b))

	b.bot.Handle("/getdb", handler.HandleGetDB(b), middleware.AdminOnly)
}
