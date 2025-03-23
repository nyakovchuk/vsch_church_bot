package event

import (
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"gopkg.in/telebot.v4"
)

// Узнать где лучше преобразовівать координаты
// в сервисе или перед ним
func HandleOnTextLocation(bm BotManager) {
	bm.TBot().Handle(telebot.OnText, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		// 1) преобразовать текст в координаты
		lat, lon, err := service.ParseCoordinates(c.Message().Text)
		if err != nil {
			return c.Send(err.Error())
		}

		text := fmt.Sprintf("Кординаты:%f, %f", lat, lon)

		// 2) Сохранить в контекст

		// 3) Вывести информацию
		// Найти церкви в радиусе (5,10, 30км)

		// 4) передать координаты в сервис
		// getNearbyChurches(координаты, радиус)
		return c.Send(text)
	})
}
