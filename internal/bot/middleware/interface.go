package middleware

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/internal/service"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	Services() *service.Service
}
