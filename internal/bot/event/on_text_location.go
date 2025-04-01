package event

import (
	"gopkg.in/telebot.v4"
)

// Узнать где лучше преобразовівать координаты
// в сервисе или перед ним
func HandleOnTextLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnText, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		coordinates, err := bm.Services().Coordinates.ParseCoordinates(c.Message().Text)
		if err != nil {
			return c.Send(err.Error())
		}

		// записать в БД
		// bm.Services().Coordinates.Create(lat, lon)
		cache["latitude"] = coordinates.Latitude
		cache["longitude"] = coordinates.Longitude

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
