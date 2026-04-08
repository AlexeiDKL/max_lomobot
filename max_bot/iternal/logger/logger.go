package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

type Config struct {
	Path  string
	Level string // debug, info, warn, error
}

func New(cfg Config) *slog.Logger {
	// Определяем уровень логирования
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Базовый writer: консоль (stderr)
	var writers []io.Writer

	if cfg.Path != "" {
		file, err := os.OpenFile(cfg.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			// Не удалось открыть файл — пишем в stderr и логируем ошибку
			consoleLogger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
			consoleLogger.Error("Failed to open log file, using stderr only",
				slog.String("path", cfg.Path),
				slog.String("error", err.Error()),
			)
			writers = append(writers, os.Stderr)
		} else {
			// Файл успешно открыт — пишем только в него
			writers = append(writers, file)
		}
	} else {
		// Файл не указан — пишем в stderr
		writers = append(writers, os.Stderr)
	}

	multiWriter := io.MultiWriter(writers...)

	handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})

	log.SetOutput(multiWriter)
	return slog.New(handler)
}
