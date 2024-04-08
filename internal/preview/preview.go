package preview

import (
	"fmt"

	"github.com/danecwalker/docmd/internal/config"
)

func Preview(dir string, port int, expose bool) error {
	configPath, ext, err := config.DiscoverConfig(dir)
	if err != nil {
		return err
	}

	switch ext {
	case config.JSON:
		return PreviewJSON(configPath, port, expose)
	default:
		return fmt.Errorf("unsupported config file type: %s", ext)
	}
}
