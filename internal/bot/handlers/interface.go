package handlers

import (
	"log/slog"

	"github.com/nyakovchuk/vsch_church_bot/config"
	"github.com/nyakovchuk/vsch_church_bot/internal/bot/commands"
	"github.com/tucnak/telebot"
)

type BotManager interface {
	Config() *config.Config
	Logger() *slog.Logger
	TBot() *telebot.Bot
	LoggerInfo(m *telebot.Message)
	Commands() commands.Commands
}
