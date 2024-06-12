package telegram

import (
	"itmo_delivery/model"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	NewOrderButtonText    = "📝 Новый заказ"
	WatchOrdersButtonText = "📦 Взять заказ"
	AboutBotButtonText    = "🤖 О боте"
	MyOrdersButtonText    = "🍟 Мои заказы"
	BackButtonText        = "🔙 Назад"
	ConfirmButtonText     = "☑️ Подтвердить"
	CancelButtonText      = "❌ Отменить"
	SupportButtonText     = "⚙️ Обратиться в поддержку"
	FeedbackButtonText    = "📲 Оставить обратную связь"
	InstructionButtonText = "📄 Инструкция пользователя"
)

func mapBuildingToKeyboardRow(building ...string) []tgbotapi.KeyboardButton {
	newArr := []tgbotapi.KeyboardButton{}
	for _, elem := range building {
		newArr = append(newArr, tgbotapi.NewKeyboardButton(string(elem)))
	}
	return tgbotapi.NewKeyboardButtonRow(newArr...)
}

var MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(NewOrderButtonText),
		tgbotapi.NewKeyboardButton(WatchOrdersButtonText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(AboutBotButtonText),
		tgbotapi.NewKeyboardButton(MyOrdersButtonText),
	),
)

var aboutBotMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BackButtonText),
		tgbotapi.NewKeyboardButton(SupportButtonText),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(InstructionButtonText),
		tgbotapi.NewKeyboardButton(FeedbackButtonText),
	),
)

var selectBuildingMenu = tgbotapi.NewReplyKeyboard(
	mapBuildingToKeyboardRow(model.AvailableBuildings...),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BackButtonText),
	),
)

var backMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(BackButtonText),
	),
)

var confirmMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(CancelButtonText),
		tgbotapi.NewKeyboardButton(ConfirmButtonText),
	),
)

var StateToKeyboard = map[model.UserState]tgbotapi.ReplyKeyboardMarkup{
	model.Main: MainMenuKeyboard,

	model.NewOrderSelectBuilding:   selectBuildingMenu,
	model.NewOrderInputDescription: backMenu,
	model.NewOrderConfirm:          confirmMenu,

	model.MyOrders: backMenu,

	model.AboutBot:    aboutBotMenu,
	model.Instruction: backMenu,
	model.Feedback:    backMenu,
	model.Support:     backMenu,

	model.CourierSelectBuilding: selectBuildingMenu,
	model.CourierActiveOrders:   backMenu,
	model.CourierConfirmOrder:   confirmMenu,
}
