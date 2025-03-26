package commandType

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type LocationEvent struct {
	tg telebot.Context
}

func NewLocationEvent(tg telebot.Context) *LocationEvent {
	return &LocationEvent{tg: tg}
}

func (m *LocationEvent) Command() string {
	return "OnLocation"
}

func (m *LocationEvent) Data() string {
	location := m.tg.Message().Location
	return fmt.Sprintf("%.5f,%.5f",
		location.Lat, location.Lng)
}
