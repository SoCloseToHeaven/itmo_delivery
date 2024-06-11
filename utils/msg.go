package utils

import "itmo_delivery/model"

const (
	StartMsg                              string = "Добро пожаловать в сервис ITMO.DELIVERY"
	ErrorMsg                              string = "Упс... Что-то пошло не так"
	UnknownAction                         string = "Неизвестное действие!"
	SelectBuilding                        string = "Выберите нужное здание в меню"
	IncorrectBuilding                            = "Неверное название здания! Попробуйте ещё раз"
	SuccessfullySelectedBuildingFormatted        = "Успешно выбрано здание: %s"
	SuccessfullyDescriptionInputFormatted        = "Успешно введено описание: %s"
	YourOrderTakenFormatted                      = "Ваш заказ: %s\nБыл взят курьером: @%s\nСвяжитесь с ним для уточнения информации."
	CourierOrderTakenFormatted                   = "Вы взяли заказ: %s\nЗаказчик: @%s\nСвяжитесь с ним для уточнения информации."
	MyOrders                              string = "Мои последние заказы"
	NoOrders                                     = "Упс... Кажется у вас ещё не было заказов!"
	ActiveOrders                          string = "Активные заказы:"
	OrderConfirmMessageFormatted          string = "Вы подтверждаете заказ? \n%s"
	InputDescription                      string = "Введите описание заказа:"
	IncorrectDescriptionSizeFormatted            = "Некорректный размер сообщения! Допустимо не больше %d символов!"
	Instruction                           string = `✅ Как оставить новый заказ ✅ 
 
* В главном меню выберете пункт "Новый заказ"
* Укажите корпус, в котором вы находитесь
* Укажите информацию о заказе (подробно опишите, что и куда доставить)

✅ Как взять заказ ✅

* В главном меню выберете пункт "Посмотреть заказы (Курьер)"
* Укажите корпус, в котором вы находитесь
* Выберете из списка возможных заказов тот, который готовы выполнить, и отправьте боту его номер
* Свяжитесь с заказчиком, чтобы узнать подробности условий заказа`
	Feedback string = `Мы постоянно стремимся улучшить наши услуги и продукты, поэтому ваш отзыв очень важен для нас — он поможет стать нам еще круче 🫶🏼

Пожалуйста, поделитесь своими мыслями, предложениями или замечаниями о вашем опыте с нами по ссылке 👇

https://forms.gle/vb3G1SvFL66FQwbc9`
	AboutBot string = "Что бы вы хотели узнать?"
	MainMenu        = `Привет!👋

Добро пожаловать в ITMO.DELIVERY - сервис доставки внутри корпусов Университета 

Здесь вы можете добавить новый заказ - от любимого айс-латте с соленой карамелью до набора цветных карандашей из ближайшего магазина канцтоваров

Здесь же вы можете взять один из имеющихся заказов и обрадовать кого-то из коллег доставленной к аудитории булочкой из Вольчека`
	Support                       = "🤥 Для обращения пишите на почту itmo.delivery.help@yandex.ru"
	NewOrderCreated               = "Успешно создан новый заказ! \n%s"
	BackButtonClicked             = "Переходим назад..."
	NoActiveOrders                = "Упс! Кажется в этом здании нет активных заказов!"
	CourierSelectOrderDescription = "Введите ID заказа, который вы хотите выбрать, или вернитесь обратно."
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
