package event

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

// Узнать где лучше преобразовівать координаты
// в сервисе или перед ним
func HandleOnTextLocation(bm BotManager, cache map[string]interface{}, radiusBtn ButtonRenderer) {
	bm.TBot().Handle(telebot.OnText, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		coords, err := bm.Services().Coordinates.ParseCoordinates(c.Message().Text)
		if err != nil {
			errText := `<b>Некорректный формат координат</b>
Пример правильного формата координат:
<code>50.4228 30.3145</code> <i>(широта, долгота через пробел)</i>
<code>50.4228, 30.3145</code> <i>(широта, долгота через запятую)</i>.
			`
			return c.Send(errText, &telebot.SendOptions{
				ParseMode: telebot.ModeHTML,
			})
		}

		externalId := utils.Int64ToString(c.Sender().ID)
		coords.ExternalID = externalId
		coords.PlatformID = bm.SharedData().Platform.ID
		coords.IsOnText = true
		savedCoords, err := bm.Services().Coordinates.Save(context.Background(), coords)
		if err != nil {
			bm.LoggerError(c, err)
			return nil
		}

		_ = fmt.Sprintf("Ваши кординаты: %f, %f", savedCoords.Latitude, savedCoords.Longitude)

		return c.Reply("Найти ближайшие церкви в радиусе:", radiusBtn.Display())
	})
}
