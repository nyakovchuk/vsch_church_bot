package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/handler"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/inline_buttons"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/reply_buttons"
	"gopkg.in/telebot.v4"
)

func (b *Bot) Handlers() {
	commandStart := b.Commands().GetByName("start")
	commandHelp := b.Commands().GetByName("help")
	commandLocation := b.Commands().GetByName("location")

	b.bot.Handle(commandStart.Route, handler.HandleStart(b))
	b.bot.Handle(commandHelp.Route, handler.HandleHelp(b))
	b.bot.Handle(commandLocation.Route, handler.HandleLocation(b))

	btns := inline_buttons.NewButtons()
	b.bot.Handle("/inline_btns", func(c telebot.Context) error {
		return c.Send("Выберите действие:", btns.Display())
	})

	b.bot.Handle("/reply_btns", func(c telebot.Context) error {
		return c.Send("📋 Выберите команду из меню:", reply_buttons.CreateMenuKeyboard())
	})
}
