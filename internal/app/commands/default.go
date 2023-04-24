package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) noCommand(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
}
