package doc

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"text/template"
)

func BuildJSON(doc *Doc, configPath string) {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read the file
	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var c Config
	err = json.Unmarshal(content, &c)

	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseGlob("./internal/templates/*.tmpl"))

	err = os.Mkdir(doc.OutputPath, 0755)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatal(err)
		}
	}

	f1, err := os.Create(doc.OutputPath + "/index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f1.Close()

	err = t.ExecuteTemplate(f1, "base", c)
	if err != nil {
		log.Fatal(err)
	}
}
