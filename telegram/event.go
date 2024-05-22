package telegram

import (
	"itmo_delivery/model"
	"itmo_delivery/utils"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Event func(handler *updateHandler, user *model.User, u tgbotapi.Update) error

type ChangeStateEvent func(handler *updateHandler, user *model.User) error

// CurrentState -> Possible Event
var CurrentEvents = map[model.UserState]Event{
	model.Main: navigationOnly,

	model.AboutBot:    navigationOnly,
	model.Support:     navigationOnly,
	model.Feedback:    navigationOnly,
	model.Instruction: navigationOnly,
	// тут какая-то хуйня с порядком сообщений, я пиздец намудрил, надо это как-то адекватнее сделать
	model.MyOrders: navigationOnly,

	model.NewOrderSelectBuilding:   navigationOnly,
	model.NewOrderInputDescription: navigationOnly,
	model.NewOrderConfirm:          navigationOnly,

	model.CourierSelectBuilding: navigationOnly,
	model.CourierActiveOrders:   navigationOnly,
	model.CourierConfirmOrder:   navigationOnly,
}

var ChangeStateEvents = map[model.UserState]ChangeStateEvent{
	model.MyOrders: sendMyOrders,
}

func navigationOnly(handler *updateHandler, user *model.User, u tgbotapi.Update) error {
	chatID := u.Message.Chat.ID
	nextState, found := Nav[user.State][u.Message.Text]

	if !found {
		return handler.sendErrMsg(user)
	}

	msgText, found := utils.StateMessage[nextState]

	if !found {
		return handler.sendErrMsg(user)
	}

	reply := tgbotapi.NewMessage(
		chatID,
		msgText,
	)

	return moveToNextState(handler, reply, user, nextState)
}

func moveToNextState(handler *updateHandler, reply tgbotapi.MessageConfig, user *model.User, newState model.UserState) error {
	if err := handler.UserService.UpdateUserState(user, newState); err != nil {
		_ = handler.sendErrMsg(user)
		return err
	}

	if event, found := ChangeStateEvents[newState]; found {
		if err := event(handler, user); err != nil {
			return err
		}
	}

	handler.setStateKeyboard(user.State, &reply)

	return handler.sendMsg(reply)
}

const orderMaxPrintCount = 5 // TODO: move to bot config

func sendMyOrders(handler *updateHandler, user *model.User) error {
	orders, err := handler.OrderService.GetLastOrderMessagesByUser(user, orderMaxPrintCount)

	if err != nil {
		_ = handler.sendErrMsg(user)
		return err
	}

	return handler.sendMsg(*orders...)
}
