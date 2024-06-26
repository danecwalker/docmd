package build

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/danecwalker/docmd/internal/colors"
	"github.com/danecwalker/docmd/internal/config"
	"github.com/danecwalker/docmd/internal/logger"
	"github.com/danecwalker/docmd/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Set []string

func (s *Set) Contains(v string) bool {
	for _, i := range *s {
		if i == v {
			return true
		}
	}
	return false
}

func (s *Set) Add(v string) {
	if !s.Contains(v) {
		*s = append(*s, v)
	}
}

type P struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Groups      []string `json:"groups"`
	Url         string   `json:"url"`
	Path        string   `json:"path"`
	Hidden      bool     `json:"hidden"`
	Download    bool     `json:"download"`
	C           *config.Config
}

type C struct {
	Theme         string
	Page          P
	SidebarGroups []string
	Sidebar       map[string][]L
	Content       string
	Hmr           bool
}

type L struct {
	Title string
	Url   string
}

func copyFile(src, dst string) error {
	s, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, s, 0644)
	if err != nil {
		return err
	}

	return nil
}

func BuildJSON(configPath string) error {
	var c config.Config
	if err := Builder(&c, configPath, false); err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("To serve the site, run:\n\n\tdocmd serve %s\n", c.InDir)
	return nil
}

