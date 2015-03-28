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
	"bytes"
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type conf_manager struct {
	O interface{}
}

type ConfigHash map[string]interface{}

var instantiated *conf_manager = nil
var config ConfigHash

func ReInitializeConf() {
	instantiated = nil
}

// New instantiates the configmanager and reads the configuration into memory.
func New(arguments ...string) *conf_manager {

	var configFilePath = "conf/environment.json"
	if len(arguments) > 0 {
		configFilePath = arguments[0]
	}

	if instantiated == nil {

		var msg bytes.Buffer
		msg.WriteString("Reading configuration from ")
		msg.WriteString(configFilePath)
		revel.TRACE.Printf("%s", msg.String())
		// Get configuration from JSON file
		var configJSON = getJSONConfig(configFilePath)

		revel.TRACE.Printf("%s", "Reading configuration from ENV")
		// Get configuration from ENV
		var configENV = getENVConfig()

		// prioritize according to source
		config = prioritizeConfig(configENV, configJSON)

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
func getJSONConfig(configFilePath string) ConfigHash {
	var data ConfigHash

	file, err := ioutil.ReadFile(configFilePath)
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
	var config ConfigHash = make(map[string]interface{})
	env := os.Environ()
	for _, str := range env {
		envarr := strings.Split(str, "=")
		k := envarr[0]
		v := envarr[1]
		config[k] = v
	}
	return config
}

func prioritizeConfig(env ConfigHash, json ConfigHash) ConfigHash {
	for k, _ := range json {
		_, keyExists := env[k]
		if keyExists {
			json[k] = env[k]
		}
	}
	return json
}
