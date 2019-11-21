package main

import (
	"{{ .ModuleName }}/pkg/logging"
	"{{ .ModuleName }}/pkg/config"
	"{{ .ModuleName }}/pkg/metrics"
)

// App contains application dependencies
type App struct {
	Metrics *Metrics
	Config  *Config
	Log     Logger
}

// NewApp constructs a new App instanace and its dependencies
func NewApp() (*App, error) {
	a := new(App)

	m, err := newMetrics()
	if err != nil {
		return nil, err
	}

	a.Metrics = m

	c, err := newConfig("/etc/{{ .Name }}/config.toml")
	if err != nil {
		return nil, err
	}

	a.Config = c

	l, err := newLogger(c.logging)
	if err != nil {
		return nil, err
	}

	a.Log = l

	a.Log.Info("Application bootstrap complete")

	return a, err
}

func newConfig(searchpaths ...string) (*Config, error) {
	return config.NewViperConfig(searchpaths)
}

func newLogger(settings logSettings) (Logger, error) {
	return logging.NewZapLogger()
}

func newMetrics() (Metrics, error) {
	return logging.NewZapLogger()
}