package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

const (
	DescriptionText = "Доступны следующие команды:"
)

func HandleHelp(bm BotManager) func(telebot.Context) error {
	return func(c telebot.Context) error {
		bm.LoggerInfo(c)

		commandText := ""
		for _, cmd := range bm.Commands().Get() {
			commandText += cmd.Route + " - " + cmd.Description + "\n"
		}

		helpText := fmt.Sprintf("%s\n%s", DescriptionText, commandText)

		return c.Send(helpText)
	}
}
