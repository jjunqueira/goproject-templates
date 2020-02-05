package config

import "github.com/spf13/viper"

// AppConfig application configs
type AppConfig struct {
	Debug   bool `mapstructure:"debug" json:"debug"`
	Logging LogSettings
}

// LogSettings logging settings
type LogSettings struct {
	Level       string   `mapstructure:"level" json:"level"`
	OutputPaths []string `mapstructure:"outputpaths" json:"outputPaths"`
}

// NewViperConfig constructs a new viper configuration loader
func NewViperConfig(searchpaths []string) (*AppConfig, error) {
	c := new(AppConfig)

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
