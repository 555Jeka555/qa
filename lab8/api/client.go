package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lab8/model"
	"net/http"
	"strconv"
	"strings"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{BaseURL: baseURL}
}

type productAddedResponse struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type productDeletedResponse struct {
	Status int `json:"status"`
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	bodyString := string(bodyBytes)

	var products []model.Product
	if err := json.Unmarshal(bodyBytes, &products); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v, raw response: %s", err, bodyString)
	}

	return products, nil
}

func (c *APIClient) AddProduct(product model.Product) (string, error) {
	body, err := json.Marshal(product)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/addproduct", c.BaseURL),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	bodyString := string(bodyBytes)

	err = c.checkOnError(bodyString)
	if err != nil {
		return "", err
	}

	var productAddedResp productAddedResponse
	if err := json.Unmarshal(bodyBytes, &productAddedResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %v, raw response: %s", err, bodyString)
	}

	return strconv.Itoa(productAddedResp.ID), nil
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

func (c *APIClient) DeleteProduct(id string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/deleteproduct?id=%s", c.BaseURL, id))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}
	bodyString := string(bodyBytes)

	err = c.checkOnError(bodyString)
	if err != nil {
		return 0, err
	}

	var productDeletedResp productDeletedResponse
	if err := json.Unmarshal(bodyBytes, &productDeletedResp); err != nil {
		return 0, fmt.Errorf("failed to unmarshal JSON: %v, raw response: %s", err, bodyString)
	}

	return productDeletedResp.Status, nil
}

func (c *APIClient) checkOnError(bodyString string) error {
	if strings.Contains(bodyString, "<h1>Произошла ошибка</h1>") {
		return fmt.Errorf("error: %s", bodyString)
	}

	return nil
}
