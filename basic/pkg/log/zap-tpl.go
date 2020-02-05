package log

import (
	"go.uber.org/zap"

	"{{ .ModuleName }}/pkg/config"
)

var logger *zap.SugaredLogger

func init() {
	zlogger, err := newDefaultZapLogger()
	if err != nil {
		panic("unable to initialize logger")
	}

	logger = zlogger
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

	logger = zapLogger.Sugar()

	return nil
}

// Info logs the input at the INFO level
func Info(args ...interface{}) {
	logger.Info(args)
}

// Infof formats the input and logs the input at the INFO level
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args)
}

// Infow logs message for formatted keys and values at the INFO level
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues)
}

// Debug logs the input at the DEBUG level
func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Debugf formats the input and logs the input at the DEBUG level
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args)
}

// Debugw logs message for formatted keys and values at the DEBUG level
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues)
}

// Error logs the input at the ERROR level
func Error(args ...interface{}) {
	logger.Error(args)
}

// Errorf formats the input and logs the input at the ERROR level
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args)
}

// Errorw logs message for formatted keys and values at the ERROR level
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues)
}
