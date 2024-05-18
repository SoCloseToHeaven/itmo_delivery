package telegram

import (
	"itmo_delivery/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// returns new state for a user
type Handler func(bot *tgbotapi.BotAPI, u tgbotapi.Update, state model.UserState) model.UserState

var stateToHandler = map[model.UserState]Handler{}

func MessageHandler(bot *tgbotapi.BotAPI, u tgbotapi.Update, state model.UserState) model.UserState {

	handler, found := stateToHandler[state]

	if !found {
		return model.Main
	}

	return handler(bot, u, state)
}
