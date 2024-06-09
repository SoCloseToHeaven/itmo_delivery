package telegram

import "itmo_delivery/model"

type ActionsMap map[string]model.UserState

type NavigationMap map[model.UserState]ActionsMap

var Nav = NavigationMap{
	model.Main: mainActionsMap,

	model.AboutBot:    aboutBotActionsMap,
	model.Support:     supportActionsMap,
	model.Feedback:    feedbackActionsMap,
	model.Instruction: instructionActionsMap,

	model.MyOrders: myOrdersActionsMap,

	model.NewOrderSelectBuilding:   newOrderSelectBuildingActionsMap,
	model.NewOrderInputDescription: newOrderInputDescriptionActionsMap,
	model.NewOrderConfirm:          newOrderConfirmActionsMap,

	model.CourierSelectBuilding: courierSelectBuildingActionsMap,
	model.CourierActiveOrders:   courierActiveOrdersActionsMap,
	model.CourierConfirmOrder:   courierConfirmOrder,
}

var mainActionsMap = ActionsMap{
	NewOrderButtonText:    model.NewOrderSelectBuilding,
	AboutBotButtonText:    model.AboutBot,
	WatchOrdersButtonText: model.CourierSelectBuilding,
	MyOrdersButtonText:    model.MyOrders,
}

var aboutBotActionsMap = ActionsMap{
	BackButtonText:        model.Main,
	SupportButtonText:     model.Support,
	FeedbackButtonText:    model.Feedback,
	InstructionButtonText: model.Instruction,
}

var supportActionsMap = ActionsMap{
	BackButtonText: model.AboutBot,
}

var feedbackActionsMap = ActionsMap{
	BackButtonText: model.AboutBot,
}

var instructionActionsMap = ActionsMap{
	BackButtonText: model.AboutBot,
}

var myOrdersActionsMap = ActionsMap{
	BackButtonText: model.Main,
}

var newOrderSelectBuildingActionsMap = mapAvailableBuildings(
	model.NewOrderInputDescription,
	model.Main,
)

// Переход на подверждение заказа через текстовое поле
var newOrderInputDescriptionActionsMap = ActionsMap{
	BackButtonText: model.NewOrderSelectBuilding,
}

var newOrderConfirmActionsMap = ActionsMap{
	CancelButtonText:  model.Main,
	ConfirmButtonText: model.Main,
}

var courierSelectBuildingActionsMap = mapAvailableBuildings(
	model.CourierActiveOrders,
	model.Main,
)

// Переход на подверждение заказа через текстовое поле
var courierActiveOrdersActionsMap = ActionsMap{
	BackButtonText: model.CourierSelectBuilding,
}

var courierConfirmOrder = ActionsMap{
	ConfirmButtonText: model.Main,
	CancelButtonText:  model.Main,
}

func mapAvailableBuildings(nextState model.UserState, prevState model.UserState) ActionsMap {
	mapped := ActionsMap{
		BackButtonText: prevState,
	}

	for _, building := range model.AvailableBuildings {
		mapped[building] = nextState
	}

	return mapped
}
