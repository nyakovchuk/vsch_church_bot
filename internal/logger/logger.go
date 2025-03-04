package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nyakovchuk/vsch_church_bot/config"
)

const (
	FileLogger   = "File"
	ScreenLogger = "Screen"
)

// SetupLogger returns logger
func SetupLogger(config *config.Config, opts *slog.HandlerOptions) (*slog.Logger, error) {
	return GetLogger(config.LogType, opts)
}

// GetLogger returns logger by name "File" or "Screen"
func GetLogger(nameLogger string, opts *slog.HandlerOptions) (*slog.Logger, error) {
	switch nameLogger {
	case FileLogger:
		return NewJSONFileLogger(os.Getenv("LOG_FILE_PATH"), opts), nil
	case ScreenLogger:
		return NewJSONScreenLogger(opts), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownLogger, nameLogger)
	}
}
