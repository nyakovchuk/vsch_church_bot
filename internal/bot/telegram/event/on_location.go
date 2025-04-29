package event

import (
	"context"

	options "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/churchsearch"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func HandleOnLocation(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnLocation, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		langCode := c.Get("lang").(string)

		printer := i18n.Printer(langCode)

		location := c.Message().Location
		if location == nil {
			return c.Send(printer.Sprintf("event.location.geolocation_error"))
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

		latStr, lngStr := savedCoords.ToGeoString()

		text := printer.Sprintf("event.location.your_coordinates", latStr, lngStr)
		c.Send(text, &telebot.SendOptions{
			ParseMode: telebot.ModeHTML,
		})

		churchsearchBtn := options.NewChurchSearchOptions(langCode)

		return c.Reply(printer.Sprintf("event.location.find_churches_in_radius"), churchsearchBtn.Display())
	})
}
