package handler

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/button/reply"
	"gopkg.in/telebot.v4"
)

func HandleStart(bm BotManager) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		btnLocation := reply.BtnLocation()

		return c.Send(descriptionStart(), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
		}, btnLocation)
	}
}

func descriptionStart() string {
	return `
🔍 <b>Найди ближайшие церкви!</b>
(работает по базе церквей сайта <b><i><a href="https://www.vsch.org">vsch.org</a></i></b>)

Просто отправь свои текущие координаты или выбери местонахождение в геолокации, и <b>бот покажет все церкви в выбранном радиусе</b>.

🌟 <b>Возможности:</b>
→ <b>Удобный поиск церквей</b> рядом с тобой
→ <b>Выбор радиуса поиска</b> <i>(5, 10, 30 км)</i>
→ <b>Информация о церквях:</b> название, конфессия, расстояние
→ Возможность <b>открыть местоположение церкви</b> на карте

🚀 <b>Как использовать?</b>
<b>1.</b> Отправь данные геолокации или введи координаты
<i>(примеры отправки геолокации <b>/help</b>)</i>
<b>2.</b> Выбери радиус поиска
<b>3.</b> Получи список ближайших церквей!

🔗 Добавь бота и всегда находи церкви рядом: <b>@vsch_church_bot</b>`
}
