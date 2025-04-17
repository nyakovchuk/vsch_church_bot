package event

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/telegram/command"
	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/shareddata"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Config() *config.Config
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	LoggerError(telebot.Context, error)
	Commands() command.Commands
	Services() *service.Service
	SharedData() shareddata.Data
}

type ButtonRenderer interface {
	Display() *telebot.ReplyMarkup
}
