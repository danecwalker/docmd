package preview

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/danecwalker/docmd/internal/config"
	"github.com/danecwalker/docmd/internal/logger"
	"github.com/danecwalker/docmd/internal/meta"
)

func PreviewJSON(configPath string, port int, expose bool) error {
	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

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

		ServeFile(w, r, c, p)

	})

	// get ip address
	ip := "localhost"

	if expose {
		netInterfaceAddresses, err := net.InterfaceAddrs()
		if err != nil {
			log.Fatal(err)
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
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), mux); err != nil {
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
