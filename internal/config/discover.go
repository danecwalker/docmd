package config

import (
	"fmt"
	"path"
	"path/filepath"
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
		switch ext := filepath.Ext(matches[0]); ext {
		case ".json":
			return matches[0], JSON, nil
		default:
			return "", ConfigExtension(ext), fmt.Errorf("unsupported config file type: %s", ext)
		}
	}
}
