package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lab8/model"
	"net/http"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{BaseURL: baseURL}
}

func (c *APIClient) GetAllProducts() ([]model.Product, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/products", c.BaseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var products []model.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (c *APIClient) AddProduct(product model.Product) (*model.Product, error) {
	body, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/addproduct", c.BaseURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("JSON resp", resp.Body)

	var createdProduct model.Product
	if err := json.NewDecoder(resp.Body).Decode(&createdProduct); err != nil {
		return nil, err
	}

	return &createdProduct, nil
}

func (c *APIClient) EditProduct(product model.Product) (*model.Product, error) {
	body, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/editproduct", c.BaseURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var updatedProduct model.Product
	if err := json.NewDecoder(resp.Body).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func (c *APIClient) DeleteProduct(id int) error {
	resp, err := http.Get(fmt.Sprintf("%s/api/deleteproduct?id=%d", c.BaseURL, id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
