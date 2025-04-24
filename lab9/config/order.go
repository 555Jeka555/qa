package config

const ProductURL = "/product/casio-mrp-700-1avef"

type OrderData struct {
	Login    string
	Password string
	Name     string
	Email    string
	Address  string
	Note     string
}

var ExistingToOrderData = OrderData{
	Login:    "login.test",
	Password: "qwerty123",
	Name:     "Вито Скаллета",
	Email:    "example@example.com",
	Address:  "Йошкар-Ола, ул. Вознесенская, 110",
	Note:     "note note note",
}
