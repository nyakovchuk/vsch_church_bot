package command

import "gopkg.in/telebot.v4"

func ToTelebotCommands() []telebot.Command {
	var cmds []telebot.Command
	for _, cmd := range commandList {
		cmds = append(cmds, telebot.Command{
			Text:        cmd.Name,
			Description: cmd.Description,
		})
	}
	return cmds
}
