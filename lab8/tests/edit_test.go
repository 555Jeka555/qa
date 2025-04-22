package tests

import (
	"lab8/api"
	"testing"
)

func TestEditProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToTestConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	// Setup: create a product to edit
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

	// Cleanup
	defer func() {
		if err := client.DeleteProduct(createdProduct.ID); err != nil {
			t.Errorf("Failed to cleanup product: %v", err)
		}
	}()

	// Test valid update
	t.Run("Edit product with valid data", func(t *testing.T) {
		updated := createdProduct
		updated.Title = "Updated Title"
		updated.Price = "200"

		result, err := client.EditProduct(updated)
		if err != nil {
			t.Fatalf("Failed to edit product: %v", err)
		}

		CompareProducts(t, updated, *result)
	})

	// Test invalid update (invalid category)
	t.Run("Edit product with invalid category", func(t *testing.T) {
		updated := createdProduct
		updated.CategoryID = "999"

		_, err := client.EditProduct(updated)
		if err == nil {
			t.Error("Expected error for invalid category, but got none")
		}
	})

	// Test invalid update (empty title)
	t.Run("Edit product with empty title", func(t *testing.T) {
		updated := createdProduct
		updated.Title = ""

		_, err := client.EditProduct(updated)
		if err == nil {
			t.Error("Expected error for empty title, but got none")
		}
	})
}
