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
		assert.NoError(t, err)

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err)

		isLoginSuccessful, err := authPage.IsLoginSuccessful()
		assert.NoError(t, err)
		assert.True(t, isLoginSuccessful)
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
		assert.NoError(t, err)

		err = authPage.Login(cfg.Login, cfg.Password)
		assert.NoError(t, err)

		isLoginFailed, err := authPage.IsLoginError()
		assert.NoError(t, err)
		assert.True(t, isLoginFailed)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}
