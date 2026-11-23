package specs

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/fullness-MFurukawa/bdd-ai-coupon-go/specs/steps"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: steps.InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			Strict:   true,
			NoColors: true,
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("シナリオが成功していません")
	}
}
