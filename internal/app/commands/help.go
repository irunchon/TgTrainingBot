package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) help(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "Write something and bot will repeat it\n"+
		"List of available commands: \n"+
		"/help — help info\n"+
		"/list — list of products\n"+
		"/get — get info\n"+
		"/delete — delete an existing entity\n"+
		"/new — create a new entity\n"+
		"/edit — edit an entity")
}
