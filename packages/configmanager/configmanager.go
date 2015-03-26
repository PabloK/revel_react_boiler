// Copyright 2015 Pablo Karlssnon. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
  Package Configmanager implements configuration management with prioritized overwrite.

  Description
  Configmanager imports configuration from file and overwrites defaults with data found in the ENV.
  To use this package the env var RRB_CONFIG_FILE needs to be set to point to a JSON file containing
  the configuration to be used. Onyl values in the JSON file will be imported to the manager.

  For an example JSON file see env.json in the rrb/conf folder.
*/
package configmanager

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type conf_manager struct {
	O interface{}
}

type ConfigHash map[string]interface{}

var instantiated *conf_manager = nil
var config ConfigHash

// New instantiates the configmanager and reads the configuration into memory.
func New() *conf_manager {
	if instantiated == nil {
		// Get configuration from JSON
		var configJSON = getJSONConfig()
		// Get configuration from ENV
		var configENV = getENVConfig()
		// prioritize according to source and store it
		config = prioritizeConfig(configJSON, configENV)

		instantiated = new(conf_manager)
	}
	return instantiated
}

// GetConfig returns the configuration as string map
func (c *conf_manager) GetConfig() ConfigHash {
	return config
}

// getJSONConfig, private method that returns the config found in
// the specified JSON file.
func getJSONConfig() ConfigHash {
	var data ConfigHash

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

func getENVConfig() ConfigHash {
	var a ConfigHash
	return a
}

func prioritizeConfig(c ConfigHash, d ConfigHash) ConfigHash {
	var a ConfigHash
	return a
}
