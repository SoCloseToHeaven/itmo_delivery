package telegram

import (
	"itmo_delivery/model"
	"itmo_delivery/service"
	"itmo_delivery/utils"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type updateHandler struct {
	db           *gorm.DB
	userService  service.UserService
	orderService service.OrderService
	bot          *tgbotapi.BotAPI
}

type UpdateHandler interface {
	Handle(u tgbotapi.Update)
	SetStateKeyboard(state model.UserState, msg *tgbotapi.MessageConfig)
	SendErrMsg(user *model.User) error
	SendMsg(messages ...tgbotapi.MessageConfig) error
	OrderService() service.OrderService
	UserService() service.UserService
	Bot() *tgbotapi.BotAPI
	DB() *gorm.DB
	SendMsgWithKeyboard(user *model.User, messages ...tgbotapi.MessageConfig) error
}

func NewUpdateHandler(db *gorm.DB, bot *tgbotapi.BotAPI) UpdateHandler {
	handler := &updateHandler{
		db:           db,
		userService:  service.NewUserService(db),
		orderService: service.NewOrderService(db),
		bot:          bot,
	}

	return handler
}

// add logging

func (r *updateHandler) Handle(u tgbotapi.Update) {
	if u.Message == nil {
		return
	}

	user, err := r.UserService().GetOrCreateUser(u)

	if err != nil {
		log.Println(err.Error())
		return
	}

	if err = CurrentEvents[user.State](r, user, u); err != nil {
		log.Println(err.Error())
	}
}

func (r *updateHandler) SetStateKeyboard(state model.UserState, msg *tgbotapi.MessageConfig) {
	keyboard := StateToKeyboard[state]
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = true

	msg.ReplyMarkup = StateToKeyboard[state]
}

func (r *updateHandler) SendErrMsg(user *model.User) error {
	chatID := user.ChatID
	state := user.State

	reply := tgbotapi.NewMessage(chatID, utils.ErrorMsg)

	r.SetStateKeyboard(state, &reply)

	return r.SendMsg(reply)
}

func (r *updateHandler) SendMsg(messages ...tgbotapi.MessageConfig) error {

	for _, msg := range messages {
		if _, err := r.Bot().Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (r *updateHandler) OrderService() service.OrderService {
	return r.orderService
}

func (r *updateHandler) UserService() service.UserService {
	return r.userService
}

func (r *updateHandler) Bot() *tgbotapi.BotAPI {
	return r.bot
}
func (r *updateHandler) DB() *gorm.DB {
	return r.db
}

func (r *updateHandler) SendMsgWithKeyboard(user *model.User, messages ...tgbotapi.MessageConfig) error {

	for _, msg := range messages {
		r.SetStateKeyboard(user.State, &msg)
		if _, err := r.Bot().Send(msg); err != nil {
			return err
		}
	}

	return nil
}
