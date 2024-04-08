package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/danecwalker/docmd/internal/config"
	"github.com/danecwalker/docmd/internal/logger"
)

func ServeJSON(configPath string, port int) error {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("GET %s\n", r.URL.Path)
		p := path.Join(c.OutDir, r.URL.Path)

		if r.URL.Path != "/" && path.Ext(r.URL.Path) == "" {
			p += ".html"
		}

		ServeFile(w, r, c, p)

	})

	fmt.Println()
	fmt.Println()
	fmt.Printf("%s%s docmd %s%s %s%s%s\n", logger.BgGreen, logger.Bold, logger.Reset, logger.BgReset, logger.Blue, "v1", logger.BgReset)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}

func ServeFile(w http.ResponseWriter, r *http.Request, c *config.Config, p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		content, err := os.ReadFile(path.Join(c.OutDir, "404.html"))
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(content)
		return
	}

	http.ServeFile(w, r, p)
}
