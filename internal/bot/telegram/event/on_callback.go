package event

import (
	"context"
	"fmt"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/language"
	nearestchurches "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/nearest_churches"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/radius"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/external"
	domainLanguage "github.com/nyakovchuk/vsch_church_bot/internal/domain/language"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

const (
	TOP3 = 3
)

func HandleOnCallback(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnCallback, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		// fmt.Println("language:", c.Get("lang"))

		// убираем время задержки у кнопок
		c.Respond()

		externalId := utils.Int64ToString(c.Sender().ID)
		external := external.ToModel(
			externalId,
			bm.SharedData().Platform,
		)

		tgHtml := ""
		data := strings.TrimSpace(c.Callback().Data)

		// Обработка кнопок выбора радиуса
		if strings.HasPrefix(data, radius.PrefixRadius) {

			radius := getRadius(data)

			// Возможно не нужно нам сохранять радиус в БД
			err := bm.Services().User.UpdateUserRadius(context.Background(), external, radius)
			if err != nil {
				return nil
			}

			html, err := searchChurchesByRadius(bm, external, radius)
			if err != nil {
				bm.LoggerError(c, err)
				return nil
			}
			tgHtml = html
		}

		// Обработка кнопки поиска ближайших церквей
		if strings.HasPrefix(data, nearestchurches.PrefixNearestChurches) {

			topN := TOP3

			html, err := findTopNNearestChurches(bm, external, topN)
			if err != nil {
				bm.LoggerError(c, err)
				return nil
			}
			tgHtml = html
		}

		// Обработка кнопки смены языка
		if strings.HasPrefix(data, language.PrefixLangCode) {

			langCode := getLangCode(data)

			langId := domainLanguage.LanguageList(bm.SharedData().Languages).
				ByCode(langCode).
				ID

			err := bm.Services().User.UpdateUserLang(context.Background(), external, langId)
			if err != nil {
				bm.LoggerError(c, err)
				return nil
			}

			c.Set("lang", langCode)

			tgHtml = fmt.Sprintf("Язык изменен на <b>%s</b>", langCode)
		}

		return c.Send(tgHtml, &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	})
}

func searchChurchesByRadius(bm BotManager, external external.External, radius int) (string, error) {
	userCoords, err := bm.Services().Coordinates.GetCoordinates(context.Background(), external)
	if err != nil {
		return "", err
	}

	findChurches := bm.Services().Distance.GetChurchesNearby(
		userCoords,
		radius,
		bm.SharedData().Churches,
	)

	html := fmt.Sprintf("⛪ В радиусе <b>%d км</b>,  найдено церквей: <b>%d</b>\n\n", radius/1000, len(findChurches))

	for i, church := range findChurches {
		html += fmt.Sprintf("<b>%d.</b> ", i+1)
		html += buildChurchesText(userCoords, church)
	}

	return html, nil
}

func findTopNNearestChurches(bm BotManager, external external.External, topN int) (string, error) {
	userCoords, err := bm.Services().Coordinates.GetCoordinates(context.Background(), external)
	if err != nil {
		return "", err
	}

	findChurches := bm.Services().Distance.FindTopNNearestChurches(
		userCoords,
		topN,
		bm.SharedData().Churches,
	)

	html := "⛪ Ближайшие церкви:\n\n"
	for i, church := range findChurches {
		html += fmt.Sprintf("<b>%d.</b> ", i+1)
		html += buildChurchesText(userCoords, church)
	}

	return html, nil
}

// return radius meters
func getRadius(key string) int {

	buttons := radius.NewButtonsMap()

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

func getLangCode(key string) string {

	buttons := language.NewButtonsMap()

	var language string

	switch key {
	case buttons.Buttons["uk"].Data:
		language = "uk"
	case buttons.Buttons["en"].Data:
		language = "en"
	case buttons.Buttons["ru"].Data:
		language = "ru"
	default:
		language = "en"
	}

	return language
}
