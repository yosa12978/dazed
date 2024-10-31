package logging

import (
	"io"
	"log/slog"
)

type Logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
}

func getLogLevel(s string) slog.Leveler {
	switch s {
	case "INFO":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	}
	return nil
}

type slogLogger struct {
	logger slog.Logger
}

func NewLogger(w io.Writer, logLevel string) Logger {
	return &slogLogger{
		logger: *slog.New(
			slog.NewJSONHandler(w, &slog.HandlerOptions{
				Level: getLogLevel(logLevel),
			})),
	}
}

func (l *slogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *slogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *slogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
func (l *slogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}
