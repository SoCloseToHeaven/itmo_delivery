package model

type User struct {
	ID     uint64
	ChatID int64
	State  UserState
	ISU    uint64
}

type UserState uint64

const (
	Main UserState = iota

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
