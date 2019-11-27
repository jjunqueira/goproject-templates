package logging

import (
	"go.uber.org/zap"

	"{{ .ModuleName }}/pkg/config"
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

// ZapLogger Logger implementation
type ZapLogger struct {
	log *zap.SugaredLogger
}

// NewZapLogger constructs a new Zap based logger instance
func NewZapLogger(c *config.Config) (*ZapLogger, error) {
	logger := new(ZapLogger)

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
		return nil, err
	}

	logger.log = zapLogger.Sugar()

	return logger, nil
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args)
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.log.Infow(msg, keysAndValues)
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.log.Debug(args)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.log.Debugf(template, args)
}

func (l *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.log.Debugw(msg, keysAndValues)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args)
}

func (l *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.log.Errorw(msg, keysAndValues)
}

