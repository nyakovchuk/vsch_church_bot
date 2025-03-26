package event

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func HandleOnLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnLocation, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		location := c.Message().Location
		if location == nil {
			return c.Send("Не удалось получить геолокацию.")
		}

		// записать в БД
		cache["latitude"] = float64(location.Lat)
		cache["longitude"] = float64(location.Lng)

		text := fmt.Sprintf("Ваши кординаты: %f, %f", location.Lat, location.Lng)
		c.Send(text)

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
