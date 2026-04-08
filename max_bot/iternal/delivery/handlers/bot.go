package handlers

import (
	"context"
	"fmt"

	file "github.com/AlexeiDKL/max_lomobot/iternal/file"
	maxbot "github.com/max-messenger/max-bot-api-client-go"
	"github.com/max-messenger/max-bot-api-client-go/schemes"
)

type BotHandler struct {
	api *maxbot.Api
	err error
}

func NewBotHandler(token string) *BotHandler {
	api, err := maxbot.New(token)
	return &BotHandler{api: api, err: err}
}

// HandleUpdate обрабатывает одно обновление (без логирования)
func (h *BotHandler) HandleUpdate(ctx context.Context, upd interface{}) error {
	switch upd := upd.(type) {
	case *schemes.MessageCreatedUpdate:
		switch upd.GetCommand() {
		case "/cats":
			{
				return h.handlerCats(ctx, upd)
			}
		default:
			{
				return h.handleUncnow(ctx, upd)
			}
		}
	default:
		return nil
	}
}

func (h *BotHandler) handlerCats(ctx context.Context, upd *schemes.MessageCreatedUpdate) error {
	msg := maxbot.NewMessage()
	chatID := upd.Message.Recipient.ChatId
	photoPath, _ := file.GetRandomCatImage()
	photo, err := h.api.Uploads.UploadPhotoFromFile(ctx, photoPath)
	if err != nil {
		return err
	}
	msg.SetChat(chatID)
	msg.AddPhoto(photo)
	err = h.api.Messages.Send(ctx, msg)
	return err

}

func (h *BotHandler) handleUncnow(ctx context.Context, upd *schemes.MessageCreatedUpdate) error {
	chatID := upd.Message.Recipient.ChatId
	fmt.Println(upd.Message.Recipient)
	msg := maxbot.NewMessage().SetChat(chatID).SetText("Не известная команда")
	err := h.api.Messages.Send(ctx, msg)
	return err
}
