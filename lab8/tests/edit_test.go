package tests

import (
	"github.com/stretchr/testify/assert"
	"lab8/api"
	"testing"
)

func TestValidEditProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToValidTestCasesConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)
	setupProduct := config.TestCases["valid_product_min_category_id"]

	id, err := client.AddProduct(setupProduct)
	if err != nil {
		t.Fatalf("Failed to setup test product: %v", err)
	}

	defer func() {
		if _, err := client.DeleteProduct(id); err != nil {
			t.Errorf("Failed to cleanup test product: %v", err)
		}
	}()

	for testcaseName, updatedProduct := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			updatedProduct.ID = id

			err = client.EditProduct(updatedProduct)
			if err != nil {
				t.Fatalf("Failed to edit product: %v", err)
			}

			updatedProducts, err := client.GetAllProducts()
			if err != nil {
				t.Fatalf("Failed to get updated products: %v", err)
			}
			resultProduct := FindProductByID(id, updatedProducts)

			CompareProducts(t, updatedProduct, resultProduct)
		})
	}
}

func TestInvalidEditProduct(t *testing.T) {
	config, err := LoadTestConfig(pathToInvalidTestCasesConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}
	validConfig, err := LoadTestConfig(pathToValidTestCasesConfig)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	client := api.NewAPIClient(config.BaseURL)
	setupProduct := validConfig.TestCases["valid_product_min_category_id"]

	id, err := client.AddProduct(setupProduct)
	if err != nil {
		t.Fatalf("Failed to setup test product: %v", err)
	}

	defer func() {
		if _, err := client.DeleteProduct(id); err != nil {
			t.Errorf("Failed to cleanup test product: %v", err)
		}
	}()

	for testcaseName, updatedProduct := range config.TestCases {
		t.Run(testcaseName, func(t *testing.T) {
			updatedProduct.ID = id

			err = client.EditProduct(updatedProduct)
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
