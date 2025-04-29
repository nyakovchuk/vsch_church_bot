package event

import (
	"context"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/handler"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/language"
	nearestchurches "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/nearest_churches"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/radius"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/external"
	domainLanguage "github.com/nyakovchuk/vsch_church_bot/internal/domain/language"
	buildText "github.com/nyakovchuk/vsch_church_bot/internal/message/build_text"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

const (
	TOP3 = 3
)

func HandleOnCallback(bm BotManager, cache map[string]interface{}) {
	bm.TBot().Handle(telebot.OnCallback, func(c telebot.Context) error {
		bm.LoggerInfo(c)

		// убираем время задержки у кнопок
		c.Respond()

		externalId := utils.Int64ToString(c.Sender().ID)
		external := external.ToModel(
			externalId,
			bm.SharedData().Platform,
		)

		langCode := c.Get("lang").(string)

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

			findChurches, userCoords, err := SearchChurchesByRadius(bm, external, radius)
			if err != nil {
				bm.LoggerError(c, err)
				return nil
			}

			html := buildText.BuildTextForSearchChurchesByRadius(userCoords, &findChurches, radius, langCode)

			tgHtml = html
		}

		// Обработка кнопки поиска ближайших церквей
		if strings.HasPrefix(data, nearestchurches.PrefixNearestChurches) {

			topN := TOP3

			findChurches, userCoords, err := FindTopNNearestChurches(bm, external, topN)
			if err != nil {
				bm.LoggerError(c, err)
				return nil
			}

			html := buildText.BuildTextForTopNNearestChurches(userCoords, &findChurches, langCode)

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

			printer := i18n.Printer(langCode)

			c.Send(
				printer.Sprintf("event.callback.language_changed_message", langCode),
				&telebot.SendOptions{
					ParseMode: telebot.ModeHTML,
				},
			)

			c.Send("/start")
			handleHelp := handler.HandleStart(bm)
			return handleHelp(c)
		}

		return c.Send(tgHtml, &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	})
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

func SearchChurchesByRadius(bm BotManager, external external.External, radius int) ([]church.DtoResponse, model.Coordinates, error) {
	userCoords, err := bm.Services().Coordinates.GetCoordinates(context.Background(), external)
	if err != nil {
		return []church.DtoResponse{}, model.Coordinates{}, err
	}

	findChurches := bm.Services().Distance.GetChurchesNearby(
		userCoords,
		radius,
		bm.SharedData().Churches,
	)

	return findChurches, userCoords, nil
}

func FindTopNNearestChurches(bm BotManager, external external.External, topN int) ([]church.DtoResponse, model.Coordinates, error) {
	userCoords, err := bm.Services().Coordinates.GetCoordinates(context.Background(), external)
	if err != nil {
		return []church.DtoResponse{}, model.Coordinates{}, err
	}

	findChurches := bm.Services().Distance.FindTopNNearestChurches(
		userCoords,
		topN,
		bm.SharedData().Churches,
	)

	return findChurches, userCoords, nil
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
