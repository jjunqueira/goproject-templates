package main

import (
	"{{ .ModuleName }}/pkg/config"
	"{{ .ModuleName }}/pkg/log"
)

// Config is the application wide configuration settings
var Config config.AppConfig

// NewApp constructs a new App instanace and its dependencies
func NewApp(configPath string) (*App, error) {

	c, err := newConfig("/etc/{{ .Name }}/", "/usr/share/{{ .Name }}", configPath)
	if err != nil {
		return nil, err
	}

	Config = c

	err = log.Configure(c)
	if err != nil {
		return err
	}

	return a, err
}

func newConfig(searchpaths ...string) (*config.Config, error) {
	return config.NewViperConfig(searchpaths)
}
