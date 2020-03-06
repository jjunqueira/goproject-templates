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

//WithCtx for application
func WithCtx(ctx context.Context) *zap.SugaredLogger {
	newLogger := logger
	if ctx != nil {
		if ctxRqID, ok := ctx.Value(RequestIDKey).(string); ok {
			newLogger = newLogger.With(zap.String("request_id", ctxRqID))
		}
	}
	return newLogger
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
	cfg.InitialFields["compiler"] = runtime.Version()
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

	logger = zapLogger.WithOptions(zap.AddCallerSkip(2)).Sugar()

	return nil
}

func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.DPanicw(msg, keysAndValues...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}
