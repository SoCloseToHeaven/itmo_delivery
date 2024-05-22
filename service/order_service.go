package service

import (
	"itmo_delivery/model"
	"itmo_delivery/repository"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type OrderService interface {
	GetLastOrderMessagesByUser(user *model.User, count uint) (*[]tgbotapi.MessageConfig, error)
}

type orderService struct {
	OrderRepository repository.OrderRepository
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{
		OrderRepository: repository.NewOrderRepository(db),
	}
}

func (r *orderService) GetLastOrderMessagesByUser(user *model.User, count uint) (*[]tgbotapi.MessageConfig, error) {
	chatID := user.ChatID

	orders, err := r.OrderRepository.GetLastOrdersByUser(user, count)

	if err != nil {
		return nil, err
	}

	var messages []tgbotapi.MessageConfig
	for _, order := range *orders {
		orderMsg := tgbotapi.NewMessage(
			chatID,
			order.ToString(),
		)

		messages = append(messages, orderMsg)
	}

	return &messages, nil
}
