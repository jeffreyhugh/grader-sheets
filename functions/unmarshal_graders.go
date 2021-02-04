package functions

import (
	"../structures"
	"encoding/json"
	"io/ioutil"
	"os"
)

func UnmarshalGraders(graderJson string) (*structures.Graders, error) {
	jsonFile, err := os.Open(graderJson)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	UnmarshaledGraders := &structures.Graders{}

	if err := json.Unmarshal(byteValue, UnmarshaledGraders); err != nil {
		return nil, err
	}

	return UnmarshaledGraders, nil
}