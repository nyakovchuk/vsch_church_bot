package event

import (
	"gopkg.in/telebot.v4"
)

// Узнать где лучше преобразовівать координаты
// в сервисе или перед ним
func HandleOnTextLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnText, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		// 1) преобразовать текст в координаты
		lat, lon, err := bm.Services().CoordinatesService.ParseCoordinates(c.Message().Text)
		if err != nil {
			return c.Send(err.Error())
		}

		cache["latitude"] = lat
		cache["longitude"] = lon

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
