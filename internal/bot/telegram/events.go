package telegram

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/event"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/radiusBtn"
)

func (b *Bot) Events() {
	radiusBtns := radiusBtn.New()

	// вместо sharedВata использовать БД
	sharedData := make(map[string]interface{})

	event.HandleOnLocation(b, sharedData, radiusBtns)
	event.HandleOnTextLocation(b, sharedData, radiusBtns)
	event.HandleOnCallback(b, sharedData)
}
