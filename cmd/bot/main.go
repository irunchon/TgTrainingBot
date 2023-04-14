package main

import (
	"TgTrainingBot/internal/app/commands"
	"TgTrainingBot/internal/service/inmemory"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN") ///////////// Check?? os.LookupEnv second arg found for checking

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//var productService service.Service
	productService := inmemory.NewService()
	commander := commands.NewCommander(bot, productService)

	commander.HandleUpdates()
}
