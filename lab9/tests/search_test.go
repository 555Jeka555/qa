package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/config"
	"lab9/page"
	"testing"
)

func TestSearchInMainPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage("")
		assert.NoError(t, err)

		err = catalogPage.SearchProduct(config.ProductNameCasio)
		assert.NoError(t, err)

		err = catalogPage.FindProduct(config.ProductNameCasio) // "Eldors Post"
		//url, _ := catalogPage.Driver.CurrentURL()
		//fmt.Println("page", url)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestSearchInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.ProductPageUrl)
		assert.NoError(t, err)

		err = catalogPage.SearchProduct(config.ProductNameRoyal)
		assert.NoError(t, err)

		err = catalogPage.FindProduct(config.ProductNameRoyal) // "Eldors Post"
		//url, _ := catalogPage.Driver.CurrentURL()
		//fmt.Println("page", url)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestSearchInCategoryPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.CategoryPageUrl)
		assert.NoError(t, err)

		err = catalogPage.SearchProduct(config.ProductNameRoyal)
		assert.NoError(t, err)

		err = catalogPage.FindProduct(config.ProductNameRoyal) // "Eldors Post"
		//url, _ := catalogPage.Driver.CurrentURL()
		//fmt.Println("page", url)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestSearchInSearchPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.SearchPageUrl)
		assert.NoError(t, err)

		err = catalogPage.SearchProduct(config.ProductNameCitizen)
		assert.NoError(t, err)

		err = catalogPage.FindProduct(config.ProductNameCitizen) // "Eldors Post"
		//url, _ := catalogPage.Driver.CurrentURL()
		//fmt.Println("page", url)
		assert.NoError(t, err)
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}
