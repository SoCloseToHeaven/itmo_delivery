package telegram

import (
	"errors"
	"itmo_delivery/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type TelegramChatID int64
type TempOrders map[TelegramChatID]model.Order // map for creating temporary orders

// var tempOrders = make(TempOrders)

type Event func(handler *updateHandler, user *model.User, u tgbotapi.Update) error

// CurrentState -> Possible Event
var Events = map[model.UserState]Event{
	model.Main: navigationOnly,

	model.AboutBot:    navigationOnly,
	model.Support:     navigationOnly,
	model.Feedback:    navigationOnly,
	model.Instruction: navigationOnly,

	// model.MyOrders: MyOrders,

	// model.NewOrderSelectBuilding:   SelectBuilding,
	// model.NewOrderInputDescription: InputDescription,
	// model.NewOrderConfirm:          OrderConfirmMessageFormatted,

	// model.CourierSelectBuilding: SelectBuilding,
	// model.CourierActiveOrders:   ActiveOrders,
	// model.CourierConfirmOrder:   OrderConfirmMessageFormatted,
}

func navigationOnly(handler *updateHandler, user *model.User, u tgbotapi.Update) error {
	chatID := u.Message.Chat.ID
	nextState, found := Nav[user.State][u.Message.Text]

	if !found {
		return errors.New("state not found in navigation rules")
	}

	msgText, found := StateMessage[nextState]

	if !found {
		return errors.New("state not found in state messages")
	}

	reply := tgbotapi.NewMessage(
		chatID,
		msgText,
	)

	return handler.next(reply, user, nextState)
}
