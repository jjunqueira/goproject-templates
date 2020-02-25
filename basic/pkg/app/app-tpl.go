package app

import (
	"{{ .ModuleName }}/pkg/config"
	"{{ .ModuleName }}/pkg/log"
)

// Config is the application wide configuration settings
var Config config.AppConfig

// Init initializes application dependencies
func Init(version string, build string, configPath string) error {
	c, err := newConfig("/etc/{{ .Name }}/", "/usr/share/{{ .Name }}", ".", configPath)
	if err != nil {
		return err
	}

	Config = *c

	err = log.Configure(version, build, c)
	if err != nil {
		return err
	}

	c.Version = version

	return err
}

func newConfig(searchpaths ...string) (*config.AppConfig, error) {
	return config.NewViperConfig(searchpaths)
}
