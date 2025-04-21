package button

import "gopkg.in/telebot.v4"

type TgBtn struct {
	Reply  *telebot.ReplyMarkup
	Button *telebot.Btn
}

func NewButton() *TgBtn {
	return &TgBtn{
		Reply:  &telebot.ReplyMarkup{},
		Button: &telebot.Btn{},
	}
}

type TgBtns struct {
	Reply   *telebot.ReplyMarkup
	Buttons map[string]telebot.Btn
}

func NewButtons() *TgBtns {
	buttons := make(map[string]telebot.Btn)
	return &TgBtns{
		Reply:   &telebot.ReplyMarkup{},
		Buttons: buttons,
	}
}
