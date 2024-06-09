package telegram

import (
	"errors"
	"itmo_delivery/model"
	"itmo_delivery/utils"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Event func(handler UpdateHandler, user *model.User, u tgbotapi.Update) error

type ChangeStateEvent func(handler UpdateHandler, user *model.User) error

// CurrentState -> Possible Event
var CurrentEvents = map[model.UserState]Event{
	model.Main: navigationOnlyEvent,

	model.AboutBot:    navigationOnlyEvent,
	model.Support:     navigationOnlyEvent,
	model.Feedback:    navigationOnlyEvent,
	model.Instruction: navigationOnlyEvent,
	// тут какая-то хуйня с порядком сообщений, я пиздец намудрил, надо это как-то адекватнее сделать
	model.MyOrders: navigationOnlyEvent,

	model.NewOrderSelectBuilding:   selectBuildingEvent,
	model.NewOrderInputDescription: InputDescriptionEvent,
	model.NewOrderConfirm:          ConfirmOrderEvent,

	model.CourierSelectBuilding: navigationOnlyEvent,
	model.CourierActiveOrders:   navigationOnlyEvent,
	model.CourierConfirmOrder:   navigationOnlyEvent,
}

var ChangeStateEvents = map[model.UserState]ChangeStateEvent{
	model.Main: sendStateMsg,

	model.AboutBot:    sendStateMsg,
	model.Support:     sendStateMsg,
	model.Feedback:    sendStateMsg,
	model.Instruction: sendStateMsg,

	model.MyOrders: sendMyOrders,

	model.NewOrderSelectBuilding:   sendStateMsg,
	model.NewOrderInputDescription: sendStateMsg,
	model.NewOrderConfirm:          sendStateMsg,

	model.CourierSelectBuilding: sendStateMsg,
	model.CourierActiveOrders:   sendStateMsg,
	model.CourierConfirmOrder:   sendStateMsg,
}

func navigationOnlyEvent(handler UpdateHandler, user *model.User, u tgbotapi.Update) error {
	nextState, found := Nav[user.State][u.Message.Text]

	if !found {
		return handler.SendErrMsg(user)
	}

	return moveToNextState(handler, nil, user, nextState)
}

func moveToNextState(handler UpdateHandler, reply *tgbotapi.MessageConfig, user *model.User, newState model.UserState) error {
	if err := handler.UserService().UpdateUserState(user, newState); err != nil {
		_ = handler.SendErrMsg(user)
		return err
	}

	if reply != nil {
		return handler.SendMsg(*reply)
	}

	if event, found := ChangeStateEvents[newState]; found {
		return event(handler, user)
	}

	return errors.New("event not found")
}

const orderMaxPrintCount = 5 // TODO: move to bot config

func sendMyOrders(handler UpdateHandler, user *model.User) error {
	orders, err := handler.OrderService().GetLastOrderMessagesByUser(user, orderMaxPrintCount)

	if err != nil {
		_ = handler.SendErrMsg(user)
		return err
	}

	if len(*orders) == 0 {
		msg := tgbotapi.NewMessage(
			user.ChatID,
			utils.NoOrders,
		)

		return handler.SendMsgWithKeyboard(user, msg)
	}

	return handler.SendMsgWithKeyboard(user, *orders...)
}

func sendStateMsg(handler UpdateHandler, user *model.User) error {
	chatID := user.ChatID

	msgText, found := utils.StateMessage[user.State]

	if !found {
		return handler.SendErrMsg(user)
	}

	reply := tgbotapi.NewMessage(
		chatID,
		msgText,
	)

	handler.SetStateKeyboard(user.State, &reply)

	return handler.SendMsg(reply)
}
