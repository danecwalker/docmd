package build

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/danecwalker/docmd/internal/doc"
)

func Build(docs *doc.Doc, docPath string) error {
	// check if docmd.config.(json|toml|yaml) exists
	matches, err := filepath.Glob(path.Join(docPath, "docmd.config.*"))
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		return fmt.Errorf("docmd config file not found")
	} else if len(matches) > 1 {
		return fmt.Errorf("multiple docmd config files found")
	} else {
		switch ext := filepath.Ext(matches[0]); ext {
		case ".json":
			doc.BuildJSON(docs, matches[0])
		default:
			return fmt.Errorf("unsupported config file type: %s", ext)
		}
	}

	return nil
}
