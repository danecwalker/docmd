package serve

import (
	"fmt"

	"github.com/danecwalker/docmd/internal/config"
)

func Serve(dir string, port int) error {
	configPath, ext, err := config.DiscoverConfig(dir)
	if err != nil {
		return err
	}

	switch ext {
	case config.JSON:
		return ServeJSON(configPath, port)
	default:
		return fmt.Errorf("unsupported config file type: %s", ext)
	}
}
