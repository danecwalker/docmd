package build

import (
	"fmt"

	"github.com/danecwalker/docmd/internal/config"
)

func Build(dir string) error {
	configPath, ext, err := config.DiscoverConfig(dir)
	if err != nil {
		return err
	}

	switch ext {
	case config.JSON:
		BuildJSON(configPath)
	default:
		return fmt.Errorf("unsupported config file type: %s", ext)
	}

	return nil
}
