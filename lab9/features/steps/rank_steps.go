package steps

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"
	"lab9/page"
)

type testContext struct {
	driver      selenium.WebDriver
	indexPage   page.Index
	summaryPage page.Summary
	t           *testing.T
}

func (ctx *testContext) iOpenMainPage() error {
	ctx.indexPage.Init(ctx.driver)
	return ctx.indexPage.OpenPage("/")
}

func (ctx *testContext) iEnterText(text string) error {
	return ctx.indexPage.InputText(text)
}

func (ctx *testContext) iSelectRegion(region string) error {
	countryCodes := map[string]string{
		"ОАЭ":      "AE",
		"Россия":   "RU",
		"Германия": "DE",
		"Франция":  "FR",
		"Индия":    "IN",
	}

	code, ok := countryCodes[region]
	if !ok {
		return fmt.Errorf("страна '%s' не найдена", region)
	}

	return ctx.indexPage.SelectRegion(code)
}

func (ctx *testContext) iSubmitForm() error {
	return ctx.indexPage.SubmitText()
}

func (ctx *testContext) iShouldSeeResults(table *godog.Table) error {
	ctx.summaryPage.Init(ctx.driver)

	time.Sleep(3 * time.Second)

	results := make(map[string]string)
	for _, row := range table.Rows[0:] {
		param := row.Cells[0].Value
		value := row.Cells[1].Value
		results[param] = value
	}

	actualText, err := ctx.summaryPage.GetResultText()
	if err != nil {
		return err
	}
	assert.Equal(ctx.t, results["Текст"], actualText)

	expectedRank, _ := strconv.ParseFloat(results["Ранг"], 64)
	actualRank, err := ctx.summaryPage.GetResultRank()
	if err != nil {
		return err
	}
	assert.InDelta(ctx.t, expectedRank, actualRank, 0.05)

	expectedSim, _ := strconv.Atoi(results["Похожесть"])
	actualSim, err := ctx.summaryPage.GetResultSimilarity()
	if err != nil {
		return err
	}
	assert.Equal(ctx.t, expectedSim, actualSim)

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext, t *testing.T) {
	tc := &testContext{t: t}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		caps := selenium.Capabilities{"browserName": "chrome"}
		driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
		if err != nil {
			t.Fatal(err)
		}
		tc.driver = driver
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if tc.driver != nil {
			tc.driver.Quit()
		}
		return ctx, nil
	})

	ctx.Step(`^я открываю главную страницу$`, tc.iOpenMainPage)
	ctx.Step(`^я ввожу текст "([^"]*)"$`, tc.iEnterText)
	ctx.Step(`^выбираю регион "([^"]*)"$`, tc.iSelectRegion)
	ctx.Step(`^отправляю форму$`, tc.iSubmitForm)
	ctx.Step(`^я должен увидеть результаты:$`, tc.iShouldSeeResults)
}
