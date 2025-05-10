package handler

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/country"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/build_text"
	"gopkg.in/telebot.v4"
)

func HandleChurchesCount(bm BotManager) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		langCode := c.Get("lang").(string)

		countries := bm.SharedData().Countries

		tgResponses := country.ToDtoResponses(&countries)
		tgHtml := build_text.ForCountryWithChurches(&tgResponses, langCode)

		return c.Send(tgHtml, &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}
