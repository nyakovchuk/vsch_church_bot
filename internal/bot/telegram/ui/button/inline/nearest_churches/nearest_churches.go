package nearestchurches

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
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
	*inline.SingleButton
}

func New(langCode string) *NearestChurchesButton {
	return &NearestChurchesButton{
		SingleButton: &inline.SingleButton{
			IButtons:   button.NewButton(),
			ButtonData: NewButtonData(langCode),
		},
	}
}

func (ncb *NearestChurchesButton) Prefix() string {
	return PrefixNearestChurches
}
