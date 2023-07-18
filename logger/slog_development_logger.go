package logger

import (
	"os"
	"time"

	"github.com/lmittmann/tint"
	"golang.org/x/exp/slog"
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() Logger {

	logger := slog.New(tint.NewHandler(os.Stderr, nil))

	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	return &slogLogger{
		logger: logger,
	}
}

func (l *slogLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debug(msg, keysAndValues...)
}

func (l *slogLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, keysAndValues...)
}

func (l *slogLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warn(msg, keysAndValues...)
}

func (l *slogLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, keysAndValues...)
}

func (l *slogLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, keysAndValues...)
	panic(msg)
}
