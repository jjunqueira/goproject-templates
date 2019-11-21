package main

import (
	"{{ .ModuleName }}/pkg/config"
	"{{ .ModuleName }}/pkg/logging"
	"{{ .ModuleName }}/pkg/metrics"
)

// App contains application dependencies
type App struct {
	Config  *config.Config
	Log     logging.Logger
	Metrics metrics.Metrics
}

// NewApp constructs a new App instanace and its dependencies
func NewApp(configPath string) (*App, error) {
	a := new(App)

	m, err := newMetrics()
	if err != nil {
		return nil, err
	}

	a.Metrics = m

	c, err := newConfig("/etc/{{ .Name }}/", "/usr/share/{{ .Name }}", configPath)
	if err != nil {
		return nil, err
	}

	a.Config = c

	l, err := newLogger(c)
	if err != nil {
		return nil, err
	}

	a.Log = l

	a.Log.Info("Application bootstrap completed with configuration %v", a.Config)

	return a, err
}

func newConfig(searchpaths ...string) (*config.Config, error) {
	return config.NewViperConfig(searchpaths)
}

func newLogger(c *config.Config) (logging.Logger, error) {
	return logging.NewZapLogger(c)
}

func newMetrics() (metrics.Metrics, error) {
	return metrics.NewPrometheusMetrics()
}
