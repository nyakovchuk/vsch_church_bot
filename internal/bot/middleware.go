package bot

import "github.com/nyakovchuk/vsch_church_bot/internal/bot/middleware"

func (b *Bot) Middleware() {
	middleware.CheckUser(b)
}
