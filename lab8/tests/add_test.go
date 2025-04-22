package tests

import (
	"lab8/api"
	"testing"
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
		err := client.AddProduct(product)
		if err != nil {
			t.Fatalf("Failed to add product: %v", err)
		}

		createdProducts, err := client.GetAllProducts()
		if err != nil {
			t.Fatalf("Failed to add product: %v", err)
		}
		createdProduct := createdProducts[len(createdProducts)-1]

		// Cleanup
		defer func() {
			if err := client.DeleteProduct(createdProduct.ID); err != nil {
				t.Errorf("Failed to cleanup product: %v", err)
			}
		}()

		// Verify fields
		CompareProducts(t, product, createdProduct)

		// Verify auto-generated fields
		if createdProduct.ID == "0" {
			t.Error("Product ID should not be 0")
		}
		if createdProduct.Alias == "" {
			t.Error("Alias should be generated")
		}
	})

	// Test invalid category
	t.Run("Add product with invalid category", func(t *testing.T) {
		product := config.Products[1] // invalid_category
		err := client.AddProduct(product)
		if err == nil {
			t.Error("Expected error for invalid category, but got none")
		}

		createdProducts, err := client.GetAllProducts()
		if err != nil {
			t.Fatalf("Failed to add product: %v", err)
		}
		createdProduct := createdProducts[len(createdProducts)-1]

		defer func() {
			if err := client.DeleteProduct(createdProduct.ID); err != nil {
				t.Errorf("Failed to cleanup product: %v", err)
			}
		}()
	})
}
