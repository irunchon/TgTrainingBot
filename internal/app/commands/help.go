package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) help(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID,
		"List of available commands: \n"+
			"/help — help info\n"+
			"/list — list of products\n"+
			"/get — get product info by ID\n"+
			"/delete — delete an existing entity by ID\n"+
			"/new — add new product\n"+
			"/edit — edit an existing product by ID\n"+
			"\n or just write something and bot will repeat it :)")
}
