package app

import (
	"go.uber.org/zap"
)

// Logger custom logging interface
type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
}

type logSettings struct {
	Level       string   `mapstructure:"level" json:"level"`
	OutputPaths []string `mapstructure:"outputpaths" json:"outputPaths"`
}

type customLogger struct {
	log *zap.SugaredLogger
}

func newLogger(settings logSettings) (*customLogger, error) {
	logger := new(customLogger)

	cfg := zap.NewDevelopmentConfig()

	switch settings.Level {
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	cfg.OutputPaths = settings.OutputPaths

	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	logger.log = zapLogger.Sugar()

	return logger, nil
}

func (l *customLogger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *customLogger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args)
}

func (l *customLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.log.Infow(msg, keysAndValues)
}

func (l *customLogger) Debug(args ...interface{}) {
	l.log.Debug(args)
}

func (l *customLogger) Debugf(template string, args ...interface{}) {
	l.log.Debugf(template, args)
}

func (l *customLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.log.Debugw(msg, keysAndValues)
}

func (l *customLogger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *customLogger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args)
}

func (l *customLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.log.Errorw(msg, keysAndValues)
}
