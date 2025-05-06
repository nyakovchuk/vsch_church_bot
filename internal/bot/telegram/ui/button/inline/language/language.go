package language

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/button/inline"
)

const (
	PrefixLangCode = "lang_code_"
)

func Order() []string {
	return []string{"uk", "en", "ru"}
}

func NewButtonsMap() *button.ButtonsMap {
	return &button.ButtonsMap{
		Order: Order(),
		Buttons: map[string]button.ButtonConfig{
			"uk": {Label: "🇺🇦 Українська",
				Data: PrefixLangCode + "uk"},
			"en": {Label: "🇬🇧 English",
				Data: PrefixLangCode + "en"},
			"ru": {Label: "🇷🇺 Русский",
				Data: PrefixLangCode + "ru"},
		},
	}
}

type LanguageButtons struct {
	*inline.ManyButtons
}

func NewButtons() *LanguageButtons {
	return &LanguageButtons{
		ManyButtons: &inline.ManyButtons{
			IButtons:   button.NewButtons(),
			ButtonsMap: NewButtonsMap(),
		},
	}
}

func (lb *LanguageButtons) Prefix() string {
	return PrefixLangCode
}
