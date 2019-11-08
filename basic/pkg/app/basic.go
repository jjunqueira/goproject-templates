package app

// App contains application dependencies
type App struct {
	Metrics *Metrics
	Config  *Config
}

// NewApp constructs a new App instanace and its dependencies
func NewApp() (*App, error) {
	a := new(App)

	m, err := newMetrics()
	if err != nil {
		return nil, err
	}

	a.Metrics = m

	c, err := newConfig()
	if err != nil {
		return nil, err
	}

	a.Config = c

	return a, err
}
