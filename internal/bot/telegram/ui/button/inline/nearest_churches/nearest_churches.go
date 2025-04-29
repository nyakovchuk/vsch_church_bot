package nearestchurches

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
	"gopkg.in/telebot.v4"
)

const (
	PrefixNearestChurches = "nearest_churches_"
)

func NewButtonData(langCode string) *button.ButtonConfig {

	printer := i18n.Printer(langCode)

	return &button.ButtonConfig{
		Label: printer.Sprintf("button.show_nearest_churches"),
		Data:  PrefixNearestChurches,
	}
}

type NearestChurchesButton struct {
	IButtons   *button.TgBtn
	ButtonData *button.ButtonConfig
}

func New(langCode string) *NearestChurchesButton {
	return &NearestChurchesButton{
		IButtons:   button.NewButton(),
		ButtonData: NewButtonData(langCode),
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
