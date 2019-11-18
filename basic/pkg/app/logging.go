package app

import (
	"encoding/json"

	"go.uber.org/zap"
)

type logSettings struct {
	level       string
	encoding    string
	outputPaths []string `mapstructure:"outputpaths"`
	initialFields map[string]string
}

type customLogger struct {
	*SugaredLogger
}

func NewLogger(settings logSettings) (customLogger, error) {
	logger := new(customLogger)
	
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	var cfg zap.Config
	if err := json.Unmarshal(settingsJSON, &cfg); err != nil {
		return nil, err
	}

	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	logger.SugaredLogger = zapLogger.Sugar()

	return logger, nil
}