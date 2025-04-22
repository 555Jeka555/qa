package tests

import (
	"lab8/api"
	"lab8/model"
	"testing"
)

func TestAddProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToTestConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	t.Run("Add valid product", func(t *testing.T) {
		product := config.Products[0] // valid_product
		id, err := client.AddProduct(product)
		if err != nil {
			t.Fatalf("Failed to add product: %v", err)
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

		CompareProducts(t, product, createdProduct)

		if createdProduct.ID == "0" {
			t.Error("Product ID should not be 0")
		}
		if createdProduct.Alias == "" {
			t.Error("Alias should be generated")
		}
	})

	t.Run("Add product with invalid category", func(t *testing.T) {
		product := config.Products[1]
		id, err := client.AddProduct(product)
		if err == nil {
			t.Fatalf("Failed to add product: %v", err)
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
	})

	t.Run("Add valid product with empty title", func(t *testing.T) {
		product := config.Products[2]
		id, err := client.AddProduct(product)
		if err == nil {
			t.Fatalf("Failed to add product: %v", err)
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
	})
}
