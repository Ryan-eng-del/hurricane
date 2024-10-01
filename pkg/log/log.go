package log

import (
	"log"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger
var _logOnce sync.Once

func setGlobalZapLogger(zapCores []zapcore.Core) {
	_logOnce.Do(func() {
		if Log == nil {
			zapCore := zapcore.NewTee(zapCores...)
			Logger := zap.New(zapCore, zap.AddCaller())
			zap.ReplaceGlobals(Logger)
			Log = Logger.Sugar()
		}
	})
}

func New(opt *Options) {
	installOptions(opt)
}

func installOptions(option *Options) {

	var (
		zapCores   = make([]zapcore.Core, 0, 1)
		enableFile = option.EnableFile
	)

	if enableFile {
		zapCores = setFileLogger(option)
		log.Println(zapCores, "zapcores")
	}

	zapCores = append(zapCores, setDisableFileLogger(option)...)
	setGlobalZapLogger(zapCores)
}
