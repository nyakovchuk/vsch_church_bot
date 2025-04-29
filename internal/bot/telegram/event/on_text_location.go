package event

import (
	"context"
	"errors"

	options "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/churchsearch"
	coordinates "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"golang.org/x/text/message"
	"gopkg.in/telebot.v4"
)

func HandleOnTextLocation(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnText, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		langCode := c.Get("lang").(string)
		printer := i18n.Printer(langCode)

		coords, err := bm.Services().Coordinates.ParseCoordinates(c.Message().Text)
		if err != nil {
			latStr, lngStr, _ := bm.Services().Coordinates.SplitCoordinates(c.Message().Text)

			html := CoordinatesParseError(printer, err, latStr, lngStr)

			return c.Send(html, &telebot.SendOptions{
				ParseMode: telebot.ModeHTML,
			})
		}

		externalId := utils.Int64ToString(c.Sender().ID)
		coords.ExternalID = externalId
		coords.PlatformID = bm.SharedData().Platform.ID
		coords.IsOnText = true
		_, err = bm.Services().Coordinates.Save(context.Background(), coords)
		if err != nil {
			bm.LoggerError(c, err)
			return nil
		}

		churchsearchBtn := options.NewChurchSearchOptions(langCode)

		return c.Reply(printer.Sprintf("event.location.find_churches_in_radius"), churchsearchBtn.Display())
	})
}

func CoordinatesParseError(printer *message.Printer, err error, latStr, lngStr string) string {
	switch {
	case errors.Is(err, coordinates.ErrCorrectLatitude):
		return printer.Sprintf("event.text_location.invalid_latitude", latStr)
	case errors.Is(err, coordinates.ErrCorrectLongitude):
		return printer.Sprintf("event.text_location.invalid_longitude", lngStr)
	default:
		return printer.Sprintf("event.text_location.invalid_coordinates_format")
	}
}
