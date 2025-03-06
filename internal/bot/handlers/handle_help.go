package handlers

import (
	"fmt"

	"github.com/tucnak/telebot"
)

const (
	DescriptionText = "Доступны следующие команды:"
)

func HandleHelp(bm BotManager) func(m *telebot.Message) {
	return func(m *telebot.Message) {
		bm.LoggerInfo(m)

		commandText := ""
		for _, cmd := range bm.Commands().Get() {
			commandText += cmd.Route + " - " + cmd.Description + "\n"
		}

		helpText := fmt.Sprintf("%s\n%s", DescriptionText, commandText)
		bm.TBot().Send(m.Chat, helpText)
	}
}
