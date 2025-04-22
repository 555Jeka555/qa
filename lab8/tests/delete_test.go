package tests

import (
	"lab8/api"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToTestConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	// Setup: create a product to delete
	product := config.Products[0] // valid_product
	err = client.AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to setup test: %v", err)
	}

	createdProducts, err := client.GetAllProducts()
	if err != nil {
		t.Fatalf("Failed to add product: %v", err)
	}
	createdProduct := createdProducts[0]

	// Test valid delete
	t.Run("Delete existing product", func(t *testing.T) {
		if err := client.DeleteProduct(createdProduct.ID); err != nil {
			t.Fatalf("Failed to delete product: %v", err)
		}

		// Verify product is deleted
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

	// Test delete non-existing product
	t.Run("Delete non-existing product", func(t *testing.T) {
		if err := client.DeleteProduct("-999999"); err == nil {
			t.Error("Expected error for non-existing product, but got none")
		}
	})
}
