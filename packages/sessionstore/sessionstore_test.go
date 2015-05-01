package sessionstore

import (
	"rrb/packages/configmanager"
	"testing"
)

var configManager configmanager.ConfigHash

// Setup
func beforeTests() {
	// |!Important| These test require an available database or the tests will fail.
	var c = configmanager.New("../../conf/environment.json")
	configManager = c.GetConfig()
}

// Taredown
func afterTests() {
	return
}

func TestThatTheSessionCanBeUsedMultipleTimes(t *testing.T) {
	beforeTests()
	var names = ConnectionTestHelper()
	if len(names) < 1 {
		t.Errorf("Expected database to be available")
	}

	names = ConnectionTestHelper()

	if len(names) < 1 {
		t.Errorf("Expected database to be available")
	}

	afterTests()
}

func ConnectionTestHelper() []string {

	var s = GetSession(configManager["DB_CONNECTION_STRING"])
	defer s.Close()
	names, _ := s.DatabaseNames()
	return names

}
