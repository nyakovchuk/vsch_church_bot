package handler

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/internal/shareddata"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	SharedData() shareddata.Data
}

type ButtonRenderer interface {
	Display() *telebot.ReplyMarkup
}
