package config

import (
	"github.com/brianvoe/gofakeit/v6"
)

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

var ValidToOrderData = OrderData{
	Login:    gofakeit.Username(),
	Password: gofakeit.Password(true, true, true, true, false, 8),
	Name:     gofakeit.Name(),
	Email:    gofakeit.Email(),
	Address:  "Йошкар-Ола, ул. Вознесенская, 110",
	Note:     "note note note",
}
