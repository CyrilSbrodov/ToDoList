package loggers

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

// SetupLogger создание нового логгера.
func SetupLogger(env string) *Logger {
	var logger *slog.Logger
	switch env {
	case "local":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return &Logger{logger: logger}
}

// LogErr обработка уровень ошибок.
func (l *Logger) LogErr(err error, msg string) {
	l.logger.Error(msg, err)
}

// LogInfo обработка уровень инфо.
func (l *Logger) LogInfo(key, msg string) {
	l.logger.Info(key, msg)
}

// LogDebug обработка уровень дебаг.
func (l *Logger) LogDebug(key, msg string) {
	l.logger.Debug(key, msg)
}
