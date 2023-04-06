package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) list(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	outputMsgText := "Here are all the products:\n\n"
	products := c.productService.List()

	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	return tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
}
