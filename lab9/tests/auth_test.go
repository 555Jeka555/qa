package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"lab9/config"
	"lab9/page"
)

func TestSuccessfulAuth(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		cfg := config.GetValidLoginData()

		authPage := page.Auth{}
		authPage.Init(driver)
		err := authPage.OpenPage(config.LoginUrl)
		assert.NoError(t, err, "Не удалось открыть страницу авторизации")

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err, "Ошибка при вводе валидных учетных данных")

		isLoginSuccessful, err := authPage.IsLoginSuccessful()
		assert.NoError(t, err, "Ошибка при проверке успешной авторизации")
		assert.True(t, isLoginSuccessful, "Ожидалась успешная авторизация, но она не прошла")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestFailedAuth(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		cfg := config.GetInvalidLoginData()

		authPage := page.Auth{}
		authPage.Init(driver)
		err := authPage.OpenPage(config.LoginUrl)
		assert.NoError(t, err, "Не удалось открыть страницу авторизации")

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err, "Ошибка при вводе невалидных учетных данных")

		isLoginFailed, err := authPage.IsLoginError()
		assert.NoError(t, err, "Ошибка при проверке сообщения об ошибке")
		assert.True(t, isLoginFailed, "Ожидалась ошибка авторизации, но она не появилась")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}
