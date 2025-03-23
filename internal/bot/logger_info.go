package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/message"
	"gopkg.in/telebot.v4"
)

func (b *Bot) LoggerInfo(c telebot.Context) {
	m := message.NewMessage(c)

	b.logger.Info(m.Data(), "type", m.Command(), "attrs", m.UserInfo())
}
