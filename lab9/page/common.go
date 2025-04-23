package page

import (
	"fmt"
	"lab9/config"
	"time"

	"github.com/tebeka/selenium"
)

type Common struct {
	Driver  selenium.WebDriver
	BaseURL string
}

func (c *Common) Init(driver selenium.WebDriver) {
	c.Driver = driver
	c.BaseURL = config.BaseUrl
}

func (c *Common) OpenPage(url string) error {
	return c.Driver.Get(c.BaseURL + url)
}

func (c *Common) WaitForElement(by, value string, timeout time.Duration) (selenium.WebElement, error) {
	var elem selenium.WebElement
	var err error

	startTime := time.Now()
	for time.Since(startTime) < timeout {
		elem, err = c.Driver.FindElement(by, value)
		if err == nil {
			return elem, nil
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil, fmt.Errorf("element not found after %v: %v", timeout, err)
}

func (c *Common) WaitForElements(by, value string, timeout time.Duration) ([]selenium.WebElement, error) {
	var elems []selenium.WebElement
	var err error

	startTime := time.Now()
	for time.Since(startTime) < timeout {
		elems, err = c.Driver.FindElements(by, value)
		if err == nil && len(elems) > 0 {
			return elems, nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil, fmt.Errorf("elements not found after %v: %v", timeout, err)
}

func (c *Common) FindElement(by, value string) (selenium.WebElement, error) {
	return c.WaitForElement(by, value, 10*time.Second)
}

func (c *Common) FindElements(by, value string) ([]selenium.WebElement, error) {
	return c.WaitForElements(by, value, 10*time.Second)
}
