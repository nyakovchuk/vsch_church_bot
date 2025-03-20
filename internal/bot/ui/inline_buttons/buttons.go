package inline_buttons

import "gopkg.in/telebot.v4"

type InlineButtons struct {
	rp      *telebot.ReplyMarkup
	buttons map[string]telebot.Btn
}

func NewButtons() *InlineButtons {
	buttons := make(map[string]telebot.Btn)
	return &InlineButtons{
		rp:      &telebot.ReplyMarkup{},
		buttons: buttons,
	}
}

func (b *InlineButtons) Display() *telebot.ReplyMarkup {
	b.createButtons()

	b.rp.Inline(
		b.rp.Row(b.buttons["start"], b.buttons["help"]),
	)

	return b.rp
}

func (b *InlineButtons) createButtons() {
	b.buttons["start"] = b.rp.Data("🏠Главное меню", "start")
	b.buttons["help"] = b.rp.Data("🆘Помощь", "help")
}
