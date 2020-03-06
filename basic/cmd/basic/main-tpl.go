package main

import (
	"flag"
	"fmt"
	"os"
	"{{ .ModuleName }}/pkg/app"
	"{{ .ModuleName }}/pkg/log"
)

//Populated via LDFLAGS on build
var version string
var build string

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var configPath = flag.String("config", "", "The path to the config file")

	flag.Parse()

	err := app.Init(version, build, *configPath)
	if err != nil {
		return err
	}

	log.Infow("Application bootstrap completed", "config", fmt.Sprintf("%v", app.Config))

	return nil
}
