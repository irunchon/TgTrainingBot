package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	token := os.Getenv("MYTOKEN")

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

		switch update.Message.Command() {
		case "help":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Write something and bot will repeat it")
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		default:
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
