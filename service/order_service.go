package service

import (
	"fmt"
	"itmo_delivery/model"
	"itmo_delivery/repository"
	"itmo_delivery/utils"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type OrderService interface {
	GetLastOrderMessagesByUser(user *model.User, count uint) (*[]tgbotapi.MessageConfig, error)
	CreateNewOrderByUser(user *model.User) (*tgbotapi.MessageConfig, error)

	GetTempOrderByUser(user *model.User) *model.TempOrderInfo
	SetTempOrderByUser(user *model.User, temp model.TempOrderInfo)

	GetCourierBuilding(user *model.User) *string
	SetCourierBuilding(user *model.User, building string)
}

type orderService struct {
	OrderRepository   repository.OrderRepository
	TempOrders        TempOrders
	CourierToBuilding map[int64]string
}

type TempOrders map[int64]model.TempOrderInfo // map for creating temporary orders

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{
		OrderRepository: repository.NewOrderRepository(db),
		TempOrders:      make(TempOrders),
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

func (r *orderService) CreateNewOrderByUser(user *model.User) (*tgbotapi.MessageConfig, error) {
	order := r.GetTempOrderByUser(user).ToOrder(user)

	if err := r.OrderRepository.Create(order); err != nil {
		return nil, err
	}

	reply := tgbotapi.NewMessage(
		order.CreatorChatID,
		fmt.Sprintf(utils.NewOrderCreated, order.ToString()),
	)

	return &reply, nil
}

func (r *orderService) GetTempOrderByUser(user *model.User) *model.TempOrderInfo {
	if temp, found := r.TempOrders[user.ChatID]; found {
		return &temp
	}
	return nil
}

func (r *orderService) SetTempOrderByUser(user *model.User, temp model.TempOrderInfo) {
	r.TempOrders[user.ChatID] = temp
}

func (r *orderService) GetCourierBuilding(user *model.User) *string {
	if building, found := r.CourierToBuilding[user.ChatID]; found {
		return &building
	}
	return nil
}

func (r *orderService) SetCourierBuilding(user *model.User, building string) {
	r.CourierToBuilding[user.ChatID] = building
}
