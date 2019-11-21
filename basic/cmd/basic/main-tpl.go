package main

import (
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
	app, err := NewApp()
	if err != nil {
		return err
	}

	fmt.Printf("app %v", app)

	return nil
}
