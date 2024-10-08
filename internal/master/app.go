package master

import (
	"hurricane/internal/master/config"
	"hurricane/internal/master/options"
	"hurricane/pkg/app"
	"hurricane/pkg/log"
)

const commandDesc = `Hurricane, a distributed cron tool in Golang. The master node is responsible for accepting tasks and distributing them to the workers using etcd as a coordination service. Upon receiving the tasks, the worker nodes execute them accordingly, ensuring efficient task management and execution in a distributed environment.

Find more hurricane information at:
    https://github.com/Ryan-eng-del/Hurricane`

// NewApp creates an App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("Hurricane Distributed CronTab Application",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.New(opts.Log)
		defer log.Sync()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
