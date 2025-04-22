package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"testing"
)

func TestValidAddProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToValidTestCasesConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	for testcaseName, product := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			id, err := client.AddProduct(product)
			if err != nil {
				t.Fatalf("Failed to add product: %v", err)
			}

			products, err := client.GetAllProducts()
			if err != nil {
				t.Fatalf("Failed to add product: %v", err)
			}
			createdProduct := FindProductByID(id, products)

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
	}
}

func TestInvalidAddProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToInvalidTestCasesConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)

	for testcaseName, product := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			id, err := client.AddProduct(product)
			assert.ErrorIs(t, err, api.ErrBadRequest)

			products, err := client.GetAllProducts()
			if err != nil {
				t.Fatalf("Failed to add product: %v", err)
			}
			createdProduct := FindProductByID(id, products)

			defer func() {
				if _, err := client.DeleteProduct(createdProduct.ID); err != nil {
					t.Errorf("Failed to cleanup product: %v", err)
				}
			}()
		})
	}
}
