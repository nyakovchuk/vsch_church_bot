package telegram

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/event"
	options "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/churchsearch"
)

func (b *Bot) Events() {
	churchsearchBtn := options.NewChurchSearchOptions()

	// вместо sharedВata использовать БД
	sharedData := make(map[string]interface{})

	event.HandleOnLocation(b, sharedData, churchsearchBtn)
	event.HandleOnTextLocation(b, sharedData, churchsearchBtn)
	event.HandleOnCallback(b, sharedData)
}
