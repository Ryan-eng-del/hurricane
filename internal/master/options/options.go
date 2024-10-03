package options

import (
	"hurricane/internal/pkg/options"
	"hurricane/pkg/log"
)

type Options struct {
	GenericServerRunOptions *options.ServerRunOptions       `json:"server"   mapstructure:"server"`
	GRPCOptions             *options.GRPCOptions            `json:"grpc"     mapstructure:"grpc"`
	InsecureServing         *options.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *options.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *options.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	RedisOptions            *options.RedisOptions           `json:"redis"    mapstructure:"redis"`
	EtcdOptions             *options.EtcdOptions            `json:"etcd"     mapstructure:"etcd"`
	Log                     *log.Options                    `json:"log"      mapstructure:"log"`
	FeatureOptions          *options.FeatureOptions         `json:"feature"  mapstructure:"feature"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: options.NewServerRunOptions(),
		GRPCOptions:             options.NewGRPCOptions(),
		InsecureServing:         options.NewInsecureServingOptions(),
		SecureServing:           options.NewSecureServingOptions(),
		MySQLOptions:            options.NewMySQLOptions(),
		RedisOptions:            options.NewRedisOptions(),
		Log:                     log.NewOption(),
		FeatureOptions:          options.NewFeatureOptions(),
	}
	return &o
}
