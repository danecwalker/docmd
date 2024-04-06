package serve

import (
	"fmt"
	"net/http"

	"github.com/danecwalker/docmd/internal/doc"
)

func Serve(docs *doc.Doc, port int) error {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(docs.OutputPath)))

	fmt.Printf("Serving docs on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		return err
	}

	return nil
}
