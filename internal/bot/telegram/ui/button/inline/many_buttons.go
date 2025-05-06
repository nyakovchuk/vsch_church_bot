package inline

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"gopkg.in/telebot.v4"
)

type ManyButtons struct {
	IButtons   *button.TgBtns
	ButtonsMap *button.ButtonsMap
}

func (mb *ManyButtons) Display() *telebot.ReplyMarkup {
	row := mb.CreateAll()

	mb.IButtons.Reply.Inline(
		mb.IButtons.Reply.Row(row...),
	)

	return mb.IButtons.Reply
}

func (mb *ManyButtons) CreateAll() []telebot.Btn {
	var btns []telebot.Btn
	for _, key := range mb.ButtonsMap.Order {
		btn := mb.ButtonsMap.Buttons[key]
		mb.IButtons.Buttons[key] = mb.IButtons.Reply.Data(btn.Label, btn.Data)

		btns = append(btns, mb.IButtons.Buttons[key])
	}

	return btns
}
