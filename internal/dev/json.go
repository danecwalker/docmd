package dev

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/danecwalker/docmd/internal/build"
	"github.com/danecwalker/docmd/internal/config"
	"github.com/danecwalker/docmd/internal/logger"
	"github.com/danecwalker/docmd/internal/preview"
	"github.com/fsnotify/fsnotify"
)

func safePanic(err error) {
	os.Setenv("DEV_ERROR", err.Error())
	logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Red, "error:", logger.Reset, err.Error())
	logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Red, "waiting for the error to be resolved...", logger.Reset, "")
}

func DevJSON(configPath string, port int, expose bool) error {
	defer (func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic in Preview:", r)
		}
	})()

	c, err := config.ParseConfigFromJsonFile(configPath)
	if err != nil {
		safePanic(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		safePanic(err)
	}
	defer watcher.Close()

	if err := build.Builder(c, configPath, true); err != nil {
		safePanic(err)
	}

	event_ch := make(chan string)

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					os.Unsetenv("DEV_ERROR")
					logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Reset, "detected changes to", logger.Bold, event.Name)
					logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Reset, "rebuilding...", logger.Reset, "")
					if err := build.Builder(c, configPath, true); err != nil {
						safePanic(err)
					} else {
						// clear the screen
						logger.ClearScreen()

						logger.PrintStatusLineKV(logger.Green, "[dev]", logger.Reset, "rebuild complete", logger.Reset, "")
						event_ch <- "{\"type\": \"reload\"}"
						preview.DisplayInfo(port, expose)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(c.InDir)
	if err != nil {
		return err
	}

	// Add a path.
	glob, err := filepath.Glob(path.Join(c.InDir, "*"))
	if err != nil {
		return err
	}

	for _, dir := range glob {
		stat, err := os.Stat(dir)
		if err != nil {
			return err
		}
		if stat.IsDir() {
			err = watcher.Add(dir)
			if err != nil {
				return err
			}
		}
	}

	if err := preview.Serve(c, port, expose, true, event_ch); err != nil {
		return err
	}

	return nil
}
