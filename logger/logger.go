package logger

import (
	"log/slog"
	"os"
)

type LogLevel string

const (
	Debug LogLevel = "DEBUG"
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
)

var logLevelMap map[LogLevel]slog.Level = map[LogLevel]slog.Level{
	Debug: slog.LevelDebug,
	Info:  slog.LevelInfo,
	Warn:  slog.LevelWarn,
	Error: slog.LevelError,
}

func NewLogger(logLevel LogLevel) *slog.Logger {
	logLevelVar := new(slog.LevelVar)
	level, ok := logLevelMap[logLevel]
	if !ok {
		level = logLevelMap["Debug"]
	}
	logLevelVar.Set(level)
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevelVar}))
}
