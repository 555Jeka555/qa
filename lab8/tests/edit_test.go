package tests

//
//import (
//	"lab8/api"
//	"lab8/model"
//	"testing"
//)
//
//func TestEditProduct(t *testing.T) {
//	config, err := LoadTestConfig(pathToTestConfig)
//	if err != nil {
//		t.Fatalf("Failed to load test config: %v", err)
//	}
//
//	client := api.NewAPIClient(config.BaseURL)
//
//	product := config.Products[0]
//	id, err := client.AddProduct(product)
//	if err != nil {
//		t.Fatalf("Failed to setup test: %v", err)
//	}
//
//	createdProducts, err := client.GetAllProducts()
//	if err != nil {
//		t.Fatalf("Failed to add product: %v", err)
//	}
//	var createdProduct model.Product
//	for _, p := range createdProducts {
//		if p.ID == id {
//			createdProduct = p
//			break
//		}
//	}
//
//	defer func() {
//		if _, err := client.DeleteProduct(createdProduct.ID); err != nil {
//			t.Errorf("Failed to cleanup product: %v", err)
//		}
//	}()
//
//	t.Run("Edit product with valid data", func(t *testing.T) {
//		updated := createdProduct
//		updated.Title = "Updated Title"
//		updated.Price = "200"
//
//		err := client.EditProduct(updated)
//		if err != nil {
//			t.Fatalf("Failed to edit product: %v", err)
//		}
//
//		products, err := client.GetAllProducts()
//		if err != nil {
//			t.Fatalf("Failed to add product: %v", err)
//		}
//		var updatedProduct model.Product
//		for _, p := range products {
//			if p.ID == id {
//				updatedProduct = p
//				break
//			}
//		}
//
//		CompareProducts(t, updated, updatedProduct)
//	})
//
//	t.Run("Edit product with invalid category", func(t *testing.T) {
//		updated := createdProduct
//		updated.CategoryID = "999"
//
//		err := client.EditProduct(updated)
//		if err == nil {
//			t.Error("Expected error for invalid category, but got none")
//		}
//	})
//}
