package build

import (
	_ "embed"
	"os"
	"path"

	"github.com/danecwalker/docmd/internal/config"
)

//go:embed templates/index.tmpl
var indexTemplate string

//go:embed templates/error.tmpl
var ErrorTemplate string

//go:embed templates/styles.css
var styles string

//go:embed templates/markdown.css
var markdownStyles string

//go:embed templates/copy.js
var copyScript string

func writeStyles(c *config.Config) error {
	f, err := os.Create(path.Join(c.OutDir, "styles.css"))
	if err != nil {
		return err
	}

	_, err = f.WriteString(styles)
	if err != nil {
		return err
	}

	f.Close()

	f, err = os.Create(path.Join(c.OutDir, "markdown.css"))
	if err != nil {
		return err
	}

	_, err = f.WriteString(markdownStyles)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func writeScripts(c *config.Config) error {
	f, err := os.Create(path.Join(c.OutDir, "copy.js"))
	if err != nil {
		return err
	}

	_, err = f.WriteString(copyScript)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
