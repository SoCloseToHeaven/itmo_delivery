package telegram

import (
	"itmo_delivery/model"
	"itmo_delivery/repository"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type EventHandler func(user *model.User, u tgbotapi.Update) tgbotapi.MessageConfig

type updateHandler struct {
	db              *gorm.DB
	UserRepository  repository.UserRepository
	OrderRepository repository.OrderRepository
	StateToHandler  map[model.UserState]EventHandler
	Bot             *tgbotapi.BotAPI
}

type MessageHandler interface {
	Handle(u tgbotapi.Update)
}

func NewMessageHandler(db *gorm.DB, bot *tgbotapi.BotAPI) MessageHandler {
	handler := &updateHandler{
		db:              db,
		UserRepository:  repository.NewUserRepository(db),
		OrderRepository: repository.NewOrderRepository(db),
		StateToHandler:  map[model.UserState]EventHandler{},
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

	reply := tgbotapi.NewMessage(user.ChatID, "privetik")
	r.setStateKeyboard(user, &reply)
	r.Bot.Send(reply)
}

func (r *updateHandler) setStateKeyboard(user *model.User, msg *tgbotapi.MessageConfig) {
	keyboard := StateToKeyboard[user.State]
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = true

	msg.ReplyMarkup = StateToKeyboard[user.State]
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
