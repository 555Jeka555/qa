package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToValidTestCasesConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	product := config.TestCases["valid_product_min_category_id"]
	id, err := client.AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}

	createdProducts, err := client.GetAllProducts()
	if err != nil {
		t.Fatalf("Failed to add product: %v", err)
	}
	createdProduct := FindProductByID(id, createdProducts)

	t.Run("Delete existing product", func(t *testing.T) {
		status, err := client.DeleteProduct(createdProduct.ID)
		if err != nil {
			t.Fatalf("Failed to delete product: %v", err)
		}
		assert.Equal(t, status, 1)

		products, err := client.GetAllProducts()
		if err != nil {
			t.Fatalf("Failed to get products: %v", err)
		}

		for _, p := range products {
			if p.ID == createdProduct.ID {
				t.Error("Product still exists after deletion")
				break
			}
		}
	})

	t.Run("Delete non-existing product", func(t *testing.T) {
		status, err := client.DeleteProduct("-999999")
		if err != nil {
			t.Fatalf("Failed to delete product: %v", err)
		}
		assert.Equal(t, status, 0)
	})
}
