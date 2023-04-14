package commands

import (
	"TgTrainingBot/internal/service/model"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) new(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()

	if args == "" {
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "Empty argument. New product cannot be added")
	}

	newProductID := c.service.Create(model.Product{Title: args})

	return tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Product with ID = %d was successfully added", newProductID),
	)
}
