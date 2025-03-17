package bot

import "github.com/nyakovchuk/vsch_church_bot/internal/bot/events"

func (b *Bot) Events() {
	events.HandleOnLocation(b)
}
