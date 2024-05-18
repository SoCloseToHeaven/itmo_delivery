package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	State          OrderState
	ContactInfo    string
	Reward         string
	Place          Building
	CreatorChatID  int64
	AssigneeChatID *int64 // intptr to allow null values
}

type OrderState uint64

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

type Building string

var AvailableBuildings = [...]Building{
	"Кронверкский пр. 49",
	"Ул. Ломоносова, 9",
	"Биржевая линия, 14-16",
	"Ул. Чайковского, 11/2",
	"Пер. Гривцова, 14-16",
}
