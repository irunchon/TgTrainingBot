package main

import (
	"TgTrainingBot/internal/app/commands"
	"TgTrainingBot/internal/service/inmemory"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(fmt.Sprintf("Error appeared while starting the bot: %v", err))
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	productService := inmemory.NewService()
	commander := commands.NewCommander(bot, productService)

	commander.HandleUpdates()
}
