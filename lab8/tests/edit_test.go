package tests

import (
	"lab8/api"
	"lab8/model"
	"testing"
)

func TestEditProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToTestConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	product := config.Products[0] // valid_product
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

	defer func() {
		if _, err := client.DeleteProduct(createdProduct.ID); err != nil {
			t.Errorf("Failed to cleanup product: %v", err)
		}
	}()

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

	t.Run("Edit product with invalid category", func(t *testing.T) {
		updated := createdProduct
		updated.CategoryID = "999"

		_, err := client.EditProduct(updated)
		if err == nil {
			t.Error("Expected error for invalid category, but got none")
		}
	})

	t.Run("Edit product with empty title", func(t *testing.T) {
		updated := createdProduct
		updated.Title = ""

		_, err := client.EditProduct(updated)
		if err == nil {
			t.Error("Expected error for empty title, but got none")
		}
	})
}
