package config

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/danecwalker/docmd/internal/logger"
)

func DiscoverConfig(dir string) (string, ConfigExtension, error) {
	// check if docmd.config.(json|toml|yaml) exists
	matches, err := filepath.Glob(path.Join(dir, "docmd.config.*"))
	if err != nil {
		return "", "", err
	}

	if len(matches) == 0 {
		return "", "", fmt.Errorf("docmd config file not found")
	} else if len(matches) > 1 {
		return "", "", fmt.Errorf("multiple docmd config files found")
	} else {
		logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, "found:", logger.Green, "\"docmd.config.json\"")
		switch ext := filepath.Ext(matches[0]); ext {
		case ".json":
			return matches[0], JSON, nil
		default:
			return "", ConfigExtension(ext), fmt.Errorf("unsupported config file type: %s", ext)
		}
	}
}
