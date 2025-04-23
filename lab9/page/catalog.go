package page

import (
	"fmt"
	"github.com/tebeka/selenium"
)

const (
	ProductNameInCatalog     = "//input[@id='typeahead']"
	ProductNameInProductPage = "//h3[contains(text(), '%s')]"
)

type Catalog struct {
	Common
}

func (c *Catalog) SearchProduct(text string) error {
	if err := c.typeInSearchInputSearchProduct(text); err != nil {
		return err
	}

	return c.submitSearchProductWithEnter()
}

func (c *Catalog) FindProduct(productName string) error {
	_, err := c.FindElement(selenium.ByXPATH, fmt.Sprintf(ProductNameInProductPage, productName))
	return err
}

func (c *Catalog) typeInSearchInputSearchProduct(text string) error {
	input, err := c.FindElement(selenium.ByXPATH, ProductNameInCatalog)
	if err != nil {
		return fmt.Errorf("failed to find search input: %v", err)
	}

	if err := input.Clear(); err != nil {
		return fmt.Errorf("failed to clear input: %v", err)
	}

	if err := input.SendKeys(text); err != nil {
		return fmt.Errorf("failed to type text: %v", err)
	}

	return nil
}

func (c *Catalog) submitSearchProductWithEnter() error {
	input, err := c.FindElement(selenium.ByXPATH, ProductNameInCatalog)
	if err != nil {
		return fmt.Errorf("failed to find search input: %v", err)
	}

	if err := input.SendKeys(selenium.EnterKey); err != nil {
		return fmt.Errorf("failed to press Enter: %v", err)
	}

	return nil
}
