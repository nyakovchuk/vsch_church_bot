package bot

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/commands"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/ui/menu"
	"gopkg.in/telebot.v4"
)

type SettingsBot interface {
	Config() *config.Config
	Logger() *slog.Logger
}

type Bot struct {
	bot      *telebot.Bot
	commands commands.CommandManager
	config   *config.Config
	logger   *slog.Logger
}

func NewBot(s SettingsBot, commands commands.CommandManager) *Bot {

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
		bot:      bot,
		commands: commands.Get(),
		config:   s.Config(),
		logger:   s.Logger(),
	}
}

func (b *Bot) Commands() commands.Commands {
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

func (b *Bot) Run() {

	menu.Create(b.bot)

	b.Handlers()

	b.Events()

	fmt.Print("DONE\n")
	b.bot.Start()
}
