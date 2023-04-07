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

	var outputMessage tgbotapi.MessageConfig

	if update.CallbackQuery != nil {
		outputMessage = tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data: "+update.CallbackQuery.Data,
		)
		_, err := c.bot.Send(outputMessage)
		if err != nil {
			log.Panic(err)
		}
		return
	}

	if update.Message == nil { // ignore any non-message updates
		return
	}

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
