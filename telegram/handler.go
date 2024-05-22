package telegram

import (
	"itmo_delivery/model"
	"itmo_delivery/service"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type updateHandler struct {
	DB           *gorm.DB
	UserService  service.UserService
	OrderService service.OrderService
	Bot          *tgbotapi.BotAPI
}

type MessageHandler interface {
	Handle(u tgbotapi.Update)
}

func NewMessageHandler(db *gorm.DB, bot *tgbotapi.BotAPI) MessageHandler {
	handler := &updateHandler{
		DB:           db,
		UserService:  service.NewUserService(db),
		OrderService: service.NewOrderService(db),
		Bot:          bot,
	}

	return handler
}

// add logging

func (r *updateHandler) Handle(u tgbotapi.Update) {
	if u.Message == nil {
		return
	}

	user, err := r.UserService.GetOrCreateUser(u)

	if err != nil {
		return
	}

	if err = CurrentEvents[user.State](r, user, u); err != nil {
		log.Println(err.Error())
	}
}

func (r *updateHandler) setStateKeyboard(state model.UserState, msg *tgbotapi.MessageConfig) {
	keyboard := StateToKeyboard[state]
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = true

	msg.ReplyMarkup = StateToKeyboard[state]
}

func (r *updateHandler) sendErrMsg(user *model.User) error {
	chatID := user.ChatID
	state := user.State

	reply := tgbotapi.NewMessage(chatID, ErrorMsg)

	r.setStateKeyboard(state, &reply)

	return r.sendMsg(reply)
}

func (r *updateHandler) sendMsg(messages ...tgbotapi.MessageConfig) error {

	for _, msg := range messages {
		if _, err := r.Bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}
