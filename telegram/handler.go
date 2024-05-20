package telegram

import (
	"itmo_delivery/model"
	"itmo_delivery/repository"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type updateHandler struct {
	DB              *gorm.DB
	UserRepository  repository.UserRepository
	OrderRepository repository.OrderRepository
	Bot             *tgbotapi.BotAPI
}

type MessageHandler interface {
	Handle(u tgbotapi.Update)
}

func NewMessageHandler(db *gorm.DB, bot *tgbotapi.BotAPI) MessageHandler {
	handler := &updateHandler{
		DB:              db,
		UserRepository:  repository.NewUserRepository(db),
		OrderRepository: repository.NewOrderRepository(db),
		Bot:             bot,
	}

	return handler
}

// add logging

func (r *updateHandler) Handle(u tgbotapi.Update) {
	if u.Message == nil {
		return
	}

	user, err := r.getUser(u)

	if err != nil {
		return
	}

	if err = CurrentEvents[user.State](r, user, u); err != nil {
		log.Println(err.Error())
	}
}

func (r *updateHandler) updateUserState(user *model.User, newState model.UserState) error {
	if user.State == newState {
		return nil
	}

	user.State = newState
	return r.UserRepository.Update(user)
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

func (r *updateHandler) sendMsg(msg tgbotapi.MessageConfig) error {
	if _, err := r.Bot.Send(msg); err != nil {
		return err
	}

	return nil
}

// TODO: Add transactions
func (r *updateHandler) getUser(u tgbotapi.Update) (*model.User, error) {
	chatID := u.Message.Chat.ID
	tgID := u.Message.From.ID

	user, err := r.UserRepository.GetByChatID(chatID)

	if err == nil {
		return user, nil
	}

	user = &model.User{
		ChatID:     chatID,
		TelegramID: tgID,
		State:      model.Main,
	}

	err = r.UserRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
