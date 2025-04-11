package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/button/inline/radiusBtn"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"gopkg.in/telebot.v4"
)

func HandleOnCallback(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnCallback, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		radiusText := strings.TrimSpace(c.Callback().Data)

		radius := getRadius(radiusText)

		// убираем время задержки у кнопок
		c.Respond()

		err := bm.Services().User.UpdateUserRadius(context.Background(), c.Sender().ID, radius)
		if err != nil {
			// залогировать
			return nil
		}

		coords, err := bm.Services().Coordinates.GetCoordinates(context.Background(), c.Sender().ID)
		if err != nil {
			// залогировать
			return nil
		}

		// работает универсально или нужно привязываться к пользователю?
		findChurches := bm.Services().Distance.GetChurchesNearby(
			coords,
			radius,
			bm.SharedData().Churches,
		)

		text := fmt.Sprintf("⛪ Найдено церквей: <b>%d</b>, в радиусе <b>%d км</b>:\n\n", len(findChurches), radius/1000)

		for i, church := range findChurches {
			text += fmt.Sprintf("%d. ", i+1)
			text += buildChurchesText(church)
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

func buildChurchesText(church church.DtoTelegram) string {
	vschUrl := fmt.Sprintf("https://www.vsch.org/church/%s", church.Alias)
	text := fmt.Sprintf("<a href=\"%s\"><b>%s</b></a> (%s) – <b>[%.2f км]</b>\n", vschUrl, church.Name, church.Confession, church.Distance/1000)
	return text
}
