package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		c.handleMessage(update)
	}
}

func (c *Commander) handleMessage(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	if update.Message == nil { // ignore any non-message updates
		return
	}

	var outputMessage tgbotapi.MessageConfig

	switch update.Message.Command() {
	case "help":
		outputMessage = c.help(update.Message)
	case "list":
		outputMessage = c.list(update.Message)
	case "get":
		outputMessage = c.get(update.Message)
	default:
		outputMessage = c.noCommand(update.Message)
	}

	_, err := c.bot.Send(outputMessage)
	if err != nil {
		log.Panic(err)
	}
}
