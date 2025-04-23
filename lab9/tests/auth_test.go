package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"lab9/config"
	"lab9/page"
)

func TestSuccessfulAuth(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	assert.NoError(t, err)

	defer driver.Quit()

	cfg := config.GetValidLoginData()

	authPage := page.Auth{}
	authPage.Init(driver)
	err = authPage.OpenPage(config.LoginUrl)
	assert.NoError(t, err)

	err = authPage.Login(cfg.Login, cfg.Password)
	assert.NoError(t, err)

	isLoginSuccessful, err := authPage.IsLoginSuccessful()
	assert.NoError(t, err)
	assert.True(t, isLoginSuccessful)
}

func TestFailedAuth(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	assert.NoError(t, err)

	defer driver.Quit()

	cfg := config.GetInvalidLoginData()

	authPage := page.Auth{}
	authPage.Init(driver)
	err = authPage.OpenPage(config.LoginUrl)
	assert.NoError(t, err)

	err = authPage.Login(cfg.Login, cfg.Password)
	assert.NoError(t, err)

	isLoginFailed, err := authPage.IsLoginError()
	assert.NoError(t, err)
	assert.True(t, isLoginFailed)
}
