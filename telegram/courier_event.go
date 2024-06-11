package telegram

import (
	"fmt"
	"itmo_delivery/model"
	"itmo_delivery/utils"
	"log"
	"strconv"

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

	_ = handler.SendMsgWithKeyboard(user, *orders...)

	reply = tgbotapi.NewMessage(
		user.ChatID,
		utils.CourierSelectOrderDescription,
	)

	return handler.SendMsgWithKeyboard(user, reply)
}

func CourierSelectOrderEvent(handler UpdateHandler, user *model.User, u tgbotapi.Update) error {
	text := u.Message.Text

	var reply tgbotapi.MessageConfig

	if nextState, found := Nav[user.State][text]; found {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.BackButtonClicked,
		)
		return moveToNextState(handler, &reply, user, nextState)
	}

	id, err := strconv.ParseUint(text, 10, 64)

	if err != nil {
		log.Println(err.Error())
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	order := handler.OrderService().AssigneeUserToOrder(user, uint(id))

	if order == nil {
		log.Println("CourierSelectEvent -> order is nil")
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	creatorChatID := order.CreatorChatID
	assigneeChatID := order.AssigneeChatID

	creatorChat, err := handler.Bot().GetChat(tgbotapi.ChatConfig{ChatID: creatorChatID})

	if err != nil {
		log.Println(err.Error())
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	assigneeChat, err := handler.Bot().GetChat(tgbotapi.ChatConfig{ChatID: *assigneeChatID})

	if err != nil {
		log.Println(err.Error())
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, user.State)
	}

	orderCreatorReply := tgbotapi.NewMessage(
		creatorChatID,
		fmt.Sprintf(utils.YourOrderTakenFormatted, order.ToStringWithID(), assigneeChat.UserName), // TODO: add checking if Username is present
	)

	creatorUser := handler.UserService().GetByChatID(creatorChatID)

	if creatorUser == nil {
		reply = tgbotapi.NewMessage(
			user.ChatID,
			utils.ErrorMsg,
		)
		return moveToNextState(handler, &reply, user, model.Main)
	}

	_ = handler.SendMsgWithKeyboard(creatorUser, orderCreatorReply)

	orderAssigneeReply := tgbotapi.NewMessage(
		*assigneeChatID,
		fmt.Sprintf(utils.CourierOrderTakenFormatted, order.ToStringWithID(), creatorChat.UserName), // TODO: add checking if Username is present
	)

	return moveToNextState(handler, &orderAssigneeReply, user, model.Main)
}
