package event

import (
	"fmt"
	"strings"

	"gopkg.in/telebot.v4"
)

func HandleOnCallback(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnCallback, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		radiusText := strings.TrimSpace(c.Callback().Data)

		var radius int
		switch radiusText {
		case "radius_five":
			radius = 5
		case "radius_ten":
			radius = 10
		case "radius_thirty":
			radius = 30
		default:
			radius = 10
		}

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
