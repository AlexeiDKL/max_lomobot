package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexeiDKL/max_lomobot/iternal/config"
	"github.com/AlexeiDKL/max_lomobot/iternal/delivery/handlers"
	"github.com/AlexeiDKL/max_lomobot/iternal/delivery/middleware"
	"github.com/AlexeiDKL/max_lomobot/iternal/logger"
	maxbot "github.com/max-messenger/max-bot-api-client-go"
)

func main() {
	cfg, _ := config.LoadConfig()

	log := logger.New(logger.Config{
		Path:  cfg.LoggerConfig.Path,
		Level: cfg.LoggerConfig.Level,
	})

	botHandler := handlers.NewBotHandler(cfg.MaxConfig.Token)

	// Оборачиваем хендлер в middleware логирования
	loggingMiddleware := middleware.BotLoggingMiddleware(log)
	wrappedHandler := loggingMiddleware(botHandler.HandleUpdate)

	// Запускаем цикл обновлений
	ctx, cancel := signalContext()
	defer cancel()

	api, err := maxbot.New(cfg.MaxConfig.Token)
	if err != nil {
		log.Error(err.Error())
	}
	updates := api.GetUpdates(ctx)

	for upd := range updates {
		// Вызываем обёрнутый хендлер
		if err := wrappedHandler(ctx, upd); err != nil {
			log.Error("Handler returned error", "error", err)
		}
	}
}

func signalContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()
	return ctx, cancel
}
