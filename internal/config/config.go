package config

import "os"

type ConfigExtension string

const (
	JSON ConfigExtension = "json"
	TOML ConfigExtension = "toml"
	YAML ConfigExtension = "yaml"
)

type Config struct {
	Title       string `json:"title"`
	LogoPath    string `json:"logoPath"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Pages       []Page `json:"pages"`
	Entry       string `json:"entry"`
	Errors      Errors `json:"errors"`
	OutDir      string `json:"outDir"`
	Theme       any    `json:"theme"`

	InDir string
}

type Page struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Path        string   `json:"path"`
	Groups      []string `json:"groups"`
}

type Errors struct {
	NotFound Page `json:"404"`
	Internal Page `json:"500"`
}

func (c *Config) OutputDirExists() bool {
	if _, err := os.Stat(c.OutDir); os.IsNotExist(err) {
		return false
	}

	return true
}
