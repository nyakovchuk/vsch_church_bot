package telegram

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/command"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/ui/menu"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"gopkg.in/telebot.v4"

	"github.com/nyakovchuk/vsch_church_bot/internal/shareddata"
)

type SettingsBot interface {
	Config() *config.Config
	Logger() *slog.Logger
}

type Bot struct {
	bot        *telebot.Bot
	commands   command.CommandManager
	config     *config.Config
	logger     *slog.Logger
	services   *service.Service
	shareddata shareddata.Data
}

func NewBot(s SettingsBot, commands command.CommandManager, services *service.Service, sharedData shareddata.Data) *Bot {

	if s.Config().TelegramBotToken == "" {
		s.Logger().Error("Retrieving the token", "err", "Token not found in environment variables")
		return nil
	}

	// Creating a new bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  s.Config().TelegramBotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		s.Logger().Error("Creating the bot", "err", "Error creating the bot")
		return nil
	}

	return &Bot{
		bot:        bot,
		commands:   commands.Get(),
		config:     s.Config(),
		logger:     s.Logger(),
		services:   services,
		shareddata: sharedData,
	}
}

func (b *Bot) Commands() command.Commands {
	return b.commands.Get()
}

func (b *Bot) Config() *config.Config {
	return b.config
}

func (b *Bot) Logger() *slog.Logger {
	return b.logger
}

func (b *Bot) TBot() *telebot.Bot {
	return b.bot
}

func (b *Bot) Services() *service.Service {
	return b.services
}

func (b *Bot) SharedData() shareddata.Data {
	return b.shareddata
}

func (b *Bot) Run(ctx context.Context) error {

	menu.Create(b.bot)
	b.Middleware()
	b.Handlers()
	b.Events()
	fmt.Print("DONE\n\n")

	go b.bot.Start()

	<-ctx.Done()

	b.logger.Info("Stopping the bot...")
	b.bot.Stop()

	return nil
}
