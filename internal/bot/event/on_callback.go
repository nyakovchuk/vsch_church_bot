package event

import (
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

		c.Respond()

		text1 := fmt.Sprintf("Вы выбрали радиус %d км.", radius)

		var lat, lon float64
		if value, ok := cache["latitude"].(float64); ok {
			lat = value
		}

		if value, ok := cache["longitude"].(float64); ok {
			lon = value
		}

		fmt.Println("lat:", lat, "lon:", lon)

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
