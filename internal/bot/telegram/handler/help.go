package handler

import (
	"gopkg.in/telebot.v4"
)

func HandleHelp(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		return c.Send(descriptionHelp(), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		})
	}
}

func descriptionHelp() string {
	return `
⛪ Бот поддерживает <b>три способа</b> передачи геолокации для поиска церквей:

<b>1.</b> <u>Кнопка "Отправить местоположение"</u> <i>(только для мобильных устройств)</i>

<b>2.</b> <u>Через меню Telegram</u> <i>(работает на всех устройствах)</i>
Нажмите на значок 📎 <b><i>"Вложения"</i></b> в поле ввода → <b><i>"Местоположение"</i></b> → Выберите точку на карте.

<b>3.</b> <u>Ручной ввод координат</u>
<b>Отправьте сообщение в формате:</b>
<code>50.4228 30.3145</code> <i>(широта, долгота через пробел)</i>
<code>50.4228, 30.3145</code> <i>(широта, долгота через запятую)</i>.
	`
}
