package bot

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/messages"
	"github.com/tucnak/telebot"
)

func (b *Bot) LoggerInfo(m *telebot.Message) {
	message := messages.Message{Tgmessage: m}

	b.logger.Info(message.Command(), "attrs", message.UserInfo())
}
