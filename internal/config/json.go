package config

import (
	"encoding/json"
	"fmt"
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

	b, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(b))

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
