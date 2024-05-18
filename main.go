package main

import (
	"log"
	"os"

	// telegram "itmo_delivery/telegram"
	model "itmo_delivery/model"
	telegram "itmo_delivery/telegram"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var userStates = make(map[int64]model.UserState)

func main() {
	botApiKey := os.Getenv("ITMO_DELIVERY_BOT_API_KEY")

	if botApiKey == "" {
		log.Panic("ITMO_DELIVERY_BOT_API_KEY ENV VARIABLE IS NOT SET")
	}

	bot, err := tgbotapi.NewBotAPI(botApiKey)

	if err != nil {
		log.Panic(err)
	}

	MustInitializeDB()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updatesChannel, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updatesChannel {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			chatID := update.Message.Chat.ID

			if _, ok := userStates[chatID]; !ok {

				// TODO: Unregistered users should be returned to register menu, registered users should be returned to main menu
				userStates[chatID] = model.Main
			}

			userState := userStates[chatID]

			newState := telegram.MessageHandler(bot, update, userState)

			userStates[chatID] = newState
		}
	}
}
