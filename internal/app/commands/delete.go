package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) delete(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()

	ID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("Wrong argument: %s (failed to convert to uint64)", args)
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "Error! Argument should be non-negative number")
	}

	err = c.service.Remove(ID)
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
