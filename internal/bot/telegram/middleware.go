package telegram

import "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/middleware"

func (b *Bot) Middleware() {
	middleware.CheckUser(b)
	middleware.GetLanguage(b)
}
