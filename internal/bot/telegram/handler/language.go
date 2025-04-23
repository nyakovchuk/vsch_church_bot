package handler

import "gopkg.in/telebot.v4"

func HandleLanguage(bm BotManager, buttons ButtonRenderer) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		return c.Send(descriptionLanguage(), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		}, buttons.Display())
	}
}

func descriptionLanguage() string {
	return `
Вы можете изменить язык интерфейса бота.
Выберите язык, которым будет пользоваться бот.
	`
}
