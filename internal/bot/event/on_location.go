package event

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"gopkg.in/telebot.v4"
)

func HandleOnLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnLocation, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		location := c.Message().Location
		if location == nil {
			return c.Send("Не удалось получить геолокацию.")
		}

		// Добавить тип координат или ontext или onlocation

		// cache["latitude"] = float64(location.Lat)
		// cache["longitude"] = float64(location.Lng)

		coords := model.ToCoordinates(float64(location.Lat), float64(location.Lng))

		// добавить username
		// user.SaveCoordinates(ctx, username, coords)
		savedCoords, err := bm.Services().Coordinates.Save(context.Background(), coords)
		if err != nil {
			return c.Send(err.Error())
		}

		text := fmt.Sprintf("Ваши кординаты: %f, %f", savedCoords.Latitude, savedCoords.Longitude)
		c.Send(text)

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
