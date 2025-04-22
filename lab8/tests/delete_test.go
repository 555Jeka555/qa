package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"lab8/model"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToTestConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	product := config.Products[0]
	id, err := client.AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}

	createdProducts, err := client.GetAllProducts()
	if err != nil {
		t.Fatalf("Failed to add product: %v", err)
	}
	var createdProduct model.Product
	for _, p := range createdProducts {
		if p.ID == id {
			createdProduct = p
			break
		}
	}

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
