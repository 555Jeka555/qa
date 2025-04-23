package page

import (
	"fmt"
	"github.com/tebeka/selenium"
)

const (
	QuantityInput             = "quantity"
	ClearButton               = "//button[text()='Очистить корзину']"
	AddToCartButton           = "//a[@data-id='%s']"
	PopupProductNameValue     = "//table//td/a[contains(text(), '%s')]"
	PopupProductPriceValue    = "//table//td[contains(text(), '%s')]"
	PopupProductQuantityValue = "//table//td[contains(text(), '%s')]"
)

type Product struct {
	Common
}

func (p *Product) AddToCart(id string) error {
	elem, err := p.FindElement(selenium.ByXPATH, fmt.Sprintf(AddToCartButton, id))
	if err != nil {
		return err
	}

	return elem.Click()
}

func (p *Product) SetProductQuantity(quantity string) error {
	input, err := p.FindElement(selenium.ByName, QuantityInput)
	if err != nil {
		return err
	}

	if err := input.Clear(); err != nil {
		return fmt.Errorf("failed to clear input: %v", err)
	}

	return input.SendKeys(quantity)
}

func (p *Product) IsProductInCart(name, price, quantity string) error {
	_, err := p.FindElement(selenium.ByXPATH, fmt.Sprintf(PopupProductNameValue, name))
	if err != nil {
		return err
	}
	_, err = p.FindElement(selenium.ByXPATH, fmt.Sprintf(PopupProductPriceValue, price))
	if err != nil {
		return err
	}
	_, err = p.FindElement(selenium.ByXPATH, fmt.Sprintf(PopupProductQuantityValue, quantity))
	if err != nil {
		return err
	}

	buttonClear, err2 := p.FindElement(selenium.ByXPATH, ClearButton)
	if err2 != nil {
		return err2
	}
	err2 = buttonClear.Click()
	if err2 != nil {
		return err2
	}

	return nil
}
