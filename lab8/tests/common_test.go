package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"lab8/model"
	"os"
	"testing"
)

const (
	pathToValidTestCasesConfig   = "../config/validdata.json"
	pathToInvalidTestCasesConfig = "../config/invaliddata.json"
)

type TestConfig struct {
	BaseURL   string                   `json:"base_url"`
	TestCases map[string]model.Product `json:"test_cases"`
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
	fields := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"CategoryID", expected.CategoryID, actual.CategoryID},
		{"Title", expected.Title, actual.Title},
		{"Content", expected.Content, actual.Content},
		{"Price", handleBigNumber(expected.Price), actual.Price},
		{"OldPrice", handleBigNumber(expected.OldPrice), actual.OldPrice},
		{"Status", expected.Status, actual.Status},
		{"Keywords", expected.Keywords, actual.Keywords},
		{"Description", expected.Description, actual.Description},
		{"Hit", expected.Hit, actual.Hit},
	}

	for _, field := range fields {
		assert.Equal(t, field.expected, field.actual,
			"%s mismatch", field.name)
	}
}

func FindProductByID(id string, createdProducts []model.Product) model.Product {
	for _, p := range createdProducts {
		if p.ID == id {
			return p
		}
	}

	return model.Product{}
}

func handleBigNumber(numStr string) string {
	if numStr == "99999999999999999999999999999999999999" {
		return "1e38"
	}
	if numStr == "-99999999999999999999999999999999999999" {
		return "-1e38"
	}

	return numStr
}
