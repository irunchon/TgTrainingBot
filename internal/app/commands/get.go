package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) get(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()
	idx, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Printf("Wrong argument: %s", args)
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "Error! Argument should be non-negative number")
	}

	product, err := c.Service.Get(idx)

	if err != nil {
		log.Printf("Fail to get inmemory with idx = %d: %v", idx, err)
		return tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Product not found",
		)
	}
	return tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)
}
