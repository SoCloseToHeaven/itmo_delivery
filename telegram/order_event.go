package telegram

import (
	"fmt"
	"itmo_delivery/model"
	"itmo_delivery/utils"
	"unicode/utf8"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func SelectBuildingEvent(handler UpdateHandler, user *model.User, u tgbotapi.Update) error {
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

	temp := model.TempOrderInfo{
		Place: *building,
	}

	handler.OrderService().SetTempOrderByUser(user, temp)

	reply = tgbotapi.NewMessage(
		user.ChatID,
		fmt.Sprintf(utils.SuccessfullySelectedBuildingFormatted, *building),
	)

	return moveToNextState(handler, &reply, user, nextState)
}

const maxDescriptionSize uint = 500 // TODO: move to config
func InputDescriptionEvent(handler UpdateHandler, user *model.User, u tgbotapi.Update) error {
	text := u.Message.Text
	var reply tgbotapi.MessageConfig

	nextState, found := Nav[user.State][u.Message.Text]

	if found {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.BackButtonClicked,
		)
		return moveToNextState(handler, &reply, user, nextState)
	}

	temp := handler.OrderService().GetTempOrderByUser(user)

	if temp == nil {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	length := utf8.RuneCountInString(text)

	if length > int(maxDescriptionSize) {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			fmt.Sprintf(utils.IncorrectDescriptionSizeFormatted, maxDescriptionSize),
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	temp.Description = text

	handler.OrderService().SetTempOrderByUser(user, *temp)

	reply = tgbotapi.NewMessage(
		user.ChatID,
		fmt.Sprintf(utils.SuccessfullyDescriptionInputFormatted, text),
	)

	return moveToNextState(handler, &reply, user, model.NewOrderConfirm)
}

func ConfirmOrderEvent(handler UpdateHandler, user *model.User, u tgbotapi.Update) error {
	text := u.Message.Text
	var reply tgbotapi.MessageConfig

	nextState, found := Nav[user.State][text]

	if !found {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.UnknownAction,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	if found && text == CancelButtonText {
		return moveToNextState(handler, nil, user, model.Main)
	}

	tempOrder := handler.OrderService().GetTempOrderByUser(user)

	if tempOrder == nil {
		return moveToNextState(handler, nil, user, model.Main)
	}

	orderMessage, err := handler.OrderService().CreateNewOrderByUser(user)

	if err != nil {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, model.Main)
	}

	return moveToNextState(handler, orderMessage, user, nextState)

}

func SendConfirmMessageEvent(handler UpdateHandler, user *model.User) error {
	chatID := user.ChatID

	tempOrder := handler.OrderService().GetTempOrderByUser(user)

	if tempOrder == nil {
		return handler.SendErrMsg(user)
	}

	fmtMessage := fmt.Sprintf(utils.OrderConfirmMessageFormatted, tempOrder.ToOrder(user).ToString())

	reply := tgbotapi.NewMessage(
		chatID,
		fmtMessage,
	)

	return handler.SendMsgWithKeyboard(user, reply)
}
