package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Credentials TODO: documentation
type Credentials struct {
	Environment string `json:"environment"`
	APIToken    string `json:"apiToken"`
}

// Read TODO: documentation
func (c Credentials) Read(filename string) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteValue, &c)
}
