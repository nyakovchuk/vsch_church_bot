package radiusBtn

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/button"
	"gopkg.in/telebot.v4"
)

func Order() []string {
	return []string{"five", "ten", "thirty"}
}

func NewButtonSet() *button.ButtonSet {
	return &button.ButtonSet{
		Order: Order(),
		Buttons: map[string]button.ButtonConfig{
			"five":   {Label: "5 км", Data: "radius_five"},
			"ten":    {Label: "10 км", Data: "radius_ten"},
			"thirty": {Label: "30 км", Data: "radius_thirty"},
		},
	}
}

type RadiusButtons struct {
	IButtons  *button.Btn
	ButtonSet *button.ButtonSet
}

func New() *RadiusButtons {
	return &RadiusButtons{
		IButtons:  button.NewButtons(),
		ButtonSet: NewButtonSet(),
	}
}

func (rb *RadiusButtons) Display() *telebot.ReplyMarkup {
	row := rb.create()

	rb.IButtons.Reply.Inline(
		rb.IButtons.Reply.Row(row...),
	)

	return rb.IButtons.Reply
}

func (rb *RadiusButtons) create() []telebot.Btn {
	var row []telebot.Btn
	for _, key := range rb.ButtonSet.Order {
		btn := rb.ButtonSet.Buttons[key]
		rb.IButtons.Buttons[key] = rb.IButtons.Reply.Data(btn.Label, btn.Data)

		row = append(row, rb.IButtons.Buttons[key])
	}

	return row
}
