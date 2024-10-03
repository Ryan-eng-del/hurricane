package master

import (
	"hurricane/internal/master/config"
	"hurricane/internal/master/options"
	"hurricane/pkg/app"
	"hurricane/pkg/log"
)

func RunServer(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.NewWithOptions(log.WithEnableColor(), log.WithDebugMode(), log.WithEnableFile())
		defer log.Sync()

		cfg, err := config.CreateConfigFromOptions(opts)

		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
