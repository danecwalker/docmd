package config

type ConfigExtension string

const (
	JSON ConfigExtension = "json"
	TOML ConfigExtension = "toml"
	YAML ConfigExtension = "yaml"
)

type Config struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Groups      []string `json:"groups"`
	Pages       []Page   `json:"pages"`
	Entry       string   `json:"entry"`
	OutDir      string   `json:"outDir"`
}

type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Path        string `json:"path"`
	Group       string `json:"group"`
}
