package bot

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/tucnak/telebot"
)

type SettingsBot interface {
	Config() *config.Config
	Logger() *slog.Logger
}

type Bot struct {
	bot    *telebot.Bot
	config *config.Config
	logger *slog.Logger
}

func NewBot(s SettingsBot) *Bot {

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
		bot:    bot,
		config: s.Config(),
		logger: s.Logger(),
	}
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
	b.Handlers()

	fmt.Print("DONE\n")
	b.bot.Start()
}
