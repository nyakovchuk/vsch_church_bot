package radius

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"gopkg.in/telebot.v4"
)

const (
	PrefixRadius = "radius_"
)

func Order() []string {
	return []string{"five", "ten", "thirty"}
}

func NewButtonsMap() *button.ButtonsMap {
	return &button.ButtonsMap{
		Order: Order(),
		Buttons: map[string]button.ButtonConfig{
			"five": {Label: "5 км",
				Data: PrefixRadius + "five"},
			"ten": {Label: "10 км",
				Data: PrefixRadius + "ten"},
			"thirty": {Label: "30 км",
				Data: PrefixRadius + "thirty"},
		},
	}
}

type RadiusButtons struct {
	IButtons   *button.TgBtns
	ButtonsMap *button.ButtonsMap
}

func New() *RadiusButtons {
	return &RadiusButtons{
		IButtons:   button.NewButtons(),
		ButtonsMap: NewButtonsMap(),
	}
}

func (rb *RadiusButtons) Prefix() string {
	return PrefixRadius
}

func (rb *RadiusButtons) Display() *telebot.ReplyMarkup {
	row := rb.CreateAll()

	rb.IButtons.Reply.Inline(
		rb.IButtons.Reply.Row(row...),
	)

	return rb.IButtons.Reply
}

func (rb *RadiusButtons) CreateAll() []telebot.Btn {
	var btns []telebot.Btn
	for _, key := range rb.ButtonsMap.Order {
		btn := rb.ButtonsMap.Buttons[key]
		rb.IButtons.Buttons[key] = rb.IButtons.Reply.Data(btn.Label, btn.Data)

		btns = append(btns, rb.IButtons.Buttons[key])
	}

	return btns
}
