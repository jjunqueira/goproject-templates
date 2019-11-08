package app

// Config application configs
type Config struct{}

func newConfig(searchpaths ...string) (*Config, error) {
	c := new(Config)
	return c, nil
}
