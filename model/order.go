package model

type Order struct {
	ID             uint64
	State          OrderState
	ContactInfo    string
	Reward         string
	CreatorChatID  int64
	AssignedChatID int64
}

type OrderState string

const (
	Active         OrderState = "Активен"
	Cancelled      OrderState = "Отменен"
	GivenToCourier OrderState = "Передан курьеру"
)

type Building string

const ()
