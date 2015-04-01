package tests

import (
	"github.com/revel/revel"
	"os"
	"rrb/packages/configmanager"
)

type ConfigManagerTest struct {
	revel.TestSuite
}

// Setup
func (t *ConfigManagerTest) Before() {
	os.Setenv("TESTING_ENV_CONFIG", "TESTING")
	configmanager.ReInitializeConf()
}

func (t *ConfigManagerTest) TestThatEnvVarsAreSetFromJsonConfig() {
	var c = configmanager.New("tests/test-conf.json")
	var conf = c.GetConfig()

	t.Assert(len(conf) >= 1)

	_, keyExists := conf["db"]
	t.Assert(keyExists)
}

// Test that if ENV var with same name as JSON var is set it overwrites the JSON var
func (t *ConfigManagerTest) TestThatEnvVarsOverwriteJSONConfig() {
	var c = configmanager.New("tests/test-conf.json")
	var conf = c.GetConfig()
	t.Assert(len(conf) >= 1)

	_, keyExists := conf["TESTING_ENV_CONFIG"]
	t.Assert(keyExists)

	t.Assert(conf["TESTING_ENV_CONFIG"] == "TESTING")
}

// Test that if env vars are unset the JSON config remains
func (t *ConfigManagerTest) TestThatJSONKeysRemainWhenNoEnvVarIsSet() {
	os.Unsetenv("TESTING_ENV_CONFIG")

	var c = configmanager.New("tests/test-conf.json")
	var conf = c.GetConfig()
	t.Assert(len(conf) >= 1)

	_, keyExists := conf["TESTING_ENV_CONFIG"]
	t.Assert(keyExists)
	t.Assert(conf["TESTING_ENV_CONFIG"] == "INVALID")
}

// Test that execution is not halted when calling with a non existing file
func (t *ConfigManagerTest) TestThatExecutinIsNotHalted() {

	var c = configmanager.New("tests/non-existing-conf.json")
	var conf = c.GetConfig()
	t.Assert(len(conf) == 0)
}

// Taredown
func (t *ConfigManagerTest) After() {
	os.Unsetenv("TESTING_ENV_CONFIG")
}
