package button

import "gopkg.in/telebot.v4"

type ButtonDisplayer interface {
	Display() *telebot.ReplyMarkup
}

type ButtonCreator interface {
	Сreate() []telebot.Btn
}

type BittonPrefixer interface {
	Prefix() string
}

type UIButton interface {
	ButtonDisplayer
	BittonPrefixer
}
