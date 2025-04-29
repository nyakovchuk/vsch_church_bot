package middleware

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/language"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
)

func GetLanguage(bm BotManager) {
	bm.TBot().Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {

			tgUser := c.Sender()

			langCode := c.Sender().LanguageCode
			if IsLanguageSelected(bm, tgUser.ID) {
				langId, _ := bm.Services().User.LanguageId(
					bm.SharedData().Platform.ID,
					utils.Int64ToString(tgUser.ID),
				)

				langCode = language.LanguageList(bm.SharedData().Languages).ByID(langId).Code
			}

			c.Set("lang", langCode)

			return next(c)
		}
	})
}

func IsLanguageSelected(bm BotManager, userid int64) bool {
	platformId := bm.SharedData().Platform.ID
	externalId := utils.Int64ToString(userid)
	isSelected, err := bm.Services().User.IsLanguageSelected(platformId, externalId)

	if err != nil {
		return false
	}

	return isSelected
}
