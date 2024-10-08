package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewStdWithOptions(options ...Option) {
	option := NewOption()

	for _, o := range options {
		o.apply(option)
	}

	consoleZapcore := setDisableFileLogger(option)
	zapCore := zapcore.NewTee(consoleZapcore...)
	Logger := zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.PanicLevel))
	StdLog = Logger.Sugar()
}
