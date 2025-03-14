package messages

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type Message struct {
	Tgmessage telebot.Context
}

func (m *Message) UserInfo() string {
	return fmt.Sprintf("from user: %s, Chat: %d", m.Tgmessage.Sender().Username, m.Tgmessage.Chat().ID)
}

func (m *Message) Command() string {
	return fmt.Sprintf("Received %s command", m.Tgmessage.Text())
}

func (m *Message) FullInfo() string {
	return m.Command() + " " + m.UserInfo()
}
