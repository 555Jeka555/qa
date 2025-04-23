package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/page"
	"testing"
)

func TestAddToCartInMainPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		productCasio := config.ProductCasio
		productPage := page.Product{}
		productPage.Init(driver)
		err := productPage.OpenPage("")
		assert.NoError(t, err)

		err = productPage.AddToCart(productCasio.ID)
		assert.NoError(t, err)

		err = productPage.IsProductInCart(productCasio.Name, productCasio.Price, config.QuantityProductsOne)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestAddOneToCartInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		productCasio := config.ProductCasio
		productPage := page.Product{}
		productPage.Init(driver)
		err := productPage.OpenPage(productCasio.URL)
		assert.NoError(t, err)

		err = productPage.AddToCart(productCasio.ID)
		assert.NoError(t, err)

		err = productPage.IsProductInCart(productCasio.Name, productCasio.Price, config.QuantityProductsOne)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestAddSeveralToCartInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		productCasio := config.ProductCasio
		productPage := page.Product{}
		productPage.Init(driver)
		err := productPage.OpenPage(productCasio.URL)
		assert.NoError(t, err)

		err = productPage.SetProductQuantity(config.QuantityProductsTen)
		assert.NoError(t, err)

		err = productPage.AddToCart(productCasio.ID)
		assert.NoError(t, err)

		err = productPage.IsProductInCart(productCasio.Name, productCasio.Price, config.QuantityProductsTen)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}
