package main

import (
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

	for update := range updates {
		if update.Message == nil { // ignore any non-message updates
			continue
		}

		var outputMessage tgbotapi.MessageConfig

		switch update.Message.Command() {
		case "help":
			outputMessage = helpCommand(update.Message)
		default:
			outputMessage = defaultBehaviour(update.Message)
		}

		_, err := bot.Send(outputMessage)
		if err != nil {
			log.Panic(err)
		}
	}
}

func helpCommand(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "Write something and bot will repeat it")
}

func defaultBehaviour(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
}
