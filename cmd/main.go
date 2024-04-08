package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/danecwalker/docmd/internal/build"
	"github.com/danecwalker/docmd/internal/logger"
	"github.com/danecwalker/docmd/internal/serve"
	"github.com/urfave/cli/v2"
)

var (
	// Version is the version of the application
	Version string = "0.0.1"
	// Revision is the git revision of the application
	Revision string = "HEAD"
)

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("version=%s revision=%s\n", cCtx.App.Version, Revision)
	}

	cli.AppHelpTemplate = `{{.Name}} - {{.Usage}}

Usage:

	{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if .Commands}}
Commands:

{{range .Commands}}{{if not .HideHelp}}  {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
Global Options:

	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`

	cli.CommandHelpTemplate = `{{.HelpName}} - {{.Usage}}
	
Usage:

	{{.HelpName}} {{if .VisibleFlags}}[command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if .VisibleFlags}}
Options:

	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}`

	app := &cli.App{
		ExitErrHandler: func(c *cli.Context, err error) {
			if err != nil {
				cli.ShowAppHelp(c)
				logger.PrintStatusLineKV(logger.Red, "[error]", logger.White, "error:", logger.Red, err.Error())
				os.Exit(1)
			}
		},
		Name:                 "docmd",
		Version:              Version,
		Copyright:            fmt.Sprintf("(c) %d %s", time.Now().Year(), "docmd"),
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Dane Walker",
				Email: "dane@danecwalker.com",
			},
		},
		Usage: "make simple docs from your markdown files",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "initialize a new docmd project",
				Action: func(c *cli.Context) error {
					fmt.Println("init called")
					return nil
				},
			},
			{
				Name:      "build",
				Aliases:   []string{"b"},
				Usage:     "build the docs",
				Args:      true,
				ArgsUsage: "./path/to/docmd.config.json",
				Action: func(c *cli.Context) error {
					mdPath := c.Args().First()
					if mdPath == "" {
						return cli.Exit("path to docmd config file is required", 1)
					}
					err := build.Build(mdPath)
					if err != nil {
						return cli.Exit(err, 1)
					}
					return nil
				},
			},
			{
				Name:      "serve",
				Aliases:   []string{"s"},
				Usage:     "serve the docs",
				Args:      true,
				ArgsUsage: "./path/to/docmd.config.json",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   4200,
						Usage:   "port to serve the docs on",
					},
				},
				Action: func(c *cli.Context) error {
					mdPath := c.Args().First()
					if mdPath == "" {
						return cli.Exit("path to docmd config file is required", 1)
					}

					err := serve.Serve(mdPath, c.Int("port"))
					if err != nil {
						return cli.Exit(err, 1)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
