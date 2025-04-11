package main

import (
	"context"
	"fmt"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/command"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/shareddata"
	"github.com/nyakovchuk/vsch_church_bot/pkg/app"
)

func main() {

	app := app.GetApp()
	fmt.Printf("Logging mode: %s\n\n", app.Config().LogType)
	defer app.DB().Close()

	// проверить наличие таблиц в БД

	cmds := command.GetCommands()

	repo := repository.New(app.DB())
	services := service.New(repo)

	churches, err := services.Church.GetAll(context.Background())
	if err != nil {
		fmt.Println("error getting churches", err)
	}
	sharedData := shareddata.Data{Churches: churches}

	fmt.Print("Starting the bot...")
	bot.NewBot(app, cmds, services, sharedData).Run()
}
