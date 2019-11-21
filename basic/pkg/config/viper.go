package app

import "github.com/spf13/viper"

// Config application configs
type Config struct {
	Logging LogSettings
}

// NewViperConfig construcs a new viper configuration loader
func NewViperConfig(searchpaths ...string) (*Config, error) {
	c := new(Config)

	viper.AutomaticEnv()
	viper.SetConfigName("config")

	for _, path := range searchpaths {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}