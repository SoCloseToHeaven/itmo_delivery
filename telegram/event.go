package telegram

import (
	"itmo_delivery/model"
	"itmo_delivery/repository"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type Handler func(u tgbotapi.Update)

type updateHandler struct {
	db              *gorm.DB
	UserRepository  repository.UserRepository
	OrderRepository repository.OrderRepository
	StateToHandler  map[model.UserState]Handler
	Bot             *tgbotapi.BotAPI
}

type MessageHandler interface {
	Handle(u tgbotapi.Update)
}

func NewMessageHandler(db *gorm.DB, bot *tgbotapi.BotAPI) MessageHandler {
	return &updateHandler{
		db:              db,
		UserRepository:  repository.NewUserRepository(db),
		OrderRepository: repository.NewOrderRepository(db),
		StateToHandler:  map[model.UserState]Handler{},
		Bot:             bot,
	}
}

func (r *updateHandler) Handle(u tgbotapi.Update) {
	// TODO: add logic

}

func (r *updateHandler) setUserState(user *model.User, newState model.UserState) error {
	user.State = newState
	if err := r.UserRepository.Update(user); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(user.ChatID, "")
	msg.ReplyMarkup = StateToKeyboard[newState]

	_, err := r.Bot.Send(msg)
	return err
}
