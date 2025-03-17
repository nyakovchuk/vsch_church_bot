package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/messages"
	"gopkg.in/telebot.v4"
)

func (b *Bot) LoggerInfo(c telebot.Context) {
	message := messages.NewMessage(c)

	b.logger.Info(message.Command(), "data", message.Data(), "attrs", message.UserInfo())
}
