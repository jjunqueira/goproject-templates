package main

import (
	"flag"
	"fmt"
	"os"
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

	app, err := NewApp(configPath)
	if err != nil {
		return err
	}

	return nil
}
