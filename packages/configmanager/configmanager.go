// Copyright 2015 Pablo Karlssnon. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file at https://github.com/PabloK/revel_react_boiler.

/*
  Package configmanager implements configuration management with prioritized overwrite from the environment.

  When instantiated configmanager imports configuration from a JSON config file and the environment and makes it
  available trough its interface. To use this package the environmnet variable RRB_CONFIG_FILE needs to point to
  a JSON config file.

  Override of JSON configuration will takeplace at instantiation if the corresponding variable is set in the environment.

  Variables not present in the JSON file will not be imported from the environment.

  For an example JSON file see env.json found in the rrb/conf folder at https://github.com/PabloK/revel_react_boiler.
*/
package configmanager

import (
	"bytes"
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"os"
	"strings"
)

type conf_manager struct {
	O interface{}
}

type ConfigHash map[string]string
type confighash map[string]interface{}

var instantiated *conf_manager = nil
var config ConfigHash

// New instantiates the configmanager and reads the configuration into memory.
func New(arguments ...string) *conf_manager {

	config = make(map[string]string)
	var configFilePath = "conf/environment.json"
	if len(arguments) > 0 {
		configFilePath = arguments[0]
	}

	if instantiated == nil {

		var msg bytes.Buffer
		msg.WriteString("Reading configuration from ")
		msg.WriteString(configFilePath)
		revel.INFO.Printf("%s", msg.String())
		// Get configuration from JSON file
		var configJSON = getJSONConfig(configFilePath)

		revel.TRACE.Printf("%s", "Reading configuration from ENV")
		// Get configuration from ENV
		var configENV = getENVConfig()

		// prioritize according to source
		var prioConf = prioritizeConfig(configENV, configJSON)

		for k, v := range prioConf {
			config[k] = v.(string)
		}

		instantiated = new(conf_manager)
	}
	return instantiated
}

// GetConfig returns the configuration as string map
func (c *conf_manager) GetConfig() ConfigHash {
	return config
}

// discardConfig throws away the current configuration
func discardConfig() {
	instantiated = nil
}

// getJSONConfig returns the config found in the specified JSON file.
func getJSONConfig(configFilePath string) confighash {
	var data confighash = make(map[string]interface{})

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		revel.ERROR.Printf("%s", err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		revel.ERROR.Printf("%s", err)
	}
	return data
}

// getENVConfig returns variables set in the environment.
func getENVConfig() confighash {
	var envconfig confighash = make(map[string]interface{})
	env := os.Environ()
	for _, str := range env {
		envarr := strings.Split(str, "=")
		k := envarr[0]
		v := envarr[1]
		envconfig[k] = v
	}
	return envconfig
}

// Overwrites JSON config with config values from the environment if corresponding variable exists in the environment config.
func prioritizeConfig(env confighash, json confighash) confighash {
	for k, _ := range json {
		_, keyExists := env[k]
		if keyExists {
			json[k] = env[k]
		}
	}
	return json
}
