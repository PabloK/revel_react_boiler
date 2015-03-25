package configmanager

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type conf_manager struct {
	O interface{}
}

var instantiated *conf_manager = nil
var config map[string]interface{}

func New() *conf_manager {
	if instantiated == nil {
		// Get configuration from JSON
		config = get_json_config()
		// Get configuration from environment
		// Overwrite JSON conf
		// Store config
		instantiated = new(conf_manager)
	}
	return instantiated
}

func (c *conf_manager) Get_Config() map[string]interface{} {
	return config
}

/* Private methods */
func get_json_config() map[string]interface{} {
	var data map[string]interface{}

	// TODO: Get config file location from env
	file, err := ioutil.ReadFile("conf/environment.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
