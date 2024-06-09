package utils

import "itmo_delivery/model"

const (
	StartMsg                              string = "Добро пожаловать в сервис ITMO.DELIVERY"
	ErrorMsg                              string = "Упс... Что-то пошло не так"
	UnknownAction                         string = "Неизвестное действие!"
	SelectBuilding                        string = "Выберите нужное здание в меню."
	IncorrectBuilding                            = "Неверное название здания! Попробуйте ещё раз."
	SuccessfullySelectedBuildingFormatted        = "Успешно выбрано здание: %s"
	SuccessfullyDescriptionInputFormatted        = "Успешно введено описание: %s"
	MyOrders                              string = "Мои последние заказы"
	ActiveOrders                          string = "Активные заказы:"
	OrderConfirmMessageFormatted          string = "Вы подтверждаете заказ? \n%s"
	InputDescription                      string = "Введите описание заказа:"
	IncorrectDescriptionSizeFormatted            = "Некорректный размер сообщения! Допустимо не больше %d символов!"
	Instruction                           string = "Инструкция"   // TODO: Нормальный текст инструкции
	Feedback                              string = "Фидбек"       // TODO: Нормальный текст фидбека
	AboutBot                              string = "О боте"       // TODO: Нормальный текст о боте
	MainMenu                                     = "Главное меню" // TODO: Нормальный текст главного меню
	Support                                      = "Поддержечка"  // TODO: Нормальный текст поддержечки
	NewOrderCreated                              = "Успешно создан новый заказ! \n%s"
	BackButtonClicked                            = "Переходим назад..."
)

var StateMessage = map[model.UserState]string{
	model.Main: MainMenu,

	model.AboutBot:    AboutBot,
	model.Support:     Support,
	model.Feedback:    Feedback,
	model.Instruction: Instruction,

	model.MyOrders: MyOrders,

	model.NewOrderSelectBuilding:   SelectBuilding,
	model.NewOrderInputDescription: InputDescription,
	model.NewOrderConfirm:          OrderConfirmMessageFormatted,

	model.CourierSelectBuilding: SelectBuilding,
	model.CourierActiveOrders:   ActiveOrders,
	model.CourierConfirmOrder:   OrderConfirmMessageFormatted,
}
