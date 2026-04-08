package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

// UpdateHandlerFunc - тип функции, обрабатывающей обновление
type UpdateHandlerFunc func(ctx context.Context, upd interface{}) error

func BotLoggingMiddleware(log *slog.Logger) func(UpdateHandlerFunc) UpdateHandlerFunc {
	return func(next UpdateHandlerFunc) UpdateHandlerFunc {
		return func(ctx context.Context, upd interface{}) error {
			start := time.Now()
			log.Debug("processing update", "type", fmt.Sprintf("%T", upd))
			err := next(ctx, upd)
			duration := time.Since(start)
			if err != nil {
				log.Error("update failed", "error", err, "duration", duration)
			} else {
				log.Info("update succeeded", "duration", duration)
			}
			return err
		}
	}
}
