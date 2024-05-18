package telegram

import (
	"itmo_delivery/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func mapBuildingToKeyboardRow(building ...model.Building) []tgbotapi.KeyboardButton {
	newArr := []tgbotapi.KeyboardButton{}
	for _, elem := range building {
		newArr = append(newArr, tgbotapi.NewKeyboardButton(string(elem)))
	}
	return tgbotapi.NewKeyboardButtonRow(newArr...)
}

var MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Новый заказ"),
		tgbotapi.NewKeyboardButton("Посмотреть заказы(Курьер)"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("О боте"),
		tgbotapi.NewKeyboardButton("Мои заказы"),
	),
)

var aboutBotMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
		tgbotapi.NewKeyboardButton("Обратиться в поддержку"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Инструкция пользователя"),
		tgbotapi.NewKeyboardButton("Оставить обратную связь"),
	),
)

var selectBuildingMenu = tgbotapi.NewReplyKeyboard(
	mapBuildingToKeyboardRow(model.AvailableBuildings...),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var backMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Отменить"),
	),
)

var confirmMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Отменить"),
		tgbotapi.NewKeyboardButton("Подтвердить"),
	),
)

var StateToKeyboard = map[model.UserState]tgbotapi.ReplyKeyboardMarkup{
	model.Main: MainMenuKeyboard,

	model.NewOrderSelectBuilding:   selectBuildingMenu,
	model.NewOrderInputDescription: backMenu,
	model.NewOrderConfirm:          confirmMenu,

	model.AboutBot:    aboutBotMenu,
	model.Instruction: aboutBotMenu,
	model.Feedback:    aboutBotMenu,
	model.Support:     aboutBotMenu,

	model.CourierSelectBuilding: selectBuildingMenu,
	model.CourierActiveOrders:   backMenu,
	model.CourierConfirmOrder:   confirmMenu,
}
