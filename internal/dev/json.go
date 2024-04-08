package dev

import (
	"fmt"
	"log"

	"github.com/danecwalker/docmd/internal/build"
	"github.com/danecwalker/docmd/internal/config"
	"github.com/danecwalker/docmd/internal/logger"
	"github.com/danecwalker/docmd/internal/preview"
	"github.com/fsnotify/fsnotify"
)

func DevJSON(configPath string, port int, expose bool) error {
	defer (func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Preview:", r)
		}
	})()

	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := build.Builder(c, configPath); err != nil {
		log.Fatal(err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Reset, "detected changes to", logger.Bold, event.Name)
					logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Reset, "rebuilding...", logger.Reset, "")
					if err := build.Builder(c, configPath); err != nil {
						log.Fatal(err)
					}
					// clear the screen
					logger.ClearScreen()

					logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Reset, "rebuild complete", logger.Reset, "")
					preview.DisplayInfo(port, expose)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add(c.InDir)
	if err != nil {
		return err
	}

	if err := preview.Serve(c, port, expose); err != nil {
		return err
	}

	return nil
}
