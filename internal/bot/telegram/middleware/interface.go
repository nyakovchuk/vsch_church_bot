package middleware

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/shareddata"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	LoggerError(telebot.Context, error)
	LoggerMessage(telebot.Context, string)
	Services() *service.Service
	SharedData() shareddata.Data
}
