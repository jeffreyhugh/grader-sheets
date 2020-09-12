package functions

import (
	"../structures"
	"encoding/json"
	"io/ioutil"
	"os"
)

const defaultGraderJson = "./graders.json"

func UnmarshalGraders(graderJson string) (*structures.Graders, error) {
	if graderJson == "" {
		graderJson = defaultGraderJson
	}

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