package commands

import (
	"TgTrainingBot/internal/service/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (c *Commander) edit(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	stringWithArguments := inputMessage.CommandArguments()

	if stringWithArguments == "" {
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "Empty argument. No changes can be done")
	}

	args := strings.Split(stringWithArguments, " ")

	ID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Printf("Wrong first argument: %s (failed to convert to uint64)", args[0])
		return tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Error! Arguments should contain ID (non-negative number) and product's name",
		)
	}
	if len(args) == 1 {
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "No product name. No changes can be done")
	}

	err = c.service.Update(ID, model.Product{Title: args[1]})
	if err != nil {
		log.Printf("Fail to get product with ID = %d: %v", ID, err)
		return tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Product not found",
		)
	}

	return tgbotapi.NewMessage(inputMessage.Chat.ID, "Product was successfully updated")
}
