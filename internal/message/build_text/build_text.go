package build_text

import (
	"fmt"
	"strings"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/country"
	"github.com/nyakovchuk/vsch_church_bot/internal/message/i18n"
)

func BuildTextForSearchChurchesByRadius(userCoords model.Coordinates, churches *[]church.DtoResponse, radius int, langCode string) string {

	printer := i18n.Printer(langCode)

	var builder strings.Builder
	builder.WriteString(printer.Sprintf("event.callback.churches_found_in_radius", radius/1000, len(*churches)))

	for i, church := range *churches {
		builder.WriteString(fmt.Sprintf("<b>%d.</b> ", i+1))
		builder.WriteString(buildChurchesText(userCoords, church, langCode))
	}

	return builder.String()
}

func BuildTextForTopNNearestChurches(userCoords model.Coordinates, churches *[]church.DtoResponse, langCode string) string {

	printer := i18n.Printer(langCode)

	var builder strings.Builder
	builder.WriteString(printer.Sprintf("event.callback.nearby_churches_title"))

	for i, church := range *churches {
		builder.WriteString(fmt.Sprintf("<b>%d.</b> ", i+1))
		builder.WriteString(buildChurchesText(userCoords, church, langCode))
	}

	return builder.String()
}

func buildChurchesText(userCoords model.Coordinates, church church.DtoResponse, langCode string) string {
	var churchName, churchConfession string
	if langCode == i18n.LangCodeRu || langCode == i18n.LangCodeUk {
		churchName = church.NameRU
		churchConfession = church.ConfessionRu
	} else {
		churchName = church.NameEN
		churchConfession = church.ConfessionEn
	}

	var builder strings.Builder

	vschUrl := fmt.Sprintf("https://www.vsch.org/church/%s", church.Alias)
	builder.WriteString(fmt.Sprintf("<a href=\"%s\"><b>%s</b></a> (%s) – <b>[%.2f км]</b> <a href=\"https://www.google.com/maps/dir/%v,%v/%v,%v\">", vschUrl, churchName, churchConfession, church.Distance/1000, userCoords.Latitude, userCoords.Longitude, church.Latitude, church.Longitude))

	printer := i18n.Printer(langCode)
	builder.WriteString(printer.Sprintf("event.callback.route"))
	builder.WriteString("</a>\n")

	return builder.String()
}

func ForCountryWithChurches(countries *[]country.DtoResponse, langCode string) string {
	printer := i18n.Printer(langCode)

	html := printer.Sprintf("command.country_churches_count_title")
	html += "\n\n"
	html += formatCountryStats(countries, langCode)
	html += "\n"
	html += printer.Sprintf("command.more_info_link")

	return html
}

// Формуємо вирівняний HTML-текст
func formatCountryStats(countries *[]country.DtoResponse, langCode string) string {

	var builder strings.Builder

	maxWidth := maxWidthChurchesCount(countries)
	for _, country := range *countries {
		widthCountry := len(fmt.Sprintf("%d", country.ChurchesCount))

		builder.WriteString(fmt.Sprintf("<b>%d</b>", country.ChurchesCount))
		for i := widthCountry; i <= maxWidth; i++ {
			builder.WriteString("  ")
		}
		builder.WriteString(country.FlagWithName(langCode))
		builder.WriteString("\n")
	}

	return builder.String()
}

// Знаходимо максимальну довжину числа
func maxWidthChurchesCount(countries *[]country.DtoResponse) int {

	maxWidth := 0
	for _, country := range *countries {
		width := len(fmt.Sprintf("%d", country.ChurchesCount))
		if width > maxWidth {
			maxWidth = width
		}
	}

	return maxWidth
}
