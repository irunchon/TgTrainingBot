package commands

import (
	"TgTrainingBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	Service service.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService service.Service) *Commander {
	return &Commander{
		bot:     bot,
		Service: productService,
	}
}

//type Commander interface {
//	Help(inputMsg *tgbotapi.Message) tgbotapi.MessageConfig
//	Get(inputMsg *tgbotapi.Message) tgbotapi.MessageConfig
//	List(inputMsg *tgbotapi.Message) tgbotapi.MessageConfig
//	Delete(inputMsg *tgbotapi.Message) tgbotapi.MessageConfig
//
//	New(inputMsg *tgbotapi.Message) tgbotapi.MessageConfig  // return error not implemented
//	Edit(inputMsg *tgbotapi.Message) tgbotapi.MessageConfig // return error not implemented
//}

//func NewCommander(bot *tgbotapi.BotAPI, service service.InMemoryService) Commander {}
