package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ChatID     int64 `gorm:"index"`
	TelegramID int   `gorm:"index"`
	State      UserState
	// ISU        uint64
}

type UserState uint64

const (
	Main UserState = iota

	NewOrderSelectBuilding
	NewOrderInputDescription
	NewOrderConfirm

	MyOrders

	AboutBot
	Instruction
	Feedback
	Support

	CourierSelectBuilding
	CourierActiveOrders
	CourierConfirmOrder
)
