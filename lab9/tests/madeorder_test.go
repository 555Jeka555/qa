package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/page"
	"testing"
)

func TestMadeOrderLoggedSuccessful(t *testing.T) {
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

	orderPage := page.Order{}
	orderPage.Init(driver)
	err = orderPage.OpenPage(config.ProductURL)
	assert.NoError(t, err)

	err = orderPage.AddToCart()
	assert.NoError(t, err)

	err = orderPage.ClickOrderButton()
	assert.NoError(t, err)

	err = orderPage.FillOrderForm(config.ExistingToOrderData.Note)
	assert.NoError(t, err)

	isSuccess, err := orderPage.IsOrderMadeSuccessful()
	assert.NoError(t, err)
	assert.True(t, isSuccess)
}

func TestMadeOrderSuccessful(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	assert.NoError(t, err)

	defer driver.Quit()

	orderPage := page.Order{}
	orderPage.Init(driver)
	err = orderPage.OpenPage(config.ProductURL)
	assert.NoError(t, err)

	err = orderPage.AddToCart()
	assert.NoError(t, err)

	err = orderPage.ClickOrderButton()
	assert.NoError(t, err)

	err = orderPage.FillFullOrderForm(config.ValidToOrderData)
	assert.NoError(t, err)

	isSuccess, err := orderPage.IsOrderMadeSuccessful()
	assert.NoError(t, err)
	assert.True(t, isSuccess)
}

func TestMadeOrderFailed(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	assert.NoError(t, err)

	defer driver.Quit()

	orderPage := page.Order{}
	orderPage.Init(driver)
	err = orderPage.OpenPage(config.ProductURL)
	assert.NoError(t, err)

	err = orderPage.AddToCart()
	assert.NoError(t, err)

	err = orderPage.ClickOrderButton()
	assert.NoError(t, err)

	err = orderPage.FillFullOrderForm(config.ExistingToOrderData)
	assert.NoError(t, err)

	isFailed, err := orderPage.IsOrderMadeFailed()
	assert.NoError(t, err)
	assert.True(t, isFailed)
}
