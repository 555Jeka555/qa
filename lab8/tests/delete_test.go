package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"lab8/api"
)

func TestDeleteProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToValidTestCasesConfig)
	assert.NoError(t, err, "Не удалось загрузить конфиг теста")

	client := api.NewAPIClient(config.BaseURL)

	product := config.TestCases["valid_product_min_category_id"]
	id, err := client.AddProduct(product)
	assert.NoError(t, err, "Не удалось создать продукт для теста")

	createdProducts, err := client.GetAllProducts()
	assert.NoError(t, err, "Не удалось получить список продуктов после добавления")

	createdProduct := FindProductByID(id, createdProducts)
	assert.NotNil(t, createdProduct, "Созданный продукт не найден в списке продуктов")

	t.Run("Delete existing product", func(t *testing.T) {
		status, err := client.DeleteProduct(createdProduct.ID)
		assert.NoError(t, err, "Ошибка при удалении существующего продукта")
		assert.Equal(t, status, 1, "Статус удаления существующего продукта должен быть 1")

		products, err := client.GetAllProducts()
		assert.NoError(t, err, "Не удалось получить список продуктов после удаления")

		found := false
		for _, p := range products {
			if p.ID == createdProduct.ID {
				found = true
				break
			}
		}
		assert.False(t, found, "Продукт всё ещё существует после удаления")
	})

	t.Run("Delete non-existing product", func(t *testing.T) {
		status, err := client.DeleteProduct("-999999")
		assert.NoError(t, err, "Ошибка при попытке удалить несуществующий продукт")
		assert.Equal(t, status, 0, "Статус удаления несуществующего продукта должен быть 0")
	})
}
