package zap

import (
	kitZap "github.com/go-kit/kit/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/valonekowd/clean-architecture/infrastructure/log"
)

type zapLogger struct {
	base *zap.Logger
}

func NewLogger(isProd bool) (log.Logger, error) {
	var baseZapLogger *zap.Logger
	var err error

	options := []zap.Option{
		zap.AddCallerSkip(1),
	}

	if isProd {
		baseZapLogger, err = zap.NewProduction(options...)
	} else {
		baseZapLogger, err = zap.NewDevelopment(options...)
	}

	if err != nil {
		return nil, err
	}

	defer baseZapLogger.Sync()

	return &zapLogger{baseZapLogger}, nil
}

func (l zapLogger) Log(keyvals ...interface{}) error {
	return l.Info(keyvals...)
}

func (l zapLogger) Debug(keyvals ...interface{}) error {
	return kitZap.NewZapSugarLogger(l.base, zapcore.DebugLevel).Log(keyvals...)
}

func (l zapLogger) Info(keyvals ...interface{}) error {
	return kitZap.NewZapSugarLogger(l.base, zapcore.InfoLevel).Log(keyvals...)
}

func (l zapLogger) Warn(keyvals ...interface{}) error {
	return kitZap.NewZapSugarLogger(l.base, zapcore.WarnLevel).Log(keyvals...)
}

func (l zapLogger) Error(keyvals ...interface{}) error {
	return kitZap.NewZapSugarLogger(l.base, zapcore.ErrorLevel).Log(keyvals...)
}

func (l zapLogger) Fatal(keyvals ...interface{}) error {
	return kitZap.NewZapSugarLogger(l.base, zapcore.FatalLevel).Log(keyvals...)
}
