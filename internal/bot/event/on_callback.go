package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/button/inline/radiusBtn"
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
			fmt.Println("error update radius:", err)
			return nil
		}

		text1 := fmt.Sprintf("Вы выбрали радиус %d км.", radius)

		// передать координаты в сервис
		// getNearbyChurches(координаты, радиус)

		return c.Send(text1)
	})
}

func getRadius(key string) int {

	buttons := radiusBtn.NewButtonSet()

	var radius int

	switch key {
	case buttons.Buttons["five"].Data:
		radius = 5
	case buttons.Buttons["ten"].Data:
		radius = 10
	case buttons.Buttons["thirty"].Data:
		radius = 30
	default:
		radius = 10
	}

	return radius
}
