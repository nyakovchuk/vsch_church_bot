package event

import (
	"context"
	"fmt"

	"gopkg.in/telebot.v4"
)

// Узнать где лучше преобразовівать координаты
// в сервисе или перед ним
func HandleOnTextLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnText, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		coords, err := bm.Services().Coordinates.ParseCoordinates(c.Message().Text)
		if err != nil {
			return c.Send(err.Error())
		}

		coords.IsOnText = true
		savedCoords, err := bm.Services().Coordinates.Save(context.Background(), coords)
		if err != nil {
			return c.Send(err.Error())
		}

		_ = fmt.Sprintf("Ваши кординаты: %f, %f", savedCoords.Latitude, savedCoords.Longitude)

		// записать в БД
		// bm.Services().Coordinates.Create(lat, lon)
		// cache["latitude"] = savedCoords.Latitude
		// cache["longitude"] = savedCoords.Longitude

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
