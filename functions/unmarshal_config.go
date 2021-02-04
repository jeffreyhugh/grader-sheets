package functions

import (
	"../structures"
	"encoding/json"
	"io/ioutil"
	"os"
)

func UnmarshalConfig(configJson string) (*structures.Config, error) {
	jsonFile, err := os.Open(configJson)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	UnmarshaledConfig := &structures.Config{}

	if err := json.Unmarshal(byteValue, UnmarshaledConfig); err != nil {
		return nil, err
	}

	_ = jsonFile.Close()

	return UnmarshaledConfig, nil
}