package features

import (
	"testing"

	"lab9/features/steps"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			steps.InitializeScenario(ctx, t)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"./"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("Тесты не прошли")
	}
}
