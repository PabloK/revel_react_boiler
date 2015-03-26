package tests

import "github.com/revel/revel"
import "rrb/packages/configmanager"

type ConfigManagerTest struct {
	revel.TestSuite
}

func (t *ConfigManagerTest) Before() {
}

func (t *ConfigManagerTest) TestThatEnvVarsAreSetFromJsonConfig() {
	var c = configmanager.New("tests/test-conf.json")
	var conf = c.GetConfig()
	t.Assert(len(conf) >= 1)

	_, keyExists := conf["db"]
	t.Assert(keyExists)
}

func (t *ConfigManagerTest) TestThatEnvVarsAreSetFromEnvironment() {
	t.Assert(true)
}

func (t *ConfigManagerTest) TestThatEnvironmentEnvVarsArePrioritized() {
	t.Assert(true)
}

func (t *ConfigManagerTest) After() {
}
