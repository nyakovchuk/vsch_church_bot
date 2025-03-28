package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/handler"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/button/reply"
	"gopkg.in/telebot.v4"
)

func (b *Bot) Handlers() {
	commandStart := b.Commands().GetByName("start")
	commandHelp := b.Commands().GetByName("help")

	b.bot.Handle(commandStart.Route, handler.HandleStart(b))
	b.bot.Handle(commandHelp.Route, handler.HandleHelp(b))

	b.bot.Handle("/reply_btns", func(c telebot.Context) error {
		return c.Send("📋 Выберите команду из меню:", reply.CreateMenuKeyboard())
	})
}
