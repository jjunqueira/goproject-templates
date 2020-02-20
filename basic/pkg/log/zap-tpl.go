package log

import (
	"go.uber.org/zap"

	"{{ .ModuleName }}/pkg/config"
)

//Logger for application
var Logger *zap.SugaredLogger

func init() {
	zlogger, err := newDefaultZapLogger()
	if err != nil {
		panic("unable to initialize logger")
	}

	Logger = zlogger
}

// newUnconfiguredZapLogger Creates a new unconfigured logger, this will be used for tests
func newDefaultZapLogger() (*zap.SugaredLogger, error) {
	cfg := zap.NewDevelopmentConfig()

	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.OutputPaths = []string{"stdout"}

	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return zapLogger.Sugar(), nil
}

// Configure configures the default logging based on the application configuration
func Configure(c *config.AppConfig) error {

	var cfg zap.Config
	if c.Debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	switch c.Logging.Level {
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	cfg.OutputPaths = c.Logging.OutputPaths

	zapLogger, err := cfg.Build()
	if err != nil {
		return err
	}

	Logger = zapLogger.Sugar()

	return nil
}
