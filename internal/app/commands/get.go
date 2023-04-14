package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) get(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Printf("Wrong argument: %s", args)
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "Argument should be a number")
	}

	product, err := c.productService.Get(idx)

	if err != nil {
		log.Printf("Fail to get inmemory with idx = %d: %v", idx, err)
		return tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Error! %v", err),
		)
	}
	return tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)
}
