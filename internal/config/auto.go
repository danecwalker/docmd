package config

import (
	"path"
	"strings"
)

func autoFillConfig(c *Config) {
	if c.Title == "" {
		c.Title = "My Docs"
	}

	if c.Description == "" {
		c.Description = "My Docs"
	}

	if c.OutDir == "" {
		c.OutDir = "dist"
	}

	if c.Entry == "" {
		c.Entry = "index.md"
	}

	if c.Errors.NotFound.Title == "" {
		c.Errors.NotFound.Title = "Not Found"
	}

	if c.Errors.NotFound.Description == "" {
		c.Errors.NotFound.Description = "The page you are looking for does not exist."
	}

	if c.Errors.NotFound.Path == "" {
		c.Errors.NotFound.Path = "# 404 - Not Found\n\nThe page you are looking for does not exist."
	}

	if c.Errors.Internal.Title == "" {
		c.Errors.Internal.Title = "Internal Server Error"
	}

	if c.Errors.Internal.Description == "" {
		c.Errors.Internal.Description = "An error occurred while processing your request."
	}

	if c.Errors.Internal.Path == "" {
		c.Errors.Internal.Path = "# 500 - Internal Server Error\n\nAn error occurred while processing your request."
	}

	c.Entry = path.Join(c.InDir, c.Entry)

	if strings.HasSuffix(c.Errors.NotFound.Path, ".md") {
		c.Errors.NotFound.Path = path.Join(c.InDir, c.Errors.NotFound.Path)
	}

	if strings.HasSuffix(c.Errors.Internal.Path, ".md") {
		c.Errors.Internal.Path = path.Join(c.InDir, c.Errors.Internal.Path)
	}

	for i := range c.Pages {
		if c.Pages[i].Path == "" {
			continue
		}

		if c.Pages[i].Url == "" {
			c.Pages[i].Url = strings.Split(c.Pages[i].Path, ".md")[0]
		}

		c.Pages[i].Groups = strings.Split(c.Pages[i].Url, "/")
		c.Pages[i].Groups = c.Pages[i].Groups[1 : len(c.Pages[i].Groups)-1]

		c.Pages[i].Path = path.Join(c.InDir, c.Pages[i].Path)
	}
}
