package telegram

import (
	"fmt"
	"itmo_delivery/model"
	"itmo_delivery/utils"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CourierSelectBuildingEvent(handler UpdateHandler, user *model.User, u tgbotapi.Update) error {
	text := u.Message.Text

	var reply tgbotapi.MessageConfig

	nextState, found := Nav[user.State][text]

	if !found {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.IncorrectBuilding,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	var building *string = nil
	for _, elem := range model.AvailableBuildings {
		if text == elem {
			building = &elem
			break
		}
	}

	if building == nil {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.BackButtonClicked,
		)
		return moveToNextState(handler, &reply, user, nextState)
	}

	handler.OrderService().SetCourierBuilding(user, *building)

	reply = tgbotapi.NewMessage(
		user.ChatID,
		fmt.Sprintf(utils.SuccessfullySelectedBuildingFormatted, *building),
	)

	return moveToNextState(handler, &reply, user, nextState)
}

func SendActiveOrdersChangeEvent(handler UpdateHandler, user *model.User) error {
	place := handler.OrderService().GetCourierBuilding(user)

	var reply tgbotapi.MessageConfig

	if place == nil {
		return handler.SendErrMsg(user)
	}

	orders, err := handler.OrderService().GetActiveOrderMessagesByPlace(user, *place)

	if err != nil || orders == nil {
		return handler.SendErrMsg(user)
	}

	if len(*orders) == 0 {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.NoActiveOrders,
		)
		return handler.SendMsgWithKeyboard(user, reply)
	}

	return handler.SendMsgWithKeyboard(user, *orders...)
}
