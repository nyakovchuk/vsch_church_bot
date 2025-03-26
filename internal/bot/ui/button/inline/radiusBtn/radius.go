package radiusBtn

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/button"
	"gopkg.in/telebot.v4"
)

type RadiusButtons struct {
	IButtons *button.Btn
}

func New() *RadiusButtons {
	return &RadiusButtons{
		IButtons: button.NewButtons(),
	}
}

func (rb *RadiusButtons) Display() *telebot.ReplyMarkup {
	rb.create()

	rb.IButtons.Reply.Inline(
		rb.IButtons.Reply.Row(
			rb.IButtons.Buttons["five"],
			rb.IButtons.Buttons["ten"],
			rb.IButtons.Buttons["thirty"],
		),
	)

	return rb.IButtons.Reply
}

func (rb *RadiusButtons) create() {
	rb.IButtons.Buttons["five"] = rb.IButtons.Reply.Data("5 км", "radius_five")
	rb.IButtons.Buttons["ten"] = rb.IButtons.Reply.Data("10 км", "radius_ten")
	rb.IButtons.Buttons["thirty"] = rb.IButtons.Reply.Data("30 км", "radius_thirty")

}
