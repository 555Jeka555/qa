package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"testing"
)

func TestValidEditProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToValidTestCasesConfig)
	assert.NoError(t, err, "Не удалось загрузить конфигурацию тестов")

	client := api.NewAPIClient(config.BaseURL)
	setupProduct := config.TestCases["valid_product_min_category_id"]

	id, err := client.AddProduct(setupProduct)
	assert.NoError(t, err, "Ошибка при создании тестового продукта")

	defer func() {
		_, err := client.DeleteProduct(id)
		assert.NoError(t, err, "Ошибка при удалении тестового продукта")
	}()

	for testcaseName, updatedProduct := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			updatedProduct.ID = id

			err = client.EditProduct(updatedProduct)
			assert.NoError(t, err, "Ошибка при редактировании продукта")

			updatedProducts, err := client.GetAllProducts()
			assert.NoError(t, err, "Ошибка при получении списка продуктов")

			resultProduct := FindProductByID(id, updatedProducts)
			assert.NotNil(t, resultProduct, "Продукт не найден после редактирования")

			CompareProducts(t, updatedProduct, resultProduct)
		})
	}
}

func TestInvalidEditProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToInvalidTestCasesConfig)
	assert.NoError(t, err, "Не удалось загрузить конфигурацию невалидных тестов")

	validConfig, err := LoadTestConfig(pathToValidTestCasesConfig)
	assert.NoError(t, err, "Не удалось загрузить валидную конфигурацию")

	client := api.NewAPIClient(config.BaseURL)
	setupProduct := validConfig.TestCases["valid_product_min_category_id"]

	id, err := client.AddProduct(setupProduct)
	assert.NoError(t, err, "Ошибка при создании тестового продукта")

	defer func() {
		_, err := client.DeleteProduct(id)
		assert.NoError(t, err, "Ошибка при удалении тестового продукта")
	}()

	for testcaseName, updatedProduct := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			updatedProduct.ID = id

			err = client.EditProduct(updatedProduct)
			assert.ErrorIs(t, err, api.ErrBadRequest, "Ожидалась ошибка BadRequest")

			products, err := client.GetAllProducts()
			assert.NoError(t, err, "Ошибка при получении списка продуктов")

			createdProduct := FindProductByID(id, products)
			assert.NotNil(t, createdProduct, "Продукт не должен быть удалён при невалидном редактировании")

			CompareProducts(t, setupProduct, createdProduct)
		})
	}
}
