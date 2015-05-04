package configmanager

import (
	"github.com/revel/revel"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// Setup
func beforeTests() {
	os.Setenv("TESTING_ENV_CONFIG", "TESTING")
	discardConfig()
	revel.INFO = log.New(ioutil.Discard, "INFO  ", log.Ldate|log.Ltime|log.Lshortfile)
	revel.WARN = log.New(ioutil.Discard, "WARN  ", log.Ldate|log.Ltime|log.Lshortfile)
	revel.ERROR = log.New(ioutil.Discard, "ERROR  ", log.Ldate|log.Ltime|log.Lshortfile)
	revel.TRACE = log.New(ioutil.Discard, "TRACE  ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Taredown
func afterTests() {
	os.Unsetenv("TESTING_ENV_CONFIG")
}

func TestThatEnvVarsAreSetFromJsonConfig(t *testing.T) {
	beforeTests()
	var c = New("test-conf.json")
	var conf = c.GetConfig()

	if len(conf) < 1 {
		t.Errorf("Unexpected config length.")
	}

	_, keyExists := conf["db"]
	if !keyExists {
		t.Errorf("Expected key: db to exist.")
	}
	afterTests()
}

// Test that if ENV var with same name as JSON var is set it overwrites the JSON var
func TestThatEnvVarsOverwriteJSONConfigt(t *testing.T) {
	beforeTests()
	var c = New("test-conf.json")
	var conf = c.GetConfig()
	if len(conf) < 1 {
		t.Errorf("Unexpected config length.")
	}

	_, keyExists := conf["TESTING_ENV_CONFIG"]
	if !keyExists {
		t.Errorf("Expected key: TESTING_ENV_CONFIG to exist.")
	}

	if conf["TESTING_ENV_CONFIG"] != "TESTING" {
		t.Errorf("Expected key: TESTING_ENV_CONFIG to contain TESTING but contained: %d.", conf["TESTING_ENV_CONFIG"])
	}
	afterTests()
}

// Test that if env vars are unset the JSON config remains
func TestThatJSONKeysRemainWhenNoEnvVarIsSet(t *testing.T) {
	beforeTests()
	os.Unsetenv("TESTING_ENV_CONFIG")

	var c = New("test-conf.json")
	var conf = c.GetConfig()
	if len(conf) < 1 {
		t.Errorf("Unexpected config length.")
	}

	_, keyExists := conf["TESTING_ENV_CONFIG"]
	if !keyExists {
		t.Errorf("Expected key: TESTING_ENV_CONFIG to exist.")
	}

	if conf["TESTING_ENV_CONFIG"] != "INVALID" {
		t.Errorf("Expected key: TESTING_ENV_CONFIG to contain TESTING but contained: %d.", conf["TESTING_ENV_CONFIG"])
	}
	afterTests()
}

// Test that execution is not halted when calling with a non existing file
func TestThatExecutinIsNotHalted(t *testing.T) {
	beforeTests()
	var c = New("non-existing-conf.json")
	var conf = c.GetConfig()
	if len(conf) != 0 {
		t.Errorf("Expected execution to continue.")
	}
	afterTests()
}