func Builder(c *config.Config, configPath string, hmr bool) error {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		return err
	}

	start1 := time.Now()
	if c.OutputDirExists() {
		os.RemoveAll(c.OutDir)
	}

	err = os.Mkdir(c.OutDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	tmp1, _ := filepath.Abs(c.OutDir)
	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, "directory:", logger.Blue, tmp1)
	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, "Collecting build info...", logger.Reset, "")

	if c.LogoPath != "" {
		err = copyFile(path.Join(c.InDir, c.LogoPath), path.Join(c.OutDir, c.LogoPath))
		if err != nil {
			return err
		}
	}

	var theme colors.Theme
	switch v := c.Theme.(type) {
	case string:
		if _, ok := colors.BuiltInThemes[v]; !ok {
			if _, err := os.Stat(path.Join(c.InDir, v)); os.IsNotExist(err) {
				return fmt.Errorf("theme file not found: %s", v)
			}

			content, err := os.ReadFile(path.Join(c.InDir, v))
			if err != nil {
				return err
			}

			err = json.Unmarshal(content, &theme)
			if err != nil {
				return err
			}
			theme = theme.Merge(colors.BuiltInThemes["default"])
			theme.Name = v
		} else {
			theme = colors.BuiltInThemes[v]
		}
	case map[string]interface{}:
		raw, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(raw, &theme)
		if err != nil {
			return err
		}
		theme = theme.Merge(colors.BuiltInThemes["default"])
		theme.Name = "custom"
	default:
		return fmt.Errorf("unsupported theme type: %T", v)
	}

	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, "theme:", logger.Blue, fmt.Sprintf("\"%s\"", theme.Name))

	themeCss, err := theme.ToCSS()
	if err != nil {
		return err
	}

	writeStyles(c)
	writeScripts(c)

	pages := make([]P, 0)

	gs := make(Set, 0)
	gss := make(map[string][]L)

	for _, page := range c.Pages {
		if !page.Hidden {
			var g string
			if len(page.Groups) < 1 {
				g = "General"
			} else {
				g = strings.Join(strings.Split(page.Groups[len(page.Groups)-1], "-"), " ")
			}
			gs.Add(g)
			if _, ok := gss[g]; !ok {
				gss[g] = make([]L, 0)
			}

			gss[g] = append(gss[g], L{
				Title: page.Title,
				Url:   page.Url,
			})
		}

		p := P{
			Title:       fmt.Sprintf("%s - %s", c.Title, page.Title),
			Description: page.Description,
			Url:         page.Url,
			Path:        page.Path,
			Groups:      page.Groups,
			Hidden:      page.Hidden,
			Download:    page.Download,
			C:           c,
		}

		pages = append(pages, p)
	}

	indexPage := P{
		Title:       c.Title,
		Description: c.Description,
		Groups:      []string{},
		Url:         "/",
		Path:        c.Entry,
		C:           c,
	}

	notFound := P{
		Title:       c.Errors.NotFound.Title,
		Description: c.Errors.NotFound.Description,
		Groups:      []string{},
		Url:         "/404",
		Path:        c.Errors.NotFound.Path,
		C:           c,
	}

	internalError := P{
		Title:       c.Errors.Internal.Title,
		Description: c.Errors.Internal.Description,
		Groups:      []string{},
		Url:         "/500",
		Path:        c.Errors.Internal.Path,
		C:           c,
	}

	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Green, fmt.Sprintf("✓ Completed in %s", utils.RoundDuration(time.Since(start1))), logger.Reset, "")
	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, "Building pages...", logger.Reset, "")
	start2 := time.Now()

	t := template.Must(template.New("base").Funcs(template.FuncMap{
		"fullUrl": func(url string) string {
			domain := c.Domain
			if domain[len(domain)-1] == '/' {
				domain = domain[:len(domain)-1]
			}
			return fmt.Sprintf("%s%s", domain, url)
		},
		"capitalize": func(s string) string {
			return cases.Title(language.English).String(s)
		},
	}).Parse(indexTemplate))

	f, err := os.Create(path.Join(c.OutDir, "index.html"))
	if err != nil {
		return err
	}

	md, err := ParseMarkdown(indexPage)
	if err != nil {
		return err
	}

	err = t.Execute(f, C{
		Theme:         themeCss,
		Page:          indexPage,
		SidebarGroups: gs,
		Sidebar:       gss,
		Content:       md,
		Hmr:           hmr,
	})
	if err != nil {
		return err
	}

	f.Close()

	f, err = os.Create(path.Join(c.OutDir, "404.html"))
	if err != nil {
		return err
	}

	if strings.HasSuffix(notFound.Path, ".md") {
		md, err = ParseMarkdown(notFound)
		if err != nil {
			return err
		}
	} else {
		md, err = ParseMarkdownString(notFound.Path)
		if err != nil {
			return err
		}
	}

	err = t.Execute(f, C{
		Theme:         themeCss,
		Page:          notFound,
		SidebarGroups: gs,
		Sidebar:       gss,
		Content:       md,
		Hmr:           hmr,
	})
	if err != nil {
		return err
	}

	f.Close()

	f, err = os.Create(path.Join(c.OutDir, "500.html"))
	if err != nil {
		return err
	}

	if strings.HasSuffix(internalError.Path, ".md") {
		md, err = ParseMarkdown(internalError)
		if err != nil {
			return err
		}
	} else {
		md, err = ParseMarkdownString(internalError.Path)
		if err != nil {
			return err
		}
	}

	err = t.Execute(f, C{
		Theme:         themeCss,
		Page:          internalError,
		SidebarGroups: gs,
		Sidebar:       gss,
		Content:       md,
		Hmr:           hmr,
	})
	if err != nil {
		return err
	}

	f.Close()

	for _, page := range pages {
		if strings.HasSuffix(page.Path, ".md") && !page.Download {

			// make group directories
			groups := path.Join(page.Groups...)
			err = os.MkdirAll(path.Join(c.OutDir, groups), 0755)
			if err != nil {
				if !os.IsExist(err) {
					return err
				}
			}

			f, err := os.Create(path.Join(c.OutDir, groups, fmt.Sprintf("%s.html", strings.Split(path.Base(page.Url), ".md")[0])))
			if err != nil {
				return err
			}

			md, err := ParseMarkdown(page)
			if err != nil {
				return err
			}

			err = t.Execute(f, C{
				Theme:         themeCss,
				Page:          page,
				SidebarGroups: gs,
				Sidebar:       gss,
				Content:       md,
				Hmr:           hmr,
			})

			if err != nil {
				return err
			}

			f.Close()
		} else if page.Download {
			r, err := os.ReadFile(page.Path)
			if err != nil {
				return err
			}

			// make group directories
			groups := path.Join(page.Groups...)
			err = os.MkdirAll(path.Join(c.OutDir, groups), 0755)
			if err != nil {
				if !os.IsExist(err) {
					return err
				}
			}

			f, err := os.Create(path.Join(c.OutDir, groups, path.Base(page.Url)))
			if err != nil {
				return err
			}

			f.Write(r)

			f.Close()
		} else {
			logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Red, fmt.Sprintf("✗ Unsupported file type: %s", page.Path), logger.Reset, "")
		}
	}

	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, fmt.Sprintf("%d page(s) built in %s%s%s", len(pages)+3, logger.Bold, utils.RoundDuration(time.Since(start2)), logger.Reset), logger.Reset, "")
	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Green, fmt.Sprintf("✓ Total time: %s", utils.RoundDuration(time.Since(start1))), logger.Reset, "")
	logger.PrintStatusLineKV(logger.Blue, "[build]", logger.Reset, fmt.Sprintf("%sComplete!%s", logger.Bold, logger.Reset), logger.Reset, "")
	return nil
}
