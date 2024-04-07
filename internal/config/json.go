package config

import (
	"encoding/json"
	"io"
	"os"
)

func ParseConfigFromJson(content []byte) (*Config, error) {
	var c Config
	err := json.Unmarshal(content, &c)
	if err != nil {
		return nil, err
	}

	autoFillConfig(&c)

	return &c, nil
}

func ParseConfigFromJsonFile(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// read the file
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return ParseConfigFromJson(content)
}
