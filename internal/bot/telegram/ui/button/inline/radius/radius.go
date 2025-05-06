package radius

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline"
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
			"five": {Label: "5 km",
				Data: PrefixRadius + "five"},
			"ten": {Label: "10 km",
				Data: PrefixRadius + "ten"},
			"thirty": {Label: "30 km",
				Data: PrefixRadius + "thirty"},
		},
	}
}

type RadiusButtons struct {
	*inline.ManyButtons
}

func New() *RadiusButtons {
	return &RadiusButtons{
		ManyButtons: &inline.ManyButtons{
			IButtons:   button.NewButtons(),
			ButtonsMap: NewButtonsMap(),
		},
	}
}

func (rb *RadiusButtons) Prefix() string {
	return PrefixRadius
}
