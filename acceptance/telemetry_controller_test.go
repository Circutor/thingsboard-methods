package acceptance_test

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

//nolint:gochecknoglobals,exhaustivestruct
var (
	godogs int
	opts   = godog.Options{Output: colors.Colored(os.Stdout)}
)

//nolint:gochecknoinits
func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		godogs = 0
	})
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	os.Exit(status)
}

func iOpenAThingsBoardSession() error {
	return nil
}

func iInitControllerTelemetry() error {
	return nil
}

func iWithTheEntityAndEntityID(arg1, arg2 string) error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I open a ThingsBoard session$`, iOpenAThingsBoardSession)
	ctx.Step(`^I init controller telemetry$`, iInitControllerTelemetry)
	ctx.Step(`^I wish the entity "([^"]*)" and entity ID "([^"]*)"$`, iWithTheEntityAndEntityID)
}
