package build

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/danecwalker/docmd/internal/config"
)

type P struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Url         string `json:"url"`
	Path        string `json:"path"`
}

type G struct {
	Name  string `json:"name"`
	Pages []P    `json:"pages"`
}

type C struct {
	Page    P
	Content string
}

func BuildJSON(configPath string) error {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		return err
	}

	err = os.Mkdir(c.OutDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	groups := make(map[string]G)

	indexPage := P{
		Title:       c.Title,
		Description: c.Description,
		Group:       "",
		Url:         "/",
		Path:        c.Entry,
	}

	for _, page := range c.Pages {
		p := P{
			Title:       fmt.Sprintf("%s - %s", c.Title, page.Title),
			Description: page.Description,
			Group:       page.Group,
			Url:         page.Url,
			Path:        page.Path,
		}

		if page.Group == "" {
			return fmt.Errorf("page group cannot be empty")
		}

		if _, ok := groups[page.Group]; !ok {
			groups[page.Group] = G{
				Name:  page.Group,
				Pages: []P{p},
			}
		} else {
			groups[page.Group] = G{
				Name:  page.Group,
				Pages: append(groups[page.Group].Pages, p),
			}
		}
	}

	fmt.Println(indexPage)
	for _, group := range groups {
		fmt.Println(group)
	}

	t := template.Must(template.New("base").Parse(indexTemplate))

	f, err := os.Create(path.Join(c.OutDir, "index.html"))
	if err != nil {
		return err
	}

	md, err := ParseMarkdown(indexPage.Path)
	if err != nil {
		return err
	}

	err = t.Execute(f, C{
		Page:    indexPage,
		Content: md,
	})
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
