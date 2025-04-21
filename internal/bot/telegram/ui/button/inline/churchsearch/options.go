package options

import (
	nearestchurches "github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/nearest_churches"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline/radius"
	"gopkg.in/telebot.v4"
)

type ChurchSearchOptions struct {
	radiusBtns         *radius.RadiusButtons
	nearestChurchesBtn *nearestchurches.NearestChurchesButton
}

func NewChurchSearchOptions() *ChurchSearchOptions {
	nearestChurchesButtons := nearestchurches.New()
	radiusButtons := radius.New()
	return &ChurchSearchOptions{
		radiusBtns:         radiusButtons,
		nearestChurchesBtn: nearestChurchesButtons,
	}
}

func (cso *ChurchSearchOptions) Display() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{}

	radiusBtns := cso.radiusBtns.CreateAll()
	nearestchurchesBtn := cso.nearestChurchesBtn.Сreate()

	menu.Inline(
		menu.Row(radiusBtns...),
		menu.Row(nearestchurchesBtn),
	)

	return menu
}

func (cso *ChurchSearchOptions) PrefixRadius() string {
	return cso.radiusBtns.Prefix()
}

func (cso *ChurchSearchOptions) PrefixNearestChurches() string {
	return cso.nearestChurchesBtn.Prefix()
}
