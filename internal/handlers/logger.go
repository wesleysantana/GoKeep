package handlers

import (
	"io"
	"log/slog"
	"strings"
	"time"
)

type Logger struct{}

func (l *Logger) GetLevelLog(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func replaceTimeFormat(group []string, a slog.Attr) slog.Attr {
	if a.Key == "time" {
		value := time.Now().Format("2006-01-02T15:04:05")
		return slog.Attr{Key: "time", Value: slog.StringValue(value)}
	}
	return a
}

func (l *Logger) NewLogger(out io.Writer, minLevel slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(out,
		&slog.HandlerOptions{
			AddSource:   true,
			Level:       minLevel,
			ReplaceAttr: replaceTimeFormat,
		}))
}
