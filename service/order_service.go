package service

import (
	"itmo_delivery/repository"

	"gorm.io/gorm"
)

type OrderService interface {
}

type orderService struct {
	OrderRepository repository.OrderRepository
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{
		OrderRepository: repository.NewOrderRepository(db),
	}
}
