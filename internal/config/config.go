package config

type ConfigExtension string

const (
	JSON ConfigExtension = "json"
	TOML ConfigExtension = "toml"
	YAML ConfigExtension = "yaml"
)

type Config struct {
	Title  string `json:"title"`
	OutDir string `json:"outDir"`
	Entry  string `json:"entry"`
}
