package middleware

import (
	"gopkg.in/telebot.v4"
)

func GetLanguage(bm BotManager) {
	bm.TBot().Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {

			// получить lang_id
			// если nil сохранить

			// получить lang_id
			// найти язык (code) из shareddata
			// и установить lang(context)
			c.Set("lang", c.Sender().LanguageCode)

			return next(c)
		}
	})
}
