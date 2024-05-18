package telegram

import (
	"itmo_delivery/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

// returns new state for a user
type Handler func(bot *tgbotapi.BotAPI, u tgbotapi.Update, state model.UserState) model.UserState
