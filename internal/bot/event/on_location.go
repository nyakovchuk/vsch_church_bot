package event

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func HandleOnLocation(bm BotManager) {
	bm.TBot().Handle(telebot.OnLocation, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		location := c.Message().Location
		if location == nil {
			return c.Send("Не удалось получить геолокацию.")
		}

		return c.Send(
			fmt.Sprintf("📍 Получены координаты:\nШирота: %.5f\nДолгота: %.5f",
				location.Lat, location.Lng),
		)
	})
}
