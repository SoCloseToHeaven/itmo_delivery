package service

import (
	"errors"
	"fmt"
	"itmo_delivery/model"
	"itmo_delivery/repository"
	"itmo_delivery/utils"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type OrderService interface {
	GetLastOrderMessagesByUser(user *model.User, count uint) (*[]tgbotapi.MessageConfig, error)
	GetActiveOrderMessagesByPlace(user *model.User, place string) (*[]tgbotapi.MessageConfig, error)
	CreateNewOrderByUser(user *model.User) (*tgbotapi.MessageConfig, error)
	AssigneeUserToOrder(user *model.User, id uint) *model.Order

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
		OrderRepository:   repository.NewOrderRepository(db),
		TempOrders:        make(TempOrders),
		CourierToBuilding: make(map[int64]string),
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
			order.ToStringWithID(),
		)

		messages = append(messages, orderMsg)
	}

	return &messages, nil
}

func (r *orderService) CreateNewOrderByUser(user *model.User) (*tgbotapi.MessageConfig, error) {
	tempOrder := r.GetTempOrderByUser(user)

	if tempOrder == nil {
		return nil, errors.New("temp order not found")
	}

	order := tempOrder.ToOrder(user)
	if err := r.OrderRepository.Create(order); err != nil {
		return nil, err
	}

	reply := tgbotapi.NewMessage(
		order.CreatorChatID,
		fmt.Sprintf(utils.NewOrderCreated, order.ToStringWithID()),
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

func (r *orderService) GetActiveOrderMessagesByPlace(user *model.User, place string) (*[]tgbotapi.MessageConfig, error) {
	chatID := user.ChatID

	orders, err := r.OrderRepository.GetByPlaceAndState(place, model.Active)

	if err != nil {
		return nil, err
	}

	var messages []tgbotapi.MessageConfig
	for _, order := range *orders {
		orderMsg := tgbotapi.NewMessage(
			chatID,
			order.ToStringWithID(),
		)

		messages = append(messages, orderMsg)
	}

	return &messages, nil
}

func (r *orderService) AssigneeUserToOrder(user *model.User, id uint) *model.Order {

	order, err := r.OrderRepository.GetByID(id)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if order == nil {
		return nil
	}

	order.AssigneeChatID = &user.ChatID
	order.State = model.GivenToCourier

	if err := r.OrderRepository.Update(order); err != nil {
		log.Println(err.Error())
		return nil
	}

	return order
}
