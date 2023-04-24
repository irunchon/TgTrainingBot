package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/irunchon/TgTrainingBot/internal/service"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service service.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService service.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: productService,
	}
}
