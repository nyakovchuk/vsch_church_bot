package telegram

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/logmessage"
	"gopkg.in/telebot.v4"
)

func (b *Bot) LoggerInfo(c telebot.Context) {
	m := logmessage.New(c)

	b.logger.Info(m.Data(), "type", m.Command(), "attrs", m.UserInfo())
}

func (b *Bot) LoggerError(c telebot.Context, err error) {
	m := logmessage.New(c)

	b.logger.Error(m.Data(), "type", m.Command(), "attrs", m.UserInfo(), "error", err)
}
