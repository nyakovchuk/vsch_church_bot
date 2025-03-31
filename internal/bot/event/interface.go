package event

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/command"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Config() *config.Config
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	Commands() command.Commands
	Services() *service.Service
}

type ButtonRenderer interface {
	Display() *telebot.ReplyMarkup
}
