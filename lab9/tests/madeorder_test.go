package tests

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/page"
	"testing"
)

func TestMadeOrderLoggedSuccessful(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		cfg := config.GetValidLoginData()

		authPage := page.Auth{}
		authPage.Init(driver)
		err := authPage.OpenPage(config.LoginUrl)
		assert.NoError(t, err, "Не удалось открыть страницу авторизации")

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err, "Ошибка при вводе логина/пароля")

		isLoginSuccessful, err := authPage.IsLoginSuccessful()
		assert.NoError(t, err, "Ошибка при проверке авторизации")
		assert.True(t, isLoginSuccessful, "Авторизация не прошла успешно")

		orderPage := page.Order{}
		orderPage.Init(driver)
		err = orderPage.OpenPage(config.ProductURL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = orderPage.AddToCart()
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = orderPage.ClickOrderButton()
		assert.NoError(t, err, "Не удалось перейти к оформлению заказа")

		err = orderPage.FillOrderForm(config.ExistingToOrderData.Note)
		assert.NoError(t, err, "Ошибка при заполнении заметки к заказу")

		isSuccess, err := orderPage.IsOrderMadeSuccessful()
		assert.NoError(t, err, "Ошибка при проверке успешности заказа")
		assert.True(t, isSuccess, "Заказ не был оформлен успешно")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestMadeOrderSuccessful(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		orderPage := page.Order{}
		orderPage.Init(driver)
		err := orderPage.OpenPage(config.ProductURL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = orderPage.AddToCart()
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = orderPage.ClickOrderButton()
		assert.NoError(t, err, "Не удалось перейти к оформлению заказа")

		validData := config.OrderData{
			Login:    gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 8),
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Address:  "Йошкар-Ола, ул. Вознесенская, 110",
			Note:     "note note note",
		}
		err = orderPage.FillFullOrderForm(validData)
		assert.NoError(t, err, "Ошибка при заполнении формы заказа")

		isSuccess, err := orderPage.IsOrderMadeSuccessful()
		assert.NoError(t, err, "Ошибка при проверке успешности заказа")
		assert.True(t, isSuccess, "Заказ не был оформлен успешно")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestMadeOrderFailed(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		orderPage := page.Order{}
		orderPage.Init(driver)
		err := orderPage.OpenPage(config.ProductURL)
		assert.NoError(t, err, "Не удалось открыть страницу товара")

		err = orderPage.AddToCart()
		assert.NoError(t, err, "Ошибка при добавлении товара в корзину")

		err = orderPage.ClickOrderButton()
		assert.NoError(t, err, "Не удалось перейти к оформлению заказа")

		err = orderPage.FillFullOrderForm(config.ExistingToOrderData)
		assert.NoError(t, err, "Ошибка при заполнении формы заказа")

		isFailed, err := orderPage.IsOrderMadeFailed()
		assert.NoError(t, err, "Ошибка при проверке неудачного заказа")
		assert.True(t, isFailed, "Ожидалась ошибка оформления, но заказ прошёл успешно")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}
