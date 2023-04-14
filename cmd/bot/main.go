package main

import (
	"TgTrainingBot/internal/app/commands"
	service "TgTrainingBot/internal/service/inmemory"
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

	productService := service.NewService()
	commander := commands.NewCommander(bot, productService)

	commander.HandleUpdates()
}
