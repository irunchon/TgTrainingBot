package commands

import (
	"TgTrainingBot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ListCommand(inputMessage *tgbotapi.Message, productService *product.Service) tgbotapi.MessageConfig {
	outputMsgText := "Here are all the products:\n\n"

	products := productService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	return tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
}
