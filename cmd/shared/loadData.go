package shared

import (
	"encoding/json"
	"os"
	"strings"
)

type AddURIStruct struct {
	ID     int32  `json:"id"`
	DB     string `json:"db"`
	Uri    string `json:"uri"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

const DataFilePath string = "json/data.json"

func LoadExistingData() ([]AddURIStruct, error) {
	file := []AddURIStruct{}

	raw, err := os.ReadFile(DataFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return file, nil
		}
		return nil, err
	}

	if len(strings.TrimSpace(string(raw))) == 0 {
		return file, nil
	}

	if err := json.Unmarshal(raw, &file); err != nil {
		return nil, err
	}

	return file, nil
}
