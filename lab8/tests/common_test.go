package tests

import (
	"encoding/json"
	"lab8/model"
	"os"
	"testing"
)

const pathToTestConfig = "../config/testdata.json"

type TestConfig struct {
	BaseURL  string          `json:"base_url"`
	Products []model.Product `json:"products"`
}

func LoadTestConfig(path string) (*TestConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config TestConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func CompareProducts(t *testing.T, expected, actual model.Product) {
	if expected.CategoryID != actual.CategoryID {
		t.Errorf("CategoryID mismatch: expected %s, got %s", expected.CategoryID, actual.CategoryID)
	}
	if expected.Title != actual.Title {
		t.Errorf("Title mismatch: expected '%s', got '%s'", expected.Title, actual.Title)
	}
	if expected.Content != actual.Content {
		t.Errorf("Content mismatch: expected '%s', got '%s'", expected.Content, actual.Content)
	}
	if expected.Price != actual.Price {
		t.Errorf("Price mismatch: expected %s, got %s", expected.Price, actual.Price)
	}
	if expected.OldPrice != actual.OldPrice {
		t.Errorf("OldPrice mismatch: expected %s, got %s", expected.OldPrice, actual.OldPrice)
	}
	if expected.Status != actual.Status {
		t.Errorf("Status mismatch: expected %s, got %s", expected.Status, actual.Status)
	}
	if expected.Keywords != actual.Keywords {
		t.Errorf("Keywords mismatch: expected '%s', got '%s'", expected.Keywords, actual.Keywords)
	}
	if expected.Description != actual.Description {
		t.Errorf("Description mismatch: expected '%s', got '%s'", expected.Description, actual.Description)
	}
	if expected.Hit != actual.Hit {
		t.Errorf("Hit mismatch: expected %s, got %s", expected.Hit, actual.Hit)
	}
}
