package page

import "github.com/tebeka/selenium"

const (
	loginField     = "login"
	passwordField  = "password"
	submitButton   = ".btn-default"
	successMessage = ".alert-success"
	errorMessage   = ".alert-danger"
)

type Auth struct {
	Common
}

func (a *Auth) Login(login, password string) error {
	loginElem, err := a.FindElement(selenium.ByName, loginField)
	if err != nil {
		return err
	}
	if err := loginElem.SendKeys(login); err != nil {
		return err
	}

	passElem, err := a.FindElement(selenium.ByName, passwordField)
	if err != nil {
		return err
	}
	if err := passElem.SendKeys(password); err != nil {
		return err
	}

	submitElem, err := a.FindElement(selenium.ByCSSSelector, submitButton)
	if err != nil {
		return err
	}

	return submitElem.Click()
}

func (a *Auth) IsLoginSuccessful() (bool, error) {
	elem, err := a.FindElement(selenium.ByCSSSelector, successMessage)
	if err != nil {
		return false, err
	}

	return elem.IsDisplayed()
}

func (a *Auth) IsLoginError() (bool, error) {
	elem, err := a.FindElement(selenium.ByCSSSelector, errorMessage)
	if err != nil {
		return false, err
	}

	return elem.IsDisplayed()
}
