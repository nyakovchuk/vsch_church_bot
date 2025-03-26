package button

import "gopkg.in/telebot.v4"

type Btn struct {
	Reply   *telebot.ReplyMarkup
	Buttons map[string]telebot.Btn
}

func NewButtons() *Btn {
	buttons := make(map[string]telebot.Btn)
	return &Btn{
		Reply:   &telebot.ReplyMarkup{},
		Buttons: buttons,
	}
}
