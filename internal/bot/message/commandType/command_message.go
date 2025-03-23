package commandType

import "gopkg.in/telebot.v4"

type CommandMessage struct {
	tg telebot.Context
}

func NewCommandMessage(tg telebot.Context) *CommandMessage {
	return &CommandMessage{tg: tg}
}

func (m *CommandMessage) Command() string {
	return "Command"
}

func (m *CommandMessage) Data() string {
	return m.tg.Message().Text
}
