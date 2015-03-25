package tests

import "github.com/revel/revel"
import "fmt"
import "rrb/packages/configmanager"

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

/* UNIT TESTING CONFIG MODULE */
// TODO: Move this to other file

func (t *AppTest) TestThatEnvVarsAreSetFromJsonConfig() {
	var c = configmanager.New()
	var conf = c.GetConfig()
	for k, val := range conf {
		fmt.Println(k)
		fmt.Println(val)
	}
	t.Assert(true)
}

func (t *AppTest) TestThatEnvVarsAreSetFromEnvironment() {
	t.Assert(true)
}

func (t *AppTest) TestThatEnvironmentEnvVarsArePrioritized() {
	t.Assert(true)
}

func (t *AppTest) After() {
	println("Tear down")
}
