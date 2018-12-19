package alchemy

import (
	"encoding/json"
	"io/ioutil"
)

// ParseConfigFile reads and unmarshalls a JSON config file into a passed pointer
func ParseConfigFile(filePath string, marshallType interface{}) error {
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	json.Unmarshal(configFile, marshallType)

	return nil
}
