package reply

import "gopkg.in/telebot.v4"

func CreateMenuKeyboard() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{}

	// Кнопки с командами
	btnStart := menu.Text("🚀 Старт")
	btnHelp := menu.Text("ℹ️ Помощь")
	btnSettings := menu.Text("⚙️ Настройки")

	// Кнопка "Меню"
	btnMenu := menu.Text("📜 Меню")

	// Организация кнопок
	menu.Reply(
		menu.Row(btnMenu),
		menu.Row(btnStart, btnHelp),
		menu.Row(btnSettings),
	)

	return menu
}
