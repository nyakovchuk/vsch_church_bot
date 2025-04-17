package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/command"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/shareddata"
	"github.com/nyakovchuk/vsch_church_bot/pkg/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := app.GetApp()
	fmt.Printf("Logging mode: %s\n\n", app.Config().LogType)
	// defer app.DB().Close()

	defer func() {
		fmt.Print("Stopping the database...")
		if err := app.DB().Close(); err != nil {
			app.Logger().Error("failed to close the database", "err", err)
		}
		fmt.Println("DONE")

		fmt.Print("Closing the log file...")
		if closer, ok := app.Logger().Handler().(interface{ Close() error }); ok {
			_ = closer.Close()
		}
		fmt.Println("DONE")

		fmt.Println("Bot stopped successfully")
	}()

	// проверить наличие таблиц в БД

	cmds := command.GetCommands()

	repo := repository.New(app.DB())
	services := service.New(repo)

	churches, err := services.Church.GetAll(ctx)
	if err != nil {
		fmt.Println("error getting churches", err)
	}

	platform, err := services.Platform.GetByName(ctx, app.Config().Platform)
	if err != nil {
		fmt.Println("error getting platform", err)
	}

	sharedData := shareddata.Data{
		Churches: churches,
		Platform: platform,
	}

	fmt.Print("Starting the bot...")
	bot := telegram.NewBot(app, cmds, services, sharedData)

	if err := bot.Run(ctx); err != nil {
		app.Logger().Error("the bot terminated with an error", "err", err)
	}
}
