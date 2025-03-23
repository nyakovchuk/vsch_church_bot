package bot

import "github.com/nyakovchuk/vsch_church_bot/internal/bot/event"

func (b *Bot) Events() {
	event.HandleOnTextLocation(b)
	event.HandleOnLocation(b)
}
