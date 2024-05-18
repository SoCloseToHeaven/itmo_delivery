package model

type User struct {
	ID     uint64
	ChatID int64
	State  UserState
}

type UserState uint64

const (
	Start UserState = iota
	Main
	NewOrderSelectBuilding
	NewOrderInputDescription
	NewOrderInputContacts
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
