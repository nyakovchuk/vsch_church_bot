package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	FLAGS       = os.O_CREATE | os.O_APPEND | os.O_WRONLY
	PERM        = 0666
	ErrFileOpen = "Error opening the log file"
)

// type JSONFileLogger struct {
// 	logFilePath string
// 	opts        *slog.HandlerOptions
// }

// NewJSONFileLogger creates a new JSON logger
// that writes to a file
func NewJSONFileLogger(logFilePath string, opts *slog.HandlerOptions) *slog.Logger {

	logFile, err := os.OpenFile(logFilePath, FLAGS, PERM)
	if err != nil {
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).
			Log(context.Background(), slog.LevelError, ErrFileOpen, slog.Any("error", err))

		os.Exit(1)
	}

	handler := slog.NewJSONHandler(logFile, opts)
	logger := slog.New(handler)

	return logger
}
