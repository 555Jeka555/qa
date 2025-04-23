package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"lab9/config"
	"lab9/page"
)

func TestSuccessfulAuth(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		t.Fatalf("Failed to open session: %v", err)
	}
	defer driver.Quit()

	cfg := config.GetValidLoginData()

	authPage := page.Auth{}
	authPage.Init(driver)
	if err := authPage.OpenPage(config.LoginUrl); err != nil {
		t.Fatalf("Failed to load page: %v", err)
	}

	if err := authPage.Login(cfg.Login, cfg.Password); err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	isLoginSuccessful, err := authPage.IsLoginSuccessful()
	if err != nil {
		t.Fatalf("Failed to get success message: %v", err)
	}
	assert.True(t, isLoginSuccessful)
}

func TestFailedAuth(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		t.Fatalf("Failed to open session: %v", err)
	}
	defer driver.Quit()

	cfg := config.GetInvalidLoginData()

	authPage := page.Auth{}
	authPage.Init(driver)
	if err := authPage.OpenPage(config.LoginUrl); err != nil {
		t.Fatalf("Failed to load page: %v", err)
	}

	if err := authPage.Login(cfg.Login, cfg.Password); err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	isLoginFailed, err := authPage.IsLoginError()
	if err != nil {
		t.Fatalf("Failed to get success message: %v", err)
	}
	assert.True(t, isLoginFailed)
}
