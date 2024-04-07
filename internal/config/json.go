package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

func ParseConfigFromJson(inDir string, content []byte) (*Config, error) {
	var c Config
	err := json.Unmarshal(content, &c)
	if err != nil {
		return nil, err
	}

	c.InDir = path.Dir(inDir)
	autoFillConfig(&c)

	b, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(b))

	return &c, nil
}

func ParseConfigFromJsonFile(inDir string) (*Config, error) {
	f, err := os.Open(inDir)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// read the file
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return ParseConfigFromJson(inDir, content)
}
