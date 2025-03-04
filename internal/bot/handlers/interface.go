package handlers

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/tucnak/telebot"
)

type Bot interface {
	Config() *config.Config
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(m *telebot.Message)
}
