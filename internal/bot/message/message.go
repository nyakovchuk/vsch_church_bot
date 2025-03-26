package message

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type Message struct {
	Tgmessage telebot.Context
	Service   CommandInfo
}

// New Message
func NewMessage(tgmessage telebot.Context) Message {
	service := GetTypeCommand(tgmessage)
	return Message{
		Tgmessage: tgmessage,
		Service:   service,
	}
}

func (m *Message) Command() string {
	return m.Service.Command()
}

func (m *Message) Data() string {
	return m.Service.Data()
}

func (m *Message) UserInfo() string {
	return fmt.Sprintf("user: %s, chat: %d", m.Tgmessage.Sender().Username, m.Tgmessage.Chat().ID)
}

func (m *Message) FullInfo() string {
	return m.Command() + " " + m.UserInfo()
}
