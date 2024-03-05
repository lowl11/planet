package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	sugar      *zap.SugaredLogger
	logChannel chan func()
}

var instance *logger

func get() *logger {
	if instance != nil {
		return instance
	}

	config := zap.NewProductionConfig()

	// set encoder
	config.Encoding = "console"
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.RFC3339TimeEncoder,
	}

	// add level
	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.Level = atom

	// build
	zapLogger, _ := config.Build()

	instance = &logger{
		sugar:      zapLogger.Sugar(),
		logChannel: make(chan func()),
	}
	instance.listen()
	return instance
}

func (logger *logger) Debug(args ...any) {
	logger.printLog(func() {
		logger.sugar.Debug(args...)
	})
}

func (logger *logger) Debugf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Debugf(template, args...)
	})
}

func (logger *logger) Info(args ...any) {
	logger.printLog(func() {
		logger.sugar.Info(args...)
	})
}

func (logger *logger) Infof(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Infof(template, args...)
	})
}

func (logger *logger) Warn(args ...any) {
	logger.printLog(func() {
		logger.sugar.Warn(args...)
	})
}

func (logger *logger) Warnf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Warnf(template, args...)
	})
}

func (logger *logger) Error(args ...any) {
	logger.printLog(func() {
		logger.sugar.Error(args...)
	})
}

func (logger *logger) Errorf(template string, args ...any) {
	logger.printLog(func() {
		logger.sugar.Errorf(template, args...)
	})
}

func (logger *logger) Fatal(args ...any) {
	logger.sugar.Fatal(args...)
}

func (logger *logger) Fatalf(template string, args ...any) {
	logger.sugar.Fatalf(template, args...)
}

func (logger *logger) listen() {
	go func() {
		for log := range logger.logChannel {
			log()
		}
	}()
}

func (logger *logger) printLog(logFunc func()) {
	logFunc()
	//logger.logChannel <- logFunc
}
