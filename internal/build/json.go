package build

import (
	"bytes"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/danecwalker/docmd/internal/config"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

type Page struct {
	Title   string
	Content string
}

func BuildJSON(configPath string) {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(c.OutDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatal(err)
		}
	}

	t := template.Must(template.ParseGlob("./internal/templates/*.tmpl"))

	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(),
		),
		goldmark.WithParser(goldmark.DefaultParser()),
		goldmark.WithRenderer(goldmark.DefaultRenderer()),
	)

	inFile, err := os.ReadFile(path.Join(path.Dir(configPath), c.Entry))
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = md.Convert(inFile, &buf)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(c.OutDir + "/index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	err = t.ExecuteTemplate(outFile, "base", Page{
		Title:   c.Title,
		Content: buf.String(),
	})
	if err != nil {
		log.Fatal(err)
	}
}
