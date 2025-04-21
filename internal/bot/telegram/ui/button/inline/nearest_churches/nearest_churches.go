package nearestchurches

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"gopkg.in/telebot.v4"
)

const (
	PrefixNearestChurches = "nearest_churches_"
)

func NewButtonData() *button.ButtonConfig {
	return &button.ButtonConfig{
		Label: "🔎 Показать ближайшие 3 церкви",
		Data:  PrefixNearestChurches,
	}
}

type NearestChurchesButton struct {
	IButtons   *button.TgBtn
	ButtonData *button.ButtonConfig
}

func New() *NearestChurchesButton {
	return &NearestChurchesButton{
		IButtons:   button.NewButton(),
		ButtonData: NewButtonData(),
	}
}

func (ncb *NearestChurchesButton) Prefix() string {
	return PrefixNearestChurches
}

func (ncb *NearestChurchesButton) Display() *telebot.ReplyMarkup {

	btnNearestChurches := ncb.Сreate()

	ncb.IButtons.Reply.Inline(
		ncb.IButtons.Reply.Row(btnNearestChurches),
	)

	return ncb.IButtons.Reply
}

func (ncb *NearestChurchesButton) Сreate() telebot.Btn {
	return ncb.IButtons.Reply.Data(
		ncb.ButtonData.Label,
		ncb.ButtonData.Data,
	)
}
