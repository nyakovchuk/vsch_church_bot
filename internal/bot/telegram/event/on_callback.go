package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/radiusBtn"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/external"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func HandleOnCallback(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnCallback, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		externalId := utils.Int64ToString(c.Sender().ID)
		external := external.ToModel(
			externalId,
			bm.SharedData().Platform,
		)

		radiusText := strings.TrimSpace(c.Callback().Data)
		radius := getRadius(radiusText)

		// убираем время задержки у кнопок
		c.Respond()

		err := bm.Services().User.UpdateUserRadius(context.Background(), external, radius)
		if err != nil {
			bm.LoggerError(c, err)
			return nil
		}

		userCoords, err := bm.Services().Coordinates.GetCoordinates(context.Background(), external)
		if err != nil {
			bm.LoggerError(c, err)
			return nil
		}

		// работает универсально или нужно привязываться к пользователю?
		findChurches := bm.Services().Distance.GetChurchesNearby(
			userCoords,
			radius,
			bm.SharedData().Churches,
		)

		text := fmt.Sprintf("⛪ В радиусе <b>%d км</b>,  найдено церквей: <b>%d</b>\n\n", radius/1000, len(findChurches))

		for i, church := range findChurches {
			text += fmt.Sprintf("<b>%d.</b> ", i+1)
			text += buildChurchesText(userCoords, church)
		}

		return c.Send(text, &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	})
}

// return radius meters
func getRadius(key string) int {

	buttons := radiusBtn.NewButtonSet()

	var radius int

	switch key {
	case buttons.Buttons["five"].Data:
		radius = 5000
	case buttons.Buttons["ten"].Data:
		radius = 10000
	case buttons.Buttons["thirty"].Data:
		radius = 30000
	default:
		radius = 0
	}

	return radius
}

func buildChurchesText(userCoords model.Coordinates, church church.DtoResponse) string {
	vschUrl := fmt.Sprintf("https://www.vsch.org/church/%s", church.Alias)
	text := fmt.Sprintf("<a href=\"%s\"><b>%s</b></a> (%s) – <b>[%.2f км]</b> <a href=\"https://www.google.com/maps/dir/%v,%v/%v,%v\">маршрут</a>\n", vschUrl, church.Name, church.Confession, church.Distance/1000, userCoords.Latitude, userCoords.Longitude, church.Latitude, church.Longitude)
	return text
}
