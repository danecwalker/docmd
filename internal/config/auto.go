package config

func autoFillConfig(c *Config) {
	if c.Title == "" {
		c.Title = "My Docs"
	}

	if c.OutDir == "" {
		c.OutDir = "dist"
	}

	if c.Entry == "" {
		c.Entry = "docs/index.md"
	}
}
