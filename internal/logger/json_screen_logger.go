package logger

import (
	"log/slog"
	"os"
)

// type JSONScreenLogger struct {
// 	w    io.Writer
// 	opts *slog.HandlerOptions
// }

// NewJSONScreenLogger creates a new JSON logger
// that writes to the standard output
func NewJSONScreenLogger(opts *slog.HandlerOptions) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
