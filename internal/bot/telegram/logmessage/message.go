package logmessage

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type LogMessage struct {
	Tgmessage telebot.Context
	Service   CommandInfo
}

// New Message
func New(tgmessage telebot.Context) LogMessage {
	service := GetTypeCommand(tgmessage)
	return LogMessage{
		Tgmessage: tgmessage,
		Service:   service,
	}
}

func (m *LogMessage) Command() string {
	return m.Service.Command()
}

func (m *LogMessage) Data() string {
	return m.Service.Data()
}

func (m *LogMessage) UserInfo() string {
	return fmt.Sprintf("user: %s, chat: %d", m.Tgmessage.Sender().Username, m.Tgmessage.Chat().ID)
}

func (m *LogMessage) FullInfo() string {
	return m.Command() + " " + m.UserInfo()
}
