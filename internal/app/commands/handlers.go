package commands

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		c.handleSingleUpdate(update)
	}
}

func (c *Commander) handleSingleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		c.handleCallback(update)
	}
	if update.Message != nil {
		c.handleMessage(update)
	}
}

func (c *Commander) handleCallback(update tgbotapi.Update) {
	outputMessage := tgbotapi.NewMessage(
		update.CallbackQuery.Message.Chat.ID,
		fmt.Sprintf("Offset: %s (functionality TBD)", update.CallbackQuery.Data),
	)
	_, err := c.bot.Send(outputMessage)
	if err != nil {
		log.Panic(err)
	}
}

func (c *Commander) handleMessage(update tgbotapi.Update) {
	var outputMessage tgbotapi.MessageConfig

	switch update.Message.Command() {
	case "help":
		outputMessage = c.help(update.Message)
	case "list":
		outputMessage = c.list(update.Message)
	case "get":
		outputMessage = c.get(update.Message)
	case "delete":
		outputMessage = c.delete(update.Message)
	case "new":
		outputMessage = c.new(update.Message)
	case "edit":
		outputMessage = c.edit(update.Message)
	default:
		outputMessage = c.noCommand(update.Message)
	}

	_, err := c.bot.Send(outputMessage)
	if err != nil {
		log.Panic(err)
	}
}
