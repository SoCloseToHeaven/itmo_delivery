package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	State       OrderState `gorm:"index"`
	Description string
	// ContactInfo    string
	// Reward         string
	Place          string `gorm:"index"`
	CreatorChatID  int64  `gorm:"index"`
	AssigneeChatID *int64 // intptr to allow null values
}

type TempOrderInfo struct {
	State       OrderState
	Description string
	// ContactInfo    string
	// Reward         string
	Place string
}

type OrderState uint64

func (temp *TempOrderInfo) ToOrder(user *User) *Order {
	return &Order{
		Description:   temp.Description,
		State:         temp.State,
		Place:         temp.Place,
		CreatorChatID: user.ChatID,
	}
}

const (
	Active OrderState = iota
	Cancelled
	GivenToCourier
)

var OrderToText = map[OrderState]string{
	Active:         "Активен",
	Cancelled:      "Отменён",
	GivenToCourier: "Передан курьеру",
}

var AvailableBuildings = []string{
	"Кронверкский пр. 49",
	"Ул. Ломоносова, 9",
	"Биржевая линия, 14-16",
	"Ул. Чайковского, 11/2",
	"Пер. Гривцова, 14-16",
}

func (r *Order) ToString() string {
	return fmt.Sprintf(
		"Здание: %s\nОписание:%s\nСостояние:%s\nДата создания:%s\n",
		r.Place,
		r.Description,
		OrderToText[r.State],
		r.CreatedAt.String(),
	)
}
