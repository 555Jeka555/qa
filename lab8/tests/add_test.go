package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"testing"
)

func TestValidAddProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToValidTestCasesConfig)
	assert.NoError(t, err, "Не удалось загрузить конфигурацию тестов")

	client := api.NewAPIClient(config.BaseURL)

	for testcaseName, product := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			id, err := client.AddProduct(product)
			assert.NoError(t, err, "Ошибка при добавлении продукта")

			products, err := client.GetAllProducts()
			assert.NoError(t, err, "Ошибка при получении списка продуктов")

			createdProduct := FindProductByID(id, products)
			assert.NotNil(t, createdProduct, "Продукт не найден после добавления")

			defer func() {
				_, err := client.DeleteProduct(createdProduct.ID)
				assert.NoError(t, err, "Ошибка при удалении тестового продукта")
			}()

			CompareProducts(t, product, createdProduct)

			assert.NotEqual(t, "0", createdProduct.ID, "ID продукта не должен быть '0'")
			assert.NotEmpty(t, createdProduct.Alias, "Alias продукта должен быть сгенерирован автоматически")
		})
	}
}

func TestInvalidAddProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToInvalidTestCasesConfig)
	assert.NoError(t, err, "Не удалось загрузить конфигурацию тестов")

	client := api.NewAPIClient(config.BaseURL)

	for testcaseName, product := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			id, err := client.AddProduct(product)
			assert.ErrorIs(t, err, api.ErrBadRequest, "Ожидалась ошибка BadRequest для невалидных данных")

			products, err := client.GetAllProducts()
			assert.NoError(t, err, "Ошибка при получении списка продуктов")

			createdProduct := FindProductByID(id, products)

			defer func() {
				if _, err := client.DeleteProduct(createdProduct.ID); err != nil {
					t.Errorf("Failed to cleanup product: %v", err)
				}
			}()
		})
	}
}
