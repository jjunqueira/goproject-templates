package log

import (
	"context"
	"os"
	"path"
	"runtime"

	"go.uber.org/zap"

	"{{ .ModuleName }}/pkg/config"
)

//CorrelationID for application logs
type CorrelationID int

const (
	//RequestIDKey for application logs
	RequestIDKey CorrelationID = iota
)

//Logger for application
var logger *zap.SugaredLogger

func init() {
	zlogger, err := newDefaultZapLogger()
	if err != nil {
		panic("unable to initialize logger")
	}

	logger = zlogger
}

//CtxLog for application
func CtxLog(ctx context.Context) *zap.SugaredLogger {
	newLogger := logger
	if ctx != nil {
		if ctxRqID, ok := ctx.Value(RequestIDKey).(string); ok {
			newLogger = newLogger.With(zap.String("request_id", ctxRqID))
		}
	}
	return newLogger
}

//Logger for application
func Logger() *zap.SugaredLogger {
	return logger
}

// newDefaultZapLogger Creates a new un-configured logger, this will be used for tests
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
func Configure(version string, build string, c *config.AppConfig) error {

	var cfg zap.Config
	if c.Debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.InitialFields = make(map[string]interface{})
	cfg.InitialFields["app"] = path.Base(os.Args[0])
	cfg.InitialFields["app_version"] = version
	cfg.InitialFields["app_build"] = build
	cfg.InitialFields["app_location"] = os.Args[0]
	cfg.InitialFields["go_version"] = runtime.Version()
	cfg.InitialFields["pid"] = os.Getpid()

	var name string
	name, err := os.Hostname()
	if err != nil {
		name = ""
	}

	if name != "" {
		cfg.InitialFields["server"] = name
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
