package reply

import "gopkg.in/telebot.v4"

const (
	btnText = "📍Отправить местоположение"
)

func BtnLocation() *telebot.ReplyMarkup {

	btn := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnLocation := btn.Location(btnText)

	btn.Reply(btn.Row(btnLocation))

	return btn
}
