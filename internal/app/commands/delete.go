package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) delete(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()
	ID, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		log.Printf("Wrong argument: %s", args)
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "Error! Argument should be non-negative number")
	}

	_, err = c.Service.Remove(ID)

	if err != nil {
		log.Printf("Fail to get product with ID = %d: %v", ID, err)
		return tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Product not found",
		)
	}
	return tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Product with ID = %s was successfully deleted", args),
	)
}
