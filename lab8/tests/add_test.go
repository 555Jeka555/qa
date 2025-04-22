package tests

import (
	"lab8/api"
	"testing"
	"time"
)

func TestAddProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToTestConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	// Test valid product
	t.Run("Add valid product", func(t *testing.T) {
		product := config.Products[0] // valid_product
		createdProduct, err := client.AddProduct(product)
		time.Sleep(time.Second)
		if err != nil {
			t.Fatalf("Failed to add product: %v", err)
		}

		// Cleanup
		defer func() {
			if err := client.DeleteProduct(createdProduct.ID); err != nil {
				t.Errorf("Failed to cleanup product: %v", err)
			}
		}()

		// Verify fields
		CompareProducts(t, product, *createdProduct)

		// Verify auto-generated fields
		if createdProduct.ID == 0 {
			t.Error("Product ID should not be 0")
		}
		if createdProduct.Alias == "" {
			t.Error("Alias should be generated")
		}
	})

	// Test invalid category
	t.Run("Add product with invalid category", func(t *testing.T) {
		product := config.Products[1] // invalid_category
		_, err := client.AddProduct(product)
		if err == nil {
			t.Error("Expected error for invalid category, but got none")
		}
	})

	// Test empty title
	t.Run("Add product with empty title", func(t *testing.T) {
		product := config.Products[2] // empty_title
		_, err := client.AddProduct(product)
		if err == nil {
			t.Error("Expected error for empty title, but got none")
		}
	})
}
