package main

import (
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/command"
	"github.com/nyakovchuk/vsch_church_bot/pkg/app"
)

func main() {

	app := app.GetApp()
	fmt.Printf("Logging mode: %s\n\n", app.Config().LogType)

	cmds := command.GetCommands()

	fmt.Print("Starting the bot...")
	bot.NewBot(app, cmds).Run()
}
