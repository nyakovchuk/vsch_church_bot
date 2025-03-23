package handler

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/command"
	"gopkg.in/telebot.v4"
)

type BotManager interface {
	Config() *config.Config
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(telebot.Context)
	Commands() command.Commands
}
