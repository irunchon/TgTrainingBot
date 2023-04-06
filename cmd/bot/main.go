package main

import (
	"TgTrainingBot/internal/app/commands"
	"TgTrainingBot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil { // ignore any non-message updates
			continue
		}

		var outputMessage tgbotapi.MessageConfig

		switch update.Message.Command() {
		case "help":
			outputMessage = commands.HelpCommand(update.Message)
		case "list":
			outputMessage = commands.ListCommand(update.Message, productService)
		default:
			outputMessage = commands.DefaultBehaviour(update.Message)
		}

		_, err := bot.Send(outputMessage)
		if err != nil {
			log.Panic(err)
		}
	}
}
