package preview

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/danecwalker/docmd/internal/build"
	"github.com/danecwalker/docmd/internal/config"
	"github.com/danecwalker/docmd/internal/logger"
	"github.com/danecwalker/docmd/internal/meta"
)

func PreviewJSON(configPath string, port int, expose bool) error {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	return Serve(c, port, expose, false, nil)
}

func Serve(c *config.Config, port int, expose bool, hmr bool, events chan string) error {
	mux := http.NewServeMux()

	if hmr {
		mux.HandleFunc("/__hmr", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			logger.PrintStatusLineKV(logger.Yellow, "[🗲 hmr]", logger.Reset, "connected", logger.Reset, "")
			for {
				select {
				case event := <-events:
					fmt.Println("event", event)
					fmt.Fprintf(w, "data: %s\n\n", event)
					// Flush the response. This is important for SSE to work correctly
					if flusher, ok := w.(http.Flusher); ok {
						flusher.Flush()
					} else {
						fmt.Println("Error: Streaming unsupported!")
						return
					}
				case <-r.Context().Done():
					logger.PrintStatusLineKV(logger.Yellow, "[✗ hmr]", logger.Reset, "disconnected", logger.Reset, "")
					return
				}
			}
		})
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		methodColor := logger.Green
		switch r.Method {
		case "GET":
			methodColor = logger.Green
		case "POST":
			methodColor = logger.Yellow
		case "PUT":
			methodColor = logger.Blue
		case "DELETE":
			methodColor = logger.Red
		}
		logger.PrintStatusLineKV(logger.Pink, "[preview]", methodColor, r.Method, logger.Bold, r.URL.Path)
		p := path.Join(c.OutDir, r.URL.Path)

		if r.URL.Path != "/" && path.Ext(r.URL.Path) == "" {
			p += ".html"
		}

		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

		if os.Getenv("DEV_ERROR") != "" {
			w.WriteHeader(http.StatusInternalServerError)
			t := template.Must(template.New("error").Parse(build.ErrorTemplate))
			var buf bytes.Buffer
			t.Execute(&buf, struct {
				Error string
			}{
				Error: os.Getenv("DEV_ERROR"),
			})

			w.Write(buf.Bytes())
		} else {
			ServeFile(w, r, c, p)
		}
	})

	ip, err := DisplayInfo(port, expose)
	if err != nil {
		return err
	}

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), mux); err != nil {
		return err
	}

	return nil
}

func DisplayInfo(port int, expose bool) (ip string, err error) {
	// get ip address
	ip = "localhost"

	if expose {
		netInterfaceAddresses, err := net.InterfaceAddrs()
		if err != nil {
			return "", err
		}

		for _, addr := range netInterfaceAddresses {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
				}
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("%s%s docmd %s%s %s%s%s\n", logger.BgBlue, logger.Bold, logger.Reset, logger.BgReset, logger.Blue, meta.Version, logger.Reset)
	fmt.Println()
	fmt.Printf("%s %s %sLocal:%s http://%s:%s%d%s/\n", logger.BgGrey, logger.BgReset, logger.Bold, logger.Reset, ip, logger.Bold, port, logger.Reset)
	fmt.Println()
	if expose {
		ip = ""
	}

	return ip, nil
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
