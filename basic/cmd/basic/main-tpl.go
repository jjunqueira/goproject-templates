package main

import (
	"flag"
	"fmt"
	"os"
	"{{ .ModuleName }}/pkg/app"
	"{{ .ModuleName }}/pkg/log"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var configPath = flag.String("config", "", "The path to the config file")

	flag.Parse()

	app, err := app.Init(*configPath)
	if err != nil {
		return err
	}

	log.Info("Application bootstrap completed with configuration %v", app.Config)

	return nil
}
