package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danecwalker/docmd/internal/config"
)

func ServeJSON(configPath string, port int) error {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(c.OutDir)))

	fmt.Printf("Serving docs on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
