package type_commands

import "gopkg.in/telebot.v4"

type CommandMessage struct {
	tg telebot.Context
}

func NewCommandMessage(tg telebot.Context) *CommandMessage {
	return &CommandMessage{tg: tg}
}

func (m *CommandMessage) Command() string {
	return "Received " + m.tg.Message().Text + " command"
}

func (m *CommandMessage) Data() string {
	return ""
}
