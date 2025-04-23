package language

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"gopkg.in/telebot.v4"
)

const (
	PrefixLangCode = "lang_code_"
)

func Order() []string {
	return []string{"uk", "en", "ru"}
}

func NewButtonsMap() *button.ButtonsMap {
	return &button.ButtonsMap{
		Order: Order(),
		Buttons: map[string]button.ButtonConfig{
			"uk": {Label: "Українська",
				Data: PrefixLangCode + "uk"},
			"en": {Label: "English",
				Data: PrefixLangCode + "en"},
			"ru": {Label: "Русский",
				Data: PrefixLangCode + "ru"},
		},
	}
}

type LanguageButtons struct {
	IButtons   *button.TgBtns
	ButtonsMap *button.ButtonsMap
}

func NewButtons() *LanguageButtons {
	return &LanguageButtons{
		IButtons:   button.NewButtons(),
		ButtonsMap: NewButtonsMap(),
	}
}

func (lb *LanguageButtons) Prefix() string {
	return PrefixLangCode
}

func (lb *LanguageButtons) Display() *telebot.ReplyMarkup {
	row := lb.CreateAll()

	lb.IButtons.Reply.Inline(
		lb.IButtons.Reply.Row(row...),
	)

	return lb.IButtons.Reply
}

func (lb *LanguageButtons) CreateAll() []telebot.Btn {
	var btns []telebot.Btn
	for _, key := range lb.ButtonsMap.Order {
		btn := lb.ButtonsMap.Buttons[key]
		lb.IButtons.Buttons[key] = lb.IButtons.Reply.Data(btn.Label, btn.Data)

		btns = append(btns, lb.IButtons.Buttons[key])
	}

	return btns
}
