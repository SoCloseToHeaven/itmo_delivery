package repository

import (
	"itmo_delivery/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *model.Order) error
	GetByID(id uint) (*model.Order, error)
	GetByState(state model.OrderState) (*model.Order, error)
	GetByPlace(place model.Building) (*model.Order, error)
	Update(order *model.Order) error
	Delete(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}
func (r *orderRepository) GetByID(id uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *orderRepository) GetByState(state model.OrderState) (*model.Order, error) {
	var order model.Order
	if err := r.db.Where("state = ?", state).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *orderRepository) GetByPlace(place model.Building) (*model.Order, error) {
	var order model.Order
	if err := r.db.Where("place = ?", place).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *orderRepository) Update(order *model.Order) error {
	return r.db.Save(order).Error
}
func (r *orderRepository) Delete(order *model.Order) error {
	return r.db.Delete(order).Error
}
