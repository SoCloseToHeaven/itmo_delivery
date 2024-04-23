package main

import (
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	botApiKey := os.Getenv("ITMO_DELIVERY_BOT_API_KEY")

	if botApiKey == "" {
		log.Panic("ITMO_DELIVERY_BOT_API_KEY ENV VARIABLE IS NOT SET")
	}

	bot, err := tgbotapi.NewBotAPI(botApiKey)

	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updatesChannel, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updatesChannel {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Test?")

			bot.Send(msg)
		}
	}
}
