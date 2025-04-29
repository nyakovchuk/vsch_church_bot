package telegram

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/event"
)

func (b *Bot) Events() {

	// вместо sharedВata использовать БД
	sharedData := make(map[string]interface{})

	event.HandleOnLocation(b, sharedData)
	event.HandleOnTextLocation(b, sharedData)
	event.HandleOnCallback(b, sharedData)
}
