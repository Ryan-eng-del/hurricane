package log

import (
	"os"
	"path/filepath"

	"go.uber.org/zap/zapcore"
)

const (
	defaultFormat      = "text"
	defaultEnableColor = false
	defaultEnableFile  = false
)

var (
	defaultDebugFilePath = getDefaultFilePath("hurricane.deb.log")
	defaultInfoFilePath  = getDefaultFilePath("hurricane.inf.log")
	defaultErrorFilePath = getDefaultFilePath("hurricane.err.log")
)

func getDefaultFilePath(name string) string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, name)
}

type Options struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `mapstructure:"max-size" json:"max-size"`
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `mapstructure:"max-age" json:"max-age"`
	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups    int    `mapstructure:"max-backups" json:"max-backups"`
	DebugFilePath string `mapstructure:"debug-file-path" json:"debug-file-path"`
	Format        string `json:"format"             mapstructure:"format"`
	Layout        string `json:"layout" mapstructure:"layout"`
	EnableFile    bool   `json:"enable-file" mapstructure:"enable-file"`
	InfoFilePath  string `mapstructure:"info-file-path" json:"info-file-path"`
	ErrorFilePath string `mapstructure:"error-file-path" json:"error-file-path"`
	EnableColor   bool   `json:"enable-color"       mapstructure:"enable-color"`
	DebugMode     bool   `json:"debug-mode" mapstructure:"debug-mode"`
}

func NewOption() *Options {
	return &Options{
		Format:      defaultFormat,
		EnableColor: defaultEnableColor,
		EnableFile:  defaultEnableFile,
	}
}

func getDefaultFileOptions(opts *Options) *Options {
	if opts.DebugMode && opts.DebugFilePath == "" {
		opts.DebugFilePath = defaultDebugFilePath
	}

	if opts.ErrorFilePath == "" {
		opts.ErrorFilePath = defaultErrorFilePath
	}

	if opts.InfoFilePath == "" {
		opts.InfoFilePath = defaultInfoFilePath
	}

	return opts
}

func setDisableFileLogger(options *Options) []zapcore.Core {
	var (
		zapCores = make([]zapcore.Core, 0, 1)
		minLevel = zapcore.DebugLevel
	)

	if !options.DebugMode {
		minLevel = zapcore.InfoLevel
	}

	stdCore := zapcore.NewCore(getDefaultEncoder(options), zapcore.Lock(os.Stdout), minLevel)
	return append(zapCores, stdCore)
}

func setFileLogger(options *Options) []zapcore.Core {
	var (
		zapCores = make([]zapcore.Core, 0, 4)
	)
	options = getDefaultFileOptions(options)

	if options.DebugMode {
		zapCores = append(zapCores, getDebugCore(options))
	}

	return append(zapCores, getInfoCore(options), getWarnCore(options))
}
