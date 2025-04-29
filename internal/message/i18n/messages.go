package i18n

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n/button"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n/command"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n/event"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Init регистрирует все переводы для поддерживаемых языков
func Init() {
	// Английские переводы
	en := language.English
	message.SetString(en, "command.help", command.EnCommandHelp())
	message.SetString(en, "command.start", command.EnCommandStart())
	message.SetString(en, "command.language", command.EnCommandLanguage())
	message.SetString(en, "button.send_location", button.EnSendLocation())
	message.SetString(en, "button.show_nearest_churches", button.EnShowNearestChurches())
	message.SetString(en, "event.callback.language_changed_message", event.EnLanguageChangedMessage())
	message.SetString(en, "event.callback.churches_found_in_radius", event.EnChurchesFoundInRadius())
	message.SetString(en, "event.callback.nearby_churches_title", event.EnNearbyChurchesTitle())
	message.SetString(en, "event.callback.route", event.EnRoute())
	message.SetString(en, "event.location.geolocation_error", event.EnGeolocationError())
	message.SetString(en, "event.location.your_coordinates", event.EnYourCoordinates())
	message.SetString(en, "event.location.find_churches_in_radius", event.EnFindChurchesInRadius())
	message.SetString(en, "event.text_location.invalid_coordinates_format", event.EnInvalidCoordinatesFormat())
	message.SetString(en, "event.text_location.invalid_latitude", event.EnInvalidLatitude())
	message.SetString(en, "event.text_location.invalid_longitude", event.EnInvalidLongitude())

	// Украинские переводы
	uk := language.Ukrainian
	message.SetString(uk, "command.help", command.UkCommandHelp())
	message.SetString(uk, "command.start", command.UkCommandStart())
	message.SetString(uk, "command.language", command.UkCommandLanguage())
	message.SetString(uk, "button.send_location", button.UkSendLocation())
	message.SetString(uk, "button.show_nearest_churches", button.UkShowNearestChurches())
	message.SetString(uk, "event.callback.language_changed_message", event.UkLanguageChangedMessage())
	message.SetString(uk, "event.callback.churches_found_in_radius", event.UkChurchesFoundInRadius())
	message.SetString(uk, "event.callback.nearby_churches_title", event.UkNearbyChurchesTitle())
	message.SetString(uk, "event.callback.route", event.UkRoute())
	message.SetString(uk, "event.location.geolocation_error", event.UkGeolocationError())
	message.SetString(uk, "event.location.your_coordinates", event.UkYourCoordinates())
	message.SetString(uk, "event.location.find_churches_in_radius", event.UkFindChurchesInRadius())
	message.SetString(uk, "event.text_location.invalid_coordinates_format", event.UkInvalidCoordinatesFormat())
	message.SetString(uk, "event.text_location.invalid_latitude", event.UkInvalidLatitude())
	message.SetString(uk, "event.text_location.invalid_longitude", event.UkInvalidLongitude())

	// Русские переводы
	ru := language.Russian
	message.SetString(ru, "command.help", command.RuCommandHelp())
	message.SetString(ru, "command.start", command.RuCommandStart())
	message.SetString(ru, "command.language", command.RuCommandLanguage())
	message.SetString(ru, "button.send_location", button.RuSendLocation())
	message.SetString(ru, "button.show_nearest_churches", button.RuShowNearestChurches())
	message.SetString(ru, "event.callback.language_changed_message", event.RuLanguageChangedMessage())
	message.SetString(ru, "event.callback.churches_found_in_radius", event.RuChurchesFoundInRadius())
	message.SetString(ru, "event.callback.nearby_churches_title", event.RuNearbyChurchesTitle())
	message.SetString(ru, "event.callback.route", event.RuRoute())
	message.SetString(ru, "event.location.geolocation_error", event.RuGeolocationError())
	message.SetString(ru, "event.location.your_coordinates", event.RuYourCoordinates())
	message.SetString(ru, "event.location.find_churches_in_radius", event.RuFindChurchesInRadius())
	message.SetString(ru, "event.text_location.invalid_coordinates_format", event.RuInvalidCoordinatesFormat())
	message.SetString(ru, "event.text_location.invalid_latitude", event.RuInvalidLatitude())
	message.SetString(ru, "event.text_location.invalid_longitude", event.RuInvalidLongitude())
}

// Printer создает message.Printer для указанного языка
func Printer(langCode string) *message.Printer {
	return message.NewPrinter(language.Make(langCode))
}
