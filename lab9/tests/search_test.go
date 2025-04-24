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
		assert.NoError(t, err, "Не удалось открыть главную страницу")

		err = catalogPage.SearchProduct(config.ProductNameCasio)
		assert.NoError(t, err, "Ошибка при поиске товара Casio")

		err = catalogPage.FindProduct(config.ProductNameCasio) // "Eldors Post"
		assert.NoError(t, err, "Товар Casio не найден на странице")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestSearchInProductPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.ProductPageUrl)
		assert.NoError(t, err, "Не удалось открыть страницу продукта")

		err = catalogPage.SearchProduct(config.ProductNameRoyal)
		assert.NoError(t, err, "Ошибка при поиске товара Royal")

		err = catalogPage.FindProduct(config.ProductNameRoyal) // "Eldors Post"
		assert.NoError(t, err, "Товар Royal не найден на странице продукта")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestSearchInCategoryPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.CategoryPageUrl)
		assert.NoError(t, err, "Не удалось открыть страницу категории")

		err = catalogPage.SearchProduct(config.ProductNameRoyal)
		assert.NoError(t, err, "Ошибка при поиске товара Royal в категории")

		err = catalogPage.FindProduct(config.ProductNameRoyal) // "Eldors Post"
		assert.NoError(t, err, "Товар Royal не найден в категории")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}

func TestSearchInSearchPage(t *testing.T) {
	testFunc := func(t *testing.T, driver selenium.WebDriver) {
		catalogPage := page.Catalog{}
		catalogPage.Init(driver)
		err := catalogPage.OpenPage(config.SearchPageUrl)
		assert.NoError(t, err, "Не удалось открыть страницу поиска")

		err = catalogPage.SearchProduct(config.ProductNameCitizen)
		assert.NoError(t, err, "Ошибка при поиске товара Citizen")

		err = catalogPage.FindProduct(config.ProductNameCitizen) // "Eldors Post"
		assert.NoError(t, err, "Товар Citizen не найден на странице поиска")
	}

	runTestForBrowser(t, "chrome", testFunc)
	runTestForBrowser(t, "firefox", testFunc)
}
