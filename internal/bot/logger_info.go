package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/logmessage"
	"gopkg.in/telebot.v4"
)

func (b *Bot) LoggerInfo(c telebot.Context) {
	m := logmessage.New(c)

	b.logger.Info(m.Data(), "type", m.Command(), "attrs", m.UserInfo())
}
