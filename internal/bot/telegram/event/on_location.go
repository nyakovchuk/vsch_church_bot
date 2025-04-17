package event

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func HandleOnLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnLocation, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		location := c.Message().Location
		if location == nil {
			return c.Send("Не удалось получить геолокацию.")
		}

		// cache["latitude"] = float64(location.Lat)
		// cache["longitude"] = float64(location.Lng)
		externalId := utils.Int64ToString(c.Sender().ID)
		coords := model.ToCoordinates(
			bm.SharedData().Platform.ID,
			externalId,
			float64(location.Lat),
			float64(location.Lng),
			false,
		)

		savedCoords, err := bm.Services().Coordinates.Save(context.Background(), coords)
		if err != nil {
			bm.LoggerError(c, err)
			return nil
		}

		text := fmt.Sprintf("Ваши кординаты: <code>%f, %f</code>", savedCoords.Latitude, savedCoords.Longitude)
		c.Send(text, &telebot.SendOptions{
			ParseMode: telebot.ModeHTML,
		})

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
